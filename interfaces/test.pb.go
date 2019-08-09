// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	EmailId              string   `protobuf:"bytes,2,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	PageNumber           int32    `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,4,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetEmailId() string {
	if m != nil {
		return m.EmailId
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

//* SearchResponse represents the response of the Search
type SearchResponse struct {
	SearchResponse       string   `protobuf:"bytes,1,opt,name=search_response,json=searchResponse,proto3" json:"search_response,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	TotalPages           int32    `protobuf:"varint,3,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetSearchResponse() string {
	if m != nil {
		return m.SearchResponse
	}
	return ""
}

func (m *SearchResponse) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchResponse) GetTotalPages() int32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0xf3, 0x50,
	0x10, 0xc5, 0xbf, 0xa4, 0x7f, 0x3e, 0x3b, 0xa5, 0x09, 0xdc, 0x8d, 0xb1, 0x20, 0x09, 0x21, 0x68,
	0xb4, 0x34, 0x42, 0xdc, 0x89, 0x1b, 0xdd, 0xb9, 0x91, 0x92, 0xee, 0x14, 0x89, 0xb7, 0xcd, 0x50,
	0x03, 0x69, 0x93, 0xce, 0xbd, 0x29, 0xa8, 0x6f, 0xd6, 0x95, 0xaf, 0xe3, 0x5b, 0x48, 0x6e, 0x0c,
	0x18, 0xdd, 0x0d, 0xbf, 0x33, 0x87, 0x39, 0x73, 0x00, 0x24, 0x0a, 0x19, 0x14, 0x94, 0xcb, 0x7c,
	0x7c, 0xb8, 0xe3, 0x59, 0x9a, 0x70, 0x89, 0x17, 0xcd, 0x50, 0x0b, 0xee, 0x5e, 0x83, 0xd1, 0x1c,
	0x39, 0x2d, 0x5f, 0x22, 0xdc, 0x96, 0x28, 0x24, 0xbb, 0x82, 0xde, 0xb6, 0x44, 0x7a, 0xb5, 0x34,
	0x47, 0xf3, 0x07, 0xb7, 0xde, 0xfe, 0xf3, 0xa3, 0x63, 0xd3, 0xb1, 0x1f, 0x86, 0x47, 0xfe, 0xe3,
	0xcd, 0xf4, 0x81, 0x4f, 0xdf, 0x9e, 0x26, 0x67, 0xce, 0x8f, 0xf9, 0xdc, 0x8b, 0x6a, 0x0b, 0xf3,
	0xe0, 0x00, 0xd7, 0x3c, 0xcd, 0xe2, 0x34, 0xb1, 0x74, 0x65, 0x1f, 0x54, 0xf6, 0x2e, 0xe9, 0xcf,
	0x5a, 0xf4, 0x5f, 0x49, 0x77, 0x09, 0xb3, 0x61, 0x58, 0xf0, 0x15, 0xc6, 0x9b, 0x72, 0xbd, 0x40,
	0xb2, 0x3a, 0x8e, 0xe6, 0xf7, 0x22, 0xa8, 0xd0, 0xbd, 0x22, 0xec, 0x04, 0x4c, 0x42, 0x51, 0x66,
	0x32, 0x2e, 0x90, 0xe2, 0x4a, 0xb0, 0xba, 0x6a, 0x69, 0x54, 0xe3, 0x19, 0xd2, 0x8c, 0xaf, 0xd0,
	0x7d, 0x07, 0xa3, 0xc9, 0x2e, 0x8a, 0x7c, 0x23, 0x90, 0x9d, 0x82, 0x29, 0x14, 0x89, 0xe9, 0x1b,
	0xd5, 0x6f, 0x44, 0x86, 0x68, 0x2f, 0xfe, 0xca, 0xa0, 0xff, 0xc9, 0x60, 0xc3, 0x50, 0xe6, 0x92,
	0x67, 0xea, 0xbc, 0x68, 0x42, 0x2a, 0x54, 0xdd, 0x16, 0xe1, 0x75, 0x53, 0xdc, 0x1c, 0x69, 0x97,
	0x2e, 0x91, 0x4d, 0xa0, 0x5f, 0x03, 0x66, 0x04, 0xad, 0x4a, 0xc7, 0x66, 0xd0, 0x8e, 0xe9, 0xfe,
	0x5b, 0xf4, 0x55, 0xfd, 0x97, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0x5b, 0x68, 0xce, 0xa5,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	// Method to Search a test service
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	// Method to Search a test service
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
