// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	Account
	Transaction
	BlockHeader
	Block
	NetBlocks
	NetBlock
*/
package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Balance []byte `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce   uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type Transaction struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From      []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Nonce     uint64 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Alg       uint32 `protobuf:"varint,9,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign      []byte `protobuf:"bytes,10,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockHeader struct {
	Hash       []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash []byte `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Nonce      uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Coinbase   []byte `protobuf:"bytes,4,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Timestamp  int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChainId    uint32 `protobuf:"varint,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	StateRoot  []byte `protobuf:"bytes,7,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	TxsRoot    []byte `protobuf:"bytes,8,opt,name=txs_root,json=txsRoot,proto3" json:"txs_root,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{2} }

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeader) GetTxsRoot() []byte {
	if m != nil {
		return m.TxsRoot
	}
	return nil
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{3} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type NetBlocks struct {
	From   string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Blocks []*Block `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *NetBlocks) Reset()                    { *m = NetBlocks{} }
func (m *NetBlocks) String() string            { return proto.CompactTextString(m) }
func (*NetBlocks) ProtoMessage()               {}
func (*NetBlocks) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{4} }

func (m *NetBlocks) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *NetBlocks) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type NetBlock struct {
	From  string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Block *Block `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *NetBlock) Reset()                    { *m = NetBlock{} }
func (m *NetBlock) String() string            { return proto.CompactTextString(m) }
func (*NetBlock) ProtoMessage()               {}
func (*NetBlock) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{5} }

func (m *NetBlock) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *NetBlock) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "corepb.Account")
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
	proto.RegisterType((*BlockHeader)(nil), "corepb.BlockHeader")
	proto.RegisterType((*Block)(nil), "corepb.Block")
	proto.RegisterType((*NetBlocks)(nil), "corepb.NetBlocks")
	proto.RegisterType((*NetBlock)(nil), "corepb.NetBlock")
}

func init() { proto.RegisterFile("block.proto", fileDescriptorBlock) }

var fileDescriptorBlock = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x25, 0x93, 0x49, 0x26, 0xa9, 0xcc, 0x8a, 0xb4, 0x1e, 0x5a, 0x51, 0x0c, 0x11, 0x61, 0x40,
	0x98, 0x83, 0x1e, 0xc4, 0xa3, 0x0a, 0xb2, 0x5e, 0x3c, 0x34, 0xde, 0x87, 0x4e, 0x4f, 0xbb, 0x13,
	0x4c, 0xba, 0x43, 0xba, 0x56, 0xf6, 0x83, 0x05, 0x7f, 0x43, 0xba, 0x3a, 0x99, 0x49, 0xdc, 0xbd,
	0x55, 0xbd, 0xaa, 0x7a, 0x95, 0x57, 0xfd, 0x02, 0x45, 0xdd, 0x5a, 0xf5, 0x6b, 0xdf, 0x0f, 0x16,
	0x2d, 0x4b, 0x95, 0x1d, 0x74, 0x5f, 0x57, 0x1f, 0x61, 0xf3, 0x49, 0x29, 0x7b, 0x6b, 0x90, 0x71,
	0xd8, 0xd4, 0xb2, 0x95, 0x46, 0x69, 0x1e, 0x95, 0xd1, 0x6e, 0x2b, 0xa6, 0x94, 0x3d, 0x85, 0xc4,
	0x58, 0x8f, 0xaf, 0xca, 0x68, 0xb7, 0x16, 0x21, 0xa9, 0xfe, 0x44, 0x50, 0xfc, 0x18, 0xa4, 0x71,
	0x52, 0x61, 0x63, 0x0d, 0x63, 0xb0, 0x3e, 0x49, 0x77, 0x1a, 0x87, 0x29, 0xf6, 0xd8, 0xcf, 0xc1,
	0x76, 0x34, 0xb8, 0x15, 0x14, 0xb3, 0x47, 0xb0, 0x42, 0xcb, 0x63, 0x42, 0x56, 0x68, 0x3d, 0xfb,
	0x6f, 0xd9, 0xde, 0x6a, 0xbe, 0x26, 0x28, 0x24, 0x97, 0x9d, 0xc9, 0x6c, 0x27, 0x7b, 0x01, 0x39,
	0x36, 0x9d, 0x76, 0x28, 0xbb, 0x9e, 0xa7, 0x65, 0xb4, 0x8b, 0xc5, 0x05, 0xf0, 0xdb, 0x8e, 0x12,
	0x25, 0xdf, 0x84, 0x6d, 0x3e, 0x66, 0xcf, 0x20, 0x53, 0x27, 0xd9, 0x98, 0x43, 0x73, 0xe4, 0x59,
	0x19, 0xed, 0xae, 0xc4, 0x86, 0xf2, 0x6f, 0x47, 0xf6, 0x18, 0x62, 0xd9, 0xde, 0xf0, 0x9c, 0x50,
	0x1f, 0x7a, 0x02, 0xd7, 0xdc, 0x18, 0x0e, 0x81, 0xc0, 0xc7, 0xd5, 0xdf, 0x08, 0x8a, 0xcf, 0xfe,
	0x72, 0xd7, 0x5a, 0x1e, 0xf5, 0xf0, 0xa0, 0xcc, 0x57, 0x50, 0xf4, 0x72, 0xd0, 0x06, 0x0f, 0x54,
	0x0a, 0x6a, 0x21, 0x40, 0xd7, 0xbe, 0xe1, 0xac, 0x26, 0x9e, 0xab, 0x79, 0x0e, 0x99, 0xb2, 0x8d,
	0xa9, 0xa5, 0x9b, 0xc4, 0x9f, 0xf3, 0xa5, 0xd2, 0xe4, 0x7f, 0xa5, 0x73, 0x55, 0xe9, 0x52, 0xd5,
	0x4b, 0x00, 0x87, 0x12, 0xf5, 0x61, 0xb0, 0x16, 0xc7, 0x53, 0xe4, 0x84, 0x08, 0x6b, 0xd1, 0x4f,
	0xe2, 0x9d, 0x0b, 0xc5, 0x2c, 0x3c, 0x33, 0xde, 0x39, 0x5f, 0xaa, 0x3a, 0x48, 0x48, 0x28, 0x7b,
	0x0b, 0xe9, 0x89, 0xc4, 0x92, 0xc8, 0xe2, 0xdd, 0x93, 0x7d, 0x70, 0xcb, 0x7e, 0x76, 0x07, 0x31,
	0xb6, 0xb0, 0x0f, 0xb0, 0xc5, 0x8b, 0x0b, 0x1c, 0x5f, 0x95, 0xf1, 0x7c, 0x64, 0xe6, 0x10, 0xb1,
	0x68, 0xac, 0xbe, 0x42, 0xfe, 0x5d, 0x23, 0x51, 0xba, 0xb3, 0x51, 0xfc, 0xc2, 0x7c, 0x34, 0xca,
	0x1b, 0x48, 0xc9, 0xb2, 0x13, 0xe7, 0xd5, 0xe2, 0x33, 0xc4, 0x58, 0xac, 0xbe, 0x40, 0x36, 0xf1,
	0x3c, 0x48, 0xf3, 0x1a, 0x12, 0xea, 0xa4, 0x67, 0xb9, 0xc7, 0x12, 0x6a, 0x75, 0x4a, 0xbf, 0xc5,
	0xfb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x98, 0x89, 0x7c, 0x25, 0x03, 0x00, 0x00,
}
