// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/auth.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email           string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone           string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Password        string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,6,opt,name=ConfirmPassword,proto3" json:"ConfirmPassword,omitempty"`
	ReferalCode     string `protobuf:"bytes,7,opt,name=ReferalCode,proto3" json:"ReferalCode,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SignupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignupRequest) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

func (x *SignupRequest) GetReferalCode() string {
	if x != nil {
		return x.ReferalCode
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone         string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	OTP           string `protobuf:"bytes,5,opt,name=OTP,proto3" json:"OTP,omitempty"`
	Token         string `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`
	IsUserExist   string `protobuf:"bytes,7,opt,name=IsUserExist,proto3" json:"IsUserExist,omitempty"`
	ReferalCode   string `protobuf:"bytes,8,opt,name=ReferalCode,proto3" json:"ReferalCode,omitempty"`
	WalletBelance uint32 `protobuf:"varint,9,opt,name=WalletBelance,proto3" json:"WalletBelance,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignupResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SignupResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignupResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignupResponse) GetOTP() string {
	if x != nil {
		return x.OTP
	}
	return ""
}

func (x *SignupResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignupResponse) GetIsUserExist() string {
	if x != nil {
		return x.IsUserExist
	}
	return ""
}

func (x *SignupResponse) GetReferalCode() string {
	if x != nil {
		return x.ReferalCode
	}
	return ""
}

func (x *SignupResponse) GetWalletBelance() uint32 {
	if x != nil {
		return x.WalletBelance
	}
	return 0
}

type OtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp            string `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
	TemperverToken string `protobuf:"bytes,2,opt,name=temperverToken,proto3" json:"temperverToken,omitempty"`
}

func (x *OtpRequest) Reset() {
	*x = OtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpRequest) ProtoMessage() {}

func (x *OtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpRequest.ProtoReflect.Descriptor instead.
func (*OtpRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{2}
}

func (x *OtpRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *OtpRequest) GetTemperverToken() string {
	if x != nil {
		return x.TemperverToken
	}
	return ""
}

type OtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *OtpResponse) Reset() {
	*x = OtpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpResponse) ProtoMessage() {}

func (x *OtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpResponse.ProtoReflect.Descriptor instead.
func (*OtpResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{3}
}

func (x *OtpResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *OtpResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone    string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type AdminLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AdminLoginRequest) Reset() {
	*x = AdminLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginRequest) ProtoMessage() {}

func (x *AdminLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginRequest.ProtoReflect.Descriptor instead.
func (*AdminLoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{6}
}

func (x *AdminLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AdminLoginResponse) Reset() {
	*x = AdminLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginResponse) ProtoMessage() {}

func (x *AdminLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginResponse.ProtoReflect.Descriptor instead.
func (*AdminLoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{7}
}

func (x *AdminLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *ValidateTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ValidateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *ValidateTokenResponse) Reset() {
	*x = ValidateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResponse) ProtoMessage() {}

func (x *ValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{9}
}

func (x *ValidateTokenResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ValidateAdminTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateAdminTokenRequest) Reset() {
	*x = ValidateAdminTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAdminTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAdminTokenRequest) ProtoMessage() {}

func (x *ValidateAdminTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAdminTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateAdminTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_proto_rawDescGZIP(), []int{10}
}

func (x *ValidateAdminTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_pkg_pb_auth_proto protoreflect.FileDescriptor

var file_pkg_pb_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0xf2, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4f, 0x54, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4f, 0x54, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x49,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x65,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x46, 0x0a, 0x0a, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x53, 0x0a,
	0x0b, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x40, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x55, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x45, 0x0a, 0x11, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5c,
	0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2f, 0x0a, 0x15,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x31, 0x0a,
	0x19, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x32, 0x99, 0x03, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x68, 0x75, 0x70, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0e, 0x4f, 0x74, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x11, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_pb_auth_proto_rawDescOnce sync.Once
	file_pkg_pb_auth_proto_rawDescData = file_pkg_pb_auth_proto_rawDesc
)

func file_pkg_pb_auth_proto_rawDescGZIP() []byte {
	file_pkg_pb_auth_proto_rawDescOnce.Do(func() {
		file_pkg_pb_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_auth_proto_rawDescData)
	})
	return file_pkg_pb_auth_proto_rawDescData
}

var file_pkg_pb_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_pb_auth_proto_goTypes = []interface{}{
	(*SignupRequest)(nil),             // 0: auth.SignupRequest
	(*SignupResponse)(nil),            // 1: auth.SignupResponse
	(*OtpRequest)(nil),                // 2: auth.OtpRequest
	(*OtpResponse)(nil),               // 3: auth.OtpResponse
	(*LoginRequest)(nil),              // 4: auth.loginRequest
	(*LoginResponse)(nil),             // 5: auth.loginResponse
	(*AdminLoginRequest)(nil),         // 6: auth.adminLoginRequest
	(*AdminLoginResponse)(nil),        // 7: auth.adminLoginResponse
	(*ValidateTokenRequest)(nil),      // 8: auth.ValidateTokenRequest
	(*ValidateTokenResponse)(nil),     // 9: auth.ValidateTokenResponse
	(*ValidateAdminTokenRequest)(nil), // 10: auth.ValidateAdminTokenRequest
	(*empty.Empty)(nil),               // 11: google.protobuf.Empty
}
var file_pkg_pb_auth_proto_depIdxs = []int32{
	0,  // 0: auth.AuthService.Sighup:input_type -> auth.SignupRequest
	2,  // 1: auth.AuthService.OtpVerifiction:input_type -> auth.OtpRequest
	4,  // 2: auth.AuthService.UserLogin:input_type -> auth.loginRequest
	6,  // 3: auth.AuthService.AdminLogin:input_type -> auth.adminLoginRequest
	8,  // 4: auth.AuthService.ValidateUserToken:input_type -> auth.ValidateTokenRequest
	10, // 5: auth.AuthService.ValidateAdminToken:input_type -> auth.ValidateAdminTokenRequest
	1,  // 6: auth.AuthService.Sighup:output_type -> auth.SignupResponse
	3,  // 7: auth.AuthService.OtpVerifiction:output_type -> auth.OtpResponse
	5,  // 8: auth.AuthService.UserLogin:output_type -> auth.loginResponse
	7,  // 9: auth.AuthService.AdminLogin:output_type -> auth.adminLoginResponse
	9,  // 10: auth.AuthService.ValidateUserToken:output_type -> auth.ValidateTokenResponse
	11, // 11: auth.AuthService.ValidateAdminToken:output_type -> google.protobuf.Empty
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_auth_proto_init() }
func file_pkg_pb_auth_proto_init() {
	if File_pkg_pb_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAdminTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_auth_proto_goTypes,
		DependencyIndexes: file_pkg_pb_auth_proto_depIdxs,
		MessageInfos:      file_pkg_pb_auth_proto_msgTypes,
	}.Build()
	File_pkg_pb_auth_proto = out.File
	file_pkg_pb_auth_proto_rawDesc = nil
	file_pkg_pb_auth_proto_goTypes = nil
	file_pkg_pb_auth_proto_depIdxs = nil
}
