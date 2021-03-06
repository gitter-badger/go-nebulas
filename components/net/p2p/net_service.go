// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package p2p

import (
	"bytes"
	"encoding/json"
	"errors"
	"hash/crc32"
	"io"
	"strings"
	"sync"

	"github.com/gogo/protobuf/proto"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	nnet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/nebulasio/go-nebulas/components/net"
	"github.com/nebulasio/go-nebulas/components/net/messages"
	"github.com/nebulasio/go-nebulas/components/net/pb"
	byteutils "github.com/nebulasio/go-nebulas/util/byteutils"
	log "github.com/sirupsen/logrus"
)

// connection state
const (
	ProtocolID     = "/neb/1.0.0"
	SNC            = -1
	SHandshaking   = 0
	SOK            = 1
	HELLO          = "hello"
	OK             = "ok"
	BYE            = "bye"
	SYNCROUTE      = "syncroute"
	SYNCROUTEREPLY = "resyncroute"
	CLIENTVERSION  = "0.2.0"
)

// MagicNumber the protocol magic number, A constant numerical or text value used to identify protocol.
var MagicNumber = []byte{0x4e, 0x45, 0x42, 0x31}

// NetService service for nebulas p2p network
type NetService struct {
	node       *Node
	quitCh     chan bool
	dispatcher *net.Dispatcher
}

// NewNetService create netService
func NewNetService(config *Config) *NetService {
	if config == nil {
		config = DefautConfig()
	}
	n, err := NewNode(config)
	if err != nil {
		log.Error("NewNetService: node create fail -> ", err)
	}
	ns := &NetService{n, make(chan bool), net.NewDispatcher()}
	return ns
}

// RegisterNetService register to Netservice
func (ns *NetService) RegisterNetService() *NetService {
	ns.node.host.SetStreamHandler(ProtocolID, ns.streamHandler)

	log.Infof("RegisterNetService: register netservice success")
	return ns
}

// Addrs get node address in string
func (ns *NetService) Addrs() ma.Multiaddr {
	len := len(ns.node.host.Addrs())
	if len > 0 {
		return ns.node.host.Addrs()[0]
	}
	return nil

}

// Node return the peer node
func (ns *NetService) Node() *Node {
	return ns.node
}

func (ns *NetService) streamHandler(s nnet.Stream) {
	go (func() {
	HandleMsg:
		for {
			select {
			case <-ns.quitCh:
				break HandleMsg
			default:
				node := ns.node
				pid := s.Conn().RemotePeer()
				addrs := s.Conn().RemoteMultiaddr()

				key := GenerateKey(addrs, pid)

				dataHeader, err := ReadUint32(s, 36)
				if err != nil {
					log.Error("streamHandler: read data header occurs error, ", err)
					ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					return
				}

				msgName := dataHeader[12:24]
				dataLength := dataHeader[24:28]
				dataChecksum := dataHeader[28:32]

				index := bytes.IndexByte(msgName, 0)
				msgNameByte := msgName[0:index]
				msgNameStr := string(msgNameByte)

				log.WithFields(log.Fields{
					"addrs":   addrs,
					"msgName": msgNameStr,
				}).Info("streamHandler:handle coming msg.")

				if !ns.verifyHeader(dataHeader) {
					ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					return
				}
				data, err := ReadUint32(s, byteutils.Uint32(dataLength))
				if err != nil {
					log.Error("streamHandler: read data occurs error, ", err)
					ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					return
				}
				dataChecksumA := crc32.ChecksumIEEE(data)
				if dataChecksumA != byteutils.Uint32(dataChecksum) {
					log.WithFields(log.Fields{
						"dataChecksumA": dataChecksumA,
						"dataChecksum":  byteutils.Uint32(dataChecksum),
					}).Error("streamHandler: data verification occurs error, dataChecksum is error, the connection will be closed.")
					ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					return
				}

				switch msgNameStr {
				case HELLO:
					if ns.handleHelloMsg(data, msgNameStr, pid, s, addrs) {
						node.stream[key] = s
						node.conn[key] = SOK
						node.routeTable.Update(pid)
					} else {
						ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					}
				case OK:
					if ns.handleOkMsg(data, msgNameStr, pid.String()) {
						node.stream[key] = s
						node.conn[key] = SOK
						node.routeTable.Update(pid)
					} else {
						ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					}
				case BYE:

				case SYNCROUTE:
					if !ns.handleSyncRouteMsg(data, msgNameStr, pid, addrs) {
						ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					}
				case SYNCROUTEREPLY:
					if !ns.handleSyncRouteReplyMsg(data, msgNameStr, pid) {
						ns.Bye(pid, []ma.Multiaddr{addrs}, s)
					}
				default:
					if node.conn[key] != SOK {
						log.Error("peer not shake hand before send message.")
						ns.Bye(pid, []ma.Multiaddr{addrs}, s)
						return
					}
					msg := messages.NewBaseMessage(msgNameStr, data)
					ns.PutMessage(msg)
				}

			}
		}
	})()

}

func (ns *NetService) handleHelloMsg(data []byte, msgNameStr string, pid peer.ID, s nnet.Stream, addrs ma.Multiaddr) bool {
	node := ns.node

	hello := new(messages.HelloMessage)
	pb := new(netpb.Hello)
	if err := proto.Unmarshal(data, pb); err != nil {
		log.Error("streamHandler: [HELLO] handle hello msg occurs error: ", err)
		return false
	}
	if err := hello.FromProto(pb); err != nil {
		log.Error("streamHandler: [HELLO] handle hello msg occurs error: ", err)
		return false
	}

	log.WithFields(log.Fields{
		"hello.NodeID":  hello.NodeID,
		"pid":           pid,
		"ClientVersion": hello.ClientVersion,
	}).Info("streamHandler: [HELLO] receive hello message.")
	if hello.NodeID == pid.String() && hello.ClientVersion == CLIENTVERSION {
		ok := messages.NewHelloMessage(node.id.String(), CLIENTVERSION)
		pbok, err := ok.ToProto()
		okdata, err := proto.Marshal(pbok)
		if err != nil {
			log.Error("streamHandler: [HELLO] send ok message occurs error, ", err)
			return false
		}

		totalData := ns.buildData(okdata, OK)

		if err := Write(s, totalData); err != nil {
			log.Error("streamHandler: [HELLO] write data occurs error, ", err)
			return false
		}
	}
	return true
}

func (ns *NetService) handleOkMsg(data []byte, msgNameStr string, pid string) bool {
	log.Info("streamHandler: [OK] handle ok message.")
	ok := new(messages.HelloMessage)
	pb := new(netpb.Hello)
	if err := proto.Unmarshal(data, pb); err != nil {
		log.Error("streamHandler: [OK] handle ok msg occurs error: ", err)
		return false
	}
	if err := ok.FromProto(pb); err != nil {
		log.Error("streamHandler: [OK] handle ok msg occurs error: ", err)
		return false
	}

	if ok.NodeID == pid && ok.ClientVersion == CLIENTVERSION {
		return true
	}

	log.Error("streamHandler: [OK] get incorrect response")
	return false

}

func (ns *NetService) handleSyncRouteMsg(data []byte, msgNameStr string, pid peer.ID, addrs ma.Multiaddr) bool {
	node := ns.node
	log.Info("streamHandler: [SYNCROUTE] handle sync route message")
	peers := node.routeTable.NearestPeers(kbucket.ConvertPeerID(pid), node.config.maxSyncNodes)
	var peerList []peerstore.PeerInfo
	for i := range peers {
		peerInfo := node.peerstore.PeerInfo(peers[i])
		if len(peerInfo.Addrs) == 0 {
			log.Warn("streamHandler: [SYNCROUTE] addrs is nil")
			continue
		}
		peerList = append(peerList, peerInfo)
	}
	log.WithFields(log.Fields{
		"peerList": peerList,
	}).Info("streamHandler: [SYNCROUTE] handle sync route request.")

	data, err := json.Marshal(peerList)
	if err != nil {
		log.Error("streamHandler: [SYNCROUTE] handle sync route occurs error...", err)
		return false
	}

	totalData := ns.buildData(data, SYNCROUTEREPLY)

	key := GenerateKey(addrs, pid)

	stream := node.stream[key]
	if stream == nil {
		log.Error("streamHandler: [SYNCROUTE] send message occrus error, stream does not exist.")
		return false
	}
	if err := Write(stream, totalData); err != nil {
		log.Error("streamHandler: [SYNCROUTE] write data occurs error, ", err)
		return false
	}
	node.routeTable.Update(pid)
	return true
}

func (ns *NetService) handleSyncRouteReplyMsg(data []byte, msgNameStr string, pid peer.ID) bool {
	node := ns.node
	log.Infof("streamHandler: [SYNCROUTEREPLY] handle sync route reply ")
	var sample []peerstore.PeerInfo

	if err := json.Unmarshal(data, &sample); err != nil {
		log.Error("streamHandler: [SYNCROUTEREPLY] handle sync route reply occurs error, ", err)
		return false
	}
	log.WithFields(log.Fields{
		"sample": sample,
	}).Info("streamHandler: [SYNCROUTEREPLY] handle sync route reply.")

	for i := range sample {
		if node.routeTable.Find(sample[i].ID) != "" || len(sample[i].Addrs) == 0 {
			log.Warnf("streamHandler: [SYNCROUTEREPLY] node %s is already exist in route table", sample[i].ID)
			continue
		}
		// Ping the peer.
		node.peerstore.AddAddr(
			sample[i].ID,
			sample[i].Addrs[0],
			peerstore.TempAddrTTL,
		)

		if err := ns.Hello(sample[i].ID); err != nil {
			log.Errorf("streamHandler: [SYNCROUTEREPLY] ping peer %s fail %s", sample[i].ID, err)
			continue
		}
		node.peerstore.AddAddr(
			sample[i].ID,
			sample[i].Addrs[0],
			peerstore.PermanentAddrTTL,
		)
		// Update the routing table.
		node.routeTable.Update(sample[i].ID)
	}
	return true
}

func (ns *NetService) verifyHeader(dataHeader []byte) bool {
	node := ns.node
	magicNumber := dataHeader[:4]
	chainID := dataHeader[4:8]
	version := []byte{dataHeader[11]}
	dataHeaderChecksum := crc32.ChecksumIEEE(dataHeader[0:32])

	if !byteutils.Equal(MagicNumber, magicNumber) {
		log.Error("verifyHeader: data verification occurs error, magic number is error, the connection will be closed.")
		return false
	}

	if node.chainID != byteutils.Uint32(chainID) {
		log.Error("verifyHeader: data verification occurs error, chainID is error, the connection will be closed.")
		return false
	}

	if !byteutils.Equal([]byte{node.version}, version) {
		log.Error("verifyHeader: data verification occurs error, version is error, the connection will be closed.")
		return false
	}

	if dataHeaderChecksum != byteutils.Uint32(dataHeader[32:36]) {
		log.Error("verifyHeader: data verification occurs error, dataHeaderChecksum is error, the connection will be closed.")
		return false
	}
	return true
}

// Bye say bye to a peer, and close connection.
func (ns *NetService) Bye(pid peer.ID, addrs []ma.Multiaddr, s nnet.Stream) {
	node := ns.node
	node.peerstore.SetAddrs(pid, addrs, 0)
	node.routeTable.Remove(pid)
	delete(node.conn, addrs[0].String())
	delete(node.stream, addrs[0].String())
	s.Close()
}

// SendMsg send message to a peer
func (ns *NetService) SendMsg(msgName string, msg []byte, addrs string) {
	node := ns.node
	// addrs := node.peerstore.PeerInfo(pid).Addrs
	log.WithFields(log.Fields{
		"addrs":   addrs,
		"msgName": msgName,
	}).Info("SendMsg: send message to a peer.")
	if len(addrs) < 0 {
		log.Error("SendMsg: wrong pid addrs")
		return
	}
	data := msg
	totalData := ns.buildData(data, msgName)

	stream := node.stream[addrs]
	if stream == nil {
		log.Error("SendMsg: send message occrus error, stream does not exist.")
		return
	}
	if err := Write(stream, totalData); err != nil {
		log.Error("SendMsg: write data occurs error, ", err)
		return
	}

}

// Hello say hello to a peer
func (ns *NetService) Hello(pid peer.ID) error {
	msgName := HELLO
	node := ns.node
	stream, err := node.host.NewStream(
		node.context,
		pid,
		ProtocolID,
	)
	addrs := node.peerstore.PeerInfo(pid).Addrs
	if err != nil {
		node.peerstore.SetAddrs(pid, addrs, 0)
		node.routeTable.Remove(pid)
		return err
	}
	if len(addrs) < 1 {
		log.Error("Hello: wrong pid addrs")
		return errors.New("wrong pid addrs")
	}

	log.Infof("Hello: say hello addrs -> %s", addrs)

	hello := messages.NewHelloMessage(node.id.String(), CLIENTVERSION)
	pb, _ := hello.ToProto()
	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	totalData := ns.buildData(data, msgName)
	if err := Write(stream, totalData); err != nil {
		log.Error("Hello: write data occurs error, ", err)
		return errors.New("Hello: write data occurs error")
	}
	ns.streamHandler(stream)
	return nil
}

// SyncRoutes sync routing table from a peer
func (ns *NetService) SyncRoutes(pid peer.ID) {
	log.Info("SyncRoutes: begin to sync route from ", pid)
	node := ns.node
	addrs := node.peerstore.PeerInfo(pid).Addrs
	if len(addrs) < 0 {
		log.Error("SyncRoutes: wrong pid addrs")
		node.peerstore.SetAddrs(pid, addrs, 0)
		node.routeTable.Remove(pid)
		return
	}
	data := []byte(SYNCROUTE)
	totalData := ns.buildData(data, SYNCROUTE)

	key := GenerateKey(addrs[0], pid)

	stream := node.stream[key]
	if stream == nil {
		log.Error("SyncRoutes: send message occrus error, stream does not exist.")
		node.peerstore.SetAddrs(pid, addrs, 0)
		node.routeTable.Remove(pid)
		return
	}

	if err := Write(stream, totalData); err != nil {
		log.Error("SyncRoutes: write data occurs error, ", err)
		node.peerstore.SetAddrs(pid, addrs, 0)
		node.routeTable.Remove(pid)
		return
	}

}

// buildHeader build header information
func buildHeader(chainID uint32, msgName string, version byte, dataLength uint32, dataChecksum uint32) []byte {
	var metaHeader = make([]byte, 32)
	msgNameByte := []byte(msgName)

	copy(metaHeader[00:], MagicNumber)
	copy(metaHeader[04:], byteutils.FromUint32(chainID))
	// 64-88 Reserved field
	copy(metaHeader[11:], []byte{version})
	copy(metaHeader[12:], msgNameByte)
	copy(metaHeader[24:], byteutils.FromUint32(dataLength))
	copy(metaHeader[28:], byteutils.FromUint32(dataChecksum))

	return metaHeader
}

func (ns *NetService) buildData(data []byte, msgName string) []byte {
	node := ns.node
	dataChecksum := crc32.ChecksumIEEE(data)
	metaHeader := buildHeader(node.chainID, msgName, node.version, uint32(len(data)), dataChecksum)
	headerChecksum := crc32.ChecksumIEEE(metaHeader)
	metaHeader = append(metaHeader[:], byteutils.FromUint32(headerChecksum)...)
	totalData := append(metaHeader[:], data...)
	return totalData
}

// Start start p2p manager.
func (ns *NetService) Start() {
	// ns.startStreamHandler()
	ns.Launch()
	ns.dispatcher.Start()

}

// Stop stop p2p manager.
func (ns *NetService) Stop() {
	ns.dispatcher.Stop()
	ns.quitCh <- true
}

// Register register the subscribers.
func (ns *NetService) Register(subscribers ...*net.Subscriber) {
	ns.dispatcher.Register(subscribers...)
}

// Deregister Deregister the subscribers.
func (ns *NetService) Deregister(subscribers ...*net.Subscriber) {
	ns.dispatcher.Deregister(subscribers...)
}

// PutMessage put message to dispatcher.
func (ns *NetService) PutMessage(msg net.Message) {
	ns.dispatcher.PutMessage(msg)
}

// Launch start netService
func (ns *NetService) Launch() error {

	node := ns.node
	log.Infof("Launch: node info {id -> %s, address -> %s}", node.id, node.host.Addrs())
	if node.running {
		return errors.New("Launch: node already running")
	}
	node.running = true
	log.Info("Launch: node start to join p2p network...")

	ns.RegisterNetService()

	var wg sync.WaitGroup
	for _, bootNode := range node.config.BootNodes {
		wg.Add(1)
		go func(bootNode ma.Multiaddr) {
			defer wg.Done()
			err := ns.SayHello(bootNode)
			if err != nil {
				log.Error("Launch: can not say hello to trusted node.", bootNode, err)
			}

		}(bootNode)
	}
	wg.Wait()

	go func() {
		ns.Discovery(node.context)
	}()

	log.Infof("Launch: node start and join to p2p network success and listening for connections on port %d... ", node.config.Port)

	return nil
}

// Write write bytes to stream
func Write(writer io.Writer, data []byte) error {
	result := make(chan error, 1)
	go func(writer io.Writer, data []byte) {
		_, err := writer.Write(data)
		result <- err
	}(writer, data)
	err := <-result
	return err
}

// ReadUint32 read bytes from a stream
func ReadUint32(reader io.Reader, n uint32) ([]byte, error) {
	data := make([]byte, n)
	result := make(chan error, 1)
	go func(reader io.Reader) {
		_, err := io.ReadFull(reader, data)
		result <- err
	}(reader)
	err := <-result
	return data, err
}

// GenerateKey generate a key
func GenerateKey(addrs ma.Multiaddr, pid peer.ID) string {
	if len(strings.Split(addrs.String(), "/")) > 2 {
		key := strings.Split(addrs.String(), "/")[2] + pid.String()
		return key
	}
	log.WithFields(log.Fields{
		"addrs": addrs,
		"pid":   pid,
	}).Error("GenerateKey: the addrs format is incorrect.")
	return ""
}
