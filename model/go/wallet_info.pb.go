// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: wallet_info.proto

package _go

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

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletID            string  `protobuf:"bytes,1,opt,name=WalletID,proto3" json:"WalletID,omitempty"`
	LastDepositAmount   float64 `protobuf:"fixed64,2,opt,name=LastDepositAmount,proto3" json:"LastDepositAmount,omitempty"`
	TwoMinuteCumulative float64 `protobuf:"fixed64,3,opt,name=TwoMinuteCumulative,proto3" json:"TwoMinuteCumulative,omitempty"`
	AboveThreshold      bool    `protobuf:"varint,4,opt,name=AboveThreshold,proto3" json:"AboveThreshold,omitempty"`
	UpdatedAt           int64   `protobuf:"varint,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_wallet_info_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *Wallet) GetLastDepositAmount() float64 {
	if x != nil {
		return x.LastDepositAmount
	}
	return 0
}

func (x *Wallet) GetTwoMinuteCumulative() float64 {
	if x != nil {
		return x.TwoMinuteCumulative
	}
	return 0
}

func (x *Wallet) GetAboveThreshold() bool {
	if x != nil {
		return x.AboveThreshold
	}
	return false
}

func (x *Wallet) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_wallet_info_proto protoreflect.FileDescriptor

var file_wallet_info_proto_rawDesc = []byte{
	0x0a, 0x11, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xca, 0x01, 0x0a, 0x06, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x4c, 0x61, 0x73,
	0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x54, 0x77, 0x6f, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x43, 0x75, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x54, 0x77, 0x6f,
	0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x41, 0x62, 0x6f, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x41, 0x62, 0x6f, 0x76, 0x65, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x67, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_info_proto_rawDescOnce sync.Once
	file_wallet_info_proto_rawDescData = file_wallet_info_proto_rawDesc
)

func file_wallet_info_proto_rawDescGZIP() []byte {
	file_wallet_info_proto_rawDescOnce.Do(func() {
		file_wallet_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_info_proto_rawDescData)
	})
	return file_wallet_info_proto_rawDescData
}

var file_wallet_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wallet_info_proto_goTypes = []interface{}{
	(*Wallet)(nil), // 0: main.Wallet
}
var file_wallet_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wallet_info_proto_init() }
func file_wallet_info_proto_init() {
	if File_wallet_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
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
			RawDescriptor: file_wallet_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wallet_info_proto_goTypes,
		DependencyIndexes: file_wallet_info_proto_depIdxs,
		MessageInfos:      file_wallet_info_proto_msgTypes,
	}.Build()
	File_wallet_info_proto = out.File
	file_wallet_info_proto_rawDesc = nil
	file_wallet_info_proto_goTypes = nil
	file_wallet_info_proto_depIdxs = nil
}
