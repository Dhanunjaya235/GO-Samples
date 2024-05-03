// ./proto/orders/order.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: orders/order.proto

package orders

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	products "rpc/grpc/examples/protogen/golang/products"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    uint64              `protobuf:"varint,1,opt,name=order_id,proto3" json:"order_id,omitempty"`
	CustomerId uint64              `protobuf:"varint,2,opt,name=customer_id,proto3" json:"customer_id,omitempty"`
	IsActive   bool                `protobuf:"varint,3,opt,name=is_active,proto3" json:"is_active,omitempty"`
	Products   []*products.Product `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	OrderDate  *date.Date          `protobuf:"bytes,5,opt,name=order_date,proto3" json:"order_date,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Order) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Order) GetProducts() []*products.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Order) GetOrderDate() *date.Date {
	if x != nil {
		return x.OrderDate
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{1}
}

type ResponseFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Status    uint64 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	OrderId   uint64 `protobuf:"varint,3,opt,name=order_id,proto3" json:"order_id,omitempty"`
}

func (x *ResponseFormat) Reset() {
	*x = ResponseFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFormat) ProtoMessage() {}

func (x *ResponseFormat) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFormat.ProtoReflect.Descriptor instead.
func (*ResponseFormat) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseFormat) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *ResponseFormat) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseFormat) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type PayloadWithSingleOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *PayloadWithSingleOrder) Reset() {
	*x = PayloadWithSingleOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadWithSingleOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadWithSingleOrder) ProtoMessage() {}

func (x *PayloadWithSingleOrder) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadWithSingleOrder.ProtoReflect.Descriptor instead.
func (*PayloadWithSingleOrder) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{3}
}

func (x *PayloadWithSingleOrder) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type PayloadWithMultipelOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *PayloadWithMultipelOrders) Reset() {
	*x = PayloadWithMultipelOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadWithMultipelOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadWithMultipelOrders) ProtoMessage() {}

func (x *PayloadWithMultipelOrders) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadWithMultipelOrders.ProtoReflect.Descriptor instead.
func (*PayloadWithMultipelOrders) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{4}
}

func (x *PayloadWithMultipelOrders) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type PayloadWithOrderID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *PayloadWithOrderID) Reset() {
	*x = PayloadWithOrderID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadWithOrderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadWithOrderID) ProtoMessage() {}

func (x *PayloadWithOrderID) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadWithOrderID.ProtoReflect.Descriptor instead.
func (*PayloadWithOrderID) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{5}
}

func (x *PayloadWithOrderID) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_orders_order_proto protoreflect.FileDescriptor

var file_orders_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x31, 0x0a,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x63, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x36,
	0x0a, 0x16, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x19, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x65, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69,
	0x74, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x32, 0xfb, 0x01, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x50, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f,
	0x76, 0x30, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x57, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x17, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x30, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x46, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x30, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_order_proto_rawDescOnce sync.Once
	file_orders_order_proto_rawDescData = file_orders_order_proto_rawDesc
)

func file_orders_order_proto_rawDescGZIP() []byte {
	file_orders_order_proto_rawDescOnce.Do(func() {
		file_orders_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_order_proto_rawDescData)
	})
	return file_orders_order_proto_rawDescData
}

var file_orders_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orders_order_proto_goTypes = []interface{}{
	(*Order)(nil),                     // 0: Order
	(*Empty)(nil),                     // 1: Empty
	(*ResponseFormat)(nil),            // 2: ResponseFormat
	(*PayloadWithSingleOrder)(nil),    // 3: PayloadWithSingleOrder
	(*PayloadWithMultipelOrders)(nil), // 4: PayloadWithMultipelOrders
	(*PayloadWithOrderID)(nil),        // 5: PayloadWithOrderID
	(*products.Product)(nil),          // 6: Product
	(*date.Date)(nil),                 // 7: google.type.Date
}
var file_orders_order_proto_depIdxs = []int32{
	6, // 0: Order.products:type_name -> Product
	7, // 1: Order.order_date:type_name -> google.type.Date
	0, // 2: PayloadWithSingleOrder.order:type_name -> Order
	0, // 3: PayloadWithMultipelOrders.orders:type_name -> Order
	3, // 4: Orders.AddOrder:input_type -> PayloadWithSingleOrder
	5, // 5: Orders.GetOrder:input_type -> PayloadWithOrderID
	1, // 6: Orders.GetAllOrders:input_type -> Empty
	2, // 7: Orders.AddOrder:output_type -> ResponseFormat
	3, // 8: Orders.GetOrder:output_type -> PayloadWithSingleOrder
	4, // 9: Orders.GetAllOrders:output_type -> PayloadWithMultipelOrders
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_orders_order_proto_init() }
func file_orders_order_proto_init() {
	if File_orders_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_orders_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_orders_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFormat); i {
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
		file_orders_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadWithSingleOrder); i {
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
		file_orders_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadWithMultipelOrders); i {
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
		file_orders_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadWithOrderID); i {
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
			RawDescriptor: file_orders_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_order_proto_goTypes,
		DependencyIndexes: file_orders_order_proto_depIdxs,
		MessageInfos:      file_orders_order_proto_msgTypes,
	}.Build()
	File_orders_order_proto = out.File
	file_orders_order_proto_rawDesc = nil
	file_orders_order_proto_goTypes = nil
	file_orders_order_proto_depIdxs = nil
}
