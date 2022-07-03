// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/post.proto

package post

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type ListLikeResponse struct {
	List                 []*LikeItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListLikeResponse) Reset()         { *m = ListLikeResponse{} }
func (m *ListLikeResponse) String() string { return proto.CompactTextString(m) }
func (*ListLikeResponse) ProtoMessage()    {}
func (*ListLikeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{0}
}

func (m *ListLikeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLikeResponse.Unmarshal(m, b)
}
func (m *ListLikeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLikeResponse.Marshal(b, m, deterministic)
}
func (m *ListLikeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLikeResponse.Merge(m, src)
}
func (m *ListLikeResponse) XXX_Size() int {
	return xxx_messageInfo_ListLikeResponse.Size(m)
}
func (m *ListLikeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLikeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLikeResponse proto.InternalMessageInfo

func (m *ListLikeResponse) GetList() []*LikeItem {
	if m != nil {
		return m.List
	}
	return nil
}

type LikeItem struct {
	TargetId             uint32   `protobuf:"varint,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	TypeId               uint32   `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LikeItem) Reset()         { *m = LikeItem{} }
func (m *LikeItem) String() string { return proto.CompactTextString(m) }
func (*LikeItem) ProtoMessage()    {}
func (*LikeItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{1}
}

func (m *LikeItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LikeItem.Unmarshal(m, b)
}
func (m *LikeItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LikeItem.Marshal(b, m, deterministic)
}
func (m *LikeItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LikeItem.Merge(m, src)
}
func (m *LikeItem) XXX_Size() int {
	return xxx_messageInfo_LikeItem.Size(m)
}
func (m *LikeItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LikeItem.DiscardUnknown(m)
}

var xxx_messageInfo_LikeItem proto.InternalMessageInfo

func (m *LikeItem) GetTargetId() uint32 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *LikeItem) GetTypeId() uint32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type LikeRequest struct {
	UserId               uint32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Item                 *LikeItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LikeRequest) Reset()         { *m = LikeRequest{} }
func (m *LikeRequest) String() string { return proto.CompactTextString(m) }
func (*LikeRequest) ProtoMessage()    {}
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{2}
}

func (m *LikeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LikeRequest.Unmarshal(m, b)
}
func (m *LikeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LikeRequest.Marshal(b, m, deterministic)
}
func (m *LikeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LikeRequest.Merge(m, src)
}
func (m *LikeRequest) XXX_Size() int {
	return xxx_messageInfo_LikeRequest.Size(m)
}
func (m *LikeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LikeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LikeRequest proto.InternalMessageInfo

func (m *LikeRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LikeRequest) GetItem() *LikeItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateCommentRequest struct {
	PostId               uint32   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	TypeId               uint32   `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	FatherId             uint32   `protobuf:"varint,3,opt,name=father_id,json=fatherId,proto3" json:"father_id,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatorId            uint32   `protobuf:"varint,5,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCommentRequest) Reset()         { *m = CreateCommentRequest{} }
func (m *CreateCommentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCommentRequest) ProtoMessage()    {}
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{3}
}

func (m *CreateCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCommentRequest.Unmarshal(m, b)
}
func (m *CreateCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCommentRequest.Marshal(b, m, deterministic)
}
func (m *CreateCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCommentRequest.Merge(m, src)
}
func (m *CreateCommentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCommentRequest.Size(m)
}
func (m *CreateCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCommentRequest proto.InternalMessageInfo

func (m *CreateCommentRequest) GetPostId() uint32 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *CreateCommentRequest) GetTypeId() uint32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *CreateCommentRequest) GetFatherId() uint32 {
	if m != nil {
		return m.FatherId
	}
	return 0
}

func (m *CreateCommentRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateCommentRequest) GetCreatorId() uint32 {
	if m != nil {
		return m.CreatorId
	}
	return 0
}

type CommentInfo struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId               uint32   `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	FatherId             uint32   `protobuf:"varint,4,opt,name=father_id,json=fatherId,proto3" json:"father_id,omitempty"`
	CreateTime           string   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreatorId            uint32   `protobuf:"varint,6,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatorName          string   `protobuf:"bytes,7,opt,name=creator_name,json=creatorName,proto3" json:"creator_name,omitempty"`
	CreatorAvatar        string   `protobuf:"bytes,8,opt,name=creator_avatar,json=creatorAvatar,proto3" json:"creator_avatar,omitempty"`
	LikeNum              uint32   `protobuf:"varint,9,opt,name=like_num,json=likeNum,proto3" json:"like_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentInfo) Reset()         { *m = CommentInfo{} }
func (m *CommentInfo) String() string { return proto.CompactTextString(m) }
func (*CommentInfo) ProtoMessage()    {}
func (*CommentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{4}
}

func (m *CommentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentInfo.Unmarshal(m, b)
}
func (m *CommentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentInfo.Marshal(b, m, deterministic)
}
func (m *CommentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentInfo.Merge(m, src)
}
func (m *CommentInfo) XXX_Size() int {
	return xxx_messageInfo_CommentInfo.Size(m)
}
func (m *CommentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CommentInfo proto.InternalMessageInfo

func (m *CommentInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CommentInfo) GetTypeId() uint32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *CommentInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CommentInfo) GetFatherId() uint32 {
	if m != nil {
		return m.FatherId
	}
	return 0
}

func (m *CommentInfo) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *CommentInfo) GetCreatorId() uint32 {
	if m != nil {
		return m.CreatorId
	}
	return 0
}

func (m *CommentInfo) GetCreatorName() string {
	if m != nil {
		return m.CreatorName
	}
	return ""
}

func (m *CommentInfo) GetCreatorAvatar() string {
	if m != nil {
		return m.CreatorAvatar
	}
	return ""
}

func (m *CommentInfo) GetLikeNum() uint32 {
	if m != nil {
		return m.LikeNum
	}
	return 0
}

type Request struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{5}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Item struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId               uint32   `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{6}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetTypeId() uint32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type UpdatePostInfoRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	UserId               uint32   `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Category             string   `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostInfoRequest) Reset()         { *m = UpdatePostInfoRequest{} }
func (m *UpdatePostInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePostInfoRequest) ProtoMessage()    {}
func (*UpdatePostInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{7}
}

func (m *UpdatePostInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostInfoRequest.Unmarshal(m, b)
}
func (m *UpdatePostInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostInfoRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePostInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostInfoRequest.Merge(m, src)
}
func (m *UpdatePostInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePostInfoRequest.Size(m)
}
func (m *UpdatePostInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostInfoRequest proto.InternalMessageInfo

func (m *UpdatePostInfoRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePostInfoRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UpdatePostInfoRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdatePostInfoRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdatePostInfoRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type CreatePostRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	TypeId               uint32   `protobuf:"varint,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Category             string   `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{8}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreatePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreatePostRequest) GetTypeId() uint32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *CreatePostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePostRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{9}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type ListPostRequest struct {
	LastId               uint32   `protobuf:"varint,1,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	TypeId               string   `protobuf:"bytes,4,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Category             string   `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostRequest) Reset()         { *m = ListPostRequest{} }
func (m *ListPostRequest) String() string { return proto.CompactTextString(m) }
func (*ListPostRequest) ProtoMessage()    {}
func (*ListPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{10}
}

func (m *ListPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostRequest.Unmarshal(m, b)
}
func (m *ListPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostRequest.Marshal(b, m, deterministic)
}
func (m *ListPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostRequest.Merge(m, src)
}
func (m *ListPostRequest) XXX_Size() int {
	return xxx_messageInfo_ListPostRequest.Size(m)
}
func (m *ListPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostRequest proto.InternalMessageInfo

func (m *ListPostRequest) GetLastId() uint32 {
	if m != nil {
		return m.LastId
	}
	return 0
}

func (m *ListPostRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListPostRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListPostRequest) GetTypeId() string {
	if m != nil {
		return m.TypeId
	}
	return ""
}

func (m *ListPostRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type ListPostResponse struct {
	List                 []*Post  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostResponse) Reset()         { *m = ListPostResponse{} }
func (m *ListPostResponse) String() string { return proto.CompactTextString(m) }
func (*ListPostResponse) ProtoMessage()    {}
func (*ListPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{11}
}

func (m *ListPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostResponse.Unmarshal(m, b)
}
func (m *ListPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostResponse.Marshal(b, m, deterministic)
}
func (m *ListPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostResponse.Merge(m, src)
}
func (m *ListPostResponse) XXX_Size() int {
	return xxx_messageInfo_ListPostResponse.Size(m)
}
func (m *ListPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostResponse proto.InternalMessageInfo

func (m *ListPostResponse) GetList() []*Post {
	if m != nil {
		return m.List
	}
	return nil
}

type Post struct {
	Id                   uint32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Time                 string         `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Category             string         `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	CreatorId            uint32         `protobuf:"varint,5,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatorName          string         `protobuf:"bytes,6,opt,name=creator_name,json=creatorName,proto3" json:"creator_name,omitempty"`
	Content              string         `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	CreatorAvatar        string         `protobuf:"bytes,8,opt,name=creator_avatar,json=creatorAvatar,proto3" json:"creator_avatar,omitempty"`
	CommentNum           uint32         `protobuf:"varint,9,opt,name=comment_num,json=commentNum,proto3" json:"comment_num,omitempty"`
	LikeNum              uint32         `protobuf:"varint,10,opt,name=like_num,json=likeNum,proto3" json:"like_num,omitempty"`
	Comments             []*CommentInfo `protobuf:"bytes,11,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{12}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Post) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Post) GetCreatorId() uint32 {
	if m != nil {
		return m.CreatorId
	}
	return 0
}

func (m *Post) GetCreatorName() string {
	if m != nil {
		return m.CreatorName
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetCreatorAvatar() string {
	if m != nil {
		return m.CreatorAvatar
	}
	return ""
}

func (m *Post) GetCommentNum() uint32 {
	if m != nil {
		return m.CommentNum
	}
	return 0
}

func (m *Post) GetLikeNum() uint32 {
	if m != nil {
		return m.LikeNum
	}
	return 0
}

func (m *Post) GetComments() []*CommentInfo {
	if m != nil {
		return m.Comments
	}
	return nil
}

func init() {
	proto.RegisterType((*ListLikeResponse)(nil), "post.ListLikeResponse")
	proto.RegisterType((*LikeItem)(nil), "post.LikeItem")
	proto.RegisterType((*LikeRequest)(nil), "post.LikeRequest")
	proto.RegisterType((*CreateCommentRequest)(nil), "post.CreateCommentRequest")
	proto.RegisterType((*CommentInfo)(nil), "post.CommentInfo")
	proto.RegisterType((*Request)(nil), "post.Request")
	proto.RegisterType((*Item)(nil), "post.Item")
	proto.RegisterType((*UpdatePostInfoRequest)(nil), "post.UpdatePostInfoRequest")
	proto.RegisterType((*CreatePostRequest)(nil), "post.CreatePostRequest")
	proto.RegisterType((*Response)(nil), "post.Response")
	proto.RegisterType((*ListPostRequest)(nil), "post.ListPostRequest")
	proto.RegisterType((*ListPostResponse)(nil), "post.ListPostResponse")
	proto.RegisterType((*Post)(nil), "post.Post")
}

func init() { proto.RegisterFile("proto/post.proto", fileDescriptor_5e57bba97c118b83) }

var fileDescriptor_5e57bba97c118b83 = []byte{
	// 766 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0xc4, 0xcd, 0x61, 0xfc, 0x25, 0x5f, 0xbb, 0xea, 0xc1, 0x4d, 0xf5, 0x7d, 0x94,
	0x95, 0x40, 0x15, 0x12, 0x2d, 0x2d, 0x12, 0x5c, 0x54, 0x88, 0x43, 0x91, 0x2a, 0xa3, 0xaa, 0x42,
	0x86, 0x5e, 0x57, 0x26, 0x9e, 0x14, 0xab, 0xb1, 0x1d, 0xec, 0x4d, 0xa5, 0xbe, 0x03, 0x5c, 0xc0,
	0x03, 0xf0, 0x04, 0x3c, 0x09, 0x4f, 0x85, 0xf6, 0x60, 0x67, 0xd7, 0x49, 0xd3, 0x72, 0xb7, 0x33,
	0x3b, 0xbb, 0xfb, 0xf3, 0x7f, 0x67, 0x66, 0x0d, 0xcb, 0xe3, 0x2c, 0x65, 0xe9, 0xde, 0x38, 0xcd,
	0xd9, 0xae, 0x18, 0x12, 0x9b, 0x8f, 0xe9, 0x33, 0x58, 0x3e, 0x89, 0x72, 0x76, 0x12, 0x5d, 0xa2,
	0x8f, 0xf9, 0x38, 0x4d, 0x72, 0x24, 0x14, 0xec, 0x51, 0x94, 0x33, 0xd7, 0xda, 0x6e, 0xec, 0x38,
	0x07, 0xbd, 0x5d, 0xb1, 0x88, 0x47, 0x78, 0x0c, 0x63, 0x5f, 0xcc, 0xd1, 0x57, 0xd0, 0x2e, 0x3c,
	0x64, 0x0b, 0x3a, 0x2c, 0xc8, 0x2e, 0x90, 0x9d, 0x47, 0xa1, 0x6b, 0x6d, 0x5b, 0x3b, 0x5d, 0xbf,
	0x2d, 0x1d, 0x5e, 0x48, 0x36, 0xa0, 0xc5, 0xae, 0xc7, 0xc8, 0xa7, 0xea, 0x62, 0xaa, 0xc9, 0x4d,
	0x2f, 0xa4, 0xef, 0xc0, 0x91, 0xa7, 0x7e, 0x99, 0x60, 0xce, 0x78, 0xdc, 0x24, 0xc7, 0x6c, 0xba,
	0x45, 0x93, 0x9b, 0x5e, 0xc8, 0x69, 0x22, 0x86, 0xb1, 0x58, 0x3d, 0x87, 0x86, 0xcf, 0xd1, 0x9f,
	0x16, 0xac, 0x1e, 0x65, 0x18, 0x30, 0x3c, 0x4a, 0xe3, 0x18, 0x13, 0xa6, 0xed, 0xca, 0xe3, 0xb5,
	0x5d, 0xb9, 0xb9, 0x00, 0x8b, 0x7f, 0xcc, 0x30, 0x60, 0x9f, 0x25, 0x49, 0x43, 0x7e, 0x8c, 0x74,
	0x78, 0x21, 0x71, 0xa1, 0x35, 0x48, 0x13, 0x86, 0x09, 0x73, 0xed, 0x6d, 0x6b, 0xa7, 0xe3, 0x17,
	0x26, 0xf9, 0x0f, 0x60, 0xc0, 0x01, 0x52, 0xb1, 0x6e, 0x49, 0xac, 0xeb, 0x28, 0x8f, 0x17, 0xd2,
	0x1f, 0x75, 0x70, 0x14, 0x9a, 0x97, 0x0c, 0x53, 0xd2, 0x83, 0x7a, 0x89, 0x54, 0x8f, 0x16, 0xe0,
	0x68, 0x27, 0x36, 0xcc, 0x13, 0x0d, 0x50, 0xbb, 0x02, 0x7a, 0x0f, 0x1c, 0x71, 0x38, 0x9e, 0xb3,
	0x28, 0x46, 0xc1, 0xd3, 0xf1, 0x25, 0x21, 0x7e, 0x8c, 0x62, 0xac, 0xf0, 0x36, 0x2b, 0xbc, 0xe4,
	0x3e, 0xfc, 0x53, 0x4c, 0x27, 0x41, 0x8c, 0x6e, 0x4b, 0x6c, 0xe0, 0x28, 0xdf, 0x69, 0x10, 0x23,
	0x79, 0x00, 0xbd, 0x22, 0x24, 0xb8, 0x0a, 0x58, 0x90, 0xb9, 0x6d, 0x11, 0xd4, 0x55, 0xde, 0xd7,
	0xc2, 0x49, 0x36, 0xa1, 0x3d, 0x8a, 0x2e, 0xf1, 0x3c, 0x99, 0xc4, 0x6e, 0x47, 0x1c, 0xd3, 0xe2,
	0xf6, 0xe9, 0x24, 0xa6, 0x9b, 0xd0, 0x2a, 0xee, 0xa9, 0xa2, 0x07, 0xdd, 0x03, 0x5b, 0xa4, 0xd6,
	0x5d, 0x75, 0xa2, 0x5f, 0x2d, 0x58, 0x3b, 0x1b, 0x87, 0x01, 0xc3, 0xf7, 0xfc, 0x82, 0x93, 0x61,
	0x7a, 0xc3, 0xd6, 0xba, 0xa2, 0x75, 0x53, 0xd1, 0x55, 0x58, 0x62, 0x11, 0x1b, 0xa1, 0x52, 0x5a,
	0x1a, 0x7a, 0x62, 0xda, 0x46, 0x62, 0xf6, 0xa1, 0x3d, 0x08, 0x18, 0x5e, 0xa4, 0xd9, 0xb5, 0x12,
	0xb8, 0xb4, 0xe9, 0x77, 0x0b, 0x56, 0x64, 0x42, 0x72, 0x9c, 0x5b, 0x73, 0xfc, 0x66, 0x26, 0xed,
	0x83, 0x1b, 0x46, 0x62, 0x94, 0xb0, 0xb6, 0x0e, 0xbb, 0x88, 0x09, 0xa0, 0x5d, 0x94, 0x38, 0xfd,
	0x66, 0xc1, 0xbf, 0xbc, 0xee, 0x2b, 0x74, 0xa3, 0xc0, 0xa8, 0x15, 0x6e, 0x7a, 0x21, 0x59, 0x87,
	0x66, 0x3a, 0x1c, 0xe6, 0xc8, 0x0a, 0xcd, 0xa5, 0xc5, 0x11, 0x46, 0x51, 0x1c, 0x31, 0x45, 0x26,
	0x0d, 0x9d, 0x58, 0xa2, 0x15, 0xc4, 0x8b, 0xd8, 0x0e, 0x64, 0x1b, 0x92, 0x38, 0xaa, 0x0d, 0xfd,
	0x6f, 0xb4, 0x21, 0x90, 0x85, 0x2f, 0x22, 0x64, 0x0b, 0xfa, 0x5d, 0x07, 0x9b, 0x9b, 0x33, 0x37,
	0x5c, 0x4a, 0x53, 0xd7, 0xa5, 0x21, 0x60, 0x8b, 0x5a, 0x90, 0x97, 0x2b, 0xc6, 0x06, 0x92, 0x6d,
	0x22, 0xdd, 0x52, 0xd1, 0x33, 0x15, 0xd2, 0x9c, 0xad, 0x10, 0xed, 0x56, 0x5b, 0xe6, 0xad, 0xde,
	0xb1, 0x76, 0x78, 0x15, 0xcb, 0xa6, 0xa1, 0x95, 0x0f, 0x28, 0xd7, 0xe9, 0x24, 0x36, 0x8a, 0x0b,
	0x8c, 0xe2, 0x22, 0x8f, 0xa1, 0xad, 0x02, 0x73, 0xd7, 0x11, 0x0a, 0xae, 0x48, 0x05, 0xb5, 0x36,
	0xe4, 0x97, 0x21, 0x07, 0xbf, 0x6c, 0x70, 0xb8, 0x98, 0x1f, 0x30, 0xbb, 0x8a, 0x06, 0x48, 0x9e,
	0x03, 0x4c, 0xf3, 0x97, 0x6c, 0xa8, 0xa5, 0xd5, 0x8c, 0xee, 0xab, 0x76, 0x5c, 0xe6, 0x55, 0x8d,
	0x1c, 0xf2, 0x87, 0x41, 0xde, 0x24, 0x59, 0x2b, 0x9a, 0xb5, 0x91, 0x68, 0xfd, 0xf5, 0xaa, 0xbb,
	0x5c, 0xfc, 0x12, 0x7a, 0x66, 0x11, 0x93, 0x2d, 0x19, 0x3b, 0xb7, 0xb4, 0xe7, 0x9c, 0xfe, 0x10,
	0x5a, 0xc7, 0x28, 0x0f, 0xef, 0x16, 0x93, 0x32, 0x56, 0xcb, 0x1f, 0x5a, 0x23, 0x4f, 0x00, 0x8e,
	0x91, 0x29, 0x29, 0xaa, 0xa1, 0xb3, 0x42, 0xd1, 0x1a, 0x79, 0x01, 0x5d, 0xe3, 0x85, 0x21, 0x7d,
	0x5d, 0x13, 0xf3, 0xd9, 0x99, 0x03, 0xf6, 0x08, 0xe0, 0x2d, 0x8e, 0x90, 0xc9, 0x17, 0x53, 0xc1,
	0xf0, 0xf1, 0x9c, 0xd8, 0xfd, 0x42, 0x7b, 0xfe, 0xca, 0x91, 0x95, 0xe9, 0x8b, 0x77, 0xf3, 0xf6,
	0xfb, 0x00, 0x3e, 0xc6, 0xe9, 0xd5, 0x5f, 0x2c, 0x39, 0x9c, 0xbe, 0xfc, 0x6f, 0xae, 0xcf, 0x64,
	0x1f, 0xaa, 0x08, 0xa1, 0x5d, 0x94, 0xfe, 0x83, 0x40, 0x6b, 0x9f, 0x9a, 0xe2, 0x1f, 0xe2, 0xe9,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x96, 0x36, 0x0c, 0x57, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PostService service

type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...client.CallOption) (*Response, error)
	ListPost(ctx context.Context, in *ListPostRequest, opts ...client.CallOption) (*ListPostResponse, error)
	UpdatePostInfo(ctx context.Context, in *UpdatePostInfoRequest, opts ...client.CallOption) (*Response, error)
	GetPost(ctx context.Context, in *Request, opts ...client.CallOption) (*Post, error)
	GetComment(ctx context.Context, in *Request, opts ...client.CallOption) (*CommentInfo, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...client.CallOption) (*Response, error)
	DeleteItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	CreateLike(ctx context.Context, in *LikeRequest, opts ...client.CallOption) (*Response, error)
	RemoveLike(ctx context.Context, in *LikeRequest, opts ...client.CallOption) (*Response, error)
	ListLikeByUserId(ctx context.Context, in *Request, opts ...client.CallOption) (*ListLikeResponse, error)
}

type postServiceClient struct {
	c           client.Client
	serviceName string
}

func NewPostServiceClient(serviceName string, c client.Client) PostServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "post"
	}
	return &postServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.CreatePost", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPost(ctx context.Context, in *ListPostRequest, opts ...client.CallOption) (*ListPostResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.ListPost", in)
	out := new(ListPostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePostInfo(ctx context.Context, in *UpdatePostInfoRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.UpdatePostInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *Request, opts ...client.CallOption) (*Post, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.GetPost", in)
	out := new(Post)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetComment(ctx context.Context, in *Request, opts ...client.CallOption) (*CommentInfo, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.GetComment", in)
	out := new(CommentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.CreateComment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeleteItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.DeleteItem", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreateLike(ctx context.Context, in *LikeRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.CreateLike", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) RemoveLike(ctx context.Context, in *LikeRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.RemoveLike", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListLikeByUserId(ctx context.Context, in *Request, opts ...client.CallOption) (*ListLikeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PostService.ListLikeByUserId", in)
	out := new(ListLikeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostService service

type PostServiceHandler interface {
	CreatePost(context.Context, *CreatePostRequest, *Response) error
	ListPost(context.Context, *ListPostRequest, *ListPostResponse) error
	UpdatePostInfo(context.Context, *UpdatePostInfoRequest, *Response) error
	GetPost(context.Context, *Request, *Post) error
	GetComment(context.Context, *Request, *CommentInfo) error
	CreateComment(context.Context, *CreateCommentRequest, *Response) error
	DeleteItem(context.Context, *Item, *Response) error
	CreateLike(context.Context, *LikeRequest, *Response) error
	RemoveLike(context.Context, *LikeRequest, *Response) error
	ListLikeByUserId(context.Context, *Request, *ListLikeResponse) error
}

func RegisterPostServiceHandler(s server.Server, hdlr PostServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PostService{hdlr}, opts...))
}

type PostService struct {
	PostServiceHandler
}

func (h *PostService) CreatePost(ctx context.Context, in *CreatePostRequest, out *Response) error {
	return h.PostServiceHandler.CreatePost(ctx, in, out)
}

func (h *PostService) ListPost(ctx context.Context, in *ListPostRequest, out *ListPostResponse) error {
	return h.PostServiceHandler.ListPost(ctx, in, out)
}

func (h *PostService) UpdatePostInfo(ctx context.Context, in *UpdatePostInfoRequest, out *Response) error {
	return h.PostServiceHandler.UpdatePostInfo(ctx, in, out)
}

func (h *PostService) GetPost(ctx context.Context, in *Request, out *Post) error {
	return h.PostServiceHandler.GetPost(ctx, in, out)
}

func (h *PostService) GetComment(ctx context.Context, in *Request, out *CommentInfo) error {
	return h.PostServiceHandler.GetComment(ctx, in, out)
}

func (h *PostService) CreateComment(ctx context.Context, in *CreateCommentRequest, out *Response) error {
	return h.PostServiceHandler.CreateComment(ctx, in, out)
}

func (h *PostService) DeleteItem(ctx context.Context, in *Item, out *Response) error {
	return h.PostServiceHandler.DeleteItem(ctx, in, out)
}

func (h *PostService) CreateLike(ctx context.Context, in *LikeRequest, out *Response) error {
	return h.PostServiceHandler.CreateLike(ctx, in, out)
}

func (h *PostService) RemoveLike(ctx context.Context, in *LikeRequest, out *Response) error {
	return h.PostServiceHandler.RemoveLike(ctx, in, out)
}

func (h *PostService) ListLikeByUserId(ctx context.Context, in *Request, out *ListLikeResponse) error {
	return h.PostServiceHandler.ListLikeByUserId(ctx, in, out)
}
