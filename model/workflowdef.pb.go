// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/workflowdef.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowDef struct {
	Name                          string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description                   string                    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Version                       int32                     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Tasks                         []*WorkflowTask           `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
	InputParameters               []string                  `protobuf:"bytes,5,rep,name=input_parameters,json=inputParameters,proto3" json:"input_parameters,omitempty"`
	OutputParameters              map[string]*_struct.Value `protobuf:"bytes,6,rep,name=output_parameters,json=outputParameters,proto3" json:"output_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FailureWorkflow               string                    `protobuf:"bytes,7,opt,name=failure_workflow,json=failureWorkflow,proto3" json:"failure_workflow,omitempty"`
	SchemaVersion                 int32                     `protobuf:"varint,8,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	Restartable                   bool                      `protobuf:"varint,9,opt,name=restartable,proto3" json:"restartable,omitempty"`
	WorkflowStatusListenerEnabled bool                      `protobuf:"varint,10,opt,name=workflow_status_listener_enabled,json=workflowStatusListenerEnabled,proto3" json:"workflow_status_listener_enabled,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                  `json:"-"`
	XXX_unrecognized              []byte                    `json:"-"`
	XXX_sizecache                 int32                     `json:"-"`
}

func (m *WorkflowDef) Reset()         { *m = WorkflowDef{} }
func (m *WorkflowDef) String() string { return proto.CompactTextString(m) }
func (*WorkflowDef) ProtoMessage()    {}
func (*WorkflowDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowdef_69bee65128423069, []int{0}
}
func (m *WorkflowDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowDef.Unmarshal(m, b)
}
func (m *WorkflowDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowDef.Marshal(b, m, deterministic)
}
func (dst *WorkflowDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowDef.Merge(dst, src)
}
func (m *WorkflowDef) XXX_Size() int {
	return xxx_messageInfo_WorkflowDef.Size(m)
}
func (m *WorkflowDef) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowDef.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowDef proto.InternalMessageInfo

func (m *WorkflowDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WorkflowDef) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkflowDef) GetTasks() []*WorkflowTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowDef) GetInputParameters() []string {
	if m != nil {
		return m.InputParameters
	}
	return nil
}

func (m *WorkflowDef) GetOutputParameters() map[string]*_struct.Value {
	if m != nil {
		return m.OutputParameters
	}
	return nil
}

func (m *WorkflowDef) GetFailureWorkflow() string {
	if m != nil {
		return m.FailureWorkflow
	}
	return ""
}

func (m *WorkflowDef) GetSchemaVersion() int32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func (m *WorkflowDef) GetRestartable() bool {
	if m != nil {
		return m.Restartable
	}
	return false
}

func (m *WorkflowDef) GetWorkflowStatusListenerEnabled() bool {
	if m != nil {
		return m.WorkflowStatusListenerEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*WorkflowDef)(nil), "conductor.proto.WorkflowDef")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.WorkflowDef.OutputParametersEntry")
}

func init() {
	proto.RegisterFile("model/workflowdef.proto", fileDescriptor_workflowdef_69bee65128423069)
}

var fileDescriptor_workflowdef_69bee65128423069 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5f, 0x6b, 0xdb, 0x30,
	0x10, 0x27, 0x4d, 0xd2, 0x36, 0x32, 0x5d, 0x32, 0xc1, 0x36, 0xd3, 0xad, 0x60, 0x06, 0x03, 0x0f,
	0x86, 0x0c, 0xe9, 0xcb, 0xe8, 0x63, 0x69, 0xd9, 0xcb, 0x60, 0xc5, 0x1b, 0x1d, 0x6c, 0x0f, 0x46,
	0x96, 0xcf, 0x8e, 0x89, 0x2c, 0x19, 0xe9, 0xd4, 0xae, 0xdf, 0x6c, 0x1f, 0x6f, 0x58, 0xb6, 0xbb,
	0x2c, 0xac, 0x6f, 0xd2, 0xef, 0xcf, 0xdd, 0xe9, 0xa7, 0x23, 0xaf, 0x1a, 0x5d, 0x80, 0x4c, 0xee,
	0xb5, 0xd9, 0x96, 0x52, 0xdf, 0x17, 0x50, 0xb2, 0xd6, 0x68, 0xd4, 0x74, 0x29, 0xb4, 0x2a, 0x9c,
	0x40, 0x6d, 0x7a, 0xe0, 0x34, 0xfc, 0x57, 0x89, 0xdc, 0x6e, 0x07, 0xe6, 0x4d, 0xa5, 0x75, 0x25,
	0x21, 0xf1, 0xb7, 0xdc, 0x95, 0x89, 0x45, 0xe3, 0x04, 0xf6, 0xec, 0xdb, 0xdf, 0x33, 0x12, 0x7c,
	0x1f, 0x4c, 0x57, 0x50, 0x52, 0x4a, 0x66, 0x8a, 0x37, 0x10, 0x4e, 0xa2, 0x49, 0xbc, 0x48, 0xfd,
	0x99, 0x46, 0x24, 0x28, 0xc0, 0x0a, 0x53, 0xb7, 0x58, 0x6b, 0x15, 0x1e, 0x78, 0x6a, 0x17, 0xa2,
	0x21, 0x39, 0xba, 0x03, 0x63, 0x3b, 0x76, 0x1a, 0x4d, 0xe2, 0x79, 0x3a, 0x5e, 0xe9, 0x39, 0x99,
	0x77, 0xb3, 0xd8, 0x70, 0x16, 0x4d, 0xe3, 0x60, 0x7d, 0xc6, 0xf6, 0x06, 0x67, 0x63, 0xf3, 0x6f,
	0xdc, 0x6e, 0xd3, 0x5e, 0x4b, 0xdf, 0x93, 0x55, 0xad, 0x5a, 0x87, 0x59, 0xcb, 0x0d, 0x6f, 0x00,
	0xc1, 0xd8, 0x70, 0x1e, 0x4d, 0xe3, 0x45, 0xba, 0xf4, 0xf8, 0xcd, 0x23, 0x4c, 0x33, 0xf2, 0x5c,
	0x3b, 0xdc, 0xd3, 0x1e, 0xfa, 0x5e, 0xeb, 0x27, 0x7b, 0x5d, 0x41, 0xc9, 0xbe, 0x78, 0xd7, 0xdf,
	0x4a, 0xd7, 0x0a, 0xcd, 0x43, 0xba, 0xd2, 0x7b, 0x70, 0x37, 0x4b, 0xc9, 0x6b, 0xe9, 0x0c, 0x64,
	0x63, 0xb8, 0xe1, 0x91, 0x4f, 0x60, 0x39, 0xe0, 0x63, 0x55, 0xfa, 0x8e, 0x3c, 0xb3, 0x62, 0x03,
	0x0d, 0xcf, 0xc6, 0x30, 0x8e, 0x7d, 0x18, 0x27, 0x3d, 0x7a, 0x3b, 0x44, 0x12, 0x91, 0xc0, 0x80,
	0x45, 0x6e, 0x90, 0xe7, 0x12, 0xc2, 0x45, 0x34, 0x89, 0x8f, 0xd3, 0x5d, 0x88, 0x7e, 0x22, 0xd1,
	0xd8, 0x2b, 0xb3, 0xc8, 0xd1, 0xd9, 0x4c, 0xd6, 0x16, 0x41, 0x81, 0xc9, 0x40, 0x75, 0x92, 0x22,
	0x24, 0xde, 0x76, 0x36, 0xea, 0xbe, 0x7a, 0xd9, 0xe7, 0x41, 0x75, 0xdd, 0x8b, 0x4e, 0x7f, 0x92,
	0x17, 0xff, 0x7d, 0x27, 0x5d, 0x91, 0xe9, 0x16, 0x1e, 0x86, 0x5f, 0xee, 0x8e, 0xf4, 0x03, 0x99,
	0xdf, 0x71, 0xe9, 0xc0, 0x7f, 0x6f, 0xb0, 0x7e, 0xc9, 0xfa, 0xb5, 0x61, 0xe3, 0xda, 0xb0, 0xdb,
	0x8e, 0x4d, 0x7b, 0xd1, 0xc5, 0xc1, 0xc7, 0xc9, 0xe5, 0x86, 0xbc, 0x16, 0xba, 0x61, 0x0a, 0xb0,
	0x94, 0xf5, 0xaf, 0xfd, 0xb0, 0x2f, 0x4f, 0x76, 0xd2, 0xbe, 0xc9, 0x7f, 0x5c, 0x54, 0x35, 0x6e,
	0x5c, 0xce, 0x84, 0x6e, 0x92, 0xc1, 0x92, 0x3c, 0x5a, 0x12, 0x21, 0x6b, 0x50, 0x98, 0x54, 0xba,
	0x32, 0xad, 0xd8, 0xc1, 0xfd, 0x56, 0xe7, 0x87, 0xbe, 0xe2, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0x26, 0xea, 0x86, 0x0f, 0x03, 0x00, 0x00,
}