// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: proto/admin.proto

package admin

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

type DeleteDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *DeleteDataRequest) Reset() {
	*x = DeleteDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDataRequest) ProtoMessage() {}

func (x *DeleteDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDataRequest.ProtoReflect.Descriptor instead.
func (*DeleteDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_admin_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteDataRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type DeleteDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDataResponse) Reset() {
	*x = DeleteDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDataResponse) ProtoMessage() {}

func (x *DeleteDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDataResponse.ProtoReflect.Descriptor instead.
func (*DeleteDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_admin_proto_rawDescGZIP(), []int{1}
}

type UsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *UsageRequest) Reset() {
	*x = UsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageRequest) ProtoMessage() {}

func (x *UsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageRequest.ProtoReflect.Descriptor instead.
func (*UsageRequest) Descriptor() ([]byte, []int) {
	return file_proto_admin_proto_rawDescGZIP(), []int{2}
}

func (x *UsageRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type UsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage map[string]*Usage `protobuf:"bytes,1,rep,name=usage,proto3" json:"usage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map of endpoints to usage
}

func (x *UsageResponse) Reset() {
	*x = UsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageResponse) ProtoMessage() {}

func (x *UsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageResponse.ProtoReflect.Descriptor instead.
func (*UsageResponse) Descriptor() ([]byte, []int) {
	return file_proto_admin_proto_rawDescGZIP(), []int{3}
}

func (x *UsageResponse) GetUsage() map[string]*Usage {
	if x != nil {
		return x.Usage
	}
	return nil
}

type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage int64  `protobuf:"varint,1,opt,name=usage,proto3" json:"usage,omitempty"` // A number that represents the usage
	Units string `protobuf:"bytes,2,opt,name=units,proto3" json:"units,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_proto_admin_proto_rawDescGZIP(), []int{4}
}

func (x *Usage) GetUsage() int64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *Usage) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

var File_proto_admin_proto protoreflect.FileDescriptor

var file_proto_admin_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x30, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2b, 0x0a, 0x0c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x8e, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x46, 0x0a, 0x0a, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x33, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x32, 0x82, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_admin_proto_rawDescOnce sync.Once
	file_proto_admin_proto_rawDescData = file_proto_admin_proto_rawDesc
)

func file_proto_admin_proto_rawDescGZIP() []byte {
	file_proto_admin_proto_rawDescOnce.Do(func() {
		file_proto_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_admin_proto_rawDescData)
	})
	return file_proto_admin_proto_rawDescData
}

var file_proto_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_admin_proto_goTypes = []interface{}{
	(*DeleteDataRequest)(nil),  // 0: admin.DeleteDataRequest
	(*DeleteDataResponse)(nil), // 1: admin.DeleteDataResponse
	(*UsageRequest)(nil),       // 2: admin.UsageRequest
	(*UsageResponse)(nil),      // 3: admin.UsageResponse
	(*Usage)(nil),              // 4: admin.Usage
	nil,                        // 5: admin.UsageResponse.UsageEntry
}
var file_proto_admin_proto_depIdxs = []int32{
	5, // 0: admin.UsageResponse.usage:type_name -> admin.UsageResponse.UsageEntry
	4, // 1: admin.UsageResponse.UsageEntry.value:type_name -> admin.Usage
	0, // 2: admin.Admin.DeleteData:input_type -> admin.DeleteDataRequest
	2, // 3: admin.Admin.Usage:input_type -> admin.UsageRequest
	1, // 4: admin.Admin.DeleteData:output_type -> admin.DeleteDataResponse
	3, // 5: admin.Admin.Usage:output_type -> admin.UsageResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_admin_proto_init() }
func file_proto_admin_proto_init() {
	if File_proto_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDataRequest); i {
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
		file_proto_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDataResponse); i {
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
		file_proto_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageRequest); i {
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
		file_proto_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageResponse); i {
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
		file_proto_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
			RawDescriptor: file_proto_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_admin_proto_goTypes,
		DependencyIndexes: file_proto_admin_proto_depIdxs,
		MessageInfos:      file_proto_admin_proto_msgTypes,
	}.Build()
	File_proto_admin_proto = out.File
	file_proto_admin_proto_rawDesc = nil
	file_proto_admin_proto_goTypes = nil
	file_proto_admin_proto_depIdxs = nil
}