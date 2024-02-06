// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/signal.proto

package service

import (
	context "context"
	fmt "fmt"
	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("flyteidl/service/signal.proto", fileDescriptor_ca211d25a1023377) }

var fileDescriptor_ca211d25a1023377 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x5f, 0xca, 0xe3, 0xc1, 0x8b, 0x88, 0x3a, 0x0b, 0x17, 0x51, 0xc1, 0x66, 0xa5, 0x05,
	0x33, 0xa8, 0xbb, 0x82, 0x08, 0x2a, 0x08, 0x22, 0x08, 0x66, 0x21, 0x74, 0x53, 0xa6, 0xc9, 0xed,
	0x38, 0x36, 0x99, 0x1b, 0x33, 0xd3, 0x56, 0x91, 0x6e, 0x04, 0x57, 0x82, 0x1b, 0xb7, 0xfe, 0x2b,
	0x5d, 0xbb, 0xf2, 0x77, 0x88, 0x64, 0xd2, 0x06, 0xad, 0x89, 0x3b, 0x77, 0xe1, 0x9e, 0x2f, 0xe7,
	0xce, 0x39, 0x5c, 0x7b, 0xa5, 0x1b, 0x5d, 0x6b, 0x10, 0x61, 0x44, 0x15, 0xa4, 0x03, 0x11, 0x00,
	0x55, 0x82, 0x4b, 0x16, 0x79, 0x49, 0x8a, 0x1a, 0xc9, 0xfc, 0x44, 0xf6, 0xc6, 0xb2, 0xb3, 0xcc,
	0x11, 0x79, 0x04, 0x94, 0x25, 0x82, 0x32, 0x29, 0x51, 0x33, 0x2d, 0x50, 0xaa, 0x9c, 0x77, 0x96,
	0x0a, 0x3b, 0x16, 0xc6, 0x42, 0x7e, 0x31, 0xdb, 0x7a, 0xfd, 0x6b, 0xcf, 0xfa, 0x66, 0xe0, 0xe7,
	0x66, 0xe4, 0xcc, 0x5e, 0x38, 0x04, 0x7d, 0x92, 0xee, 0xa7, 0xc0, 0x34, 0xe4, 0x1a, 0x59, 0xf3,
	0x8a, 0xa5, 0xc6, 0xc4, 0xcb, 0xe7, 0x9f, 0xc0, 0x53, 0xb8, 0xec, 0x83, 0xd2, 0xce, 0x62, 0x39,
	0xe9, 0xfe, 0x21, 0x2f, 0x35, 0x7b, 0xe6, 0x58, 0x28, 0x9d, 0x0f, 0x14, 0xa9, 0x97, 0x93, 0x19,
	0x32, 0x31, 0x73, 0xaa, 0x11, 0xf7, 0xae, 0x76, 0xfb, 0xfc, 0xf6, 0x58, 0x7b, 0xb7, 0x48, 0xcf,
	0x04, 0x1f, 0x6c, 0x8e, 0x93, 0x29, 0x7a, 0x33, 0xc4, 0xb4, 0xd7, 0x8d, 0x70, 0xd8, 0x86, 0x2b,
	0x08, 0xfa, 0x59, 0x19, 0x6d, 0x11, 0x66, 0x89, 0x2f, 0x20, 0xd0, 0xa3, 0x2a, 0x3d, 0xc4, 0x98,
	0x09, 0x59, 0x29, 0x4b, 0x16, 0xc3, 0xa8, 0xf5, 0x60, 0x91, 0x7b, 0x6b, 0x7a, 0x21, 0xa6, 0xbc,
	0xea, 0x2f, 0x4c, 0xf9, 0xe8, 0x37, 0x1f, 0x44, 0x9e, 0x2c, 0xfb, 0xbf, 0x0f, 0xe3, 0x56, 0xc9,
	0x6a, 0x79, 0x63, 0x3e, 0x14, 0x9d, 0xd6, 0x7f, 0x20, 0x54, 0x82, 0x52, 0x81, 0x7b, 0x64, 0x9a,
	0x3d, 0x70, 0xe7, 0xa6, 0x72, 0x36, 0xad, 0x46, 0xcb, 0x73, 0xd7, 0x4b, 0xd3, 0x8b, 0xd0, 0xfb,
	0x96, 0xbb, 0x69, 0x35, 0xf6, 0x76, 0x5b, 0x3b, 0x5c, 0xe8, 0xf3, 0x7e, 0xc7, 0x0b, 0x30, 0xa6,
	0x66, 0x75, 0xc6, 0x9b, 0x0f, 0x5a, 0x5c, 0x26, 0x07, 0x49, 0x93, 0xce, 0x06, 0x47, 0x3a, 0x7d,
	0xfb, 0x9d, 0x7f, 0xe6, 0x50, 0xb7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0xf5, 0x8e, 0x13,
	0x16, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignalServiceClient is the client API for SignalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignalServiceClient interface {
	// Fetches or creates a :ref:`ref_flyteidl.admin.Signal`.
	GetOrCreateSignal(ctx context.Context, in *admin.SignalGetOrCreateRequest, opts ...grpc.CallOption) (*admin.Signal, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Signal` definitions.
	ListSignals(ctx context.Context, in *admin.SignalListRequest, opts ...grpc.CallOption) (*admin.SignalList, error)
	// Sets the value on a :ref:`ref_flyteidl.admin.Signal` definition
	SetSignal(ctx context.Context, in *admin.SignalSetRequest, opts ...grpc.CallOption) (*admin.SignalSetResponse, error)
}

type signalServiceClient struct {
	cc *grpc.ClientConn
}

func NewSignalServiceClient(cc *grpc.ClientConn) SignalServiceClient {
	return &signalServiceClient{cc}
}

func (c *signalServiceClient) GetOrCreateSignal(ctx context.Context, in *admin.SignalGetOrCreateRequest, opts ...grpc.CallOption) (*admin.Signal, error) {
	out := new(admin.Signal)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/GetOrCreateSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) ListSignals(ctx context.Context, in *admin.SignalListRequest, opts ...grpc.CallOption) (*admin.SignalList, error) {
	out := new(admin.SignalList)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/ListSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) SetSignal(ctx context.Context, in *admin.SignalSetRequest, opts ...grpc.CallOption) (*admin.SignalSetResponse, error) {
	out := new(admin.SignalSetResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/SetSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalServiceServer is the server API for SignalService service.
type SignalServiceServer interface {
	// Fetches or creates a :ref:`ref_flyteidl.admin.Signal`.
	GetOrCreateSignal(context.Context, *admin.SignalGetOrCreateRequest) (*admin.Signal, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Signal` definitions.
	ListSignals(context.Context, *admin.SignalListRequest) (*admin.SignalList, error)
	// Sets the value on a :ref:`ref_flyteidl.admin.Signal` definition
	SetSignal(context.Context, *admin.SignalSetRequest) (*admin.SignalSetResponse, error)
}

// UnimplementedSignalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSignalServiceServer struct {
}

func (*UnimplementedSignalServiceServer) GetOrCreateSignal(ctx context.Context, req *admin.SignalGetOrCreateRequest) (*admin.Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateSignal not implemented")
}
func (*UnimplementedSignalServiceServer) ListSignals(ctx context.Context, req *admin.SignalListRequest) (*admin.SignalList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSignals not implemented")
}
func (*UnimplementedSignalServiceServer) SetSignal(ctx context.Context, req *admin.SignalSetRequest) (*admin.SignalSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSignal not implemented")
}

func RegisterSignalServiceServer(s *grpc.Server, srv SignalServiceServer) {
	s.RegisterService(&_SignalService_serviceDesc, srv)
}

func _SignalService_GetOrCreateSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalGetOrCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetOrCreateSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/GetOrCreateSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetOrCreateSignal(ctx, req.(*admin.SignalGetOrCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_ListSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).ListSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/ListSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).ListSignals(ctx, req.(*admin.SignalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_SetSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).SetSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/SetSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).SetSignal(ctx, req.(*admin.SignalSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.SignalService",
	HandlerType: (*SignalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrCreateSignal",
			Handler:    _SignalService_GetOrCreateSignal_Handler,
		},
		{
			MethodName: "ListSignals",
			Handler:    _SignalService_ListSignals_Handler,
		},
		{
			MethodName: "SetSignal",
			Handler:    _SignalService_SetSignal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/signal.proto",
}