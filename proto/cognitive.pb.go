// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: cognitive.proto

package cognitive

import (
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

type LangType int32

const (
	LangType_ENGLISH  LangType = 0
	LangType_RUSSIA   LangType = 1
	LangType_FRENCH   LangType = 2
	LangType_SPANISH  LangType = 3
	LangType_PORTUGAL LangType = 4
	LangType_JAPANESE LangType = 5
	LangType_KOREAN   LangType = 6
	LangType_ARABIC   LangType = 7
)

// Enum value maps for LangType.
var (
	LangType_name = map[int32]string{
		0: "ENGLISH",
		1: "RUSSIA",
		2: "FRENCH",
		3: "SPANISH",
		4: "PORTUGAL",
		5: "JAPANESE",
		6: "KOREAN",
		7: "ARABIC",
	}
	LangType_value = map[string]int32{
		"ENGLISH":  0,
		"RUSSIA":   1,
		"FRENCH":   2,
		"SPANISH":  3,
		"PORTUGAL": 4,
		"JAPANESE": 5,
		"KOREAN":   6,
		"ARABIC":   7,
	}
)

func (x LangType) Enum() *LangType {
	p := new(LangType)
	*p = x
	return p
}

func (x LangType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LangType) Descriptor() protoreflect.EnumDescriptor {
	return file_cognitive_proto_enumTypes[0].Descriptor()
}

func (LangType) Type() protoreflect.EnumType {
	return &file_cognitive_proto_enumTypes[0]
}

func (x LangType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LangType.Descriptor instead.
func (LangType) EnumDescriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{0}
}

type ReqHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReqHead) Reset() {
	*x = ReqHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqHead) ProtoMessage() {}

func (x *ReqHead) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqHead.ProtoReflect.Descriptor instead.
func (*ReqHead) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{0}
}

func (x *ReqHead) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RspHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RspHead) Reset() {
	*x = RspHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspHead) ProtoMessage() {}

func (x *RspHead) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspHead.ProtoReflect.Descriptor instead.
func (*RspHead) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{1}
}

func (x *RspHead) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RspHead) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head  *ReqHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Title string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text  string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Lang  int32    `protobuf:"varint,4,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{2}
}

func (x *ChatRequest) GetHead() *ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *ChatRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChatRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ChatRequest) GetLang() int32 {
	if x != nil {
		return x.Lang
	}
	return 0
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head      *RspHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	VideoItem *Video   `protobuf:"bytes,2,opt,name=video_item,json=videoItem,proto3" json:"video_item,omitempty"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{3}
}

func (x *ChatResponse) GetHead() *RspHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *ChatResponse) GetVideoItem() *Video {
	if x != nil {
		return x.VideoItem
	}
	return nil
}

type FetchVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head    *ReqHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Session string   `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *FetchVideoRequest) Reset() {
	*x = FetchVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchVideoRequest) ProtoMessage() {}

func (x *FetchVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchVideoRequest.ProtoReflect.Descriptor instead.
func (*FetchVideoRequest) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{4}
}

func (x *FetchVideoRequest) GetHead() *ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *FetchVideoRequest) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src    string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{5}
}

func (x *Video) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type FetchVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head      *RspHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	VideoList []*Video `protobuf:"bytes,2,rep,name=videoList,proto3" json:"videoList,omitempty"`
}

func (x *FetchVideoResponse) Reset() {
	*x = FetchVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchVideoResponse) ProtoMessage() {}

func (x *FetchVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchVideoResponse.ProtoReflect.Descriptor instead.
func (*FetchVideoResponse) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{6}
}

func (x *FetchVideoResponse) GetHead() *RspHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *FetchVideoResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type ResetVoiceGenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *ReqHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *ResetVoiceGenRequest) Reset() {
	*x = ResetVoiceGenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetVoiceGenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetVoiceGenRequest) ProtoMessage() {}

func (x *ResetVoiceGenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetVoiceGenRequest.ProtoReflect.Descriptor instead.
func (*ResetVoiceGenRequest) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{7}
}

func (x *ResetVoiceGenRequest) GetHead() *ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

type ResetVoiceGenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *RspHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *ResetVoiceGenResponse) Reset() {
	*x = ResetVoiceGenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cognitive_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetVoiceGenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetVoiceGenResponse) ProtoMessage() {}

func (x *ResetVoiceGenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cognitive_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetVoiceGenResponse.ProtoReflect.Descriptor instead.
func (*ResetVoiceGenResponse) Descriptor() ([]byte, []int) {
	return file_cognitive_proto_rawDescGZIP(), []int{8}
}

func (x *ResetVoiceGenResponse) GetHead() *RspHead {
	if x != nil {
		return x.Head
	}
	return nil
}

var File_cognitive_proto protoreflect.FileDescriptor

var file_cognitive_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x22, 0x19, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x07, 0x52, 0x73, 0x70, 0x48, 0x65,
	0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x73, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x67, 0x0a,
	0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x52,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x55, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x68,
	0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68,
	0x65, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a,
	0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x22, 0x3f, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x52,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x2a, 0x70, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x52, 0x55, 0x53, 0x53, 0x49, 0x41, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52,
	0x45, 0x4e, 0x43, 0x48, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x41, 0x4e, 0x49, 0x53,
	0x48, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4f, 0x52, 0x54, 0x55, 0x47, 0x41, 0x4c, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x41, 0x50, 0x41, 0x4e, 0x45, 0x53, 0x45, 0x10, 0x05, 0x12,
	0x0a, 0x0a, 0x06, 0x4b, 0x4f, 0x52, 0x45, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x52, 0x41, 0x42, 0x49, 0x43, 0x10, 0x07, 0x32, 0xfc, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74,
	0x32, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x12, 0x46, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x32,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x54, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x6e, 0x79, 0x68, 0x61, 0x6e, 0x6b, 0x6b, 0x2f,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cognitive_proto_rawDescOnce sync.Once
	file_cognitive_proto_rawDescData = file_cognitive_proto_rawDesc
)

func file_cognitive_proto_rawDescGZIP() []byte {
	file_cognitive_proto_rawDescOnce.Do(func() {
		file_cognitive_proto_rawDescData = protoimpl.X.CompressGZIP(file_cognitive_proto_rawDescData)
	})
	return file_cognitive_proto_rawDescData
}

var file_cognitive_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cognitive_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cognitive_proto_goTypes = []interface{}{
	(LangType)(0),                 // 0: cognitive.LangType
	(*ReqHead)(nil),               // 1: cognitive.ReqHead
	(*RspHead)(nil),               // 2: cognitive.RspHead
	(*ChatRequest)(nil),           // 3: cognitive.ChatRequest
	(*ChatResponse)(nil),          // 4: cognitive.ChatResponse
	(*FetchVideoRequest)(nil),     // 5: cognitive.FetchVideoRequest
	(*Video)(nil),                 // 6: cognitive.Video
	(*FetchVideoResponse)(nil),    // 7: cognitive.FetchVideoResponse
	(*ResetVoiceGenRequest)(nil),  // 8: cognitive.ResetVoiceGenRequest
	(*ResetVoiceGenResponse)(nil), // 9: cognitive.ResetVoiceGenResponse
}
var file_cognitive_proto_depIdxs = []int32{
	1,  // 0: cognitive.ChatRequest.head:type_name -> cognitive.ReqHead
	2,  // 1: cognitive.ChatResponse.head:type_name -> cognitive.RspHead
	6,  // 2: cognitive.ChatResponse.video_item:type_name -> cognitive.Video
	1,  // 3: cognitive.FetchVideoRequest.head:type_name -> cognitive.ReqHead
	2,  // 4: cognitive.FetchVideoResponse.head:type_name -> cognitive.RspHead
	6,  // 5: cognitive.FetchVideoResponse.videoList:type_name -> cognitive.Video
	1,  // 6: cognitive.ResetVoiceGenRequest.head:type_name -> cognitive.ReqHead
	2,  // 7: cognitive.ResetVoiceGenResponse.head:type_name -> cognitive.RspHead
	3,  // 8: cognitive.Text2Speech.Chat2VoiceRequest:input_type -> cognitive.ChatRequest
	5,  // 9: cognitive.Text2Speech.FetchVideoList:input_type -> cognitive.FetchVideoRequest
	8,  // 10: cognitive.Text2Speech.ResetGenVoice:input_type -> cognitive.ResetVoiceGenRequest
	4,  // 11: cognitive.Text2Speech.Chat2VoiceRequest:output_type -> cognitive.ChatResponse
	7,  // 12: cognitive.Text2Speech.FetchVideoList:output_type -> cognitive.FetchVideoResponse
	9,  // 13: cognitive.Text2Speech.ResetGenVoice:output_type -> cognitive.ResetVoiceGenResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cognitive_proto_init() }
func file_cognitive_proto_init() {
	if File_cognitive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cognitive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqHead); i {
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
		file_cognitive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspHead); i {
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
		file_cognitive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_cognitive_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
		file_cognitive_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchVideoRequest); i {
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
		file_cognitive_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_cognitive_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchVideoResponse); i {
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
		file_cognitive_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetVoiceGenRequest); i {
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
		file_cognitive_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetVoiceGenResponse); i {
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
			RawDescriptor: file_cognitive_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cognitive_proto_goTypes,
		DependencyIndexes: file_cognitive_proto_depIdxs,
		EnumInfos:         file_cognitive_proto_enumTypes,
		MessageInfos:      file_cognitive_proto_msgTypes,
	}.Build()
	File_cognitive_proto = out.File
	file_cognitive_proto_rawDesc = nil
	file_cognitive_proto_goTypes = nil
	file_cognitive_proto_depIdxs = nil
}
