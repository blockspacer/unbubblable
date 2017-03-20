// Code generated by protoc-gen-go.
// source: UnProtos/Messaging/Responses/ApplyControllerResponse.proto
// DO NOT EDIT!

/*
Package UnProtos_Messaging_Responses is a generated protocol buffer package.

It is generated from these files:
	UnProtos/Messaging/Responses/ApplyControllerResponse.proto
	UnProtos/Messaging/Responses/GetChankResponse.proto
	UnProtos/Messaging/Responses/GetTerrainResponse.proto

It has these top-level messages:
	ApplyControllerResponse
	GetChunkResponse
	GetTerrainResponse
*/
package UnProtos_Messaging_Responses

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import UnProtos_Data1 "UnProtos/Data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApplyControllerResponse_Result int32

const (
	ApplyControllerResponse_UNSET   ApplyControllerResponse_Result = 0
	ApplyControllerResponse_SUCCESS ApplyControllerResponse_Result = 1
)

var ApplyControllerResponse_Result_name = map[int32]string{
	0: "UNSET",
	1: "SUCCESS",
}
var ApplyControllerResponse_Result_value = map[string]int32{
	"UNSET":   0,
	"SUCCESS": 1,
}

func (x ApplyControllerResponse_Result) String() string {
	return proto.EnumName(ApplyControllerResponse_Result_name, int32(x))
}
func (ApplyControllerResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// ApplyControllerResponse определяет позицию и движение игрока
type ApplyControllerResponse struct {
	Result   ApplyControllerResponse_Result `protobuf:"varint,1,opt,name=result,enum=UnProtos.Messaging.Responses.ApplyControllerResponse_Result" json:"result,omitempty"`
	Movement *UnProtos_Data1.Movement       `protobuf:"bytes,2,opt,name=movement" json:"movement,omitempty"`
}

func (m *ApplyControllerResponse) Reset()                    { *m = ApplyControllerResponse{} }
func (m *ApplyControllerResponse) String() string            { return proto.CompactTextString(m) }
func (*ApplyControllerResponse) ProtoMessage()               {}
func (*ApplyControllerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApplyControllerResponse) GetResult() ApplyControllerResponse_Result {
	if m != nil {
		return m.Result
	}
	return ApplyControllerResponse_UNSET
}

func (m *ApplyControllerResponse) GetMovement() *UnProtos_Data1.Movement {
	if m != nil {
		return m.Movement
	}
	return nil
}

func init() {
	proto.RegisterType((*ApplyControllerResponse)(nil), "UnProtos.Messaging.Responses.ApplyControllerResponse")
	proto.RegisterEnum("UnProtos.Messaging.Responses.ApplyControllerResponse_Result", ApplyControllerResponse_Result_name, ApplyControllerResponse_Result_value)
}

func init() {
	proto.RegisterFile("UnProtos/Messaging/Responses/ApplyControllerResponse.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x0a, 0xcd, 0x0b, 0x28,
	0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0xcf, 0xcc, 0x4b, 0xd7, 0x0f,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x2d, 0xd6, 0x77, 0x2c, 0x28, 0xc8, 0xa9, 0x74, 0xce,
	0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0x82, 0x49, 0xe8, 0x15, 0x80, 0xb4, 0x08, 0xc9,
	0xc0, 0xf4, 0xea, 0xc1, 0xf5, 0xea, 0xc1, 0xf5, 0x4a, 0xc1, 0x65, 0xf5, 0x5d, 0x12, 0x4b, 0x12,
	0xf5, 0x7d, 0xf3, 0xcb, 0x52, 0x73, 0x53, 0xf3, 0x4a, 0x20, 0x7a, 0x95, 0x8e, 0x31, 0x72, 0x89,
	0xe3, 0x30, 0x5d, 0x28, 0x84, 0x8b, 0xad, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x44, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xcf, 0xc8, 0x46, 0x0f, 0x9f, 0x45, 0x7a, 0xb8, 0x1c, 0x19, 0x04, 0x36, 0x23, 0x08,
	0x6a, 0x96, 0x90, 0x31, 0x17, 0x47, 0x2e, 0xd4, 0x0d, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0xe2, 0x08, 0x73, 0x41, 0x4e, 0xd4, 0x83, 0x39, 0x31, 0x08, 0xae, 0x50, 0x49, 0x81, 0x8b, 0x0d,
	0x62, 0x8c, 0x10, 0x27, 0x17, 0x6b, 0xa8, 0x5f, 0xb0, 0x6b, 0x88, 0x00, 0x83, 0x10, 0x37, 0x17,
	0x7b, 0x70, 0xa8, 0xb3, 0xb3, 0x6b, 0x70, 0xb0, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x3f, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0xcf, 0xc8, 0xa9, 0x49, 0x01, 0x00, 0x00,
}
