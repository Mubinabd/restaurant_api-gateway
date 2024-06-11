// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: restaurant_protos/payment.proto

package genproto

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

type PaymentCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ReservationId string `protobuf:"bytes,2,opt,name=ReservationId,proto3" json:"ReservationId,omitempty"`
	PaymentMethod string `protobuf:"bytes,3,opt,name=PaymentMethod,proto3" json:"PaymentMethod,omitempty"`
	PaymentStatus string `protobuf:"bytes,4,opt,name=PaymentStatus,proto3" json:"PaymentStatus,omitempty"`
}

func (x *PaymentCreate) Reset() {
	*x = PaymentCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCreate) ProtoMessage() {}

func (x *PaymentCreate) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCreate.ProtoReflect.Descriptor instead.
func (*PaymentCreate) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentCreate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaymentCreate) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

func (x *PaymentCreate) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *PaymentCreate) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ReservationId string  `protobuf:"bytes,2,opt,name=ReservationId,proto3" json:"ReservationId,omitempty"`
	Amount        float32 `protobuf:"fixed32,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	PaymentMethod string  `protobuf:"bytes,4,opt,name=PaymentMethod,proto3" json:"PaymentMethod,omitempty"`
	PaymentStatus string  `protobuf:"bytes,5,opt,name=PaymentStatus,proto3" json:"PaymentStatus,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payment) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

func (x *Payment) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *Payment) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

type PaymentFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethod string  `protobuf:"bytes,1,opt,name=PaymentMethod,proto3" json:"PaymentMethod,omitempty"`
	AmountFrom    float32 `protobuf:"fixed32,2,opt,name=AmountFrom,proto3" json:"AmountFrom,omitempty"`
	AmountTo      float32 `protobuf:"fixed32,3,opt,name=AmountTo,proto3" json:"AmountTo,omitempty"`
	PaymentStatus string  `protobuf:"bytes,4,opt,name=PaymentStatus,proto3" json:"PaymentStatus,omitempty"`
}

func (x *PaymentFilter) Reset() {
	*x = PaymentFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentFilter) ProtoMessage() {}

func (x *PaymentFilter) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentFilter.ProtoReflect.Descriptor instead.
func (*PaymentFilter) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentFilter) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *PaymentFilter) GetAmountFrom() float32 {
	if x != nil {
		return x.AmountFrom
	}
	return 0
}

func (x *PaymentFilter) GetAmountTo() float32 {
	if x != nil {
		return x.AmountTo
	}
	return 0
}

func (x *PaymentFilter) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

type Payments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payments []*Payment `protobuf:"bytes,1,rep,name=Payments,proto3" json:"Payments,omitempty"`
}

func (x *Payments) Reset() {
	*x = Payments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_protos_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payments) ProtoMessage() {}

func (x *Payments) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_protos_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payments.ProtoReflect.Descriptor instead.
func (*Payments) Descriptor() ([]byte, []int) {
	return file_restaurant_protos_payment_proto_rawDescGZIP(), []int{3}
}

func (x *Payments) GetPayments() []*Payment {
	if x != nil {
		return x.Payments
	}
	return nil
}

var File_restaurant_protos_payment_proto protoreflect.FileDescriptor

var file_restaurant_protos_payment_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x97,
	0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x42, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x8c, 0x03, 0x0a,
	0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restaurant_protos_payment_proto_rawDescOnce sync.Once
	file_restaurant_protos_payment_proto_rawDescData = file_restaurant_protos_payment_proto_rawDesc
)

func file_restaurant_protos_payment_proto_rawDescGZIP() []byte {
	file_restaurant_protos_payment_proto_rawDescOnce.Do(func() {
		file_restaurant_protos_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_protos_payment_proto_rawDescData)
	})
	return file_restaurant_protos_payment_proto_rawDescData
}

var file_restaurant_protos_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_restaurant_protos_payment_proto_goTypes = []interface{}{
	(*PaymentCreate)(nil), // 0: restaurant_protos.PaymentCreate
	(*Payment)(nil),       // 1: restaurant_protos.Payment
	(*PaymentFilter)(nil), // 2: restaurant_protos.PaymentFilter
	(*Payments)(nil),      // 3: restaurant_protos.Payments
	(*ById)(nil),          // 4: restaurant_protos.ById
	(*Void)(nil),          // 5: restaurant_protos.Void
}
var file_restaurant_protos_payment_proto_depIdxs = []int32{
	1, // 0: restaurant_protos.Payments.Payments:type_name -> restaurant_protos.Payment
	0, // 1: restaurant_protos.PaymentService.CreatePayment:input_type -> restaurant_protos.PaymentCreate
	4, // 2: restaurant_protos.PaymentService.GetPayment:input_type -> restaurant_protos.ById
	0, // 3: restaurant_protos.PaymentService.UpdatePayment:input_type -> restaurant_protos.PaymentCreate
	4, // 4: restaurant_protos.PaymentService.DeletePayment:input_type -> restaurant_protos.ById
	2, // 5: restaurant_protos.PaymentService.GetPayments:input_type -> restaurant_protos.PaymentFilter
	1, // 6: restaurant_protos.PaymentService.CreatePayment:output_type -> restaurant_protos.Payment
	1, // 7: restaurant_protos.PaymentService.GetPayment:output_type -> restaurant_protos.Payment
	1, // 8: restaurant_protos.PaymentService.UpdatePayment:output_type -> restaurant_protos.Payment
	5, // 9: restaurant_protos.PaymentService.DeletePayment:output_type -> restaurant_protos.Void
	3, // 10: restaurant_protos.PaymentService.GetPayments:output_type -> restaurant_protos.Payments
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_restaurant_protos_payment_proto_init() }
func file_restaurant_protos_payment_proto_init() {
	if File_restaurant_protos_payment_proto != nil {
		return
	}
	file_restaurant_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_restaurant_protos_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCreate); i {
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
		file_restaurant_protos_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_restaurant_protos_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentFilter); i {
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
		file_restaurant_protos_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payments); i {
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
			RawDescriptor: file_restaurant_protos_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restaurant_protos_payment_proto_goTypes,
		DependencyIndexes: file_restaurant_protos_payment_proto_depIdxs,
		MessageInfos:      file_restaurant_protos_payment_proto_msgTypes,
	}.Build()
	File_restaurant_protos_payment_proto = out.File
	file_restaurant_protos_payment_proto_rawDesc = nil
	file_restaurant_protos_payment_proto_goTypes = nil
	file_restaurant_protos_payment_proto_depIdxs = nil
}
