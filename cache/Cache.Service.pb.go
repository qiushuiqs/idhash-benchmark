// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Cache.Service.proto

package cache

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("Cache.Service.proto", fileDescriptor_b47aea17f626d001) }

var fileDescriptor_b47aea17f626d001 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x76, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x0b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4d, 0x06, 0x09, 0x4a, 0x41, 0xe5, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x20, 0x72,
	0x52, 0x22, 0x30, 0xc1, 0xe2, 0x82, 0xfc, 0xbc, 0x62, 0xa8, 0x0e, 0x23, 0x67, 0x2e, 0x1e, 0xb0,
	0x38, 0xd4, 0x1c, 0x21, 0x63, 0x2e, 0xbe, 0x80, 0xa2, 0xfc, 0xe4, 0xd4, 0xe2, 0x62, 0xa8, 0x6e,
	0x21, 0x3e, 0xbd, 0x64, 0x64, 0xd3, 0xa4, 0xf8, 0xe1, 0x7c, 0x88, 0x41, 0x4a, 0x0c, 0x49, 0x6c,
	0x60, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x9d, 0x91, 0xeb, 0x94, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheServiceClient interface {
	ProcessRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type cacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewCacheServiceClient(cc *grpc.ClientConn) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) ProcessRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cache.CacheService/ProcessRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
type CacheServiceServer interface {
	ProcessRequest(context.Context, *Request) (*Response, error)
}

// UnimplementedCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (*UnimplementedCacheServiceServer) ProcessRequest(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessRequest not implemented")
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_ProcessRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).ProcessRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.CacheService/ProcessRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).ProcessRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cache.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessRequest",
			Handler:    _CacheService_ProcessRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Cache.Service.proto",
}
