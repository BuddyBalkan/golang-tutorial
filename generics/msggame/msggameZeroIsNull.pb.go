// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: msggameZeroIsNull.proto

package msggame

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

// 工程师游戏 棋盘节点
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XLocation int32 `protobuf:"varint,1,opt,name=XLocation,proto3" json:"XLocation,omitempty"`
	YLocation int32 `protobuf:"varint,2,opt,name=YLocation,proto3" json:"YLocation,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msggameZeroIsNull_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_msggameZeroIsNull_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_msggameZeroIsNull_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetXLocation() int32 {
	if x != nil {
		return x.XLocation
	}
	return 0
}

func (x *Point) GetYLocation() int32 {
	if x != nil {
		return x.YLocation
	}
	return 0
}

// 工程师游戏 棋盘链路
type ConnectionPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair *TerminalPair `protobuf:"bytes,1,opt,name=Pair,proto3" json:"Pair,omitempty"`
	Path []*Point      `protobuf:"bytes,2,rep,name=Path,proto3" json:"Path,omitempty"`
}

func (x *ConnectionPath) Reset() {
	*x = ConnectionPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msggameZeroIsNull_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionPath) ProtoMessage() {}

func (x *ConnectionPath) ProtoReflect() protoreflect.Message {
	mi := &file_msggameZeroIsNull_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionPath.ProtoReflect.Descriptor instead.
func (*ConnectionPath) Descriptor() ([]byte, []int) {
	return file_msggameZeroIsNull_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionPath) GetPair() *TerminalPair {
	if x != nil {
		return x.Pair
	}
	return nil
}

func (x *ConnectionPath) GetPath() []*Point {
	if x != nil {
		return x.Path
	}
	return nil
}

// 工程师游戏 棋盘上的一对点
type TerminalPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointA *Point `protobuf:"bytes,1,opt,name=PointA,proto3" json:"PointA,omitempty"`
	PointB *Point `protobuf:"bytes,2,opt,name=PointB,proto3" json:"PointB,omitempty"`
}

func (x *TerminalPair) Reset() {
	*x = TerminalPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msggameZeroIsNull_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalPair) ProtoMessage() {}

func (x *TerminalPair) ProtoReflect() protoreflect.Message {
	mi := &file_msggameZeroIsNull_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalPair.ProtoReflect.Descriptor instead.
func (*TerminalPair) Descriptor() ([]byte, []int) {
	return file_msggameZeroIsNull_proto_rawDescGZIP(), []int{2}
}

func (x *TerminalPair) GetPointA() *Point {
	if x != nil {
		return x.PointA
	}
	return nil
}

func (x *TerminalPair) GetPointB() *Point {
	if x != nil {
		return x.PointB
	}
	return nil
}

// 工程师游戏棋盘问题结构
type ConnectionQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SideTileCount    int32           `protobuf:"varint,1,opt,name=SideTileCount,proto3" json:"SideTileCount,omitempty"`
	TerminalPairList []*TerminalPair `protobuf:"bytes,2,rep,name=TerminalPairList,proto3" json:"TerminalPairList,omitempty"`
	CrossLocations   []*Point        `protobuf:"bytes,3,rep,name=CrossLocations,proto3" json:"CrossLocations,omitempty"`
}

func (x *ConnectionQuestion) Reset() {
	*x = ConnectionQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msggameZeroIsNull_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionQuestion) ProtoMessage() {}

func (x *ConnectionQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_msggameZeroIsNull_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionQuestion.ProtoReflect.Descriptor instead.
func (*ConnectionQuestion) Descriptor() ([]byte, []int) {
	return file_msggameZeroIsNull_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionQuestion) GetSideTileCount() int32 {
	if x != nil {
		return x.SideTileCount
	}
	return 0
}

func (x *ConnectionQuestion) GetTerminalPairList() []*TerminalPair {
	if x != nil {
		return x.TerminalPairList
	}
	return nil
}

func (x *ConnectionQuestion) GetCrossLocations() []*Point {
	if x != nil {
		return x.CrossLocations
	}
	return nil
}

var File_msggameZeroIsNull_proto protoreflect.FileDescriptor

var file_msggameZeroIsNull_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x73, 0x67, 0x67, 0x61, 0x6d, 0x65, 0x5a, 0x65, 0x72, 0x6f, 0x49, 0x73, 0x4e,
	0x75, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x45, 0x54, 0x22, 0x43, 0x0a,
	0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x58, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x58, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x59, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x59, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x45, 0x54, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x54, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x22, 0x54, 0x0a, 0x0c, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x12, 0x21, 0x0a, 0x06, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x54, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x12, 0x21, 0x0a, 0x06,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45,
	0x54, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x22,
	0xab, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x69, 0x64, 0x65, 0x54, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x53,
	0x69, 0x64, 0x65, 0x54, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x10,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x45, 0x54, 0x2e, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x52, 0x10, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0e, 0x43, 0x72,
	0x6f, 0x73, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x54, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0e, 0x43,
	0x72, 0x6f, 0x73, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x6d, 0x73, 0x67, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_msggameZeroIsNull_proto_rawDescOnce sync.Once
	file_msggameZeroIsNull_proto_rawDescData = file_msggameZeroIsNull_proto_rawDesc
)

func file_msggameZeroIsNull_proto_rawDescGZIP() []byte {
	file_msggameZeroIsNull_proto_rawDescOnce.Do(func() {
		file_msggameZeroIsNull_proto_rawDescData = protoimpl.X.CompressGZIP(file_msggameZeroIsNull_proto_rawDescData)
	})
	return file_msggameZeroIsNull_proto_rawDescData
}

var file_msggameZeroIsNull_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_msggameZeroIsNull_proto_goTypes = []interface{}{
	(*Point)(nil),              // 0: ET.Point
	(*ConnectionPath)(nil),     // 1: ET.ConnectionPath
	(*TerminalPair)(nil),       // 2: ET.TerminalPair
	(*ConnectionQuestion)(nil), // 3: ET.ConnectionQuestion
}
var file_msggameZeroIsNull_proto_depIdxs = []int32{
	2, // 0: ET.ConnectionPath.Pair:type_name -> ET.TerminalPair
	0, // 1: ET.ConnectionPath.Path:type_name -> ET.Point
	0, // 2: ET.TerminalPair.PointA:type_name -> ET.Point
	0, // 3: ET.TerminalPair.PointB:type_name -> ET.Point
	2, // 4: ET.ConnectionQuestion.TerminalPairList:type_name -> ET.TerminalPair
	0, // 5: ET.ConnectionQuestion.CrossLocations:type_name -> ET.Point
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_msggameZeroIsNull_proto_init() }
func file_msggameZeroIsNull_proto_init() {
	if File_msggameZeroIsNull_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msggameZeroIsNull_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_msggameZeroIsNull_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionPath); i {
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
		file_msggameZeroIsNull_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalPair); i {
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
		file_msggameZeroIsNull_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionQuestion); i {
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
			RawDescriptor: file_msggameZeroIsNull_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msggameZeroIsNull_proto_goTypes,
		DependencyIndexes: file_msggameZeroIsNull_proto_depIdxs,
		MessageInfos:      file_msggameZeroIsNull_proto_msgTypes,
	}.Build()
	File_msggameZeroIsNull_proto = out.File
	file_msggameZeroIsNull_proto_rawDesc = nil
	file_msggameZeroIsNull_proto_goTypes = nil
	file_msggameZeroIsNull_proto_depIdxs = nil
}
