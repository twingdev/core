// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pb/autonat.proto

package pb

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

type Message_MessageType int32

const (
	Message_DIAL          Message_MessageType = 0
	Message_DIAL_RESPONSE Message_MessageType = 1
)

// Enum value maps for Message_MessageType.
var (
	Message_MessageType_name = map[int32]string{
		0: "DIAL",
		1: "DIAL_RESPONSE",
	}
	Message_MessageType_value = map[string]int32{
		"DIAL":          0,
		"DIAL_RESPONSE": 1,
	}
)

func (x Message_MessageType) Enum() *Message_MessageType {
	p := new(Message_MessageType)
	*p = x
	return p
}

func (x Message_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_autonat_proto_enumTypes[0].Descriptor()
}

func (Message_MessageType) Type() protoreflect.EnumType {
	return &file_pb_autonat_proto_enumTypes[0]
}

func (x Message_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Message_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Message_MessageType(num)
	return nil
}

// Deprecated: Use Message_MessageType.Descriptor instead.
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_pb_autonat_proto_rawDescGZIP(), []int{0, 0}
}

type Message_ResponseStatus int32

const (
	Message_OK               Message_ResponseStatus = 0
	Message_E_DIAL_ERROR     Message_ResponseStatus = 100
	Message_E_DIAL_REFUSED   Message_ResponseStatus = 101
	Message_E_BAD_REQUEST    Message_ResponseStatus = 200
	Message_E_INTERNAL_ERROR Message_ResponseStatus = 300
)

// Enum value maps for Message_ResponseStatus.
var (
	Message_ResponseStatus_name = map[int32]string{
		0:   "OK",
		100: "E_DIAL_ERROR",
		101: "E_DIAL_REFUSED",
		200: "E_BAD_REQUEST",
		300: "E_INTERNAL_ERROR",
	}
	Message_ResponseStatus_value = map[string]int32{
		"OK":               0,
		"E_DIAL_ERROR":     100,
		"E_DIAL_REFUSED":   101,
		"E_BAD_REQUEST":    200,
		"E_INTERNAL_ERROR": 300,
	}
)

func (x Message_ResponseStatus) Enum() *Message_ResponseStatus {
	p := new(Message_ResponseStatus)
	*p = x
	return p
}

func (x Message_ResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_autonat_proto_enumTypes[1].Descriptor()
}

func (Message_ResponseStatus) Type() protoreflect.EnumType {
	return &file_pb_autonat_proto_enumTypes[1]
}

func (x Message_ResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Message_ResponseStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Message_ResponseStatus(num)
	return nil
}

// Deprecated: Use Message_ResponseStatus.Descriptor instead.
func (Message_ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_pb_autonat_proto_rawDescGZIP(), []int{0, 1}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         *Message_MessageType  `protobuf:"varint,1,opt,name=type,enum=autonat.pb.Message_MessageType" json:"type,omitempty"`
	Dial         *Message_Dial         `protobuf:"bytes,2,opt,name=dial" json:"dial,omitempty"`
	DialResponse *Message_DialResponse `protobuf:"bytes,3,opt,name=dialResponse" json:"dialResponse,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_autonat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pb_autonat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pb_autonat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() Message_MessageType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Message_DIAL
}

func (x *Message) GetDial() *Message_Dial {
	if x != nil {
		return x.Dial
	}
	return nil
}

func (x *Message) GetDialResponse() *Message_DialResponse {
	if x != nil {
		return x.DialResponse
	}
	return nil
}

type Message_PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    []byte   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Addrs [][]byte `protobuf:"bytes,2,rep,name=addrs" json:"addrs,omitempty"`
}

func (x *Message_PeerInfo) Reset() {
	*x = Message_PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_autonat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_PeerInfo) ProtoMessage() {}

func (x *Message_PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_autonat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_PeerInfo.ProtoReflect.Descriptor instead.
func (*Message_PeerInfo) Descriptor() ([]byte, []int) {
	return file_pb_autonat_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Message_PeerInfo) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Message_PeerInfo) GetAddrs() [][]byte {
	if x != nil {
		return x.Addrs
	}
	return nil
}

type Message_Dial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer *Message_PeerInfo `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
}

func (x *Message_Dial) Reset() {
	*x = Message_Dial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_autonat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_Dial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Dial) ProtoMessage() {}

func (x *Message_Dial) ProtoReflect() protoreflect.Message {
	mi := &file_pb_autonat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Dial.ProtoReflect.Descriptor instead.
func (*Message_Dial) Descriptor() ([]byte, []int) {
	return file_pb_autonat_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Message_Dial) GetPeer() *Message_PeerInfo {
	if x != nil {
		return x.Peer
	}
	return nil
}

type Message_DialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *Message_ResponseStatus `protobuf:"varint,1,opt,name=status,enum=autonat.pb.Message_ResponseStatus" json:"status,omitempty"`
	StatusText *string                 `protobuf:"bytes,2,opt,name=statusText" json:"statusText,omitempty"`
	Addr       []byte                  `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
}

func (x *Message_DialResponse) Reset() {
	*x = Message_DialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_autonat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_DialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_DialResponse) ProtoMessage() {}

func (x *Message_DialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_autonat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_DialResponse.ProtoReflect.Descriptor instead.
func (*Message_DialResponse) Descriptor() ([]byte, []int) {
	return file_pb_autonat_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Message_DialResponse) GetStatus() Message_ResponseStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Message_OK
}

func (x *Message_DialResponse) GetStatusText() string {
	if x != nil && x.StatusText != nil {
		return *x.StatusText
	}
	return ""
}

func (x *Message_DialResponse) GetAddr() []byte {
	if x != nil {
		return x.Addr
	}
	return nil
}

var File_pb_autonat_proto protoreflect.FileDescriptor

var file_pb_autonat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6e, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x6e, 0x61, 0x74, 0x2e, 0x70, 0x62, 0x22, 0xb5,
	0x04, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6e,
	0x61, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x64, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6e, 0x61, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x69, 0x61, 0x6c, 0x12, 0x44, 0x0a,
	0x0c, 0x64, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6e, 0x61, 0x74, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x30, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05,
	0x61, 0x64, 0x64, 0x72, 0x73, 0x1a, 0x38, 0x0a, 0x04, 0x44, 0x69, 0x61, 0x6c, 0x12, 0x30, 0x0a,
	0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6e, 0x61, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x1a,
	0x7e, 0x0a, 0x0c, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6e, 0x61, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22,
	0x2a, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x49, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x49, 0x41, 0x4c,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x22, 0x69, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x64, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x5f, 0x44, 0x49, 0x41,
	0x4c, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x44, 0x10, 0x65, 0x12, 0x12, 0x0a, 0x0d, 0x45,
	0x5f, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xc8, 0x01, 0x12,
	0x15, 0x0a, 0x10, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xac, 0x02,
}

var (
	file_pb_autonat_proto_rawDescOnce sync.Once
	file_pb_autonat_proto_rawDescData = file_pb_autonat_proto_rawDesc
)

func file_pb_autonat_proto_rawDescGZIP() []byte {
	file_pb_autonat_proto_rawDescOnce.Do(func() {
		file_pb_autonat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_autonat_proto_rawDescData)
	})
	return file_pb_autonat_proto_rawDescData
}

var file_pb_autonat_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_autonat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_autonat_proto_goTypes = []interface{}{
	(Message_MessageType)(0),     // 0: autonat.pb.Message.MessageType
	(Message_ResponseStatus)(0),  // 1: autonat.pb.Message.ResponseStatus
	(*Message)(nil),              // 2: autonat.pb.Message
	(*Message_PeerInfo)(nil),     // 3: autonat.pb.Message.PeerInfo
	(*Message_Dial)(nil),         // 4: autonat.pb.Message.Dial
	(*Message_DialResponse)(nil), // 5: autonat.pb.Message.DialResponse
}
var file_pb_autonat_proto_depIdxs = []int32{
	0, // 0: autonat.pb.Message.type:type_name -> autonat.pb.Message.MessageType
	4, // 1: autonat.pb.Message.dial:type_name -> autonat.pb.Message.Dial
	5, // 2: autonat.pb.Message.dialResponse:type_name -> autonat.pb.Message.DialResponse
	3, // 3: autonat.pb.Message.Dial.peer:type_name -> autonat.pb.Message.PeerInfo
	1, // 4: autonat.pb.Message.DialResponse.status:type_name -> autonat.pb.Message.ResponseStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pb_autonat_proto_init() }
func file_pb_autonat_proto_init() {
	if File_pb_autonat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_autonat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pb_autonat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_PeerInfo); i {
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
		file_pb_autonat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_Dial); i {
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
		file_pb_autonat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_DialResponse); i {
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
			RawDescriptor: file_pb_autonat_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_autonat_proto_goTypes,
		DependencyIndexes: file_pb_autonat_proto_depIdxs,
		EnumInfos:         file_pb_autonat_proto_enumTypes,
		MessageInfos:      file_pb_autonat_proto_msgTypes,
	}.Build()
	File_pb_autonat_proto = out.File
	file_pb_autonat_proto_rawDesc = nil
	file_pb_autonat_proto_goTypes = nil
	file_pb_autonat_proto_depIdxs = nil
}
