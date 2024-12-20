// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: rides/rides.proto

package rides

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rides_StartRide_FullMethodName     = "/showcase.Rides/StartRide"
	Rides_EndRide_FullMethodName       = "/showcase.Rides/EndRide"
	Rides_GetRide_FullMethodName       = "/showcase.Rides/GetRide"
	Rides_SetInvoiceUlr_FullMethodName = "/showcase.Rides/SetInvoiceUlr"
)

// RidesClient is the client API for Rides service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RidesClient interface {
	StartRide(ctx context.Context, in *StartRideRequest, opts ...grpc.CallOption) (*RideReply, error)
	EndRide(ctx context.Context, in *EndRideRequest, opts ...grpc.CallOption) (*RideReply, error)
	GetRide(ctx context.Context, in *GetRideRequest, opts ...grpc.CallOption) (*RideReply, error)
	SetInvoiceUlr(ctx context.Context, in *SetInvoiceUrlRequest, opts ...grpc.CallOption) (*RideReply, error)
}

type ridesClient struct {
	cc grpc.ClientConnInterface
}

func NewRidesClient(cc grpc.ClientConnInterface) RidesClient {
	return &ridesClient{cc}
}

func (c *ridesClient) StartRide(ctx context.Context, in *StartRideRequest, opts ...grpc.CallOption) (*RideReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideReply)
	err := c.cc.Invoke(ctx, Rides_StartRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ridesClient) EndRide(ctx context.Context, in *EndRideRequest, opts ...grpc.CallOption) (*RideReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideReply)
	err := c.cc.Invoke(ctx, Rides_EndRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ridesClient) GetRide(ctx context.Context, in *GetRideRequest, opts ...grpc.CallOption) (*RideReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideReply)
	err := c.cc.Invoke(ctx, Rides_GetRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ridesClient) SetInvoiceUlr(ctx context.Context, in *SetInvoiceUrlRequest, opts ...grpc.CallOption) (*RideReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RideReply)
	err := c.cc.Invoke(ctx, Rides_SetInvoiceUlr_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RidesServer is the server API for Rides service.
// All implementations must embed UnimplementedRidesServer
// for forward compatibility.
type RidesServer interface {
	StartRide(context.Context, *StartRideRequest) (*RideReply, error)
	EndRide(context.Context, *EndRideRequest) (*RideReply, error)
	GetRide(context.Context, *GetRideRequest) (*RideReply, error)
	SetInvoiceUlr(context.Context, *SetInvoiceUrlRequest) (*RideReply, error)
	mustEmbedUnimplementedRidesServer()
}

// UnimplementedRidesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRidesServer struct{}

func (UnimplementedRidesServer) StartRide(context.Context, *StartRideRequest) (*RideReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRide not implemented")
}
func (UnimplementedRidesServer) EndRide(context.Context, *EndRideRequest) (*RideReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndRide not implemented")
}
func (UnimplementedRidesServer) GetRide(context.Context, *GetRideRequest) (*RideReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRide not implemented")
}
func (UnimplementedRidesServer) SetInvoiceUlr(context.Context, *SetInvoiceUrlRequest) (*RideReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInvoiceUlr not implemented")
}
func (UnimplementedRidesServer) mustEmbedUnimplementedRidesServer() {}
func (UnimplementedRidesServer) testEmbeddedByValue()               {}

// UnsafeRidesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RidesServer will
// result in compilation errors.
type UnsafeRidesServer interface {
	mustEmbedUnimplementedRidesServer()
}

func RegisterRidesServer(s grpc.ServiceRegistrar, srv RidesServer) {
	// If the following call pancis, it indicates UnimplementedRidesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rides_ServiceDesc, srv)
}

func _Rides_StartRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RidesServer).StartRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rides_StartRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RidesServer).StartRide(ctx, req.(*StartRideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rides_EndRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RidesServer).EndRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rides_EndRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RidesServer).EndRide(ctx, req.(*EndRideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rides_GetRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RidesServer).GetRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rides_GetRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RidesServer).GetRide(ctx, req.(*GetRideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rides_SetInvoiceUlr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInvoiceUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RidesServer).SetInvoiceUlr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rides_SetInvoiceUlr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RidesServer).SetInvoiceUlr(ctx, req.(*SetInvoiceUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rides_ServiceDesc is the grpc.ServiceDesc for Rides service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rides_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "showcase.Rides",
	HandlerType: (*RidesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartRide",
			Handler:    _Rides_StartRide_Handler,
		},
		{
			MethodName: "EndRide",
			Handler:    _Rides_EndRide_Handler,
		},
		{
			MethodName: "GetRide",
			Handler:    _Rides_GetRide_Handler,
		},
		{
			MethodName: "SetInvoiceUlr",
			Handler:    _Rides_SetInvoiceUlr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rides/rides.proto",
}
