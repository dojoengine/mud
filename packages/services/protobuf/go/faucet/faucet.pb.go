// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: proto/faucet.proto

package faucet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LinkedTwitterPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *LinkedTwitterPair) Reset() {
	*x = LinkedTwitterPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedTwitterPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedTwitterPair) ProtoMessage() {}

func (x *LinkedTwitterPair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedTwitterPair.ProtoReflect.Descriptor instead.
func (*LinkedTwitterPair) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{0}
}

func (x *LinkedTwitterPair) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LinkedTwitterPair) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type FaucetStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressToUsername map[string]string `protobuf:"bytes,1,rep,name=addressToUsername,proto3" json:"addressToUsername,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UsernameToAddress map[string]string `protobuf:"bytes,2,rep,name=usernameToAddress,proto3" json:"usernameToAddress,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Username to timestamp of latest drip.
	LatestDrip map[string]int64 `protobuf:"bytes,3,rep,name=latestDrip,proto3" json:"latestDrip,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Global drip counter.
	TotalDripCount uint64 `protobuf:"varint,4,opt,name=totalDripCount,proto3" json:"totalDripCount,omitempty"`
}

func (x *FaucetStore) Reset() {
	*x = FaucetStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaucetStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaucetStore) ProtoMessage() {}

func (x *FaucetStore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaucetStore.ProtoReflect.Descriptor instead.
func (*FaucetStore) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{1}
}

func (x *FaucetStore) GetAddressToUsername() map[string]string {
	if x != nil {
		return x.AddressToUsername
	}
	return nil
}

func (x *FaucetStore) GetUsernameToAddress() map[string]string {
	if x != nil {
		return x.UsernameToAddress
	}
	return nil
}

func (x *FaucetStore) GetLatestDrip() map[string]int64 {
	if x != nil {
		return x.LatestDrip
	}
	return nil
}

func (x *FaucetStore) GetTotalDripCount() uint64 {
	if x != nil {
		return x.TotalDripCount
	}
	return 0
}

// Request for drip after a successful DripVerifyTweet RPC.
type DripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DripRequest) Reset() {
	*x = DripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DripRequest) ProtoMessage() {}

func (x *DripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DripRequest.ProtoReflect.Descriptor instead.
func (*DripRequest) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{2}
}

func (x *DripRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DripRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Request for initial drip via DripVerifyTweet RPC that requires verifying a tweet
type DripVerifyTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DripVerifyTweetRequest) Reset() {
	*x = DripVerifyTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DripVerifyTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DripVerifyTweetRequest) ProtoMessage() {}

func (x *DripVerifyTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DripVerifyTweetRequest.ProtoReflect.Descriptor instead.
func (*DripVerifyTweetRequest) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{3}
}

func (x *DripVerifyTweetRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DripVerifyTweetRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Request for drip to any address when running in dev mode.
type DripDevRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DripDevRequest) Reset() {
	*x = DripDevRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DripDevRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DripDevRequest) ProtoMessage() {}

func (x *DripDevRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DripDevRequest.ProtoReflect.Descriptor instead.
func (*DripDevRequest) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{4}
}

func (x *DripDevRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Response for drip request that contains the transaction hash of the drip tx.
type DripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *DripResponse) Reset() {
	*x = DripResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DripResponse) ProtoMessage() {}

func (x *DripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DripResponse.ProtoReflect.Descriptor instead.
func (*DripResponse) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{5}
}

func (x *DripResponse) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

// Response for the time until next drip request.
type TimeUntilDripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeUntilDripMinutes float64 `protobuf:"fixed64,1,opt,name=timeUntilDripMinutes,proto3" json:"timeUntilDripMinutes,omitempty"`
	TimeUntilDripSeconds float64 `protobuf:"fixed64,2,opt,name=timeUntilDripSeconds,proto3" json:"timeUntilDripSeconds,omitempty"`
}

func (x *TimeUntilDripResponse) Reset() {
	*x = TimeUntilDripResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeUntilDripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeUntilDripResponse) ProtoMessage() {}

func (x *TimeUntilDripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeUntilDripResponse.ProtoReflect.Descriptor instead.
func (*TimeUntilDripResponse) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{6}
}

func (x *TimeUntilDripResponse) GetTimeUntilDripMinutes() float64 {
	if x != nil {
		return x.TimeUntilDripMinutes
	}
	return 0
}

func (x *TimeUntilDripResponse) GetTimeUntilDripSeconds() float64 {
	if x != nil {
		return x.TimeUntilDripSeconds
	}
	return 0
}

type GetLinkedTwittersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLinkedTwittersRequest) Reset() {
	*x = GetLinkedTwittersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkedTwittersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkedTwittersRequest) ProtoMessage() {}

func (x *GetLinkedTwittersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkedTwittersRequest.ProtoReflect.Descriptor instead.
func (*GetLinkedTwittersRequest) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{7}
}

type GetLinkedTwittersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkedTwitters []*LinkedTwitterPair `protobuf:"bytes,1,rep,name=linkedTwitters,proto3" json:"linkedTwitters,omitempty"`
}

func (x *GetLinkedTwittersResponse) Reset() {
	*x = GetLinkedTwittersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkedTwittersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkedTwittersResponse) ProtoMessage() {}

func (x *GetLinkedTwittersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkedTwittersResponse.ProtoReflect.Descriptor instead.
func (*GetLinkedTwittersResponse) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{8}
}

func (x *GetLinkedTwittersResponse) GetLinkedTwitters() []*LinkedTwitterPair {
	if x != nil {
		return x.LinkedTwitters
	}
	return nil
}

type LinkedTwitterForAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *LinkedTwitterForAddressRequest) Reset() {
	*x = LinkedTwitterForAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedTwitterForAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedTwitterForAddressRequest) ProtoMessage() {}

func (x *LinkedTwitterForAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedTwitterForAddressRequest.ProtoReflect.Descriptor instead.
func (*LinkedTwitterForAddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{9}
}

func (x *LinkedTwitterForAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type LinkedTwitterForAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *LinkedTwitterForAddressResponse) Reset() {
	*x = LinkedTwitterForAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedTwitterForAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedTwitterForAddressResponse) ProtoMessage() {}

func (x *LinkedTwitterForAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedTwitterForAddressResponse.ProtoReflect.Descriptor instead.
func (*LinkedTwitterForAddressResponse) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{10}
}

func (x *LinkedTwitterForAddressResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type LinkedAddressForTwitterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *LinkedAddressForTwitterRequest) Reset() {
	*x = LinkedAddressForTwitterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedAddressForTwitterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedAddressForTwitterRequest) ProtoMessage() {}

func (x *LinkedAddressForTwitterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedAddressForTwitterRequest.ProtoReflect.Descriptor instead.
func (*LinkedAddressForTwitterRequest) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{11}
}

func (x *LinkedAddressForTwitterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type LinkedAddressForTwitterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *LinkedAddressForTwitterResponse) Reset() {
	*x = LinkedAddressForTwitterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_faucet_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedAddressForTwitterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedAddressForTwitterResponse) ProtoMessage() {}

func (x *LinkedAddressForTwitterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_faucet_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedAddressForTwitterResponse.ProtoReflect.Descriptor instead.
func (*LinkedAddressForTwitterResponse) Descriptor() ([]byte, []int) {
	return file_proto_faucet_proto_rawDescGZIP(), []int{12}
}

func (x *LinkedAddressForTwitterResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_proto_faucet_proto protoreflect.FileDescriptor

var file_proto_faucet_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x22, 0x49, 0x0a, 0x11,
	0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xf9, 0x03, 0x0a, 0x0b, 0x46, 0x61, 0x75, 0x63,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x46, 0x61, 0x75, 0x63,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x58, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66,
	0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x70,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x72, 0x69, 0x70, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44,
	0x72, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x44, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x44,
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72,
	0x69, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x43, 0x0a, 0x0b, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4e, 0x0a, 0x16, 0x44, 0x72, 0x69, 0x70,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2a, 0x0a, 0x0e, 0x44, 0x72, 0x69, 0x70,
	0x44, 0x65, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x7f, 0x0a, 0x15,
	0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x44, 0x72, 0x69, 0x70, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x44, 0x72,
	0x69, 0x70, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x74, 0x69, 0x6d,
	0x65, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x44, 0x72, 0x69, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x44, 0x72, 0x69, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x1a, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x1e, 0x4c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3d, 0x0a, 0x1f, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54,
	0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x1e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x1f, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32,
	0xcf, 0x04, 0x0a, 0x0d, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x04, 0x44, 0x72, 0x69, 0x70, 0x12, 0x13, 0x2e, 0x66, 0x61, 0x75, 0x63,
	0x65, 0x74, 0x2e, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x07, 0x44, 0x72, 0x69, 0x70, 0x44, 0x65,
	0x76, 0x12, 0x16, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x44, 0x72, 0x69, 0x70, 0x44,
	0x65, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x61, 0x75, 0x63,
	0x65, 0x74, 0x2e, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0f, 0x44, 0x72, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x44, 0x72,
	0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x44, 0x72,
	0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d,
	0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x44, 0x72, 0x69, 0x70, 0x12, 0x13, 0x2e,
	0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x44, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6f, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x2e,
	0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6f, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x26,
	0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e,
	0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x6f, 0x72,
	0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f,
	0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_faucet_proto_rawDescOnce sync.Once
	file_proto_faucet_proto_rawDescData = file_proto_faucet_proto_rawDesc
)

func file_proto_faucet_proto_rawDescGZIP() []byte {
	file_proto_faucet_proto_rawDescOnce.Do(func() {
		file_proto_faucet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_faucet_proto_rawDescData)
	})
	return file_proto_faucet_proto_rawDescData
}

var file_proto_faucet_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proto_faucet_proto_goTypes = []interface{}{
	(*LinkedTwitterPair)(nil),               // 0: faucet.LinkedTwitterPair
	(*FaucetStore)(nil),                     // 1: faucet.FaucetStore
	(*DripRequest)(nil),                     // 2: faucet.DripRequest
	(*DripVerifyTweetRequest)(nil),          // 3: faucet.DripVerifyTweetRequest
	(*DripDevRequest)(nil),                  // 4: faucet.DripDevRequest
	(*DripResponse)(nil),                    // 5: faucet.DripResponse
	(*TimeUntilDripResponse)(nil),           // 6: faucet.TimeUntilDripResponse
	(*GetLinkedTwittersRequest)(nil),        // 7: faucet.GetLinkedTwittersRequest
	(*GetLinkedTwittersResponse)(nil),       // 8: faucet.GetLinkedTwittersResponse
	(*LinkedTwitterForAddressRequest)(nil),  // 9: faucet.LinkedTwitterForAddressRequest
	(*LinkedTwitterForAddressResponse)(nil), // 10: faucet.LinkedTwitterForAddressResponse
	(*LinkedAddressForTwitterRequest)(nil),  // 11: faucet.LinkedAddressForTwitterRequest
	(*LinkedAddressForTwitterResponse)(nil), // 12: faucet.LinkedAddressForTwitterResponse
	nil,                                     // 13: faucet.FaucetStore.AddressToUsernameEntry
	nil,                                     // 14: faucet.FaucetStore.UsernameToAddressEntry
	nil,                                     // 15: faucet.FaucetStore.LatestDripEntry
}
var file_proto_faucet_proto_depIdxs = []int32{
	13, // 0: faucet.FaucetStore.addressToUsername:type_name -> faucet.FaucetStore.AddressToUsernameEntry
	14, // 1: faucet.FaucetStore.usernameToAddress:type_name -> faucet.FaucetStore.UsernameToAddressEntry
	15, // 2: faucet.FaucetStore.latestDrip:type_name -> faucet.FaucetStore.LatestDripEntry
	0,  // 3: faucet.GetLinkedTwittersResponse.linkedTwitters:type_name -> faucet.LinkedTwitterPair
	2,  // 4: faucet.FaucetService.Drip:input_type -> faucet.DripRequest
	4,  // 5: faucet.FaucetService.DripDev:input_type -> faucet.DripDevRequest
	3,  // 6: faucet.FaucetService.DripVerifyTweet:input_type -> faucet.DripVerifyTweetRequest
	2,  // 7: faucet.FaucetService.TimeUntilDrip:input_type -> faucet.DripRequest
	7,  // 8: faucet.FaucetService.GetLinkedTwitters:input_type -> faucet.GetLinkedTwittersRequest
	9,  // 9: faucet.FaucetService.GetLinkedTwitterForAddress:input_type -> faucet.LinkedTwitterForAddressRequest
	11, // 10: faucet.FaucetService.GetLinkedAddressForTwitter:input_type -> faucet.LinkedAddressForTwitterRequest
	5,  // 11: faucet.FaucetService.Drip:output_type -> faucet.DripResponse
	5,  // 12: faucet.FaucetService.DripDev:output_type -> faucet.DripResponse
	5,  // 13: faucet.FaucetService.DripVerifyTweet:output_type -> faucet.DripResponse
	6,  // 14: faucet.FaucetService.TimeUntilDrip:output_type -> faucet.TimeUntilDripResponse
	8,  // 15: faucet.FaucetService.GetLinkedTwitters:output_type -> faucet.GetLinkedTwittersResponse
	10, // 16: faucet.FaucetService.GetLinkedTwitterForAddress:output_type -> faucet.LinkedTwitterForAddressResponse
	12, // 17: faucet.FaucetService.GetLinkedAddressForTwitter:output_type -> faucet.LinkedAddressForTwitterResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_faucet_proto_init() }
func file_proto_faucet_proto_init() {
	if File_proto_faucet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_faucet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedTwitterPair); i {
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
		file_proto_faucet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaucetStore); i {
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
		file_proto_faucet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DripRequest); i {
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
		file_proto_faucet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DripVerifyTweetRequest); i {
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
		file_proto_faucet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DripDevRequest); i {
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
		file_proto_faucet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DripResponse); i {
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
		file_proto_faucet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeUntilDripResponse); i {
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
		file_proto_faucet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkedTwittersRequest); i {
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
		file_proto_faucet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkedTwittersResponse); i {
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
		file_proto_faucet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedTwitterForAddressRequest); i {
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
		file_proto_faucet_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedTwitterForAddressResponse); i {
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
		file_proto_faucet_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedAddressForTwitterRequest); i {
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
		file_proto_faucet_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedAddressForTwitterResponse); i {
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
			RawDescriptor: file_proto_faucet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_faucet_proto_goTypes,
		DependencyIndexes: file_proto_faucet_proto_depIdxs,
		MessageInfos:      file_proto_faucet_proto_msgTypes,
	}.Build()
	File_proto_faucet_proto = out.File
	file_proto_faucet_proto_rawDesc = nil
	file_proto_faucet_proto_goTypes = nil
	file_proto_faucet_proto_depIdxs = nil
}
