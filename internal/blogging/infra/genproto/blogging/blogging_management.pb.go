// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: blogging/blogging_management.proto

package blogging

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDashboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDashboardRequest) Reset() {
	*x = GetDashboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_blogging_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardRequest) ProtoMessage() {}

func (x *GetDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_blogging_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardRequest.ProtoReflect.Descriptor instead.
func (*GetDashboardRequest) Descriptor() ([]byte, []int) {
	return file_blogging_blogging_management_proto_rawDescGZIP(), []int{0}
}

type GetDashboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *GetDashboardResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDashboardResponse) Reset() {
	*x = GetDashboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_blogging_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardResponse) ProtoMessage() {}

func (x *GetDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_blogging_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardResponse.ProtoReflect.Descriptor instead.
func (*GetDashboardResponse) Descriptor() ([]byte, []int) {
	return file_blogging_blogging_management_proto_rawDescGZIP(), []int{1}
}

func (x *GetDashboardResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetDashboardResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetDashboardResponse) GetData() *GetDashboardResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPoliciesRequest) Reset() {
	*x = GetPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_blogging_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoliciesRequest) ProtoMessage() {}

func (x *GetPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_blogging_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoliciesRequest.ProtoReflect.Descriptor instead.
func (*GetPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_blogging_blogging_management_proto_rawDescGZIP(), []int{2}
}

type GetPoliciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *GetPoliciesResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPoliciesResponse) Reset() {
	*x = GetPoliciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_blogging_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoliciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoliciesResponse) ProtoMessage() {}

func (x *GetPoliciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_blogging_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoliciesResponse.ProtoReflect.Descriptor instead.
func (*GetPoliciesResponse) Descriptor() ([]byte, []int) {
	return file_blogging_blogging_management_proto_rawDescGZIP(), []int{3}
}

func (x *GetPoliciesResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetPoliciesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetPoliciesResponse) GetData() *GetPoliciesResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetDashboardResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalBlogs            int32           `protobuf:"varint,1,opt,name=total_blogs,json=totalBlogs,proto3" json:"total_blogs,omitempty"`
	TotalComments         int32           `protobuf:"varint,2,opt,name=total_comments,json=totalComments,proto3" json:"total_comments,omitempty"`
	TotalUsers            int32           `protobuf:"varint,3,opt,name=total_users,json=totalUsers,proto3" json:"total_users,omitempty"`
	TotalCategories       int32           `protobuf:"varint,4,opt,name=total_categories,json=totalCategories,proto3" json:"total_categories,omitempty"`
	BlogsIncreaseInDay    int32           `protobuf:"varint,5,opt,name=blogs_increase_in_day,json=blogsIncreaseInDay,proto3" json:"blogs_increase_in_day,omitempty"`
	CommentsIncreaseInDay int32           `protobuf:"varint,6,opt,name=comments_increase_in_day,json=commentsIncreaseInDay,proto3" json:"comments_increase_in_day,omitempty"`
	UsersIncreaseInDay    int32           `protobuf:"varint,7,opt,name=users_increase_in_day,json=usersIncreaseInDay,proto3" json:"users_increase_in_day,omitempty"`
	LatestBlogs           []*BlogMetadata `protobuf:"bytes,8,rep,name=latest_blogs,json=latestBlogs,proto3" json:"latest_blogs,omitempty"`
	LatestComments        []*Comment      `protobuf:"bytes,9,rep,name=latest_comments,json=latestComments,proto3" json:"latest_comments,omitempty"`
}

func (x *GetDashboardResponse_Data) Reset() {
	*x = GetDashboardResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_blogging_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDashboardResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardResponse_Data) ProtoMessage() {}

func (x *GetDashboardResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_blogging_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardResponse_Data.ProtoReflect.Descriptor instead.
func (*GetDashboardResponse_Data) Descriptor() ([]byte, []int) {
	return file_blogging_blogging_management_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetDashboardResponse_Data) GetTotalBlogs() int32 {
	if x != nil {
		return x.TotalBlogs
	}
	return 0
}

func (x *GetDashboardResponse_Data) GetTotalComments() int32 {
	if x != nil {
		return x.TotalComments
	}
	return 0
}

func (x *GetDashboardResponse_Data) GetTotalUsers() int32 {
	if x != nil {
		return x.TotalUsers
	}
	return 0
}

func (x *GetDashboardResponse_Data) GetTotalCategories() int32 {
	if x != nil {
		return x.TotalCategories
	}
	return 0
}

func (x *GetDashboardResponse_Data) GetBlogsIncreaseInDay() int32 {
	if x != nil {
		return x.BlogsIncreaseInDay
	}
	return 0
}

func (x *GetDashboardResponse_Data) GetCommentsIncreaseInDay() int32 {
	if x != nil {
		return x.CommentsIncreaseInDay
	}
	return 0
}

func (x *GetDashboardResponse_Data) GetUsersIncreaseInDay() int32 {
	if x != nil {
		return x.UsersIncreaseInDay
	}
	return 0
}

func (x *GetDashboardResponse_Data) GetLatestBlogs() []*BlogMetadata {
	if x != nil {
		return x.LatestBlogs
	}
	return nil
}

func (x *GetDashboardResponse_Data) GetLatestComments() []*Comment {
	if x != nil {
		return x.LatestComments
	}
	return nil
}

type GetPoliciesResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *GetPoliciesResponse_Data) Reset() {
	*x = GetPoliciesResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_blogging_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoliciesResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoliciesResponse_Data) ProtoMessage() {}

func (x *GetPoliciesResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_blogging_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoliciesResponse_Data.ProtoReflect.Descriptor instead.
func (*GetPoliciesResponse_Data) Descriptor() ([]byte, []int) {
	return file_blogging_blogging_management_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetPoliciesResponse_Data) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

var File_blogging_blogging_management_proto protoreflect.FileDescriptor

var file_blogging_blogging_management_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb0, 0x04, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xb0, 0x03, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x10,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x62, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x12, 0x37, 0x0a, 0x18, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x44, 0x61, 0x79, 0x12, 0x31, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x75, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x73, 0x12, 0x3a, 0x0a, 0x0f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x14, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x34, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x32, 0xff, 0x01, 0x0a, 0x19, 0x42, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x6d, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogging_blogging_management_proto_rawDescOnce sync.Once
	file_blogging_blogging_management_proto_rawDescData = file_blogging_blogging_management_proto_rawDesc
)

func file_blogging_blogging_management_proto_rawDescGZIP() []byte {
	file_blogging_blogging_management_proto_rawDescOnce.Do(func() {
		file_blogging_blogging_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogging_blogging_management_proto_rawDescData)
	})
	return file_blogging_blogging_management_proto_rawDescData
}

var file_blogging_blogging_management_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_blogging_blogging_management_proto_goTypes = []interface{}{
	(*GetDashboardRequest)(nil),       // 0: blogging.GetDashboardRequest
	(*GetDashboardResponse)(nil),      // 1: blogging.GetDashboardResponse
	(*GetPoliciesRequest)(nil),        // 2: blogging.GetPoliciesRequest
	(*GetPoliciesResponse)(nil),       // 3: blogging.GetPoliciesResponse
	(*GetDashboardResponse_Data)(nil), // 4: blogging.GetDashboardResponse.Data
	(*GetPoliciesResponse_Data)(nil),  // 5: blogging.GetPoliciesResponse.Data
	(*BlogMetadata)(nil),              // 6: blogging.BlogMetadata
	(*Comment)(nil),                   // 7: blogging.Comment
	(*Policy)(nil),                    // 8: blogging.Policy
}
var file_blogging_blogging_management_proto_depIdxs = []int32{
	4, // 0: blogging.GetDashboardResponse.data:type_name -> blogging.GetDashboardResponse.Data
	5, // 1: blogging.GetPoliciesResponse.data:type_name -> blogging.GetPoliciesResponse.Data
	6, // 2: blogging.GetDashboardResponse.Data.latest_blogs:type_name -> blogging.BlogMetadata
	7, // 3: blogging.GetDashboardResponse.Data.latest_comments:type_name -> blogging.Comment
	8, // 4: blogging.GetPoliciesResponse.Data.policies:type_name -> blogging.Policy
	0, // 5: blogging.BloggingManagementService.GetDashboard:input_type -> blogging.GetDashboardRequest
	2, // 6: blogging.BloggingManagementService.GetPolicies:input_type -> blogging.GetPoliciesRequest
	1, // 7: blogging.BloggingManagementService.GetDashboard:output_type -> blogging.GetDashboardResponse
	3, // 8: blogging.BloggingManagementService.GetPolicies:output_type -> blogging.GetPoliciesResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_blogging_blogging_management_proto_init() }
func file_blogging_blogging_management_proto_init() {
	if File_blogging_blogging_management_proto != nil {
		return
	}
	file_blogging_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blogging_blogging_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardRequest); i {
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
		file_blogging_blogging_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardResponse); i {
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
		file_blogging_blogging_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoliciesRequest); i {
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
		file_blogging_blogging_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoliciesResponse); i {
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
		file_blogging_blogging_management_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDashboardResponse_Data); i {
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
		file_blogging_blogging_management_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoliciesResponse_Data); i {
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
			RawDescriptor: file_blogging_blogging_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogging_blogging_management_proto_goTypes,
		DependencyIndexes: file_blogging_blogging_management_proto_depIdxs,
		MessageInfos:      file_blogging_blogging_management_proto_msgTypes,
	}.Build()
	File_blogging_blogging_management_proto = out.File
	file_blogging_blogging_management_proto_rawDesc = nil
	file_blogging_blogging_management_proto_goTypes = nil
	file_blogging_blogging_management_proto_depIdxs = nil
}
