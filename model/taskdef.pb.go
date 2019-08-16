// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/taskdef.proto

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

type TaskDef_RetryLogic int32

const (
	TaskDef_FIXED               TaskDef_RetryLogic = 0
	TaskDef_EXPONENTIAL_BACKOFF TaskDef_RetryLogic = 1
)

var TaskDef_RetryLogic_name = map[int32]string{
	0: "FIXED",
	1: "EXPONENTIAL_BACKOFF",
}
var TaskDef_RetryLogic_value = map[string]int32{
	"FIXED":               0,
	"EXPONENTIAL_BACKOFF": 1,
}

func (x TaskDef_RetryLogic) String() string {
	return proto.EnumName(TaskDef_RetryLogic_name, int32(x))
}
func (TaskDef_RetryLogic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_taskdef_5b295c16f2293610, []int{0, 0}
}

type TaskDef_TimeoutPolicy int32

const (
	TaskDef_RETRY       TaskDef_TimeoutPolicy = 0
	TaskDef_TIME_OUT_WF TaskDef_TimeoutPolicy = 1
	TaskDef_ALERT_ONLY  TaskDef_TimeoutPolicy = 2
)

var TaskDef_TimeoutPolicy_name = map[int32]string{
	0: "RETRY",
	1: "TIME_OUT_WF",
	2: "ALERT_ONLY",
}
var TaskDef_TimeoutPolicy_value = map[string]int32{
	"RETRY":       0,
	"TIME_OUT_WF": 1,
	"ALERT_ONLY":  2,
}

func (x TaskDef_TimeoutPolicy) String() string {
	return proto.EnumName(TaskDef_TimeoutPolicy_name, int32(x))
}
func (TaskDef_TimeoutPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_taskdef_5b295c16f2293610, []int{0, 1}
}

type TaskDef struct {
	Name                        string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description                 string                    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	RetryCount                  int32                     `protobuf:"varint,3,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	TimeoutSeconds              int64                     `protobuf:"varint,4,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	InputKeys                   []string                  `protobuf:"bytes,5,rep,name=input_keys,json=inputKeys,proto3" json:"input_keys,omitempty"`
	OutputKeys                  []string                  `protobuf:"bytes,6,rep,name=output_keys,json=outputKeys,proto3" json:"output_keys,omitempty"`
	TimeoutPolicy               TaskDef_TimeoutPolicy     `protobuf:"varint,7,opt,name=timeout_policy,json=timeoutPolicy,proto3,enum=conductor.proto.TaskDef_TimeoutPolicy" json:"timeout_policy,omitempty"`
	RetryLogic                  TaskDef_RetryLogic        `protobuf:"varint,8,opt,name=retry_logic,json=retryLogic,proto3,enum=conductor.proto.TaskDef_RetryLogic" json:"retry_logic,omitempty"`
	RetryDelaySeconds           int32                     `protobuf:"varint,9,opt,name=retry_delay_seconds,json=retryDelaySeconds,proto3" json:"retry_delay_seconds,omitempty"`
	ResponseTimeoutSeconds      int64                     `protobuf:"varint,10,opt,name=response_timeout_seconds,json=responseTimeoutSeconds,proto3" json:"response_timeout_seconds,omitempty"`
	ConcurrentExecLimit         int32                     `protobuf:"varint,11,opt,name=concurrent_exec_limit,json=concurrentExecLimit,proto3" json:"concurrent_exec_limit,omitempty"`
	InputTemplate               map[string]*_struct.Value `protobuf:"bytes,12,rep,name=input_template,json=inputTemplate,proto3" json:"input_template,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RateLimitPerFrequency       int32                     `protobuf:"varint,14,opt,name=rate_limit_per_frequency,json=rateLimitPerFrequency,proto3" json:"rate_limit_per_frequency,omitempty"`
	RateLimitFrequencyInSeconds int32                     `protobuf:"varint,15,opt,name=rate_limit_frequency_in_seconds,json=rateLimitFrequencyInSeconds,proto3" json:"rate_limit_frequency_in_seconds,omitempty"`
	IsolationGroupId            string                    `protobuf:"bytes,16,opt,name=isolation_group_id,json=isolationGroupId,proto3" json:"isolation_group_id,omitempty"`
	ExecutionNameSpace          string                    `protobuf:"bytes,17,opt,name=execution_name_space,json=executionNameSpace,proto3" json:"execution_name_space,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                  `json:"-"`
	XXX_unrecognized            []byte                    `json:"-"`
	XXX_sizecache               int32                     `json:"-"`
}

func (m *TaskDef) Reset()         { *m = TaskDef{} }
func (m *TaskDef) String() string { return proto.CompactTextString(m) }
func (*TaskDef) ProtoMessage()    {}
func (*TaskDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskdef_5b295c16f2293610, []int{0}
}
func (m *TaskDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDef.Unmarshal(m, b)
}
func (m *TaskDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDef.Marshal(b, m, deterministic)
}
func (dst *TaskDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDef.Merge(dst, src)
}
func (m *TaskDef) XXX_Size() int {
	return xxx_messageInfo_TaskDef.Size(m)
}
func (m *TaskDef) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDef.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDef proto.InternalMessageInfo

func (m *TaskDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskDef) GetRetryCount() int32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *TaskDef) GetTimeoutSeconds() int64 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *TaskDef) GetInputKeys() []string {
	if m != nil {
		return m.InputKeys
	}
	return nil
}

func (m *TaskDef) GetOutputKeys() []string {
	if m != nil {
		return m.OutputKeys
	}
	return nil
}

func (m *TaskDef) GetTimeoutPolicy() TaskDef_TimeoutPolicy {
	if m != nil {
		return m.TimeoutPolicy
	}
	return TaskDef_RETRY
}

func (m *TaskDef) GetRetryLogic() TaskDef_RetryLogic {
	if m != nil {
		return m.RetryLogic
	}
	return TaskDef_FIXED
}

func (m *TaskDef) GetRetryDelaySeconds() int32 {
	if m != nil {
		return m.RetryDelaySeconds
	}
	return 0
}

func (m *TaskDef) GetResponseTimeoutSeconds() int64 {
	if m != nil {
		return m.ResponseTimeoutSeconds
	}
	return 0
}

func (m *TaskDef) GetConcurrentExecLimit() int32 {
	if m != nil {
		return m.ConcurrentExecLimit
	}
	return 0
}

func (m *TaskDef) GetInputTemplate() map[string]*_struct.Value {
	if m != nil {
		return m.InputTemplate
	}
	return nil
}

func (m *TaskDef) GetRateLimitPerFrequency() int32 {
	if m != nil {
		return m.RateLimitPerFrequency
	}
	return 0
}

func (m *TaskDef) GetRateLimitFrequencyInSeconds() int32 {
	if m != nil {
		return m.RateLimitFrequencyInSeconds
	}
	return 0
}

func (m *TaskDef) GetIsolationGroupId() string {
	if m != nil {
		return m.IsolationGroupId
	}
	return ""
}

func (m *TaskDef) GetExecutionNameSpace() string {
	if m != nil {
		return m.ExecutionNameSpace
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskDef)(nil), "conductor.proto.TaskDef")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.TaskDef.InputTemplateEntry")
	proto.RegisterEnum("conductor.proto.TaskDef_RetryLogic", TaskDef_RetryLogic_name, TaskDef_RetryLogic_value)
	proto.RegisterEnum("conductor.proto.TaskDef_TimeoutPolicy", TaskDef_TimeoutPolicy_name, TaskDef_TimeoutPolicy_value)
}

func init() { proto.RegisterFile("model/taskdef.proto", fileDescriptor_taskdef_5b295c16f2293610) }

var fileDescriptor_taskdef_5b295c16f2293610 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xc1, 0x4f, 0xdb, 0x30,
	0x14, 0xc6, 0x09, 0xa5, 0xb0, 0xbe, 0x8e, 0x36, 0xb8, 0x83, 0x45, 0xb0, 0x89, 0x8a, 0x49, 0x5b,
	0xa5, 0xa1, 0x14, 0x75, 0x87, 0x21, 0x76, 0x02, 0x9a, 0x4e, 0x15, 0xa5, 0xad, 0x42, 0xb6, 0xc1,
	0x2e, 0x56, 0xea, 0xb8, 0x5d, 0xd4, 0x24, 0xce, 0x1c, 0x67, 0x22, 0xff, 0xf9, 0x8e, 0x93, 0x9d,
	0x36, 0x14, 0x26, 0x6e, 0xee, 0xf7, 0xfb, 0xde, 0xb3, 0xfd, 0xe5, 0xb9, 0xd0, 0x08, 0x99, 0x47,
	0x83, 0xb6, 0x70, 0x93, 0xb9, 0x47, 0xa7, 0x66, 0xcc, 0x99, 0x60, 0xa8, 0x4e, 0x58, 0xe4, 0xa5,
	0x44, 0x30, 0x9e, 0x0b, 0xfb, 0x6f, 0x66, 0x8c, 0xcd, 0x02, 0xda, 0x56, 0xbf, 0x26, 0xe9, 0xb4,
	0x9d, 0x08, 0x9e, 0x12, 0x91, 0xd3, 0xa3, 0xbf, 0x5b, 0xb0, 0xe5, 0xb8, 0xc9, 0xbc, 0x4b, 0xa7,
	0x08, 0xc1, 0x46, 0xe4, 0x86, 0xd4, 0xd0, 0x9a, 0x5a, 0xab, 0x62, 0xab, 0x35, 0x6a, 0x42, 0xd5,
	0xa3, 0x09, 0xe1, 0x7e, 0x2c, 0x7c, 0x16, 0x19, 0xeb, 0x0a, 0xad, 0x4a, 0xe8, 0x10, 0xaa, 0x9c,
	0x0a, 0x9e, 0x61, 0xc2, 0xd2, 0x48, 0x18, 0xa5, 0xa6, 0xd6, 0x2a, 0xdb, 0xa0, 0xa4, 0x4b, 0xa9,
	0xa0, 0x0f, 0x50, 0x17, 0x7e, 0x48, 0x59, 0x2a, 0x70, 0x42, 0xe5, 0xe9, 0x12, 0x63, 0xa3, 0xa9,
	0xb5, 0x4a, 0x76, 0x6d, 0x21, 0xdf, 0xe4, 0x2a, 0x7a, 0x0b, 0xe0, 0x47, 0x71, 0x2a, 0xf0, 0x9c,
	0x66, 0x89, 0x51, 0x6e, 0x96, 0x5a, 0x15, 0xbb, 0xa2, 0x94, 0x2b, 0x9a, 0x25, 0x72, 0x23, 0x96,
	0x8a, 0x82, 0x6f, 0x2a, 0x0e, 0xb9, 0xa4, 0x0c, 0xd7, 0xb0, 0xec, 0x88, 0x63, 0x16, 0xf8, 0x24,
	0x33, 0xb6, 0x9a, 0x5a, 0xab, 0xd6, 0x79, 0x6f, 0x3e, 0xc9, 0xc4, 0x5c, 0xdc, 0xd8, 0x74, 0x72,
	0xfb, 0x58, 0xb9, 0xed, 0x6d, 0xb1, 0xfa, 0x13, 0x75, 0x97, 0x17, 0x0b, 0xd8, 0xcc, 0x27, 0xc6,
	0x0b, 0xd5, 0xeb, 0xdd, 0xb3, 0xbd, 0x6c, 0xe9, 0x1d, 0x48, 0xeb, 0xe2, 0xf6, 0x6a, 0x8d, 0x4c,
	0x68, 0xe4, 0x5d, 0x3c, 0x1a, 0xb8, 0x59, 0x91, 0x40, 0x45, 0xc5, 0xb4, 0xa3, 0x50, 0x57, 0x92,
	0x65, 0x08, 0xa7, 0x60, 0x70, 0x9a, 0xc4, 0x2c, 0x4a, 0x28, 0x7e, 0x1a, 0x1b, 0xa8, 0xd8, 0xf6,
	0x96, 0xdc, 0x79, 0x1c, 0x5f, 0x07, 0x76, 0x09, 0x8b, 0x48, 0xca, 0x39, 0x8d, 0x04, 0xa6, 0xf7,
	0x94, 0xe0, 0xc0, 0x0f, 0x7d, 0x61, 0x54, 0xd5, 0x5e, 0x8d, 0x07, 0x68, 0xdd, 0x53, 0x32, 0x90,
	0x08, 0xd9, 0x50, 0xcb, 0x23, 0x17, 0x34, 0x8c, 0x03, 0x57, 0x50, 0xe3, 0x65, 0xb3, 0xd4, 0xaa,
	0x76, 0x3e, 0x3e, 0x7b, 0xcd, 0xbe, 0xb4, 0x3b, 0x0b, 0xb7, 0x15, 0x09, 0x9e, 0xd9, 0xdb, 0xfe,
	0xaa, 0x86, 0x3e, 0x83, 0xc1, 0x5d, 0x41, 0xf3, 0xcd, 0x71, 0x4c, 0x39, 0x9e, 0x72, 0xfa, 0x3b,
	0xa5, 0x11, 0xc9, 0x8c, 0x9a, 0x3a, 0xca, 0xae, 0xe4, 0xea, 0x00, 0x63, 0xca, 0x7b, 0x4b, 0x88,
	0xba, 0x70, 0xb8, 0x52, 0x58, 0x14, 0x61, 0x3f, 0x2a, 0x12, 0xa8, 0xab, 0xfa, 0x83, 0xa2, 0xbe,
	0x28, 0xee, 0x47, 0xcb, 0x18, 0x8e, 0x01, 0xf9, 0x09, 0x0b, 0x5c, 0x39, 0x9c, 0x78, 0xc6, 0x59,
	0x1a, 0x63, 0xdf, 0x33, 0x74, 0x35, 0xb8, 0x7a, 0x41, 0xbe, 0x4a, 0xd0, 0xf7, 0xd0, 0x09, 0xbc,
	0x92, 0x49, 0xa5, 0xca, 0x2d, 0x27, 0x1e, 0x27, 0xb1, 0x4b, 0xa8, 0xb1, 0xa3, 0xfc, 0xa8, 0x60,
	0x43, 0x37, 0xa4, 0x37, 0x92, 0xec, 0xdf, 0x02, 0xfa, 0x3f, 0x03, 0xa4, 0x43, 0x69, 0x4e, 0xb3,
	0xc5, 0xd3, 0x91, 0x4b, 0x74, 0x0c, 0xe5, 0x3f, 0x6e, 0x90, 0x52, 0xf5, 0x66, 0xaa, 0x9d, 0x3d,
	0x33, 0x7f, 0x87, 0xe6, 0xf2, 0x1d, 0x9a, 0xdf, 0x25, 0xb5, 0x73, 0xd3, 0xd9, 0xfa, 0xa9, 0x76,
	0x74, 0x02, 0xf0, 0x30, 0x44, 0xa8, 0x02, 0xe5, 0x5e, 0xff, 0xd6, 0xea, 0xea, 0x6b, 0xe8, 0x35,
	0x34, 0xac, 0xdb, 0xf1, 0x68, 0x68, 0x0d, 0x9d, 0xfe, 0xf9, 0x00, 0x5f, 0x9c, 0x5f, 0x5e, 0x8d,
	0x7a, 0x3d, 0x5d, 0x3b, 0xfa, 0x02, 0xdb, 0x8f, 0x46, 0x58, 0x16, 0xd9, 0x96, 0x63, 0xdf, 0xe9,
	0x6b, 0xa8, 0x0e, 0x55, 0xa7, 0x7f, 0x6d, 0xe1, 0xd1, 0x37, 0x07, 0xff, 0xe8, 0xe9, 0x1a, 0xaa,
	0x01, 0x9c, 0x0f, 0x2c, 0xdb, 0xc1, 0xa3, 0xe1, 0xe0, 0x4e, 0x5f, 0xbf, 0xf0, 0xe0, 0x80, 0xb0,
	0xd0, 0x8c, 0xa8, 0x98, 0x06, 0xfe, 0xfd, 0xd3, 0x0f, 0x7e, 0x51, 0x59, 0x7c, 0xf1, 0xf1, 0xe4,
	0xe7, 0xd9, 0xcc, 0x17, 0xbf, 0xd2, 0x89, 0x49, 0x58, 0xd8, 0x5e, 0xd8, 0xdb, 0x85, 0xbd, 0x4d,
	0x02, 0x9f, 0x46, 0xa2, 0x3d, 0x63, 0x33, 0x1e, 0x93, 0x15, 0x5d, 0xfd, 0x3b, 0x4d, 0x36, 0x55,
	0xb7, 0x4f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x66, 0x28, 0x43, 0xc9, 0xad, 0x04, 0x00, 0x00,
}
