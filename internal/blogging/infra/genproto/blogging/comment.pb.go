// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: blogging/comment.proto

package blogging

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetBlogCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	BlogId   string `protobuf:"bytes,3,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
}

func (x *GetBlogCommentsRequest) Reset() {
	*x = GetBlogCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogCommentsRequest) ProtoMessage() {}

func (x *GetBlogCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetBlogCommentsRequest) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlogCommentsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetBlogCommentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetBlogCommentsRequest) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

type GetBlogCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *GetBlogCommentsResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetBlogCommentsResponse) Reset() {
	*x = GetBlogCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogCommentsResponse) ProtoMessage() {}

func (x *GetBlogCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetBlogCommentsResponse) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlogCommentsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetBlogCommentsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetBlogCommentsResponse) GetData() *GetBlogCommentsResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateBlogCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId         string                  `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	Content        string                  `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	ReplyCommentId *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=reply_comment_id,json=replyCommentId,proto3" json:"reply_comment_id,omitempty"`
}

func (x *CreateBlogCommentsRequest) Reset() {
	*x = CreateBlogCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogCommentsRequest) ProtoMessage() {}

func (x *CreateBlogCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogCommentsRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogCommentsRequest) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBlogCommentsRequest) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

func (x *CreateBlogCommentsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateBlogCommentsRequest) GetReplyCommentId() *wrapperspb.StringValue {
	if x != nil {
		return x.ReplyCommentId
	}
	return nil
}

type CreateBlogCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *CreateBlogCommentsResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateBlogCommentsResponse) Reset() {
	*x = CreateBlogCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogCommentsResponse) ProtoMessage() {}

func (x *CreateBlogCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogCommentsResponse.ProtoReflect.Descriptor instead.
func (*CreateBlogCommentsResponse) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBlogCommentsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateBlogCommentsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateBlogCommentsResponse) GetData() *CreateBlogCommentsResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCommentsByParamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32                   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int32                   `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	SearchName *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=searchName,proto3" json:"searchName,omitempty"`
	SortBy     *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	IsToxicity *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=isToxicity,proto3" json:"isToxicity,omitempty"`
	UserIds    *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *GetCommentsByParamRequest) Reset() {
	*x = GetCommentsByParamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByParamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByParamRequest) ProtoMessage() {}

func (x *GetCommentsByParamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByParamRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsByParamRequest) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentsByParamRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetCommentsByParamRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetCommentsByParamRequest) GetSearchName() *wrapperspb.StringValue {
	if x != nil {
		return x.SearchName
	}
	return nil
}

func (x *GetCommentsByParamRequest) GetSortBy() *wrapperspb.StringValue {
	if x != nil {
		return x.SortBy
	}
	return nil
}

func (x *GetCommentsByParamRequest) GetIsToxicity() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsToxicity
	}
	return nil
}

func (x *GetCommentsByParamRequest) GetUserIds() *wrapperspb.StringValue {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetCommentsByParamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *GetCommentsByParamResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCommentsByParamResponse) Reset() {
	*x = GetCommentsByParamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByParamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByParamResponse) ProtoMessage() {}

func (x *GetCommentsByParamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByParamResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsByParamResponse) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommentsByParamResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetCommentsByParamResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetCommentsByParamResponse) GetData() *GetCommentsByParamResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetBlogCommentsResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments   []*CommentWithReplies `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Pagination *Pagination           `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetBlogCommentsResponse_Data) Reset() {
	*x = GetBlogCommentsResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogCommentsResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogCommentsResponse_Data) ProtoMessage() {}

func (x *GetBlogCommentsResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogCommentsResponse_Data.ProtoReflect.Descriptor instead.
func (*GetBlogCommentsResponse_Data) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetBlogCommentsResponse_Data) GetComments() []*CommentWithReplies {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetBlogCommentsResponse_Data) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreateBlogCommentsResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateBlogCommentsResponse_Data) Reset() {
	*x = CreateBlogCommentsResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogCommentsResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogCommentsResponse_Data) ProtoMessage() {}

func (x *CreateBlogCommentsResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogCommentsResponse_Data.ProtoReflect.Descriptor instead.
func (*CreateBlogCommentsResponse_Data) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CreateBlogCommentsResponse_Data) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type GetCommentsByParamResponse_CommentItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentInfo    *Comment                `protobuf:"bytes,1,opt,name=comment_info,json=commentInfo,proto3" json:"comment_info,omitempty"`
	BlogId         string                  `protobuf:"bytes,2,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	ReplyCommentId *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=reply_comment_id,json=replyCommentId,proto3" json:"reply_comment_id,omitempty"`
}

func (x *GetCommentsByParamResponse_CommentItem) Reset() {
	*x = GetCommentsByParamResponse_CommentItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByParamResponse_CommentItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByParamResponse_CommentItem) ProtoMessage() {}

func (x *GetCommentsByParamResponse_CommentItem) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByParamResponse_CommentItem.ProtoReflect.Descriptor instead.
func (*GetCommentsByParamResponse_CommentItem) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetCommentsByParamResponse_CommentItem) GetCommentInfo() *Comment {
	if x != nil {
		return x.CommentInfo
	}
	return nil
}

func (x *GetCommentsByParamResponse_CommentItem) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

func (x *GetCommentsByParamResponse_CommentItem) GetReplyCommentId() *wrapperspb.StringValue {
	if x != nil {
		return x.ReplyCommentId
	}
	return nil
}

type GetCommentsByParamResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments   []*GetCommentsByParamResponse_CommentItem `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Pagination *Pagination                               `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetCommentsByParamResponse_Data) Reset() {
	*x = GetCommentsByParamResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByParamResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByParamResponse_Data) ProtoMessage() {}

func (x *GetCommentsByParamResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByParamResponse_Data.ProtoReflect.Descriptor instead.
func (*GetCommentsByParamResponse_Data) Descriptor() ([]byte, []int) {
	return file_blogging_comment_proto_rawDescGZIP(), []int{5, 1}
}

func (x *GetCommentsByParamResponse_Data) GetComments() []*GetCommentsByParamResponse_CommentItem {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetCommentsByParamResponse_Data) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_blogging_comment_proto protoreflect.FileDescriptor

var file_blogging_comment_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0xfb, 0x01, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x76, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x10, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0xbe, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x33,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0xb4, 0x02, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x69, 0x73, 0x54, 0x6f, 0x78, 0x69, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x6f, 0x78, 0x69, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x37, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xbd, 0x03, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xa4, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07,
	0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x8a, 0x01,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x96, 0x03, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x20, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x7b, 0x62, 0x6c, 0x6f,
	0x67, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x89, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x7b, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x78, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x23, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogging_comment_proto_rawDescOnce sync.Once
	file_blogging_comment_proto_rawDescData = file_blogging_comment_proto_rawDesc
)

func file_blogging_comment_proto_rawDescGZIP() []byte {
	file_blogging_comment_proto_rawDescOnce.Do(func() {
		file_blogging_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogging_comment_proto_rawDescData)
	})
	return file_blogging_comment_proto_rawDescData
}

var file_blogging_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_blogging_comment_proto_goTypes = []interface{}{
	(*GetBlogCommentsRequest)(nil),                 // 0: blogging.GetBlogCommentsRequest
	(*GetBlogCommentsResponse)(nil),                // 1: blogging.GetBlogCommentsResponse
	(*CreateBlogCommentsRequest)(nil),              // 2: blogging.CreateBlogCommentsRequest
	(*CreateBlogCommentsResponse)(nil),             // 3: blogging.CreateBlogCommentsResponse
	(*GetCommentsByParamRequest)(nil),              // 4: blogging.GetCommentsByParamRequest
	(*GetCommentsByParamResponse)(nil),             // 5: blogging.GetCommentsByParamResponse
	(*GetBlogCommentsResponse_Data)(nil),           // 6: blogging.GetBlogCommentsResponse.Data
	(*CreateBlogCommentsResponse_Data)(nil),        // 7: blogging.CreateBlogCommentsResponse.Data
	(*GetCommentsByParamResponse_CommentItem)(nil), // 8: blogging.GetCommentsByParamResponse.CommentItem
	(*GetCommentsByParamResponse_Data)(nil),        // 9: blogging.GetCommentsByParamResponse.Data
	(*wrapperspb.StringValue)(nil),                 // 10: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),                   // 11: google.protobuf.BoolValue
	(*CommentWithReplies)(nil),                     // 12: blogging.CommentWithReplies
	(*Pagination)(nil),                             // 13: blogging.Pagination
	(*Comment)(nil),                                // 14: blogging.Comment
}
var file_blogging_comment_proto_depIdxs = []int32{
	6,  // 0: blogging.GetBlogCommentsResponse.data:type_name -> blogging.GetBlogCommentsResponse.Data
	10, // 1: blogging.CreateBlogCommentsRequest.reply_comment_id:type_name -> google.protobuf.StringValue
	7,  // 2: blogging.CreateBlogCommentsResponse.data:type_name -> blogging.CreateBlogCommentsResponse.Data
	10, // 3: blogging.GetCommentsByParamRequest.searchName:type_name -> google.protobuf.StringValue
	10, // 4: blogging.GetCommentsByParamRequest.sortBy:type_name -> google.protobuf.StringValue
	11, // 5: blogging.GetCommentsByParamRequest.isToxicity:type_name -> google.protobuf.BoolValue
	10, // 6: blogging.GetCommentsByParamRequest.user_ids:type_name -> google.protobuf.StringValue
	9,  // 7: blogging.GetCommentsByParamResponse.data:type_name -> blogging.GetCommentsByParamResponse.Data
	12, // 8: blogging.GetBlogCommentsResponse.Data.comments:type_name -> blogging.CommentWithReplies
	13, // 9: blogging.GetBlogCommentsResponse.Data.pagination:type_name -> blogging.Pagination
	14, // 10: blogging.CreateBlogCommentsResponse.Data.comment:type_name -> blogging.Comment
	14, // 11: blogging.GetCommentsByParamResponse.CommentItem.comment_info:type_name -> blogging.Comment
	10, // 12: blogging.GetCommentsByParamResponse.CommentItem.reply_comment_id:type_name -> google.protobuf.StringValue
	8,  // 13: blogging.GetCommentsByParamResponse.Data.comments:type_name -> blogging.GetCommentsByParamResponse.CommentItem
	13, // 14: blogging.GetCommentsByParamResponse.Data.pagination:type_name -> blogging.Pagination
	0,  // 15: blogging.CommentService.GetBlogComments:input_type -> blogging.GetBlogCommentsRequest
	2,  // 16: blogging.CommentService.CreateBlogComment:input_type -> blogging.CreateBlogCommentsRequest
	4,  // 17: blogging.CommentService.GetCommentsByParam:input_type -> blogging.GetCommentsByParamRequest
	1,  // 18: blogging.CommentService.GetBlogComments:output_type -> blogging.GetBlogCommentsResponse
	3,  // 19: blogging.CommentService.CreateBlogComment:output_type -> blogging.CreateBlogCommentsResponse
	5,  // 20: blogging.CommentService.GetCommentsByParam:output_type -> blogging.GetCommentsByParamResponse
	18, // [18:21] is the sub-list for method output_type
	15, // [15:18] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_blogging_comment_proto_init() }
func file_blogging_comment_proto_init() {
	if File_blogging_comment_proto != nil {
		return
	}
	file_blogging_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blogging_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogCommentsRequest); i {
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
		file_blogging_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogCommentsResponse); i {
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
		file_blogging_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogCommentsRequest); i {
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
		file_blogging_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogCommentsResponse); i {
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
		file_blogging_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByParamRequest); i {
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
		file_blogging_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByParamResponse); i {
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
		file_blogging_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogCommentsResponse_Data); i {
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
		file_blogging_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogCommentsResponse_Data); i {
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
		file_blogging_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByParamResponse_CommentItem); i {
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
		file_blogging_comment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByParamResponse_Data); i {
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
			RawDescriptor: file_blogging_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogging_comment_proto_goTypes,
		DependencyIndexes: file_blogging_comment_proto_depIdxs,
		MessageInfos:      file_blogging_comment_proto_msgTypes,
	}.Build()
	File_blogging_comment_proto = out.File
	file_blogging_comment_proto_rawDesc = nil
	file_blogging_comment_proto_goTypes = nil
	file_blogging_comment_proto_depIdxs = nil
}
