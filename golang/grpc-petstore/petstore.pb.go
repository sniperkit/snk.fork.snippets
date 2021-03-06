// Code generated by protoc-gen-go. DO NOT EDIT.
// source: petstore.proto

package petstore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OrderStatus int32

const (
	OrderStatus_OrderStatusUnknown   OrderStatus = 0
	OrderStatus_OrderStatusPlaced    OrderStatus = 1
	OrderStatus_OrderStatusApproved  OrderStatus = 2
	OrderStatus_OrderStatusDelivered OrderStatus = 3
)

var OrderStatus_name = map[int32]string{
	0: "OrderStatusUnknown",
	1: "OrderStatusPlaced",
	2: "OrderStatusApproved",
	3: "OrderStatusDelivered",
}
var OrderStatus_value = map[string]int32{
	"OrderStatusUnknown":   0,
	"OrderStatusPlaced":    1,
	"OrderStatusApproved":  2,
	"OrderStatusDelivered": 3,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{0}
}

type PetStatus int32

const (
	PetStatus_PetStatusUnknown   PetStatus = 0
	PetStatus_PetStatusAvailable PetStatus = 1
	PetStatus_PetStatusPending   PetStatus = 2
	PetStatus_PetStatusSold      PetStatus = 3
)

var PetStatus_name = map[int32]string{
	0: "PetStatusUnknown",
	1: "PetStatusAvailable",
	2: "PetStatusPending",
	3: "PetStatusSold",
}
var PetStatus_value = map[string]int32{
	"PetStatusUnknown":   0,
	"PetStatusAvailable": 1,
	"PetStatusPending":   2,
	"PetStatusSold":      3,
}

func (x PetStatus) String() string {
	return proto.EnumName(PetStatus_name, int32(x))
}
func (PetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{1}
}

type Tag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{0}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Category struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{1}
}
func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (dst *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(dst, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type Order struct {
	Id                   uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PetId                int64       `protobuf:"varint,2,opt,name=petId,proto3" json:"petId,omitempty"`
	Quantity             int32       `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Status               OrderStatus `protobuf:"varint,4,opt,name=status,proto3,enum=petstore.OrderStatus" json:"status,omitempty"`
	Complete             bool        `protobuf:"varint,5,opt,name=complete,proto3" json:"complete,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{3}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetPetId() int64 {
	if m != nil {
		return m.PetId
	}
	return 0
}

func (m *Order) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Order) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_OrderStatusUnknown
}

func (m *Order) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

type Pet struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category             *Category `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PhotoUrls            []string  `protobuf:"bytes,4,rep,name=photoUrls,proto3" json:"photoUrls,omitempty"`
	Tags                 []*Tag    `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Status               PetStatus `protobuf:"varint,6,opt,name=status,proto3,enum=petstore.PetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}
func (*Pet) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{4}
}
func (m *Pet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pet.Unmarshal(m, b)
}
func (m *Pet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pet.Marshal(b, m, deterministic)
}
func (dst *Pet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pet.Merge(dst, src)
}
func (m *Pet) XXX_Size() int {
	return xxx_messageInfo_Pet.Size(m)
}
func (m *Pet) XXX_DiscardUnknown() {
	xxx_messageInfo_Pet.DiscardUnknown(m)
}

var xxx_messageInfo_Pet proto.InternalMessageInfo

func (m *Pet) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pet) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *Pet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pet) GetPhotoUrls() []string {
	if m != nil {
		return m.PhotoUrls
	}
	return nil
}

func (m *Pet) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Pet) GetStatus() PetStatus {
	if m != nil {
		return m.Status
	}
	return PetStatus_PetStatusUnknown
}

type ReadPetRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPetRequest) Reset()         { *m = ReadPetRequest{} }
func (m *ReadPetRequest) String() string { return proto.CompactTextString(m) }
func (*ReadPetRequest) ProtoMessage()    {}
func (*ReadPetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{5}
}
func (m *ReadPetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPetRequest.Unmarshal(m, b)
}
func (m *ReadPetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPetRequest.Marshal(b, m, deterministic)
}
func (dst *ReadPetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPetRequest.Merge(dst, src)
}
func (m *ReadPetRequest) XXX_Size() int {
	return xxx_messageInfo_ReadPetRequest.Size(m)
}
func (m *ReadPetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPetRequest proto.InternalMessageInfo

func (m *ReadPetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2d59206d057ccdb5, []int{6}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Tag)(nil), "petstore.Tag")
	proto.RegisterType((*Category)(nil), "petstore.Category")
	proto.RegisterType((*User)(nil), "petstore.User")
	proto.RegisterType((*Order)(nil), "petstore.Order")
	proto.RegisterType((*Pet)(nil), "petstore.Pet")
	proto.RegisterType((*ReadPetRequest)(nil), "petstore.ReadPetRequest")
	proto.RegisterType((*Empty)(nil), "petstore.Empty")
	proto.RegisterEnum("petstore.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("petstore.PetStatus", PetStatus_name, PetStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PetstoreServiceClient is the client API for PetstoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PetstoreServiceClient interface {
	CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error)
	ReadPet(ctx context.Context, in *ReadPetRequest, opts ...grpc.CallOption) (PetstoreService_ReadPetClient, error)
	UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error)
	DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error)
}

type petstoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewPetstoreServiceClient(cc *grpc.ClientConn) PetstoreServiceClient {
	return &petstoreServiceClient{cc}
}

func (c *petstoreServiceClient) CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/CreatePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) ReadPet(ctx context.Context, in *ReadPetRequest, opts ...grpc.CallOption) (PetstoreService_ReadPetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PetstoreService_serviceDesc.Streams[0], "/petstore.PetstoreService/ReadPet", opts...)
	if err != nil {
		return nil, err
	}
	x := &petstoreServiceReadPetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PetstoreService_ReadPetClient interface {
	Recv() (*Pet, error)
	grpc.ClientStream
}

type petstoreServiceReadPetClient struct {
	grpc.ClientStream
}

func (x *petstoreServiceReadPetClient) Recv() (*Pet, error) {
	m := new(Pet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *petstoreServiceClient) UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/UpdatePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/DeletePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetstoreServiceServer is the server API for PetstoreService service.
type PetstoreServiceServer interface {
	CreatePet(context.Context, *Pet) (*Empty, error)
	ReadPet(*ReadPetRequest, PetstoreService_ReadPetServer) error
	UpdatePet(context.Context, *Pet) (*Empty, error)
	DeletePet(context.Context, *Pet) (*Empty, error)
}

func RegisterPetstoreServiceServer(s *grpc.Server, srv PetstoreServiceServer) {
	s.RegisterService(&_PetstoreService_serviceDesc, srv)
}

func _PetstoreService_CreatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).CreatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/CreatePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).CreatePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_ReadPet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadPetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PetstoreServiceServer).ReadPet(m, &petstoreServiceReadPetServer{stream})
}

type PetstoreService_ReadPetServer interface {
	Send(*Pet) error
	grpc.ServerStream
}

type petstoreServiceReadPetServer struct {
	grpc.ServerStream
}

func (x *petstoreServiceReadPetServer) Send(m *Pet) error {
	return x.ServerStream.SendMsg(m)
}

func _PetstoreService_UpdatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).UpdatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/UpdatePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).UpdatePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/DeletePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).DeletePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

var _PetstoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "petstore.PetstoreService",
	HandlerType: (*PetstoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePet",
			Handler:    _PetstoreService_CreatePet_Handler,
		},
		{
			MethodName: "UpdatePet",
			Handler:    _PetstoreService_UpdatePet_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _PetstoreService_DeletePet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadPet",
			Handler:       _PetstoreService_ReadPet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "petstore.proto",
}

func init() { proto.RegisterFile("petstore.proto", fileDescriptor_petstore_2d59206d057ccdb5) }

var fileDescriptor_petstore_2d59206d057ccdb5 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x6c, 0x27, 0x69, 0x72, 0xab, 0x24, 0xee, 0xed, 0xcf, 0x67, 0x85, 0x2e, 0x52, 0xaf,
	0x4a, 0x10, 0x09, 0x84, 0x1d, 0x2b, 0x4a, 0xdb, 0x05, 0x1b, 0x88, 0xdc, 0x46, 0x48, 0xec, 0xa6,
	0xf1, 0xc5, 0xb5, 0x70, 0x3c, 0xee, 0x78, 0x12, 0x54, 0x21, 0x36, 0xbc, 0x02, 0x2c, 0x78, 0x13,
	0x1e, 0x81, 0x07, 0xe0, 0x15, 0x78, 0x10, 0x34, 0x93, 0x64, 0x6c, 0xab, 0x9b, 0xec, 0x7c, 0xee,
	0x39, 0x73, 0xe6, 0xcc, 0xb9, 0x51, 0xa0, 0x93, 0x91, 0xcc, 0x25, 0x17, 0x34, 0xcc, 0x04, 0x97,
	0x1c, 0x9b, 0x1b, 0xdc, 0x3b, 0x8e, 0x38, 0x8f, 0x12, 0x1a, 0xb1, 0x2c, 0x1e, 0xb1, 0x34, 0xe5,
	0x92, 0xc9, 0x98, 0xa7, 0xf9, 0x4a, 0xe7, 0x3f, 0x06, 0xe7, 0x9a, 0x45, 0xd8, 0x01, 0x3b, 0x0e,
	0x3d, 0xab, 0x6f, 0x9d, 0x3a, 0x81, 0x1d, 0x87, 0x88, 0x50, 0x4b, 0xd9, 0x9c, 0x3c, 0xbb, 0x6f,
	0x9d, 0xb6, 0x02, 0xfd, 0xed, 0x0f, 0xa1, 0x79, 0xce, 0x24, 0x45, 0x5c, 0xdc, 0x6f, 0xa5, 0xff,
	0x69, 0x41, 0x6d, 0x9a, 0x93, 0x78, 0x20, 0xee, 0x41, 0x73, 0x91, 0x93, 0x28, 0x1d, 0x30, 0x18,
	0x8f, 0xa1, 0xf5, 0x31, 0x16, 0xb9, 0x7c, 0xab, 0x48, 0x47, 0x93, 0xc5, 0x40, 0x9d, 0x4c, 0xd8,
	0x9a, 0xac, 0xad, 0x4e, 0x6e, 0x30, 0x1e, 0x40, 0x9d, 0xe6, 0x2c, 0x4e, 0xbc, 0xba, 0x26, 0x56,
	0x40, 0x4d, 0xb3, 0x5b, 0x9e, 0x92, 0xd7, 0x58, 0x4d, 0x35, 0xf0, 0x7f, 0x58, 0x50, 0x7f, 0x27,
	0xc2, 0x4a, 0xb6, 0x9a, 0xce, 0xa6, 0xf4, 0x24, 0xdf, 0x84, 0x3a, 0x98, 0x13, 0xac, 0x80, 0xba,
	0xf7, 0x6e, 0xc1, 0x52, 0x19, 0xcb, 0x7b, 0x1d, 0xaa, 0x1e, 0x18, 0x8c, 0x4f, 0xa1, 0x91, 0x4b,
	0x26, 0x17, 0xb9, 0x4e, 0xd4, 0x19, 0x1f, 0x0e, 0xcd, 0x2a, 0xf4, 0x15, 0x57, 0x9a, 0x0c, 0xd6,
	0x22, 0x65, 0x35, 0xe3, 0xf3, 0x2c, 0x21, 0x49, 0x3a, 0x69, 0x33, 0x30, 0xd8, 0xff, 0x6d, 0x81,
	0x33, 0x21, 0xf9, 0xa0, 0xb0, 0x21, 0x34, 0x67, 0xeb, 0xe6, 0x75, 0xae, 0xdd, 0x31, 0x16, 0x97,
	0x6c, 0x76, 0x12, 0x18, 0x8d, 0xd9, 0x86, 0x53, 0x6c, 0x43, 0x15, 0x9b, 0xdd, 0x72, 0xc9, 0xa7,
	0x22, 0x51, 0x49, 0x1d, 0x55, 0xac, 0x19, 0xe0, 0x09, 0xd4, 0x24, 0x8b, 0x72, 0xaf, 0xde, 0x77,
	0x4e, 0x77, 0xc7, 0xed, 0xc2, 0xfd, 0x9a, 0x45, 0x81, 0xa6, 0xf0, 0x89, 0x79, 0x67, 0x43, 0xbf,
	0x73, 0xbf, 0x10, 0x4d, 0x48, 0x56, 0x5f, 0xe9, 0xf7, 0xa1, 0x13, 0x10, 0x0b, 0x27, 0x24, 0x03,
	0xba, 0x5b, 0x50, 0xfe, 0xe0, 0x4d, 0xfe, 0x0e, 0xd4, 0x2f, 0xe7, 0x99, 0xbc, 0x1f, 0x70, 0xd8,
	0x2d, 0xf5, 0x84, 0x47, 0x80, 0x25, 0x38, 0x4d, 0x3f, 0xa5, 0xfc, 0x73, 0xea, 0xfe, 0x87, 0x87,
	0xb0, 0x57, 0x9a, 0x4f, 0x12, 0x36, 0xa3, 0xd0, 0xb5, 0xf0, 0x7f, 0xd8, 0x2f, 0x8d, 0xcf, 0xb2,
	0x4c, 0xf0, 0x25, 0x85, 0xae, 0x8d, 0x1e, 0x1c, 0x94, 0x88, 0x0b, 0x4a, 0xe2, 0x25, 0x09, 0x0a,
	0x5d, 0x67, 0x70, 0x03, 0x2d, 0x13, 0x18, 0x0f, 0xc0, 0x35, 0xa0, 0xb8, 0xec, 0x08, 0xd0, 0x4c,
	0xcf, 0x96, 0x2c, 0x4e, 0xd8, 0x4d, 0x42, 0xae, 0x55, 0x51, 0x4f, 0x28, 0x0d, 0xe3, 0x34, 0x72,
	0x6d, 0xdc, 0x83, 0xb6, 0x99, 0x5e, 0xf1, 0x24, 0x74, 0x9d, 0xf1, 0x2f, 0x1b, 0xba, 0x93, 0x75,
	0x3d, 0x57, 0x24, 0x96, 0xf1, 0x8c, 0xf0, 0x15, 0xb4, 0xce, 0x05, 0x31, 0x49, 0x6a, 0xc5, 0xed,
	0x4a, 0x7b, 0xbd, 0x6e, 0x01, 0x75, 0x2b, 0xfe, 0xfe, 0xb7, 0x3f, 0x7f, 0xbf, 0xdb, 0x6d, 0xbf,
	0x39, 0x5a, 0x3e, 0x1f, 0x29, 0xee, 0xa5, 0x35, 0xc0, 0xf7, 0xb0, 0xb3, 0x6e, 0x15, 0xbd, 0xe2,
	0x40, 0xb5, 0xe8, 0x5e, 0xd5, 0xd9, 0x3f, 0xd1, 0x46, 0x8f, 0xd0, 0x18, 0x7d, 0xe8, 0x62, 0x7b,
	0xf3, 0x3d, 0xfa, 0x12, 0x87, 0x5f, 0x9f, 0x59, 0x78, 0x09, 0xad, 0x69, 0x16, 0x6e, 0x19, 0xcd,
	0xd3, 0x8e, 0xd8, 0xab, 0xba, 0xa8, 0x7c, 0xaf, 0xa1, 0x75, 0x41, 0xea, 0x97, 0xbc, 0x8d, 0xcd,
	0xa1, 0xb6, 0xe9, 0x0e, 0xaa, 0x36, 0x37, 0x0d, 0xfd, 0xbf, 0xf4, 0xe2, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0xa7, 0xfd, 0x77, 0xd1, 0x04, 0x00, 0x00,
}
