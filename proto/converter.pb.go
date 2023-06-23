// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/converter.proto

package proto

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type StorageRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetImageResponse   *GetImageResponse   `protobuf:"bytes,1,opt,name=getImageResponse,proto3" json:"getImageResponse,omitempty"`
	StoreImageResponse *StoreImageResponse `protobuf:"bytes,2,opt,name=storeImageResponse,proto3" json:"storeImageResponse,omitempty"`
}

func (x *StorageRespone) Reset() {
	*x = StorageRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageRespone) ProtoMessage() {}

func (x *StorageRespone) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageRespone.ProtoReflect.Descriptor instead.
func (*StorageRespone) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{1}
}

func (x *StorageRespone) GetGetImageResponse() *GetImageResponse {
	if x != nil {
		return x.GetImageResponse
	}
	return nil
}

func (x *StorageRespone) GetStoreImageResponse() *StoreImageResponse {
	if x != nil {
		return x.StoreImageResponse
	}
	return nil
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ReplyTo string `protobuf:"bytes,3,opt,name=replyTo,proto3" json:"replyTo,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{2}
}

func (x *GetImageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetImageRequest) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

type GetImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image  *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetImageResponse) Reset() {
	*x = GetImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageResponse) ProtoMessage() {}

func (x *GetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageResponse.ProtoReflect.Descriptor instead.
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{3}
}

func (x *GetImageResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *GetImageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type StoreImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Image   *Image `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	ReplyTo string `protobuf:"bytes,3,opt,name=replyTo,proto3" json:"replyTo,omitempty"`
}

func (x *StoreImageRequest) Reset() {
	*x = StoreImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreImageRequest) ProtoMessage() {}

func (x *StoreImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreImageRequest.ProtoReflect.Descriptor instead.
func (*StoreImageRequest) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{4}
}

func (x *StoreImageRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *StoreImageRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *StoreImageRequest) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

type StoreImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StoreImageResponse) Reset() {
	*x = StoreImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreImageResponse) ProtoMessage() {}

func (x *StoreImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreImageResponse.ProtoReflect.Descriptor instead.
func (*StoreImageResponse) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{5}
}

func (x *StoreImageResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *StoreImageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_converter_proto protoreflect.FileDescriptor

var file_proto_converter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x47, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x67, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x12, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x22,
	0x5a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6f, 0x0a, 0x11, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x22, 0x3e, 0x0a, 0x12,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_converter_proto_rawDescOnce sync.Once
	file_proto_converter_proto_rawDescData = file_proto_converter_proto_rawDesc
)

func file_proto_converter_proto_rawDescGZIP() []byte {
	file_proto_converter_proto_rawDescOnce.Do(func() {
		file_proto_converter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_converter_proto_rawDescData)
	})
	return file_proto_converter_proto_rawDescData
}

var file_proto_converter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_converter_proto_goTypes = []interface{}{
	(*Image)(nil),              // 0: converter_service.Image
	(*StorageRespone)(nil),     // 1: converter_service.StorageRespone
	(*GetImageRequest)(nil),    // 2: converter_service.GetImageRequest
	(*GetImageResponse)(nil),   // 3: converter_service.GetImageResponse
	(*StoreImageRequest)(nil),  // 4: converter_service.StoreImageRequest
	(*StoreImageResponse)(nil), // 5: converter_service.StoreImageResponse
}
var file_proto_converter_proto_depIdxs = []int32{
	3, // 0: converter_service.StorageRespone.getImageResponse:type_name -> converter_service.GetImageResponse
	5, // 1: converter_service.StorageRespone.storeImageResponse:type_name -> converter_service.StoreImageResponse
	0, // 2: converter_service.GetImageResponse.image:type_name -> converter_service.Image
	0, // 3: converter_service.StoreImageRequest.image:type_name -> converter_service.Image
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_converter_proto_init() }
func file_proto_converter_proto_init() {
	if File_proto_converter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_converter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_proto_converter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageRespone); i {
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
		file_proto_converter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageRequest); i {
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
		file_proto_converter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageResponse); i {
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
		file_proto_converter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreImageRequest); i {
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
		file_proto_converter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreImageResponse); i {
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
			RawDescriptor: file_proto_converter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_converter_proto_goTypes,
		DependencyIndexes: file_proto_converter_proto_depIdxs,
		MessageInfos:      file_proto_converter_proto_msgTypes,
	}.Build()
	File_proto_converter_proto = out.File
	file_proto_converter_proto_rawDesc = nil
	file_proto_converter_proto_goTypes = nil
	file_proto_converter_proto_depIdxs = nil
}
