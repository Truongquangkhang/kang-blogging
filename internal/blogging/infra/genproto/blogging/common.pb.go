// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: blogging/common.proto

package blogging

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo        *UserInfo           `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	UserFollowed    []*UserInfoMetadata `protobuf:"bytes,2,rep,name=user_followed,json=userFollowed,proto3" json:"user_followed,omitempty"`
	UserFollowingMe []*UserInfoMetadata `protobuf:"bytes,3,rep,name=user_following_me,json=userFollowingMe,proto3" json:"user_following_me,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *UserInfo) GetUserFollowed() []*UserInfoMetadata {
	if x != nil {
		return x.UserFollowed
	}
	return nil
}

func (x *UserInfo) GetUserFollowingMe() []*UserInfoMetadata {
	if x != nil {
		return x.UserFollowingMe
	}
	return nil
}

type UserInfoMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string                  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	DisplayName string                  `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	TotalBlogs  int32                   `protobuf:"varint,5,opt,name=total_blogs,json=totalBlogs,proto3" json:"total_blogs,omitempty"`
	Gender      *wrapperspb.BoolValue   `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Avatar      *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *UserInfoMetadata) Reset() {
	*x = UserInfoMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoMetadata) ProtoMessage() {}

func (x *UserInfoMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoMetadata.ProtoReflect.Descriptor instead.
func (*UserInfoMetadata) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfoMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfoMetadata) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoMetadata) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UserInfoMetadata) GetTotalBlogs() int32 {
	if x != nil {
		return x.TotalBlogs
	}
	return 0
}

func (x *UserInfoMetadata) GetGender() *wrapperspb.BoolValue {
	if x != nil {
		return x.Gender
	}
	return nil
}

func (x *UserInfoMetadata) GetAvatar() *wrapperspb.StringValue {
	if x != nil {
		return x.Avatar
	}
	return nil
}

type BlogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogInfo *BlogMetadata           `protobuf:"bytes,1,opt,name=blog_info,json=blogInfo,proto3" json:"blog_info,omitempty"`
	Content  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *BlogInfo) Reset() {
	*x = BlogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogInfo) ProtoMessage() {}

func (x *BlogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogInfo.ProtoReflect.Descriptor instead.
func (*BlogInfo) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{2}
}

func (x *BlogInfo) GetBlogInfo() *BlogMetadata {
	if x != nil {
		return x.BlogInfo
	}
	return nil
}

func (x *BlogInfo) GetContent() *wrapperspb.StringValue {
	if x != nil {
		return x.Content
	}
	return nil
}

type BlogMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description       string                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Categories        []*Category             `protobuf:"bytes,4,rep,name=categories,proto3" json:"categories,omitempty"`
	Thumbnail         *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	CreatedAt         int64                   `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Author            *UserInfoMetadata       `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	TotalBlogComments int32                   `protobuf:"varint,8,opt,name=total_blog_comments,json=totalBlogComments,proto3" json:"total_blog_comments,omitempty"`
}

func (x *BlogMetadata) Reset() {
	*x = BlogMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogMetadata) ProtoMessage() {}

func (x *BlogMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogMetadata.ProtoReflect.Descriptor instead.
func (*BlogMetadata) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{3}
}

func (x *BlogMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BlogMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlogMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BlogMetadata) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *BlogMetadata) GetThumbnail() *wrapperspb.StringValue {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *BlogMetadata) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *BlogMetadata) GetAuthor() *UserInfoMetadata {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *BlogMetadata) GetTotalBlogComments() int32 {
	if x != nil {
		return x.TotalBlogComments
	}
	return 0
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{4}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CategoryMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BlogCount int32  `protobuf:"varint,3,opt,name=blog_count,json=blogCount,proto3" json:"blog_count,omitempty"`
}

func (x *CategoryMetadata) Reset() {
	*x = CategoryMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryMetadata) ProtoMessage() {}

func (x *CategoryMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryMetadata.ProtoReflect.Descriptor instead.
func (*CategoryMetadata) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryMetadata) GetBlogCount() int32 {
	if x != nil {
		return x.BlogCount
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content    string            `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	IsToxicity bool              `protobuf:"varint,3,opt,name=is_toxicity,json=isToxicity,proto3" json:"is_toxicity,omitempty"`
	CreatedAt  int64             `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  int64             `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	User       *UserInfoMetadata `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{6}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetIsToxicity() bool {
	if x != nil {
		return x.IsToxicity
	}
	return false
}

func (x *Comment) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Comment) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Comment) GetUser() *UserInfoMetadata {
	if x != nil {
		return x.User
	}
	return nil
}

type CommentWithReplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment   `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Replies []*Comment `protobuf:"bytes,2,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *CommentWithReplies) Reset() {
	*x = CommentWithReplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentWithReplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentWithReplies) ProtoMessage() {}

func (x *CommentWithReplies) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentWithReplies.ProtoReflect.Descriptor instead.
func (*CommentWithReplies) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{7}
}

func (x *CommentWithReplies) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *CommentWithReplies) GetReplies() []*Comment {
	if x != nil {
		return x.Replies
	}
	return nil
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total    int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_blogging_common_proto_rawDescGZIP(), []int{8}
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_blogging_common_proto protoreflect.FileDescriptor

var file_blogging_common_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x3f, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x12, 0x46, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x77, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x33, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x62, 0x6c,
	0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xc7,
	0x02, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x67,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x10, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xc2, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x6f, 0x78, 0x69,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x6f,
	0x78, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogging_common_proto_rawDescOnce sync.Once
	file_blogging_common_proto_rawDescData = file_blogging_common_proto_rawDesc
)

func file_blogging_common_proto_rawDescGZIP() []byte {
	file_blogging_common_proto_rawDescOnce.Do(func() {
		file_blogging_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogging_common_proto_rawDescData)
	})
	return file_blogging_common_proto_rawDescData
}

var file_blogging_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_blogging_common_proto_goTypes = []interface{}{
	(*UserInfo)(nil),               // 0: blogging.UserInfo
	(*UserInfoMetadata)(nil),       // 1: blogging.UserInfoMetadata
	(*BlogInfo)(nil),               // 2: blogging.BlogInfo
	(*BlogMetadata)(nil),           // 3: blogging.BlogMetadata
	(*Category)(nil),               // 4: blogging.Category
	(*CategoryMetadata)(nil),       // 5: blogging.CategoryMetadata
	(*Comment)(nil),                // 6: blogging.Comment
	(*CommentWithReplies)(nil),     // 7: blogging.CommentWithReplies
	(*Pagination)(nil),             // 8: blogging.Pagination
	(*wrapperspb.BoolValue)(nil),   // 9: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
}
var file_blogging_common_proto_depIdxs = []int32{
	0,  // 0: blogging.UserInfo.user_info:type_name -> blogging.UserInfo
	1,  // 1: blogging.UserInfo.user_followed:type_name -> blogging.UserInfoMetadata
	1,  // 2: blogging.UserInfo.user_following_me:type_name -> blogging.UserInfoMetadata
	9,  // 3: blogging.UserInfoMetadata.gender:type_name -> google.protobuf.BoolValue
	10, // 4: blogging.UserInfoMetadata.avatar:type_name -> google.protobuf.StringValue
	3,  // 5: blogging.BlogInfo.blog_info:type_name -> blogging.BlogMetadata
	10, // 6: blogging.BlogInfo.content:type_name -> google.protobuf.StringValue
	4,  // 7: blogging.BlogMetadata.categories:type_name -> blogging.Category
	10, // 8: blogging.BlogMetadata.thumbnail:type_name -> google.protobuf.StringValue
	1,  // 9: blogging.BlogMetadata.author:type_name -> blogging.UserInfoMetadata
	1,  // 10: blogging.Comment.user:type_name -> blogging.UserInfoMetadata
	6,  // 11: blogging.CommentWithReplies.comment:type_name -> blogging.Comment
	6,  // 12: blogging.CommentWithReplies.replies:type_name -> blogging.Comment
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_blogging_common_proto_init() }
func file_blogging_common_proto_init() {
	if File_blogging_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blogging_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_blogging_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoMetadata); i {
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
		file_blogging_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogInfo); i {
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
		file_blogging_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogMetadata); i {
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
		file_blogging_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_blogging_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryMetadata); i {
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
		file_blogging_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_blogging_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentWithReplies); i {
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
		file_blogging_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			RawDescriptor: file_blogging_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blogging_common_proto_goTypes,
		DependencyIndexes: file_blogging_common_proto_depIdxs,
		MessageInfos:      file_blogging_common_proto_msgTypes,
	}.Build()
	File_blogging_common_proto = out.File
	file_blogging_common_proto_rawDesc = nil
	file_blogging_common_proto_goTypes = nil
	file_blogging_common_proto_depIdxs = nil
}
