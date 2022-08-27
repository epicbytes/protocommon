// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParserOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fiber     bool   `protobuf:"varint,1,opt,name=fiber,proto3" json:"fiber"`
	Swag      bool   `protobuf:"varint,2,opt,name=swag,proto3" json:"swag"`
	Paging    bool   `protobuf:"varint,3,opt,name=paging,proto3" json:"paging"`
	List      bool   `protobuf:"varint,4,opt,name=list,proto3" json:"list"`
	Merge     bool   `protobuf:"varint,5,opt,name=merge,proto3" json:"merge"`
	MergeFrom string `protobuf:"bytes,6,opt,name=merge_from,json=mergeFrom,proto3" json:"merge_from"`
	Pick      bool   `protobuf:"varint,7,opt,name=pick,proto3" json:"pick"`
	PickWith  string `protobuf:"bytes,8,opt,name=pick_with,json=pickWith,proto3" json:"pick_with"`
}

func (x *ParserOption) Reset() {
	*x = ParserOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserOption) ProtoMessage() {}

func (x *ParserOption) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserOption.ProtoReflect.Descriptor instead.
func (*ParserOption) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *ParserOption) GetFiber() bool {
	if x != nil {
		return x.Fiber
	}
	return false
}

func (x *ParserOption) GetSwag() bool {
	if x != nil {
		return x.Swag
	}
	return false
}

func (x *ParserOption) GetPaging() bool {
	if x != nil {
		return x.Paging
	}
	return false
}

func (x *ParserOption) GetList() bool {
	if x != nil {
		return x.List
	}
	return false
}

func (x *ParserOption) GetMerge() bool {
	if x != nil {
		return x.Merge
	}
	return false
}

func (x *ParserOption) GetMergeFrom() string {
	if x != nil {
		return x.MergeFrom
	}
	return ""
}

func (x *ParserOption) GetPick() bool {
	if x != nil {
		return x.Pick
	}
	return false
}

func (x *ParserOption) GetPickWith() string {
	if x != nil {
		return x.PickWith
	}
	return ""
}

type ModelFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keeper    bool   `protobuf:"varint,1,opt,name=keeper,proto3" json:"keeper"`
	KeeperKey string `protobuf:"bytes,2,opt,name=keeper_key,json=keeperKey,proto3" json:"keeper_key"`
}

func (x *ModelFeature) Reset() {
	*x = ModelFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelFeature) ProtoMessage() {}

func (x *ModelFeature) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelFeature.ProtoReflect.Descriptor instead.
func (*ModelFeature) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *ModelFeature) GetKeeper() bool {
	if x != nil {
		return x.Keeper
	}
	return false
}

func (x *ModelFeature) GetKeeperKey() string {
	if x != nil {
		return x.KeeperKey
	}
	return ""
}

type ModelFieldOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Merged *bool `protobuf:"varint,1,opt,name=merged,proto3,oneof" json:"merged"`
	Picked *bool `protobuf:"varint,2,opt,name=picked,proto3,oneof" json:"picked"`
}

func (x *ModelFieldOption) Reset() {
	*x = ModelFieldOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelFieldOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelFieldOption) ProtoMessage() {}

func (x *ModelFieldOption) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelFieldOption.ProtoReflect.Descriptor instead.
func (*ModelFieldOption) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *ModelFieldOption) GetMerged() bool {
	if x != nil && x.Merged != nil {
		return *x.Merged
	}
	return false
}

func (x *ModelFieldOption) GetPicked() bool {
	if x != nil && x.Picked != nil {
		return *x.Picked
	}
	return false
}

// swagger:model AvailableProvider
type AvailableProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required: true
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label"`
	// required: true
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	// required: true
	ProviderType string `protobuf:"bytes,3,opt,name=provider_type,json=providerType,proto3" json:"provider_type"`
}

func (x *AvailableProvider) Reset() {
	*x = AvailableProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableProvider) ProtoMessage() {}

func (x *AvailableProvider) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableProvider.ProtoReflect.Descriptor instead.
func (*AvailableProvider) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *AvailableProvider) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AvailableProvider) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AvailableProvider) GetProviderType() string {
	if x != nil {
		return x.ProviderType
	}
	return ""
}

// swagger:model Pagination
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit"`
	Skip       int64 `protobuf:"varint,2,opt,name=skip,proto3" json:"skip"`
	TotalItems int64 `protobuf:"varint,3,opt,name=total_items,json=totalItems,proto3" json:"total_items"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
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
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *Pagination) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *Pagination) GetTotalItems() int64 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

// swagger:model CommentedResponse
type CommentedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required: true
	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result"`
	// required: true
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment"`
}

func (x *CommentedResponse) Reset() {
	*x = CommentedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentedResponse) ProtoMessage() {}

func (x *CommentedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentedResponse.ProtoReflect.Descriptor instead.
func (*CommentedResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *CommentedResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CommentedResponse) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

// swagger:model FileResponse
type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required: true
	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file"`
	// required: true
	MimeType string `protobuf:"bytes,2,opt,name=mimeType,proto3" json:"mimeType"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *FileResponse) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FileResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

// swagger:parameters FileRequest
type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required: true
	// In: formData
	// swagger:file
	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file"` // @gotags: validate:"required"
	// In: formData
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path"` // @gotags: validate:"required"
	// In: formData
	FileName string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *FileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var file_common_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ParserOption)(nil),
		Field:         50000,
		Name:          "common.parser",
		Tag:           "bytes,50000,opt,name=parser",
		Filename:      "common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ModelFeature)(nil),
		Field:         50001,
		Name:          "common.model_feature",
		Tag:           "bytes,50001,opt,name=model_feature",
		Filename:      "common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*ModelFieldOption)(nil),
		Field:         50002,
		Name:          "common.field_option",
		Tag:           "bytes,50002,opt,name=field_option",
		Filename:      "common.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional common.ParserOption parser = 50000;
	E_Parser = &file_common_proto_extTypes[0]
	// optional common.ModelFeature model_feature = 50001;
	E_ModelFeature = &file_common_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional common.ModelFieldOption field_option = 50002;
	E_FieldOption = &file_common_proto_extTypes[2]
)

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x62, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x77, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73,
	0x77, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x70, 0x69, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x69, 0x63, 0x6b,
	0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69, 0x63,
	0x6b, 0x57, 0x69, 0x74, 0x68, 0x22, 0x45, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x22, 0x62, 0x0a, 0x10,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x06, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x70, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52,
	0x06, 0x70, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x65, 0x64,
	0x22, 0x64, 0x0a, 0x11, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x45, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x52, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x3a, 0x5f, 0x0a,
	0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0c, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x5f,
	0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x70,
	0x69, 0x63, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_proto_goTypes = []interface{}{
	(*ParserOption)(nil),                // 0: common.ParserOption
	(*ModelFeature)(nil),                // 1: common.ModelFeature
	(*ModelFieldOption)(nil),            // 2: common.ModelFieldOption
	(*AvailableProvider)(nil),           // 3: common.AvailableProvider
	(*Pagination)(nil),                  // 4: common.Pagination
	(*CommentedResponse)(nil),           // 5: common.CommentedResponse
	(*FileResponse)(nil),                // 6: common.FileResponse
	(*FileRequest)(nil),                 // 7: common.FileRequest
	(*descriptorpb.MessageOptions)(nil), // 8: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 9: google.protobuf.FieldOptions
}
var file_common_proto_depIdxs = []int32{
	8, // 0: common.parser:extendee -> google.protobuf.MessageOptions
	8, // 1: common.model_feature:extendee -> google.protobuf.MessageOptions
	9, // 2: common.field_option:extendee -> google.protobuf.FieldOptions
	0, // 3: common.parser:type_name -> common.ParserOption
	1, // 4: common.model_feature:type_name -> common.ModelFeature
	2, // 5: common.field_option:type_name -> common.ModelFieldOption
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserOption); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelFeature); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelFieldOption); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableProvider); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentedResponse); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
	file_common_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
		ExtensionInfos:    file_common_proto_extTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
