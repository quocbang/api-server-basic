// Code generated by protoc-gen-go. DO NOT EDIT.
// source: func_generate.proto

package function

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// FuntionOperationID is list of existing function in service.
type FunctionOperationID int32

const (
	FunctionOperationID_CREATE_ACCOUNT  FunctionOperationID = 0
	FunctionOperationID_DELETE_ACCOUNTS FunctionOperationID = 1
	FunctionOperationID_GET_TASKS       FunctionOperationID = 2
	FunctionOperationID_CREATE_TASKS    FunctionOperationID = 3
	FunctionOperationID_UPDATE_TASK     FunctionOperationID = 4
	FunctionOperationID_DELETE_TASKS    FunctionOperationID = 5
)

var FunctionOperationID_name = map[int32]string{
	0: "CREATE_ACCOUNT",
	1: "DELETE_ACCOUNTS",
	2: "GET_TASKS",
	3: "CREATE_TASKS",
	4: "UPDATE_TASK",
	5: "DELETE_TASKS",
}

var FunctionOperationID_value = map[string]int32{
	"CREATE_ACCOUNT":  0,
	"DELETE_ACCOUNTS": 1,
	"GET_TASKS":       2,
	"CREATE_TASKS":    3,
	"UPDATE_TASK":     4,
	"DELETE_TASKS":    5,
}

func (x FunctionOperationID) String() string {
	return proto.EnumName(FunctionOperationID_name, int32(x))
}

func (FunctionOperationID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c854d1f7de7de144, []int{0}
}

func init() {
	proto.RegisterEnum("function.FunctionOperationID", FunctionOperationID_name, FunctionOperationID_value)
}

func init() { proto.RegisterFile("func_generate.proto", fileDescriptor_c854d1f7de7de144) }

var fileDescriptor_c854d1f7de7de144 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2b, 0xcd, 0x4b,
	0x8e, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x00, 0x09, 0x96, 0x64, 0xe6, 0xe7, 0x69, 0x35, 0x31, 0x72, 0x09, 0xbb, 0x41, 0x39, 0xfe,
	0x05, 0x20, 0x35, 0x99, 0xf9, 0x79, 0x9e, 0x2e, 0x42, 0x42, 0x5c, 0x7c, 0xce, 0x41, 0xae, 0x8e,
	0x21, 0xae, 0xf1, 0x8e, 0xce, 0xce, 0xfe, 0xa1, 0x7e, 0x21, 0x02, 0x0c, 0x42, 0xc2, 0x5c, 0xfc,
	0x2e, 0xae, 0x3e, 0xae, 0x08, 0xb1, 0x60, 0x01, 0x46, 0x21, 0x5e, 0x2e, 0x4e, 0x77, 0xd7, 0x90,
	0xf8, 0x10, 0xc7, 0x60, 0xef, 0x60, 0x01, 0x26, 0x21, 0x01, 0x2e, 0x1e, 0xa8, 0x3e, 0x88, 0x08,
	0xb3, 0x10, 0x3f, 0x17, 0x77, 0x68, 0x80, 0x0b, 0x4c, 0x44, 0x80, 0x05, 0xa4, 0x04, 0x6a, 0x0c,
	0x44, 0x09, 0x6b, 0x12, 0x1b, 0xd8, 0x55, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0xdb,
	0xc5, 0xde, 0xac, 0x00, 0x00, 0x00,
}
