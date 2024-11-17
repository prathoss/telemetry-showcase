// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: bikes/bikes.proto

package bikes

import (
	context "context"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Bikes_GetBikeById_FullMethodName      = "/showcase.Bikes/GetBikeById"
	Bikes_ListBikes_FullMethodName        = "/showcase.Bikes/ListBikes"
	Bikes_SetBikeReserved_FullMethodName  = "/showcase.Bikes/SetBikeReserved"
	Bikes_SetBikeAvailable_FullMethodName = "/showcase.Bikes/SetBikeAvailable"
)

// BikesClient is the client API for Bikes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BikesClient interface {
	GetBikeById(ctx context.Context, in *GetBikeByIdRequest, opts ...grpc.CallOption) (*BikeReply, error)
	ListBikes(ctx context.Context, in *latlng.LatLng, opts ...grpc.CallOption) (*ListBikesReply, error)
	SetBikeReserved(ctx context.Context, in *SetBikeReservedRequest, opts ...grpc.CallOption) (*SetBikeReservedReply, error)
	SetBikeAvailable(ctx context.Context, in *SetBikeAvailableRequest, opts ...grpc.CallOption) (*SetBikeAvailableReply, error)
}

type bikesClient struct {
	cc grpc.ClientConnInterface
}

func NewBikesClient(cc grpc.ClientConnInterface) BikesClient {
	return &bikesClient{cc}
}

func (c *bikesClient) GetBikeById(ctx context.Context, in *GetBikeByIdRequest, opts ...grpc.CallOption) (*BikeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BikeReply)
	err := c.cc.Invoke(ctx, Bikes_GetBikeById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesClient) ListBikes(ctx context.Context, in *latlng.LatLng, opts ...grpc.CallOption) (*ListBikesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBikesReply)
	err := c.cc.Invoke(ctx, Bikes_ListBikes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesClient) SetBikeReserved(ctx context.Context, in *SetBikeReservedRequest, opts ...grpc.CallOption) (*SetBikeReservedReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetBikeReservedReply)
	err := c.cc.Invoke(ctx, Bikes_SetBikeReserved_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesClient) SetBikeAvailable(ctx context.Context, in *SetBikeAvailableRequest, opts ...grpc.CallOption) (*SetBikeAvailableReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetBikeAvailableReply)
	err := c.cc.Invoke(ctx, Bikes_SetBikeAvailable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BikesServer is the server API for Bikes service.
// All implementations must embed UnimplementedBikesServer
// for forward compatibility.
type BikesServer interface {
	GetBikeById(context.Context, *GetBikeByIdRequest) (*BikeReply, error)
	ListBikes(context.Context, *latlng.LatLng) (*ListBikesReply, error)
	SetBikeReserved(context.Context, *SetBikeReservedRequest) (*SetBikeReservedReply, error)
	SetBikeAvailable(context.Context, *SetBikeAvailableRequest) (*SetBikeAvailableReply, error)
	mustEmbedUnimplementedBikesServer()
}

// UnimplementedBikesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBikesServer struct{}

func (UnimplementedBikesServer) GetBikeById(context.Context, *GetBikeByIdRequest) (*BikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBikeById not implemented")
}
func (UnimplementedBikesServer) ListBikes(context.Context, *latlng.LatLng) (*ListBikesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBikes not implemented")
}
func (UnimplementedBikesServer) SetBikeReserved(context.Context, *SetBikeReservedRequest) (*SetBikeReservedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBikeReserved not implemented")
}
func (UnimplementedBikesServer) SetBikeAvailable(context.Context, *SetBikeAvailableRequest) (*SetBikeAvailableReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBikeAvailable not implemented")
}
func (UnimplementedBikesServer) mustEmbedUnimplementedBikesServer() {}
func (UnimplementedBikesServer) testEmbeddedByValue()               {}

// UnsafeBikesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BikesServer will
// result in compilation errors.
type UnsafeBikesServer interface {
	mustEmbedUnimplementedBikesServer()
}

func RegisterBikesServer(s grpc.ServiceRegistrar, srv BikesServer) {
	// If the following call pancis, it indicates UnimplementedBikesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Bikes_ServiceDesc, srv)
}

func _Bikes_GetBikeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBikeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesServer).GetBikeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bikes_GetBikeById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesServer).GetBikeById(ctx, req.(*GetBikeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bikes_ListBikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(latlng.LatLng)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesServer).ListBikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bikes_ListBikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesServer).ListBikes(ctx, req.(*latlng.LatLng))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bikes_SetBikeReserved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBikeReservedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesServer).SetBikeReserved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bikes_SetBikeReserved_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesServer).SetBikeReserved(ctx, req.(*SetBikeReservedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bikes_SetBikeAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBikeAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesServer).SetBikeAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bikes_SetBikeAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesServer).SetBikeAvailable(ctx, req.(*SetBikeAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bikes_ServiceDesc is the grpc.ServiceDesc for Bikes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bikes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "showcase.Bikes",
	HandlerType: (*BikesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBikeById",
			Handler:    _Bikes_GetBikeById_Handler,
		},
		{
			MethodName: "ListBikes",
			Handler:    _Bikes_ListBikes_Handler,
		},
		{
			MethodName: "SetBikeReserved",
			Handler:    _Bikes_SetBikeReserved_Handler,
		},
		{
			MethodName: "SetBikeAvailable",
			Handler:    _Bikes_SetBikeAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bikes/bikes.proto",
}
