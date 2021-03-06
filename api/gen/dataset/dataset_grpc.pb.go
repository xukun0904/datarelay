// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: dataset.proto

package dataset

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	response "jhr.com/datarelay/api/gen/response"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataSetServiceClient is the client API for DataSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSetServiceClient interface {
	QueryForList(ctx context.Context, in *ConnectionInfo, opts ...grpc.CallOption) (*response.ResponseResult, error)
}

type dataSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSetServiceClient(cc grpc.ClientConnInterface) DataSetServiceClient {
	return &dataSetServiceClient{cc}
}

func (c *dataSetServiceClient) QueryForList(ctx context.Context, in *ConnectionInfo, opts ...grpc.CallOption) (*response.ResponseResult, error) {
	out := new(response.ResponseResult)
	err := c.cc.Invoke(ctx, "/dataset.DataSetService/QueryForList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSetServiceServer is the server API for DataSetService service.
// All implementations should embed UnimplementedDataSetServiceServer
// for forward compatibility
type DataSetServiceServer interface {
	QueryForList(context.Context, *ConnectionInfo) (*response.ResponseResult, error)
}

// UnimplementedDataSetServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDataSetServiceServer struct {
}

func (UnimplementedDataSetServiceServer) QueryForList(context.Context, *ConnectionInfo) (*response.ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryForList not implemented")
}

// UnsafeDataSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataSetServiceServer will
// result in compilation errors.
type UnsafeDataSetServiceServer interface {
	mustEmbedUnimplementedDataSetServiceServer()
}

func RegisterDataSetServiceServer(s grpc.ServiceRegistrar, srv DataSetServiceServer) {
	s.RegisterService(&DataSetService_ServiceDesc, srv)
}

func _DataSetService_QueryForList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSetServiceServer).QueryForList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataset.DataSetService/QueryForList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSetServiceServer).QueryForList(ctx, req.(*ConnectionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// DataSetService_ServiceDesc is the grpc.ServiceDesc for DataSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dataset.DataSetService",
	HandlerType: (*DataSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryForList",
			Handler:    _DataSetService_QueryForList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataset.proto",
}
