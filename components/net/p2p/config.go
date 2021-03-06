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
	"time"

	"github.com/multiformats/go-multiaddr"
)

/*
Config config is used to start a local node.
*/
type Config struct {
	bucketsize   int
	latency      time.Duration
	BootNodes    []multiaddr.Multiaddr
	IP           string
	Port         uint
	Randseed     int64
	maxSyncNodes int
}

// DefautConfig defautConfig is the p2p network defaut config
func DefautConfig() *Config {
	//bootNode, err:= multiaddr.NewMultiaddr("/ip4/192.168.2.148/tcp/9999/ipfs/QmYiAecoMfkTroY87LkPFXfNJ2tpJ7M1PHPdPNhonXEBLm")
	//if err != nil {
	//	return nil
	//}
	//return &Config{
	//	30, 10, []multiaddr.Multiaddr{bootNode}, "127.0.0.1", 20000, 1896599, 16,
	//}
	return &Config{
		30, 10, []multiaddr.Multiaddr{}, "127.0.0.1", 9999, 12345, 16,
	}
}
