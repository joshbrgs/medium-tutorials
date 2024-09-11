// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: nemesis.proto

package gen

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateNemesisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Power string `protobuf:"bytes,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *CreateNemesisRequest) Reset() {
	*x = CreateNemesisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nemesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNemesisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNemesisRequest) ProtoMessage() {}

func (x *CreateNemesisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nemesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNemesisRequest.ProtoReflect.Descriptor instead.
func (*CreateNemesisRequest) Descriptor() ([]byte, []int) {
	return file_nemesis_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNemesisRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNemesisRequest) GetPower() string {
	if x != nil {
		return x.Power
	}
	return ""
}

type GetNemesisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNemesisRequest) Reset() {
	*x = GetNemesisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nemesis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNemesisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNemesisRequest) ProtoMessage() {}

func (x *GetNemesisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nemesis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNemesisRequest.ProtoReflect.Descriptor instead.
func (*GetNemesisRequest) Descriptor() ([]byte, []int) {
	return file_nemesis_proto_rawDescGZIP(), []int{1}
}

func (x *GetNemesisRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateNemesisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Power string `protobuf:"bytes,3,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *UpdateNemesisRequest) Reset() {
	*x = UpdateNemesisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nemesis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNemesisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNemesisRequest) ProtoMessage() {}

func (x *UpdateNemesisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nemesis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNemesisRequest.ProtoReflect.Descriptor instead.
func (*UpdateNemesisRequest) Descriptor() ([]byte, []int) {
	return file_nemesis_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNemesisRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateNemesisRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNemesisRequest) GetPower() string {
	if x != nil {
		return x.Power
	}
	return ""
}

type DeleteNemesisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNemesisRequest) Reset() {
	*x = DeleteNemesisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nemesis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNemesisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNemesisRequest) ProtoMessage() {}

func (x *DeleteNemesisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nemesis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNemesisRequest.ProtoReflect.Descriptor instead.
func (*DeleteNemesisRequest) Descriptor() ([]byte, []int) {
	return file_nemesis_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNemesisRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NemesisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Power string `protobuf:"bytes,3,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *NemesisResponse) Reset() {
	*x = NemesisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nemesis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NemesisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NemesisResponse) ProtoMessage() {}

func (x *NemesisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nemesis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NemesisResponse.ProtoReflect.Descriptor instead.
func (*NemesisResponse) Descriptor() ([]byte, []int) {
	return file_nemesis_proto_rawDescGZIP(), []int{4}
}

func (x *NemesisResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NemesisResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NemesisResponse) GetPower() string {
	if x != nil {
		return x.Power
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nemesis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_nemesis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_nemesis_proto_rawDescGZIP(), []int{5}
}

type ListNemesisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nemeses []*NemesisResponse `protobuf:"bytes,1,rep,name=nemeses,proto3" json:"nemeses,omitempty"`
}

func (x *ListNemesisResponse) Reset() {
	*x = ListNemesisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nemesis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNemesisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNemesisResponse) ProtoMessage() {}

func (x *ListNemesisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nemesis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNemesisResponse.ProtoReflect.Descriptor instead.
func (*ListNemesisResponse) Descriptor() ([]byte, []int) {
	return file_nemesis_proto_rawDescGZIP(), []int{6}
}

func (x *ListNemesisResponse) GetNemeses() []*NemesisResponse {
	if x != nil {
		return x.Nemeses
	}
	return nil
}

var File_nemesis_proto protoreflect.FileDescriptor

var file_nemesis_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x22,
	0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x0f, 0x4e, 0x65, 0x6d, 0x65, 0x73,
	0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x49, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e,
	0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x65, 0x73, 0x32, 0xce, 0x03, 0x0a, 0x0e, 0x4e, 0x65, 0x6d,
	0x65, 0x73, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x12, 0x1d, 0x2e, 0x6e,
	0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x6d,
	0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65,
	0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x08, 0x2f,
	0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x12, 0x59, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x6d, 0x65, 0x73, 0x69, 0x73, 0x12, 0x1a, 0x2e, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x4e, 0x65, 0x6d, 0x65,
	0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x6d, 0x65,
	0x73, 0x69, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x4e, 0x65, 0x6d,
	0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x32, 0x0d, 0x2f, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x55, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x6d,
	0x65, 0x73, 0x69, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x6e, 0x65,
	0x6d, 0x65, 0x73, 0x69, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4d, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x12, 0x0e, 0x2e, 0x6e, 0x65, 0x6d, 0x65,
	0x73, 0x69, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x6e, 0x65, 0x6d, 0x65,
	0x73, 0x69, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x2d, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x62, 0x61, 0x64, 0x2d, 0x69, 0x6e, 0x63,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x6e, 0x65, 0x6d, 0x65, 0x73, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nemesis_proto_rawDescOnce sync.Once
	file_nemesis_proto_rawDescData = file_nemesis_proto_rawDesc
)

func file_nemesis_proto_rawDescGZIP() []byte {
	file_nemesis_proto_rawDescOnce.Do(func() {
		file_nemesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_nemesis_proto_rawDescData)
	})
	return file_nemesis_proto_rawDescData
}

var file_nemesis_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nemesis_proto_goTypes = []any{
	(*CreateNemesisRequest)(nil), // 0: nemesis.CreateNemesisRequest
	(*GetNemesisRequest)(nil),    // 1: nemesis.GetNemesisRequest
	(*UpdateNemesisRequest)(nil), // 2: nemesis.UpdateNemesisRequest
	(*DeleteNemesisRequest)(nil), // 3: nemesis.DeleteNemesisRequest
	(*NemesisResponse)(nil),      // 4: nemesis.NemesisResponse
	(*Empty)(nil),                // 5: nemesis.Empty
	(*ListNemesisResponse)(nil),  // 6: nemesis.ListNemesisResponse
}
var file_nemesis_proto_depIdxs = []int32{
	4, // 0: nemesis.ListNemesisResponse.nemeses:type_name -> nemesis.NemesisResponse
	0, // 1: nemesis.NemesisService.CreateNemesis:input_type -> nemesis.CreateNemesisRequest
	1, // 2: nemesis.NemesisService.GetNemesis:input_type -> nemesis.GetNemesisRequest
	2, // 3: nemesis.NemesisService.UpdateNemesis:input_type -> nemesis.UpdateNemesisRequest
	3, // 4: nemesis.NemesisService.DeleteNemesis:input_type -> nemesis.DeleteNemesisRequest
	5, // 5: nemesis.NemesisService.ListNemesis:input_type -> nemesis.Empty
	4, // 6: nemesis.NemesisService.CreateNemesis:output_type -> nemesis.NemesisResponse
	4, // 7: nemesis.NemesisService.GetNemesis:output_type -> nemesis.NemesisResponse
	4, // 8: nemesis.NemesisService.UpdateNemesis:output_type -> nemesis.NemesisResponse
	5, // 9: nemesis.NemesisService.DeleteNemesis:output_type -> nemesis.Empty
	6, // 10: nemesis.NemesisService.ListNemesis:output_type -> nemesis.ListNemesisResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nemesis_proto_init() }
func file_nemesis_proto_init() {
	if File_nemesis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nemesis_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateNemesisRequest); i {
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
		file_nemesis_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetNemesisRequest); i {
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
		file_nemesis_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateNemesisRequest); i {
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
		file_nemesis_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteNemesisRequest); i {
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
		file_nemesis_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*NemesisResponse); i {
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
		file_nemesis_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_nemesis_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListNemesisResponse); i {
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
			RawDescriptor: file_nemesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nemesis_proto_goTypes,
		DependencyIndexes: file_nemesis_proto_depIdxs,
		MessageInfos:      file_nemesis_proto_msgTypes,
	}.Build()
	File_nemesis_proto = out.File
	file_nemesis_proto_rawDesc = nil
	file_nemesis_proto_goTypes = nil
	file_nemesis_proto_depIdxs = nil
}
