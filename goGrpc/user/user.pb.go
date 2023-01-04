// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// UserProfile contains the user personal information
type UserProfile struct {
	// Id is the user id which uniquely identify a user
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Email is the unique email of the user
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// FirstName is the user first name,it does not have any min length as it can
	// be empty too.
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// LastName is the user last Name , it does not have any min length as it can
	// be empty too.
	LastName string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// BirthDate is the user`s date of birth
	BirthDate *timestamp.Timestamp `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	// Telephones are the user`s contact numbers and it should be saved with
	// country code.
	Telephones           []string `protobuf:"bytes,7,rep,name=telephones,proto3" json:"telephones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (m *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(m, src)
}
func (m *UserProfile) XXX_Size() int {
	return xxx_messageInfo_UserProfile.Size(m)
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserProfile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserProfile) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserProfile) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserProfile) GetBirthDate() *timestamp.Timestamp {
	if m != nil {
		return m.BirthDate
	}
	return nil
}

func (m *UserProfile) GetTelephones() []string {
	if m != nil {
		return m.Telephones
	}
	return nil
}

type CreateUserProfileRequest struct {
	UserProfile          *UserProfile `protobuf:"bytes,1,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateUserProfileRequest) Reset()         { *m = CreateUserProfileRequest{} }
func (m *CreateUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserProfileRequest) ProtoMessage()    {}
func (*CreateUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CreateUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserProfileRequest.Unmarshal(m, b)
}
func (m *CreateUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserProfileRequest.Merge(m, src)
}
func (m *CreateUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserProfileRequest.Size(m)
}
func (m *CreateUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserProfileRequest proto.InternalMessageInfo

func (m *CreateUserProfileRequest) GetUserProfile() *UserProfile {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetUserProfileRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserProfileRequest) Reset()         { *m = GetUserProfileRequest{} }
func (m *GetUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserProfileRequest) ProtoMessage()    {}
func (*GetUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GetUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserProfileRequest.Unmarshal(m, b)
}
func (m *GetUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserProfileRequest.Merge(m, src)
}
func (m *GetUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserProfileRequest.Size(m)
}
func (m *GetUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserProfileRequest proto.InternalMessageInfo

func (m *GetUserProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteUserProfileRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserProfileRequest) Reset()         { *m = DeleteUserProfileRequest{} }
func (m *DeleteUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserProfileRequest) ProtoMessage()    {}
func (*DeleteUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *DeleteUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserProfileRequest.Unmarshal(m, b)
}
func (m *DeleteUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserProfileRequest.Merge(m, src)
}
func (m *DeleteUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserProfileRequest.Size(m)
}
func (m *DeleteUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserProfileRequest proto.InternalMessageInfo

func (m *DeleteUserProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateUserProfileRequest struct {
	UserProfile          *UserProfile `protobuf:"bytes,1,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateUserProfileRequest) Reset()         { *m = UpdateUserProfileRequest{} }
func (m *UpdateUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileRequest) ProtoMessage()    {}
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UpdateUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserProfileRequest.Unmarshal(m, b)
}
func (m *UpdateUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserProfileRequest.Merge(m, src)
}
func (m *UpdateUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserProfileRequest.Size(m)
}
func (m *UpdateUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserProfileRequest proto.InternalMessageInfo

func (m *UpdateUserProfileRequest) GetUserProfile() *UserProfile {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type ListUsersProfilesRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersProfilesRequest) Reset()         { *m = ListUsersProfilesRequest{} }
func (m *ListUsersProfilesRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersProfilesRequest) ProtoMessage()    {}
func (*ListUsersProfilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *ListUsersProfilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersProfilesRequest.Unmarshal(m, b)
}
func (m *ListUsersProfilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersProfilesRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersProfilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersProfilesRequest.Merge(m, src)
}
func (m *ListUsersProfilesRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersProfilesRequest.Size(m)
}
func (m *ListUsersProfilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersProfilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersProfilesRequest proto.InternalMessageInfo

func (m *ListUsersProfilesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type ListUsersProfilesResponse struct {
	Profiles             []*UserProfile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListUsersProfilesResponse) Reset()         { *m = ListUsersProfilesResponse{} }
func (m *ListUsersProfilesResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersProfilesResponse) ProtoMessage()    {}
func (*ListUsersProfilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *ListUsersProfilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersProfilesResponse.Unmarshal(m, b)
}
func (m *ListUsersProfilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersProfilesResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersProfilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersProfilesResponse.Merge(m, src)
}
func (m *ListUsersProfilesResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersProfilesResponse.Size(m)
}
func (m *ListUsersProfilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersProfilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersProfilesResponse proto.InternalMessageInfo

func (m *ListUsersProfilesResponse) GetProfiles() []*UserProfile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func init() {
	proto.RegisterType((*UserProfile)(nil), "user.UserProfile")
	proto.RegisterType((*CreateUserProfileRequest)(nil), "user.CreateUserProfileRequest")
	proto.RegisterType((*GetUserProfileRequest)(nil), "user.GetUserProfileRequest")
	proto.RegisterType((*DeleteUserProfileRequest)(nil), "user.DeleteUserProfileRequest")
	proto.RegisterType((*UpdateUserProfileRequest)(nil), "user.UpdateUserProfileRequest")
	proto.RegisterType((*ListUsersProfilesRequest)(nil), "user.ListUsersProfilesRequest")
	proto.RegisterType((*ListUsersProfilesResponse)(nil), "user.ListUsersProfilesResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x7d, 0x49, 0x5b, 0x7d, 0xb9, 0x79, 0x3c, 0xc8, 0xf0, 0x84, 0x31, 0x0f, 0xdb, 0x90, 0x8d,
	0x45, 0x30, 0x95, 0xea, 0xc6, 0x6d, 0x2d, 0x58, 0x44, 0xa4, 0x04, 0xeb, 0xc2, 0x4d, 0x99, 0x90,
	0xdb, 0x76, 0x20, 0x69, 0xd2, 0x99, 0xc9, 0xc2, 0x9f, 0xf3, 0x03, 0xfc, 0x2a, 0xe9, 0x4c, 0xa3,
	0xb1, 0x49, 0x76, 0xee, 0x7a, 0xcf, 0x39, 0x3d, 0xf7, 0xcc, 0xb9, 0x01, 0xa8, 0x24, 0x8a, 0xa8,
	0x14, 0x85, 0x2a, 0xc8, 0xf0, 0xfc, 0xdb, 0x9f, 0xec, 0x8b, 0x62, 0x9f, 0xe1, 0x4c, 0x63, 0x49,
	0xb5, 0x9b, 0x29, 0x9e, 0xa3, 0x54, 0x2c, 0x2f, 0x8d, 0x2c, 0xfc, 0x65, 0x81, 0xbb, 0x91, 0x28,
	0xd6, 0xa2, 0xd8, 0xf1, 0x0c, 0xc9, 0x3d, 0xd8, 0x3c, 0xa5, 0x56, 0x60, 0x4d, 0x9d, 0xd8, 0xe6,
	0x29, 0x79, 0x80, 0x11, 0xe6, 0x8c, 0x67, 0xd4, 0xd6, 0x90, 0x19, 0xc8, 0x0b, 0x80, 0x1d, 0x17,
	0x52, 0x6d, 0x8f, 0x2c, 0x47, 0x3a, 0xd0, 0x94, 0xa3, 0x91, 0x2f, 0x2c, 0x47, 0xf2, 0x08, 0x4e,
	0xc6, 0x6a, 0x76, 0xa8, 0xd9, 0xdb, 0x33, 0xa0, 0xc9, 0xf7, 0x00, 0x09, 0x17, 0xea, 0xb0, 0x4d,
	0x99, 0x42, 0x3a, 0x0a, 0xac, 0xa9, 0x3b, 0xf7, 0x23, 0x93, 0x33, 0xaa, 0x73, 0x46, 0x5f, 0xeb,
	0x9c, 0xb1, 0xa3, 0xd5, 0x4b, 0xa6, 0x90, 0x8c, 0x01, 0x14, 0x66, 0x58, 0x1e, 0x8a, 0x23, 0x4a,
	0xfa, 0x34, 0x18, 0x4c, 0x9d, 0xb8, 0x81, 0x84, 0x6b, 0xa0, 0x1f, 0x04, 0x32, 0x85, 0x8d, 0x17,
	0xc5, 0x78, 0xaa, 0x50, 0x2a, 0xf2, 0x0e, 0xee, 0xce, 0x8d, 0x6c, 0x4b, 0x03, 0xeb, 0x27, 0xba,
	0x73, 0x2f, 0xd2, 0x95, 0x35, 0xf5, 0x6e, 0xf5, 0x77, 0x08, 0x5f, 0xc2, 0xb3, 0x8f, 0xa8, 0x3a,
	0xec, 0xae, 0x7a, 0x0a, 0x5f, 0x01, 0x5d, 0x62, 0x86, 0x9d, 0xab, 0xaf, 0xb5, 0x6b, 0xa0, 0x9b,
	0x32, 0xfd, 0x9f, 0x31, 0xdf, 0x00, 0xfd, 0xcc, 0xa5, 0xce, 0x29, 0x2f, 0x98, 0xac, 0x1d, 0x1f,
	0x60, 0x74, 0xaa, 0x50, 0xfc, 0xb8, 0x04, 0x30, 0x43, 0xf8, 0x09, 0x9e, 0x77, 0xfc, 0x43, 0x96,
	0xc5, 0x51, 0x22, 0x79, 0x0d, 0xb7, 0x97, 0xfd, 0x92, 0x5a, 0xc1, 0xa0, 0x3b, 0xc0, 0x1f, 0xc9,
	0xfc, 0xa7, 0x0d, 0x77, 0x0d, 0x46, 0x92, 0x15, 0x78, 0xad, 0x3b, 0x90, 0xb1, 0xb1, 0xe8, 0x3b,
	0x90, 0xdf, 0x5e, 0x11, 0xde, 0x90, 0x05, 0xdc, 0xff, 0xdb, 0x3f, 0x79, 0x34, 0xb2, 0xce, 0xab,
	0x74, 0x7b, 0xac, 0xc0, 0x6b, 0xd5, 0x5d, 0xa7, 0xe9, 0xbb, 0x43, 0xb7, 0xd3, 0x37, 0xf0, 0x5a,
	0xa5, 0xd5, 0x4e, 0x7d, 0xfd, 0xfb, 0x93, 0x5e, 0xde, 0xb4, 0x1d, 0xde, 0x2c, 0x86, 0xdf, 0xed,
	0x32, 0x49, 0x9e, 0xe8, 0x8f, 0xff, 0xed, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xe8, 0x8f,
	0xca, 0xc6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserProfilesClient is the client API for UserProfiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserProfilesClient interface {
	// CreateProfile creates new user profile.
	CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	// GetUserProfile returns the user profile by its unique user id.
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	ListUsersProfiles(ctx context.Context, in *ListUsersProfilesRequest, opts ...grpc.CallOption) (*ListUsersProfilesResponse, error)
}

type userProfilesClient struct {
	cc *grpc.ClientConn
}

func NewUserProfilesClient(cc *grpc.ClientConn) UserProfilesClient {
	return &userProfilesClient{cc}
}

func (c *userProfilesClient) CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/user.UserProfiles/CreateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfilesClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/user.UserProfiles/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfilesClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/user.UserProfiles/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfilesClient) ListUsersProfiles(ctx context.Context, in *ListUsersProfilesRequest, opts ...grpc.CallOption) (*ListUsersProfilesResponse, error) {
	out := new(ListUsersProfilesResponse)
	err := c.cc.Invoke(ctx, "/user.UserProfiles/ListUsersProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfilesServer is the server API for UserProfiles service.
type UserProfilesServer interface {
	// CreateProfile creates new user profile.
	CreateUserProfile(context.Context, *CreateUserProfileRequest) (*UserProfile, error)
	// GetUserProfile returns the user profile by its unique user id.
	GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfile, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfile, error)
	ListUsersProfiles(context.Context, *ListUsersProfilesRequest) (*ListUsersProfilesResponse, error)
}

// UnimplementedUserProfilesServer can be embedded to have forward compatible implementations.
type UnimplementedUserProfilesServer struct {
}

func (*UnimplementedUserProfilesServer) CreateUserProfile(ctx context.Context, req *CreateUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserProfile not implemented")
}
func (*UnimplementedUserProfilesServer) GetUserProfile(ctx context.Context, req *GetUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (*UnimplementedUserProfilesServer) UpdateUserProfile(ctx context.Context, req *UpdateUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (*UnimplementedUserProfilesServer) ListUsersProfiles(ctx context.Context, req *ListUsersProfilesRequest) (*ListUsersProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersProfiles not implemented")
}

func RegisterUserProfilesServer(s *grpc.Server, srv UserProfilesServer) {
	s.RegisterService(&_UserProfiles_serviceDesc, srv)
}

func _UserProfiles_CreateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).CreateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserProfiles/CreateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).CreateUserProfile(ctx, req.(*CreateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfiles_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserProfiles/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfiles_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserProfiles/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfiles_ListUsersProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).ListUsersProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserProfiles/ListUsersProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).ListUsersProfiles(ctx, req.(*ListUsersProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserProfiles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserProfiles",
	HandlerType: (*UserProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserProfile",
			Handler:    _UserProfiles_CreateUserProfile_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserProfiles_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserProfiles_UpdateUserProfile_Handler,
		},
		{
			MethodName: "ListUsersProfiles",
			Handler:    _UserProfiles_ListUsersProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
