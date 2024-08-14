// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.3
// source: proto/stats.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// StaticInfo is the single node register request with the hub
type StaticInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`                         // the ID of the node
	GnoVersion string `protobuf:"bytes,2,opt,name=gno_version,json=gnoVersion,proto3" json:"gno_version,omitempty"` // the gno version of the node
	OsVersion  string `protobuf:"bytes,3,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`    // the OS info of the node
	Location   string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`                       // the location of the node (Country)
}

func (x *StaticInfo) Reset() {
	*x = StaticInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticInfo) ProtoMessage() {}

func (x *StaticInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticInfo.ProtoReflect.Descriptor instead.
func (*StaticInfo) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{0}
}

func (x *StaticInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StaticInfo) GetGnoVersion() string {
	if x != nil {
		return x.GnoVersion
	}
	return ""
}

func (x *StaticInfo) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *StaticInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

// DynamicInfo is the single node data push
type DynamicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`                             // the ID of the node
	Moniker     string     `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`                             // the moniker of the node
	IsValidator bool       `protobuf:"varint,3,opt,name=is_validator,json=isValidator,proto3" json:"is_validator,omitempty"` // flag indicating valset inclusion
	NetInfo     *NetInfo   `protobuf:"bytes,4,opt,name=net_info,json=netInfo,proto3" json:"net_info,omitempty"`              // the active p2p information
	PendingTxs  uint64     `protobuf:"varint,5,opt,name=pending_txs,json=pendingTxs,proto3" json:"pending_txs,omitempty"`    // the number of currently pending txs
	BlockInfo   *BlockInfo `protobuf:"bytes,6,opt,name=block_info,json=blockInfo,proto3" json:"block_info,omitempty"`        // the latest block information
}

func (x *DynamicInfo) Reset() {
	*x = DynamicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicInfo) ProtoMessage() {}

func (x *DynamicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicInfo.ProtoReflect.Descriptor instead.
func (*DynamicInfo) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{1}
}

func (x *DynamicInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DynamicInfo) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *DynamicInfo) GetIsValidator() bool {
	if x != nil {
		return x.IsValidator
	}
	return false
}

func (x *DynamicInfo) GetNetInfo() *NetInfo {
	if x != nil {
		return x.NetInfo
	}
	return nil
}

func (x *DynamicInfo) GetPendingTxs() uint64 {
	if x != nil {
		return x.PendingTxs
	}
	return 0
}

func (x *DynamicInfo) GetBlockInfo() *BlockInfo {
	if x != nil {
		return x.BlockInfo
	}
	return nil
}

// BlockInfo is the latest node block information
type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number    uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`                        // the latest block height
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                  // the latest block timestamp (unix)
	GasUsed   uint64 `protobuf:"varint,3,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`       // the gas used in the block
	GasWanted uint64 `protobuf:"varint,4,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"` // the gas wanted in the block (limit)
	Proposer  string `protobuf:"bytes,5,opt,name=proposer,proto3" json:"proposer,omitempty"`                     // the proposer of the block
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{2}
}

func (x *BlockInfo) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BlockInfo) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockInfo) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *BlockInfo) GetGasWanted() uint64 {
	if x != nil {
		return x.GasWanted
	}
	return 0
}

func (x *BlockInfo) GetProposer() string {
	if x != nil {
		return x.Proposer
	}
	return ""
}

// NetInfo is the latest node p2p information
type NetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P2PAddress string      `protobuf:"bytes,1,opt,name=p2p_address,json=p2pAddress,proto3" json:"p2p_address,omitempty"` // the network p2p address of the node
	Peers      []*PeerInfo `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`                             // the peer information
}

func (x *NetInfo) Reset() {
	*x = NetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetInfo) ProtoMessage() {}

func (x *NetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetInfo.ProtoReflect.Descriptor instead.
func (*NetInfo) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{3}
}

func (x *NetInfo) GetP2PAddress() string {
	if x != nil {
		return x.P2PAddress
	}
	return ""
}

func (x *NetInfo) GetPeers() []*PeerInfo {
	if x != nil {
		return x.Peers
	}
	return nil
}

// PeerInfo is information relating to a single peer
type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P2PAddress string `protobuf:"bytes,1,opt,name=p2p_address,json=p2pAddress,proto3" json:"p2p_address,omitempty"` // the p2p address of the peer
	Moniker    string `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`                         // the moniker of the peer
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{4}
}

func (x *PeerInfo) GetP2PAddress() string {
	if x != nil {
		return x.P2PAddress
	}
	return ""
}

func (x *PeerInfo) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

// DataPoint is the newest data point for a specific node,
// that wraps the dynamic and static node info
type DataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DynamicInfo *DynamicInfo `protobuf:"bytes,1,opt,name=dynamic_info,json=dynamicInfo,proto3" json:"dynamic_info,omitempty"` // the node's dynamic info
	StaticInfo  *StaticInfo  `protobuf:"bytes,2,opt,name=static_info,json=staticInfo,proto3" json:"static_info,omitempty"`    // the node's static info
}

func (x *DataPoint) Reset() {
	*x = DataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPoint) ProtoMessage() {}

func (x *DataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPoint.ProtoReflect.Descriptor instead.
func (*DataPoint) Descriptor() ([]byte, []int) {
	return file_proto_stats_proto_rawDescGZIP(), []int{5}
}

func (x *DataPoint) GetDynamicInfo() *DynamicInfo {
	if x != nil {
		return x.DynamicInfo
	}
	return nil
}

func (x *DataPoint) GetStaticInfo() *StaticInfo {
	if x != nil {
		return x.StaticInfo
	}
	return nil
}

var File_proto_stats_proto protoreflect.FileDescriptor

var file_proto_stats_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6e, 0x6f,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x6e, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x73,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd5, 0x01, 0x0a, 0x0b, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x08,
	0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x78, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x78, 0x73, 0x12, 0x29, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x97, 0x01,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x67, 0x61, 0x73, 0x5f, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x67, 0x61, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x32, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x32, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x22, 0x45, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x32, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x32, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x22, 0x6a, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xa1, 0x01, 0x0a, 0x03, 0x48, 0x75, 0x62, 0x12,
	0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x0b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stats_proto_rawDescOnce sync.Once
	file_proto_stats_proto_rawDescData = file_proto_stats_proto_rawDesc
)

func file_proto_stats_proto_rawDescGZIP() []byte {
	file_proto_stats_proto_rawDescOnce.Do(func() {
		file_proto_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stats_proto_rawDescData)
	})
	return file_proto_stats_proto_rawDescData
}

var file_proto_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_stats_proto_goTypes = []interface{}{
	(*StaticInfo)(nil),    // 0: StaticInfo
	(*DynamicInfo)(nil),   // 1: DynamicInfo
	(*BlockInfo)(nil),     // 2: BlockInfo
	(*NetInfo)(nil),       // 3: NetInfo
	(*PeerInfo)(nil),      // 4: PeerInfo
	(*DataPoint)(nil),     // 5: DataPoint
	(*emptypb.Empty)(nil), // 6: google.protobuf.Empty
}
var file_proto_stats_proto_depIdxs = []int32{
	3, // 0: DynamicInfo.net_info:type_name -> NetInfo
	2, // 1: DynamicInfo.block_info:type_name -> BlockInfo
	4, // 2: NetInfo.peers:type_name -> PeerInfo
	1, // 3: DataPoint.dynamic_info:type_name -> DynamicInfo
	0, // 4: DataPoint.static_info:type_name -> StaticInfo
	6, // 5: Hub.GetDataStream:input_type -> google.protobuf.Empty
	0, // 6: Hub.Register:input_type -> StaticInfo
	1, // 7: Hub.PushData:input_type -> DynamicInfo
	5, // 8: Hub.GetDataStream:output_type -> DataPoint
	6, // 9: Hub.Register:output_type -> google.protobuf.Empty
	6, // 10: Hub.PushData:output_type -> google.protobuf.Empty
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_stats_proto_init() }
func file_proto_stats_proto_init() {
	if File_proto_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_stats_proto_goTypes,
		DependencyIndexes: file_proto_stats_proto_depIdxs,
		MessageInfos:      file_proto_stats_proto_msgTypes,
	}.Build()
	File_proto_stats_proto = out.File
	file_proto_stats_proto_rawDesc = nil
	file_proto_stats_proto_goTypes = nil
	file_proto_stats_proto_depIdxs = nil
}
