// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: einride/example/freight/v1/shipper.proto

package freightv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A shipper is a supplier or owner of goods to be transported.
type Shipper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the shipper.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The creation timestamp of the shipper.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the shipper.
	//
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the shipper.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The display name of the shipper.
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Shipper) Reset() {
	*x = Shipper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_example_freight_v1_shipper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipper) ProtoMessage() {}

func (x *Shipper) ProtoReflect() protoreflect.Message {
	mi := &file_einride_example_freight_v1_shipper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipper.ProtoReflect.Descriptor instead.
func (*Shipper) Descriptor() ([]byte, []int) {
	return file_einride_example_freight_v1_shipper_proto_rawDescGZIP(), []int{0}
}

func (x *Shipper) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shipper) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Shipper) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Shipper) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Shipper) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_einride_example_freight_v1_shipper_proto protoreflect.FileDescriptor

var file_einride_example_freight_v1_shipper_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x66, 0x72, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x07, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x50, 0xea, 0x41, 0x4d, 0x0a, 0x24, 0x66, 0x72, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x12, 0x12, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x7d, 0x2a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x73, 0x32, 0x07,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x42, 0xfd, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x61, 0x69, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x45, 0x45, 0x46, 0xaa, 0x02, 0x1a, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1a, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5c, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x26, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5c, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x46, 0x72, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_example_freight_v1_shipper_proto_rawDescOnce sync.Once
	file_einride_example_freight_v1_shipper_proto_rawDescData = file_einride_example_freight_v1_shipper_proto_rawDesc
)

func file_einride_example_freight_v1_shipper_proto_rawDescGZIP() []byte {
	file_einride_example_freight_v1_shipper_proto_rawDescOnce.Do(func() {
		file_einride_example_freight_v1_shipper_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_example_freight_v1_shipper_proto_rawDescData)
	})
	return file_einride_example_freight_v1_shipper_proto_rawDescData
}

var file_einride_example_freight_v1_shipper_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_example_freight_v1_shipper_proto_goTypes = []interface{}{
	(*Shipper)(nil),               // 0: einride.example.freight.v1.Shipper
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_einride_example_freight_v1_shipper_proto_depIdxs = []int32{
	1, // 0: einride.example.freight.v1.Shipper.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: einride.example.freight.v1.Shipper.update_time:type_name -> google.protobuf.Timestamp
	1, // 2: einride.example.freight.v1.Shipper.delete_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_einride_example_freight_v1_shipper_proto_init() }
func file_einride_example_freight_v1_shipper_proto_init() {
	if File_einride_example_freight_v1_shipper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_example_freight_v1_shipper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shipper); i {
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
			RawDescriptor: file_einride_example_freight_v1_shipper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_example_freight_v1_shipper_proto_goTypes,
		DependencyIndexes: file_einride_example_freight_v1_shipper_proto_depIdxs,
		MessageInfos:      file_einride_example_freight_v1_shipper_proto_msgTypes,
	}.Build()
	File_einride_example_freight_v1_shipper_proto = out.File
	file_einride_example_freight_v1_shipper_proto_rawDesc = nil
	file_einride_example_freight_v1_shipper_proto_goTypes = nil
	file_einride_example_freight_v1_shipper_proto_depIdxs = nil
}
