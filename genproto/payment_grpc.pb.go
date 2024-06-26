// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: payment.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PaymentService_CreatePayment_FullMethodName = "/restaurant_protos.PaymentService/CreatePayment"
	PaymentService_GetPayment_FullMethodName    = "/restaurant_protos.PaymentService/GetPayment"
	PaymentService_UpdatePayment_FullMethodName = "/restaurant_protos.PaymentService/UpdatePayment"
	PaymentService_DeletePayment_FullMethodName = "/restaurant_protos.PaymentService/DeletePayment"
	PaymentService_GetPayments_FullMethodName   = "/restaurant_protos.PaymentService/GetPayments"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	CreatePayment(ctx context.Context, in *PaymentCreate, opts ...grpc.CallOption) (*Payment, error)
	GetPayment(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Payment, error)
	UpdatePayment(ctx context.Context, in *PaymentCreate, opts ...grpc.CallOption) (*Payment, error)
	DeletePayment(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetPayments(ctx context.Context, in *PaymentFilter, opts ...grpc.CallOption) (*Payments, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *PaymentCreate, opts ...grpc.CallOption) (*Payment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Payment)
	err := c.cc.Invoke(ctx, PaymentService_CreatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPayment(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Payment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Payment)
	err := c.cc.Invoke(ctx, PaymentService_GetPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePayment(ctx context.Context, in *PaymentCreate, opts ...grpc.CallOption) (*Payment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Payment)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DeletePayment(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, PaymentService_DeletePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPayments(ctx context.Context, in *PaymentFilter, opts ...grpc.CallOption) (*Payments, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Payments)
	err := c.cc.Invoke(ctx, PaymentService_GetPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	CreatePayment(context.Context, *PaymentCreate) (*Payment, error)
	GetPayment(context.Context, *ById) (*Payment, error)
	UpdatePayment(context.Context, *PaymentCreate) (*Payment, error)
	DeletePayment(context.Context, *ById) (*Void, error)
	GetPayments(context.Context, *PaymentFilter) (*Payments, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) CreatePayment(context.Context, *PaymentCreate) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetPayment(context.Context, *ById) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePayment(context.Context, *PaymentCreate) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayment not implemented")
}
func (UnimplementedPaymentServiceServer) DeletePayment(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetPayments(context.Context, *PaymentFilter) (*Payments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayments not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*PaymentCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPayment(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePayment(ctx, req.(*PaymentCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DeletePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DeletePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_DeletePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DeletePayment(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPayments(ctx, req.(*PaymentFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant_protos.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentService_CreatePayment_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _PaymentService_GetPayment_Handler,
		},
		{
			MethodName: "UpdatePayment",
			Handler:    _PaymentService_UpdatePayment_Handler,
		},
		{
			MethodName: "DeletePayment",
			Handler:    _PaymentService_DeletePayment_Handler,
		},
		{
			MethodName: "GetPayments",
			Handler:    _PaymentService_GetPayments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
