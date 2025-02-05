// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: chromadb/proto/coordinator.proto

package coordinatorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SysDB_CreateDatabase_FullMethodName                 = "/chroma.SysDB/CreateDatabase"
	SysDB_GetDatabase_FullMethodName                    = "/chroma.SysDB/GetDatabase"
	SysDB_CreateTenant_FullMethodName                   = "/chroma.SysDB/CreateTenant"
	SysDB_GetTenant_FullMethodName                      = "/chroma.SysDB/GetTenant"
	SysDB_CreateSegment_FullMethodName                  = "/chroma.SysDB/CreateSegment"
	SysDB_DeleteSegment_FullMethodName                  = "/chroma.SysDB/DeleteSegment"
	SysDB_GetSegments_FullMethodName                    = "/chroma.SysDB/GetSegments"
	SysDB_UpdateSegment_FullMethodName                  = "/chroma.SysDB/UpdateSegment"
	SysDB_CreateCollection_FullMethodName               = "/chroma.SysDB/CreateCollection"
	SysDB_DeleteCollection_FullMethodName               = "/chroma.SysDB/DeleteCollection"
	SysDB_GetCollections_FullMethodName                 = "/chroma.SysDB/GetCollections"
	SysDB_CheckCollections_FullMethodName               = "/chroma.SysDB/CheckCollections"
	SysDB_UpdateCollection_FullMethodName               = "/chroma.SysDB/UpdateCollection"
	SysDB_ResetState_FullMethodName                     = "/chroma.SysDB/ResetState"
	SysDB_GetLastCompactionTimeForTenant_FullMethodName = "/chroma.SysDB/GetLastCompactionTimeForTenant"
	SysDB_SetLastCompactionTimeForTenant_FullMethodName = "/chroma.SysDB/SetLastCompactionTimeForTenant"
	SysDB_FlushCollectionCompaction_FullMethodName      = "/chroma.SysDB/FlushCollectionCompaction"
)

// SysDBClient is the client API for SysDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysDBClient interface {
	CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error)
	GetDatabase(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*GetDatabaseResponse, error)
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error)
	CreateSegment(ctx context.Context, in *CreateSegmentRequest, opts ...grpc.CallOption) (*CreateSegmentResponse, error)
	DeleteSegment(ctx context.Context, in *DeleteSegmentRequest, opts ...grpc.CallOption) (*DeleteSegmentResponse, error)
	GetSegments(ctx context.Context, in *GetSegmentsRequest, opts ...grpc.CallOption) (*GetSegmentsResponse, error)
	UpdateSegment(ctx context.Context, in *UpdateSegmentRequest, opts ...grpc.CallOption) (*UpdateSegmentResponse, error)
	CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error)
	GetCollections(ctx context.Context, in *GetCollectionsRequest, opts ...grpc.CallOption) (*GetCollectionsResponse, error)
	CheckCollections(ctx context.Context, in *CheckCollectionsRequest, opts ...grpc.CallOption) (*CheckCollectionsResponse, error)
	UpdateCollection(ctx context.Context, in *UpdateCollectionRequest, opts ...grpc.CallOption) (*UpdateCollectionResponse, error)
	ResetState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResetStateResponse, error)
	GetLastCompactionTimeForTenant(ctx context.Context, in *GetLastCompactionTimeForTenantRequest, opts ...grpc.CallOption) (*GetLastCompactionTimeForTenantResponse, error)
	SetLastCompactionTimeForTenant(ctx context.Context, in *SetLastCompactionTimeForTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FlushCollectionCompaction(ctx context.Context, in *FlushCollectionCompactionRequest, opts ...grpc.CallOption) (*FlushCollectionCompactionResponse, error)
}

type sysDBClient struct {
	cc grpc.ClientConnInterface
}

func NewSysDBClient(cc grpc.ClientConnInterface) SysDBClient {
	return &sysDBClient{cc}
}

func (c *sysDBClient) CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDatabaseResponse)
	err := c.cc.Invoke(ctx, SysDB_CreateDatabase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) GetDatabase(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*GetDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDatabaseResponse)
	err := c.cc.Invoke(ctx, SysDB_GetDatabase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTenantResponse)
	err := c.cc.Invoke(ctx, SysDB_CreateTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, SysDB_GetTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) CreateSegment(ctx context.Context, in *CreateSegmentRequest, opts ...grpc.CallOption) (*CreateSegmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSegmentResponse)
	err := c.cc.Invoke(ctx, SysDB_CreateSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) DeleteSegment(ctx context.Context, in *DeleteSegmentRequest, opts ...grpc.CallOption) (*DeleteSegmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSegmentResponse)
	err := c.cc.Invoke(ctx, SysDB_DeleteSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) GetSegments(ctx context.Context, in *GetSegmentsRequest, opts ...grpc.CallOption) (*GetSegmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSegmentsResponse)
	err := c.cc.Invoke(ctx, SysDB_GetSegments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) UpdateSegment(ctx context.Context, in *UpdateSegmentRequest, opts ...grpc.CallOption) (*UpdateSegmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSegmentResponse)
	err := c.cc.Invoke(ctx, SysDB_UpdateSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, SysDB_CreateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCollectionResponse)
	err := c.cc.Invoke(ctx, SysDB_DeleteCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) GetCollections(ctx context.Context, in *GetCollectionsRequest, opts ...grpc.CallOption) (*GetCollectionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCollectionsResponse)
	err := c.cc.Invoke(ctx, SysDB_GetCollections_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) CheckCollections(ctx context.Context, in *CheckCollectionsRequest, opts ...grpc.CallOption) (*CheckCollectionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckCollectionsResponse)
	err := c.cc.Invoke(ctx, SysDB_CheckCollections_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) UpdateCollection(ctx context.Context, in *UpdateCollectionRequest, opts ...grpc.CallOption) (*UpdateCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCollectionResponse)
	err := c.cc.Invoke(ctx, SysDB_UpdateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) ResetState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResetStateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetStateResponse)
	err := c.cc.Invoke(ctx, SysDB_ResetState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) GetLastCompactionTimeForTenant(ctx context.Context, in *GetLastCompactionTimeForTenantRequest, opts ...grpc.CallOption) (*GetLastCompactionTimeForTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLastCompactionTimeForTenantResponse)
	err := c.cc.Invoke(ctx, SysDB_GetLastCompactionTimeForTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) SetLastCompactionTimeForTenant(ctx context.Context, in *SetLastCompactionTimeForTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SysDB_SetLastCompactionTimeForTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDBClient) FlushCollectionCompaction(ctx context.Context, in *FlushCollectionCompactionRequest, opts ...grpc.CallOption) (*FlushCollectionCompactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlushCollectionCompactionResponse)
	err := c.cc.Invoke(ctx, SysDB_FlushCollectionCompaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysDBServer is the server API for SysDB service.
// All implementations must embed UnimplementedSysDBServer
// for forward compatibility.
type SysDBServer interface {
	CreateDatabase(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error)
	GetDatabase(context.Context, *GetDatabaseRequest) (*GetDatabaseResponse, error)
	CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error)
	GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error)
	CreateSegment(context.Context, *CreateSegmentRequest) (*CreateSegmentResponse, error)
	DeleteSegment(context.Context, *DeleteSegmentRequest) (*DeleteSegmentResponse, error)
	GetSegments(context.Context, *GetSegmentsRequest) (*GetSegmentsResponse, error)
	UpdateSegment(context.Context, *UpdateSegmentRequest) (*UpdateSegmentResponse, error)
	CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error)
	GetCollections(context.Context, *GetCollectionsRequest) (*GetCollectionsResponse, error)
	CheckCollections(context.Context, *CheckCollectionsRequest) (*CheckCollectionsResponse, error)
	UpdateCollection(context.Context, *UpdateCollectionRequest) (*UpdateCollectionResponse, error)
	ResetState(context.Context, *emptypb.Empty) (*ResetStateResponse, error)
	GetLastCompactionTimeForTenant(context.Context, *GetLastCompactionTimeForTenantRequest) (*GetLastCompactionTimeForTenantResponse, error)
	SetLastCompactionTimeForTenant(context.Context, *SetLastCompactionTimeForTenantRequest) (*emptypb.Empty, error)
	FlushCollectionCompaction(context.Context, *FlushCollectionCompactionRequest) (*FlushCollectionCompactionResponse, error)
	mustEmbedUnimplementedSysDBServer()
}

// UnimplementedSysDBServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSysDBServer struct{}

func (UnimplementedSysDBServer) CreateDatabase(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabase not implemented")
}
func (UnimplementedSysDBServer) GetDatabase(context.Context, *GetDatabaseRequest) (*GetDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatabase not implemented")
}
func (UnimplementedSysDBServer) CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedSysDBServer) GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedSysDBServer) CreateSegment(context.Context, *CreateSegmentRequest) (*CreateSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSegment not implemented")
}
func (UnimplementedSysDBServer) DeleteSegment(context.Context, *DeleteSegmentRequest) (*DeleteSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSegment not implemented")
}
func (UnimplementedSysDBServer) GetSegments(context.Context, *GetSegmentsRequest) (*GetSegmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSegments not implemented")
}
func (UnimplementedSysDBServer) UpdateSegment(context.Context, *UpdateSegmentRequest) (*UpdateSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSegment not implemented")
}
func (UnimplementedSysDBServer) CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedSysDBServer) DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollection not implemented")
}
func (UnimplementedSysDBServer) GetCollections(context.Context, *GetCollectionsRequest) (*GetCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollections not implemented")
}
func (UnimplementedSysDBServer) CheckCollections(context.Context, *CheckCollectionsRequest) (*CheckCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCollections not implemented")
}
func (UnimplementedSysDBServer) UpdateCollection(context.Context, *UpdateCollectionRequest) (*UpdateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollection not implemented")
}
func (UnimplementedSysDBServer) ResetState(context.Context, *emptypb.Empty) (*ResetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetState not implemented")
}
func (UnimplementedSysDBServer) GetLastCompactionTimeForTenant(context.Context, *GetLastCompactionTimeForTenantRequest) (*GetLastCompactionTimeForTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastCompactionTimeForTenant not implemented")
}
func (UnimplementedSysDBServer) SetLastCompactionTimeForTenant(context.Context, *SetLastCompactionTimeForTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLastCompactionTimeForTenant not implemented")
}
func (UnimplementedSysDBServer) FlushCollectionCompaction(context.Context, *FlushCollectionCompactionRequest) (*FlushCollectionCompactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushCollectionCompaction not implemented")
}
func (UnimplementedSysDBServer) mustEmbedUnimplementedSysDBServer() {}
func (UnimplementedSysDBServer) testEmbeddedByValue()               {}

// UnsafeSysDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysDBServer will
// result in compilation errors.
type UnsafeSysDBServer interface {
	mustEmbedUnimplementedSysDBServer()
}

func RegisterSysDBServer(s grpc.ServiceRegistrar, srv SysDBServer) {
	// If the following call pancis, it indicates UnimplementedSysDBServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SysDB_ServiceDesc, srv)
}

func _SysDB_CreateDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).CreateDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_CreateDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).CreateDatabase(ctx, req.(*CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_GetDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).GetDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_GetDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).GetDatabase(ctx, req.(*GetDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_CreateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_GetTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_CreateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).CreateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_CreateSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).CreateSegment(ctx, req.(*CreateSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_DeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).DeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_DeleteSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).DeleteSegment(ctx, req.(*DeleteSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_GetSegments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).GetSegments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_GetSegments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).GetSegments(ctx, req.(*GetSegmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_UpdateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).UpdateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_UpdateSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).UpdateSegment(ctx, req.(*UpdateSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_CreateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).CreateCollection(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_DeleteCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).DeleteCollection(ctx, req.(*DeleteCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_GetCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).GetCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_GetCollections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).GetCollections(ctx, req.(*GetCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_CheckCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).CheckCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_CheckCollections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).CheckCollections(ctx, req.(*CheckCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_UpdateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).UpdateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_UpdateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).UpdateCollection(ctx, req.(*UpdateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_ResetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).ResetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_ResetState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).ResetState(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_GetLastCompactionTimeForTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastCompactionTimeForTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).GetLastCompactionTimeForTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_GetLastCompactionTimeForTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).GetLastCompactionTimeForTenant(ctx, req.(*GetLastCompactionTimeForTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_SetLastCompactionTimeForTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLastCompactionTimeForTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).SetLastCompactionTimeForTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_SetLastCompactionTimeForTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).SetLastCompactionTimeForTenant(ctx, req.(*SetLastCompactionTimeForTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDB_FlushCollectionCompaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushCollectionCompactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDBServer).FlushCollectionCompaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDB_FlushCollectionCompaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDBServer).FlushCollectionCompaction(ctx, req.(*FlushCollectionCompactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SysDB_ServiceDesc is the grpc.ServiceDesc for SysDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chroma.SysDB",
	HandlerType: (*SysDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDatabase",
			Handler:    _SysDB_CreateDatabase_Handler,
		},
		{
			MethodName: "GetDatabase",
			Handler:    _SysDB_GetDatabase_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _SysDB_CreateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _SysDB_GetTenant_Handler,
		},
		{
			MethodName: "CreateSegment",
			Handler:    _SysDB_CreateSegment_Handler,
		},
		{
			MethodName: "DeleteSegment",
			Handler:    _SysDB_DeleteSegment_Handler,
		},
		{
			MethodName: "GetSegments",
			Handler:    _SysDB_GetSegments_Handler,
		},
		{
			MethodName: "UpdateSegment",
			Handler:    _SysDB_UpdateSegment_Handler,
		},
		{
			MethodName: "CreateCollection",
			Handler:    _SysDB_CreateCollection_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _SysDB_DeleteCollection_Handler,
		},
		{
			MethodName: "GetCollections",
			Handler:    _SysDB_GetCollections_Handler,
		},
		{
			MethodName: "CheckCollections",
			Handler:    _SysDB_CheckCollections_Handler,
		},
		{
			MethodName: "UpdateCollection",
			Handler:    _SysDB_UpdateCollection_Handler,
		},
		{
			MethodName: "ResetState",
			Handler:    _SysDB_ResetState_Handler,
		},
		{
			MethodName: "GetLastCompactionTimeForTenant",
			Handler:    _SysDB_GetLastCompactionTimeForTenant_Handler,
		},
		{
			MethodName: "SetLastCompactionTimeForTenant",
			Handler:    _SysDB_SetLastCompactionTimeForTenant_Handler,
		},
		{
			MethodName: "FlushCollectionCompaction",
			Handler:    _SysDB_FlushCollectionCompaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chromadb/proto/coordinator.proto",
}
