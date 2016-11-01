// Code generated by protoc-gen-go.
// source: proto/protocol.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	proto/protocol.proto

It has these top-level messages:
	Controller
	Vec3
	PlayerPosition
	Terrain
	MessageItem
	MessageContainer
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MessageType определяет типы сообщений
type MessageType int32

const (
	MessageType_MsgController     MessageType = 0
	MessageType_MsgPlayerPosition MessageType = 1
	MessageType_MsgTerrain        MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "MsgController",
	1: "MsgPlayerPosition",
	2: "MsgTerrain",
}
var MessageType_value = map[string]int32{
	"MsgController":     0,
	"MsgPlayerPosition": 1,
	"MsgTerrain":        2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Controller определяет клавиши управления
type Controller struct {
	MoveForward  bool                  `protobuf:"varint,1,opt,name=MoveForward" json:"MoveForward,omitempty"`
	MoveBackward bool                  `protobuf:"varint,2,opt,name=MoveBackward" json:"MoveBackward,omitempty"`
	MoveLeft     bool                  `protobuf:"varint,3,opt,name=MoveLeft" json:"MoveLeft,omitempty"`
	MoveRight    bool                  `protobuf:"varint,4,opt,name=MoveRight" json:"MoveRight,omitempty"`
	RotateLeft   bool                  `protobuf:"varint,5,opt,name=RotateLeft" json:"RotateLeft,omitempty"`
	RotateRight  bool                  `protobuf:"varint,6,opt,name=RotateRight" json:"RotateRight,omitempty"`
	Modifiers    *Controller_Modifiers `protobuf:"bytes,7,opt,name=modifiers" json:"modifiers,omitempty"`
}

func (m *Controller) Reset()                    { *m = Controller{} }
func (m *Controller) String() string            { return proto.CompactTextString(m) }
func (*Controller) ProtoMessage()               {}
func (*Controller) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Controller) GetModifiers() *Controller_Modifiers {
	if m != nil {
		return m.Modifiers
	}
	return nil
}

// Modifiers определяет клавиши-модификаторы
type Controller_Modifiers struct {
	Shift bool `protobuf:"varint,1,opt,name=Shift" json:"Shift,omitempty"`
	Ctrl  bool `protobuf:"varint,2,opt,name=Ctrl" json:"Ctrl,omitempty"`
	Alt   bool `protobuf:"varint,3,opt,name=Alt" json:"Alt,omitempty"`
	Meta  bool `protobuf:"varint,4,opt,name=Meta" json:"Meta,omitempty"`
}

func (m *Controller_Modifiers) Reset()                    { *m = Controller_Modifiers{} }
func (m *Controller_Modifiers) String() string            { return proto.CompactTextString(m) }
func (*Controller_Modifiers) ProtoMessage()               {}
func (*Controller_Modifiers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Vec3 struct {
	X float64 `protobuf:"fixed64,1,opt,name=X" json:"X,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=Y" json:"Y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=Z" json:"Z,omitempty"`
}

func (m *Vec3) Reset()                    { *m = Vec3{} }
func (m *Vec3) String() string            { return proto.CompactTextString(m) }
func (*Vec3) ProtoMessage()               {}
func (*Vec3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// PlayerPosition определяет позицию и движение игрока
type PlayerPosition struct {
	// Position определяет положение игрока
	Position *Vec3 `protobuf:"bytes,1,opt,name=Position" json:"Position,omitempty"`
	// Motion определяет движение игрока
	Motion *Vec3 `protobuf:"bytes,2,opt,name=Motion" json:"Motion,omitempty"`
	// Angle определяет направление игрока (угол между положительным направленим оси Y и направлением игрока по часовой стрелке)
	Angle float64 `protobuf:"fixed64,3,opt,name=Angle" json:"Angle,omitempty"`
	// Slew определяет поворот игрока
	Slew float64 `protobuf:"fixed64,4,opt,name=Slew" json:"Slew,omitempty"`
}

func (m *PlayerPosition) Reset()                    { *m = PlayerPosition{} }
func (m *PlayerPosition) String() string            { return proto.CompactTextString(m) }
func (*PlayerPosition) ProtoMessage()               {}
func (*PlayerPosition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PlayerPosition) GetPosition() *Vec3 {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *PlayerPosition) GetMotion() *Vec3 {
	if m != nil {
		return m.Motion
	}
	return nil
}

// Map определяет карту
type Terrain struct {
	// Width определяет ширину
	Width int32 `protobuf:"varint,1,opt,name=Width" json:"Width,omitempty"`
	// Height определяет высоту
	Height int32 `protobuf:"varint,2,opt,name=Height" json:"Height,omitempty"`
	// Seed определяет зерно генератора
	Seed int64 `protobuf:"varint,3,opt,name=Seed" json:"Seed,omitempty"`
	// Map определяет карту
	Map []*Terrain_Tile `protobuf:"bytes,4,rep,name=Map" json:"Map,omitempty"`
}

func (m *Terrain) Reset()                    { *m = Terrain{} }
func (m *Terrain) String() string            { return proto.CompactTextString(m) }
func (*Terrain) ProtoMessage()               {}
func (*Terrain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Terrain) GetMap() []*Terrain_Tile {
	if m != nil {
		return m.Map
	}
	return nil
}

// Tile определяет тайл карты
type Terrain_Tile struct {
	Ground int32 `protobuf:"varint,1,opt,name=Ground" json:"Ground,omitempty"`
	Block  int32 `protobuf:"varint,2,opt,name=Block" json:"Block,omitempty"`
	Detail int32 `protobuf:"varint,3,opt,name=Detail" json:"Detail,omitempty"`
}

func (m *Terrain_Tile) Reset()                    { *m = Terrain_Tile{} }
func (m *Terrain_Tile) String() string            { return proto.CompactTextString(m) }
func (*Terrain_Tile) ProtoMessage()               {}
func (*Terrain_Tile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// MessageItem определяет одно сообщение
type MessageItem struct {
	Type MessageType `protobuf:"varint,1,opt,name=Type,enum=protocol.MessageType" json:"Type,omitempty"`
	Body []byte      `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *MessageItem) Reset()                    { *m = MessageItem{} }
func (m *MessageItem) String() string            { return proto.CompactTextString(m) }
func (*MessageItem) ProtoMessage()               {}
func (*MessageItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// MessageContainer определяет несколько сообщений
type MessageContainer struct {
	Messages []*MessageItem `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *MessageContainer) Reset()                    { *m = MessageContainer{} }
func (m *MessageContainer) String() string            { return proto.CompactTextString(m) }
func (*MessageContainer) ProtoMessage()               {}
func (*MessageContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MessageContainer) GetMessages() []*MessageItem {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*Controller)(nil), "protocol.Controller")
	proto.RegisterType((*Controller_Modifiers)(nil), "protocol.Controller.Modifiers")
	proto.RegisterType((*Vec3)(nil), "protocol.Vec3")
	proto.RegisterType((*PlayerPosition)(nil), "protocol.PlayerPosition")
	proto.RegisterType((*Terrain)(nil), "protocol.Terrain")
	proto.RegisterType((*Terrain_Tile)(nil), "protocol.Terrain.Tile")
	proto.RegisterType((*MessageItem)(nil), "protocol.MessageItem")
	proto.RegisterType((*MessageContainer)(nil), "protocol.MessageContainer")
	proto.RegisterEnum("protocol.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("proto/protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xd1, 0x6e, 0x12, 0x41,
	0x14, 0x75, 0x60, 0xa1, 0xf4, 0x82, 0x84, 0x4e, 0xda, 0x66, 0x43, 0x4c, 0x43, 0xf6, 0xc1, 0x60,
	0x1f, 0x50, 0xe9, 0xab, 0x2f, 0xa5, 0x6a, 0x35, 0x61, 0x93, 0x66, 0x4a, 0xb4, 0xad, 0x4f, 0x2b,
	0x3b, 0x2c, 0x93, 0x4e, 0x19, 0x32, 0x3b, 0xda, 0xf0, 0x13, 0xfe, 0x94, 0x5f, 0xe3, 0x5f, 0x98,
	0x7b, 0x77, 0xd8, 0xa5, 0xd1, 0x17, 0x72, 0xcf, 0xbd, 0xe7, 0x9e, 0x73, 0xcf, 0x32, 0x70, 0xb8,
	0xb6, 0xc6, 0x99, 0xd7, 0xf4, 0x3b, 0x37, 0x7a, 0x44, 0x05, 0x6f, 0x6d, 0x71, 0xf4, 0xa7, 0x06,
	0x70, 0x61, 0x56, 0xce, 0x1a, 0xad, 0xa5, 0xe5, 0x03, 0x68, 0xc7, 0xe6, 0xa7, 0xfc, 0x68, 0xec,
	0x63, 0x62, 0xd3, 0x90, 0x0d, 0xd8, 0xb0, 0x25, 0x76, 0x5b, 0x3c, 0x82, 0x0e, 0xc2, 0x49, 0x32,
	0xbf, 0x27, 0x4a, 0x8d, 0x28, 0x4f, 0x7a, 0xbc, 0x0f, 0x2d, 0xc4, 0x53, 0xb9, 0x70, 0x61, 0x9d,
	0xe6, 0x25, 0xe6, 0x2f, 0x60, 0x1f, 0x6b, 0xa1, 0xb2, 0xa5, 0x0b, 0x03, 0x1a, 0x56, 0x0d, 0x7e,
	0x02, 0x20, 0x8c, 0x4b, 0x5c, 0xb1, 0xdb, 0xa0, 0xf1, 0x4e, 0x07, 0xef, 0x2b, 0x50, 0xb1, 0xdf,
	0x2c, 0xee, 0xdb, 0x69, 0xf1, 0x77, 0xb0, 0xff, 0x60, 0x52, 0xb5, 0x50, 0xd2, 0xe6, 0xe1, 0xde,
	0x80, 0x0d, 0xdb, 0xe3, 0x93, 0x51, 0x19, 0xbf, 0x8a, 0x3a, 0x8a, 0xb7, 0x2c, 0x51, 0x2d, 0xf4,
	0xbf, 0xe1, 0x75, 0x1e, 0xf0, 0x43, 0x68, 0x5c, 0x2f, 0xd5, 0xc2, 0xf9, 0xcf, 0x50, 0x00, 0xce,
	0x21, 0xb8, 0x70, 0x56, 0xfb, 0xe0, 0x54, 0xf3, 0x1e, 0xd4, 0xcf, 0xf5, 0x36, 0x2b, 0x96, 0xc8,
	0x8a, 0xa5, 0x4b, 0x7c, 0x42, 0xaa, 0xa3, 0x37, 0x10, 0x7c, 0x91, 0xf3, 0x33, 0xde, 0x01, 0x76,
	0x43, 0x9a, 0x4c, 0xb0, 0x1b, 0x44, 0xb7, 0x24, 0xc6, 0x04, 0xbb, 0x45, 0x74, 0x47, 0x3a, 0x4c,
	0xb0, 0xbb, 0xe8, 0x17, 0x83, 0xee, 0x95, 0x4e, 0x36, 0xd2, 0x5e, 0x99, 0x5c, 0x39, 0x65, 0x56,
	0xfc, 0x14, 0x5a, 0xdb, 0x9a, 0x34, 0xda, 0xe3, 0x6e, 0x15, 0x0f, 0xe5, 0x45, 0x39, 0xe7, 0x2f,
	0xa1, 0x19, 0x1b, 0x62, 0xd6, 0xfe, 0xcb, 0xf4, 0x53, 0x0c, 0x7a, 0xbe, 0xca, 0xb4, 0xf4, 0xc6,
	0x05, 0xc0, 0x08, 0xd7, 0x5a, 0x3e, 0x52, 0x04, 0x26, 0xa8, 0x8e, 0x7e, 0x33, 0xd8, 0x9b, 0x49,
	0x6b, 0x13, 0x45, 0x5b, 0x5f, 0x55, 0xea, 0x96, 0x74, 0x46, 0x43, 0x14, 0x80, 0x1f, 0x43, 0xf3,
	0x93, 0xa4, 0x3f, 0xa7, 0x46, 0x6d, 0x8f, 0x48, 0x4d, 0xca, 0x94, 0x2c, 0xea, 0x82, 0x6a, 0x3e,
	0x84, 0x7a, 0x9c, 0xac, 0xc3, 0x60, 0x50, 0x1f, 0xb6, 0xc7, 0xc7, 0xd5, 0x71, 0xde, 0x61, 0x34,
	0x53, 0x5a, 0x0a, 0xa4, 0xf4, 0xa7, 0x10, 0x20, 0x40, 0xf5, 0x4b, 0x6b, 0x7e, 0xac, 0x52, 0x6f,
	0xea, 0x11, 0xde, 0x32, 0xd1, 0x66, 0x7e, 0xef, 0x4d, 0x0b, 0x80, 0xec, 0xf7, 0xd2, 0x25, 0x4a,
	0x93, 0x6b, 0x43, 0x78, 0x14, 0x4d, 0xa1, 0x1d, 0xcb, 0x3c, 0x4f, 0x32, 0xf9, 0xd9, 0xc9, 0x07,
	0xfe, 0x0a, 0x82, 0xd9, 0x66, 0x2d, 0x49, 0xb2, 0x3b, 0x3e, 0xaa, 0xee, 0xf0, 0x24, 0x1c, 0x0a,
	0xa2, 0x60, 0x8a, 0x89, 0x49, 0x37, 0x64, 0xd3, 0x11, 0x54, 0x47, 0x1f, 0xa0, 0xe7, 0x89, 0xf8,
	0xba, 0x12, 0xb5, 0x92, 0x96, 0xbf, 0x85, 0x96, 0xef, 0xe5, 0x21, 0xa3, 0x78, 0xff, 0xca, 0xa2,
	0xb7, 0x28, 0x69, 0xa7, 0x97, 0xe5, 0x51, 0xe4, 0x74, 0x00, 0xcf, 0xe3, 0x3c, 0xab, 0xde, 0x6b,
	0xef, 0x19, 0x3f, 0x82, 0x83, 0x38, 0xcf, 0x9e, 0xbe, 0x87, 0x1e, 0xe3, 0x5d, 0x80, 0x38, 0xcf,
	0xfc, 0x37, 0xeb, 0xd5, 0xbe, 0x37, 0xc9, 0xe8, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0x82, 0xcb, 0x04, 0xfb, 0x03, 0x00, 0x00,
}
