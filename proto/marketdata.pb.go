// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: marketdata.proto

package proto

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

// Message representing a full snapshot of market data.
type MarketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol          string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	BidPrice        float64 `protobuf:"fixed64,2,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	AskPrice        float64 `protobuf:"fixed64,3,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	LastTradedPrice float64 `protobuf:"fixed64,4,opt,name=last_traded_price,json=lastTradedPrice,proto3" json:"last_traded_price,omitempty"`
	OpenPrice       float64 `protobuf:"fixed64,5,opt,name=open_price,json=openPrice,proto3" json:"open_price,omitempty"`
	HighPrice       float64 `protobuf:"fixed64,6,opt,name=high_price,json=highPrice,proto3" json:"high_price,omitempty"`
	LowPrice        float64 `protobuf:"fixed64,7,opt,name=low_price,json=lowPrice,proto3" json:"low_price,omitempty"`
	ClosePrice      float64 `protobuf:"fixed64,8,opt,name=close_price,json=closePrice,proto3" json:"close_price,omitempty"`
	Volume          int64   `protobuf:"varint,9,opt,name=volume,proto3" json:"volume,omitempty"`
	Exchange        string  `protobuf:"bytes,10,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Timestamp       int64   `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MarketData) Reset() {
	*x = MarketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketData) ProtoMessage() {}

func (x *MarketData) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketData.ProtoReflect.Descriptor instead.
func (*MarketData) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{0}
}

func (x *MarketData) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *MarketData) GetBidPrice() float64 {
	if x != nil {
		return x.BidPrice
	}
	return 0
}

func (x *MarketData) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *MarketData) GetLastTradedPrice() float64 {
	if x != nil {
		return x.LastTradedPrice
	}
	return 0
}

func (x *MarketData) GetOpenPrice() float64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *MarketData) GetHighPrice() float64 {
	if x != nil {
		return x.HighPrice
	}
	return 0
}

func (x *MarketData) GetLowPrice() float64 {
	if x != nil {
		return x.LowPrice
	}
	return 0
}

func (x *MarketData) GetClosePrice() float64 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *MarketData) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *MarketData) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *MarketData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Update for the bid price and bid volume.
type BidUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	BidPrice  float64 `protobuf:"fixed64,2,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	BidVolume int64   `protobuf:"varint,3,opt,name=bid_volume,json=bidVolume,proto3" json:"bid_volume,omitempty"` // Volume for bid.
	Timestamp int64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BidUpdate) Reset() {
	*x = BidUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidUpdate) ProtoMessage() {}

func (x *BidUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidUpdate.ProtoReflect.Descriptor instead.
func (*BidUpdate) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{1}
}

func (x *BidUpdate) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *BidUpdate) GetBidPrice() float64 {
	if x != nil {
		return x.BidPrice
	}
	return 0
}

func (x *BidUpdate) GetBidVolume() int64 {
	if x != nil {
		return x.BidVolume
	}
	return 0
}

func (x *BidUpdate) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Update for the ask price and ask volume.
type AskUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	AskPrice  float64 `protobuf:"fixed64,2,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	AskVolume int64   `protobuf:"varint,3,opt,name=ask_volume,json=askVolume,proto3" json:"ask_volume,omitempty"` // Volume for ask.
	Timestamp int64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AskUpdate) Reset() {
	*x = AskUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskUpdate) ProtoMessage() {}

func (x *AskUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskUpdate.ProtoReflect.Descriptor instead.
func (*AskUpdate) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{2}
}

func (x *AskUpdate) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *AskUpdate) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *AskUpdate) GetAskVolume() int64 {
	if x != nil {
		return x.AskVolume
	}
	return 0
}

func (x *AskUpdate) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Update for the last traded price.
type LastTradeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol          string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	LastTradedPrice float64 `protobuf:"fixed64,2,opt,name=last_traded_price,json=lastTradedPrice,proto3" json:"last_traded_price,omitempty"`
	Volume          int64   `protobuf:"varint,3,opt,name=volume,proto3" json:"volume,omitempty"` // Volume for this trade.
	Timestamp       int64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LastTradeUpdate) Reset() {
	*x = LastTradeUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastTradeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastTradeUpdate) ProtoMessage() {}

func (x *LastTradeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastTradeUpdate.ProtoReflect.Descriptor instead.
func (*LastTradeUpdate) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{3}
}

func (x *LastTradeUpdate) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *LastTradeUpdate) GetLastTradedPrice() float64 {
	if x != nil {
		return x.LastTradedPrice
	}
	return 0
}

func (x *LastTradeUpdate) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *LastTradeUpdate) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Update for the overall trading volume.
type VolumeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Volume    int64  `protobuf:"varint,2,opt,name=volume,proto3" json:"volume,omitempty"` // Total volume for the trading day.
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *VolumeUpdate) Reset() {
	*x = VolumeUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeUpdate) ProtoMessage() {}

func (x *VolumeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeUpdate.ProtoReflect.Descriptor instead.
func (*VolumeUpdate) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{4}
}

func (x *VolumeUpdate) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *VolumeUpdate) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *VolumeUpdate) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Unified update message that can contain any of the specific updates.
type MarketDataUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Types that are assignable to Update:
	//	*MarketDataUpdate_BidUpdate
	//	*MarketDataUpdate_AskUpdate
	//	*MarketDataUpdate_LastTradeUpdate
	//	*MarketDataUpdate_VolumeUpdate
	Update isMarketDataUpdate_Update `protobuf_oneof:"update"`
}

func (x *MarketDataUpdate) Reset() {
	*x = MarketDataUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketDataUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketDataUpdate) ProtoMessage() {}

func (x *MarketDataUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketDataUpdate.ProtoReflect.Descriptor instead.
func (*MarketDataUpdate) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{5}
}

func (x *MarketDataUpdate) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (m *MarketDataUpdate) GetUpdate() isMarketDataUpdate_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (x *MarketDataUpdate) GetBidUpdate() *BidUpdate {
	if x, ok := x.GetUpdate().(*MarketDataUpdate_BidUpdate); ok {
		return x.BidUpdate
	}
	return nil
}

func (x *MarketDataUpdate) GetAskUpdate() *AskUpdate {
	if x, ok := x.GetUpdate().(*MarketDataUpdate_AskUpdate); ok {
		return x.AskUpdate
	}
	return nil
}

func (x *MarketDataUpdate) GetLastTradeUpdate() *LastTradeUpdate {
	if x, ok := x.GetUpdate().(*MarketDataUpdate_LastTradeUpdate); ok {
		return x.LastTradeUpdate
	}
	return nil
}

func (x *MarketDataUpdate) GetVolumeUpdate() *VolumeUpdate {
	if x, ok := x.GetUpdate().(*MarketDataUpdate_VolumeUpdate); ok {
		return x.VolumeUpdate
	}
	return nil
}

type isMarketDataUpdate_Update interface {
	isMarketDataUpdate_Update()
}

type MarketDataUpdate_BidUpdate struct {
	BidUpdate *BidUpdate `protobuf:"bytes,2,opt,name=bid_update,json=bidUpdate,proto3,oneof"`
}

type MarketDataUpdate_AskUpdate struct {
	AskUpdate *AskUpdate `protobuf:"bytes,3,opt,name=ask_update,json=askUpdate,proto3,oneof"`
}

type MarketDataUpdate_LastTradeUpdate struct {
	LastTradeUpdate *LastTradeUpdate `protobuf:"bytes,4,opt,name=last_trade_update,json=lastTradeUpdate,proto3,oneof"`
}

type MarketDataUpdate_VolumeUpdate struct {
	VolumeUpdate *VolumeUpdate `protobuf:"bytes,5,opt,name=volume_update,json=volumeUpdate,proto3,oneof"`
}

func (*MarketDataUpdate_BidUpdate) isMarketDataUpdate_Update() {}

func (*MarketDataUpdate_AskUpdate) isMarketDataUpdate_Update() {}

func (*MarketDataUpdate_LastTradeUpdate) isMarketDataUpdate_Update() {}

func (*MarketDataUpdate_VolumeUpdate) isMarketDataUpdate_Update() {}

// Message representing a request for market data for a specific symbol.
type MarketDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"` // The ticker symbol to request data for.
}

func (x *MarketDataRequest) Reset() {
	*x = MarketDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketDataRequest) ProtoMessage() {}

func (x *MarketDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketDataRequest.ProtoReflect.Descriptor instead.
func (*MarketDataRequest) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{6}
}

func (x *MarketDataRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

// Message representing a response to a market data request.
type MarketDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *MarketData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // Market data for the requested symbol.
}

func (x *MarketDataResponse) Reset() {
	*x = MarketDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketDataResponse) ProtoMessage() {}

func (x *MarketDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketDataResponse.ProtoReflect.Descriptor instead.
func (*MarketDataResponse) Descriptor() ([]byte, []int) {
	return file_marketdata_proto_rawDescGZIP(), []int{7}
}

func (x *MarketDataResponse) GetData() *MarketData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_marketdata_proto protoreflect.FileDescriptor

var file_marketdata_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd8,
	0x02, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x69,
	0x67, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x77,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x6f,
	0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x7d, 0x0a, 0x09, 0x42, 0x69, 0x64,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x69, 0x64, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x62, 0x69, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x7d, 0x0a, 0x09, 0x41, 0x73, 0x6b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73,
	0x6b, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x61, 0x73, 0x6b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x8b, 0x01, 0x0a, 0x0f, 0x4c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x5c, 0x0a, 0x0c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0xb0, 0x02, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x36, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x42, 0x69, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x62,
	0x69, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x61, 0x73, 0x6b, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x73, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x49, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0c,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x22, 0x40, 0x0a, 0x12, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb6, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x30, 0x01, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x65,
	0x69, 0x6c, 0x68, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_marketdata_proto_rawDescOnce sync.Once
	file_marketdata_proto_rawDescData = file_marketdata_proto_rawDesc
)

func file_marketdata_proto_rawDescGZIP() []byte {
	file_marketdata_proto_rawDescOnce.Do(func() {
		file_marketdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_marketdata_proto_rawDescData)
	})
	return file_marketdata_proto_rawDescData
}

var file_marketdata_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_marketdata_proto_goTypes = []interface{}{
	(*MarketData)(nil),         // 0: marketdata.MarketData
	(*BidUpdate)(nil),          // 1: marketdata.BidUpdate
	(*AskUpdate)(nil),          // 2: marketdata.AskUpdate
	(*LastTradeUpdate)(nil),    // 3: marketdata.LastTradeUpdate
	(*VolumeUpdate)(nil),       // 4: marketdata.VolumeUpdate
	(*MarketDataUpdate)(nil),   // 5: marketdata.MarketDataUpdate
	(*MarketDataRequest)(nil),  // 6: marketdata.MarketDataRequest
	(*MarketDataResponse)(nil), // 7: marketdata.MarketDataResponse
}
var file_marketdata_proto_depIdxs = []int32{
	1, // 0: marketdata.MarketDataUpdate.bid_update:type_name -> marketdata.BidUpdate
	2, // 1: marketdata.MarketDataUpdate.ask_update:type_name -> marketdata.AskUpdate
	3, // 2: marketdata.MarketDataUpdate.last_trade_update:type_name -> marketdata.LastTradeUpdate
	4, // 3: marketdata.MarketDataUpdate.volume_update:type_name -> marketdata.VolumeUpdate
	0, // 4: marketdata.MarketDataResponse.data:type_name -> marketdata.MarketData
	6, // 5: marketdata.MarketDataService.GetMarketData:input_type -> marketdata.MarketDataRequest
	6, // 6: marketdata.MarketDataService.StreamMarketData:input_type -> marketdata.MarketDataRequest
	7, // 7: marketdata.MarketDataService.GetMarketData:output_type -> marketdata.MarketDataResponse
	5, // 8: marketdata.MarketDataService.StreamMarketData:output_type -> marketdata.MarketDataUpdate
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_marketdata_proto_init() }
func file_marketdata_proto_init() {
	if File_marketdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_marketdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketData); i {
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
		file_marketdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidUpdate); i {
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
		file_marketdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskUpdate); i {
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
		file_marketdata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastTradeUpdate); i {
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
		file_marketdata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeUpdate); i {
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
		file_marketdata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketDataUpdate); i {
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
		file_marketdata_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketDataRequest); i {
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
		file_marketdata_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketDataResponse); i {
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
	file_marketdata_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*MarketDataUpdate_BidUpdate)(nil),
		(*MarketDataUpdate_AskUpdate)(nil),
		(*MarketDataUpdate_LastTradeUpdate)(nil),
		(*MarketDataUpdate_VolumeUpdate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_marketdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_marketdata_proto_goTypes,
		DependencyIndexes: file_marketdata_proto_depIdxs,
		MessageInfos:      file_marketdata_proto_msgTypes,
	}.Build()
	File_marketdata_proto = out.File
	file_marketdata_proto_rawDesc = nil
	file_marketdata_proto_goTypes = nil
	file_marketdata_proto_depIdxs = nil
}