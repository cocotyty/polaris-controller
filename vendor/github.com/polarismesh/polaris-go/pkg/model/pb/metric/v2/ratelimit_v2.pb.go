// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ratelimit_v2.proto

package v2

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

// 命令字
type RateLimitCmd int32

const (
	RateLimitCmd_INIT    RateLimitCmd = 0
	RateLimitCmd_ACQUIRE RateLimitCmd = 1
)

var RateLimitCmd_name = map[int32]string{
	0: "INIT",
	1: "ACQUIRE",
}
var RateLimitCmd_value = map[string]int32{
	"INIT":    0,
	"ACQUIRE": 1,
}

func (x RateLimitCmd) String() string {
	return proto.EnumName(RateLimitCmd_name, int32(x))
}
func (RateLimitCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{0}
}

// 限频模式
type Mode int32

const (
	// 自适应模式，根据历史流量自动调整
	Mode_ADAPTIVE Mode = 0
	// 批量抢占模式，客户端进行拉取，Server返回全量剩余配额
	Mode_BATCH_OCCUPY Mode = 1
	// 批量分摊模式，客户端进行拉取，Server按比例进行分摊
	Mode_BATCH_SHARE Mode = 2
)

var Mode_name = map[int32]string{
	0: "ADAPTIVE",
	1: "BATCH_OCCUPY",
	2: "BATCH_SHARE",
}
var Mode_value = map[string]int32{
	"ADAPTIVE":     0,
	"BATCH_OCCUPY": 1,
	"BATCH_SHARE":  2,
}

func (x Mode) String() string {
	return proto.EnumName(Mode_name, int32(x))
}
func (Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{1}
}

// 阈值模式
type QuotaMode int32

const (
	// 整体阈值
	QuotaMode_WHOLE QuotaMode = 0
	// 单机均分阈值
	QuotaMode_DIVIDE QuotaMode = 1
)

var QuotaMode_name = map[int32]string{
	0: "WHOLE",
	1: "DIVIDE",
}
var QuotaMode_value = map[string]int32{
	"WHOLE":  0,
	"DIVIDE": 1,
}

func (x QuotaMode) String() string {
	return proto.EnumName(QuotaMode_name, int32(x))
}
func (QuotaMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{2}
}

// 限流请求
type RateLimitRequest struct {
	// 命令字
	Cmd RateLimitCmd `protobuf:"varint,1,opt,name=cmd,proto3,enum=polaris.metric.v2.RateLimitCmd" json:"cmd,omitempty"`
	// 初始化请求
	RateLimitInitRequest *RateLimitInitRequest `protobuf:"bytes,2,opt,name=rateLimitInitRequest,proto3" json:"rateLimitInitRequest,omitempty"`
	// 上报请求
	RateLimitReportRequest *RateLimitReportRequest `protobuf:"bytes,3,opt,name=rateLimitReportRequest,proto3" json:"rateLimitReportRequest,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{0}
}
func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (dst *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(dst, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetCmd() RateLimitCmd {
	if m != nil {
		return m.Cmd
	}
	return RateLimitCmd_INIT
}

func (m *RateLimitRequest) GetRateLimitInitRequest() *RateLimitInitRequest {
	if m != nil {
		return m.RateLimitInitRequest
	}
	return nil
}

func (m *RateLimitRequest) GetRateLimitReportRequest() *RateLimitReportRequest {
	if m != nil {
		return m.RateLimitReportRequest
	}
	return nil
}

// 限流应答
type RateLimitResponse struct {
	// 命令字
	Cmd RateLimitCmd `protobuf:"varint,1,opt,name=cmd,proto3,enum=polaris.metric.v2.RateLimitCmd" json:"cmd,omitempty"`
	// 初始化请求
	RateLimitInitResponse *RateLimitInitResponse `protobuf:"bytes,2,opt,name=rateLimitInitResponse,proto3" json:"rateLimitInitResponse,omitempty"`
	// 上报请求
	RateLimitReportResponse *RateLimitReportResponse `protobuf:"bytes,3,opt,name=rateLimitReportResponse,proto3" json:"rateLimitReportResponse,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{1}
}
func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (dst *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(dst, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetCmd() RateLimitCmd {
	if m != nil {
		return m.Cmd
	}
	return RateLimitCmd_INIT
}

func (m *RateLimitResponse) GetRateLimitInitResponse() *RateLimitInitResponse {
	if m != nil {
		return m.RateLimitInitResponse
	}
	return nil
}

func (m *RateLimitResponse) GetRateLimitReportResponse() *RateLimitReportResponse {
	if m != nil {
		return m.RateLimitReportResponse
	}
	return nil
}

// 初始化请求
type RateLimitInitRequest struct {
	// 限流目标对象数据
	Target *LimitTarget `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// 客户端唯一标识
	ClientId string `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// 限流规则信息
	Totals []*QuotaTotal `protobuf:"bytes,3,rep,name=totals,proto3" json:"totals,omitempty"`
	// 客户端可指定滑窗数，不指定用默认值
	SlideCount uint32 `protobuf:"varint,4,opt,name=slideCount,proto3" json:"slideCount,omitempty"`
	// 限流模式
	Mode                 Mode     `protobuf:"varint,5,opt,name=mode,proto3,enum=polaris.metric.v2.Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitInitRequest) Reset()         { *m = RateLimitInitRequest{} }
func (m *RateLimitInitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitInitRequest) ProtoMessage()    {}
func (*RateLimitInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{2}
}
func (m *RateLimitInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitInitRequest.Unmarshal(m, b)
}
func (m *RateLimitInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitInitRequest.Marshal(b, m, deterministic)
}
func (dst *RateLimitInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitInitRequest.Merge(dst, src)
}
func (m *RateLimitInitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitInitRequest.Size(m)
}
func (m *RateLimitInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitInitRequest proto.InternalMessageInfo

func (m *RateLimitInitRequest) GetTarget() *LimitTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *RateLimitInitRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *RateLimitInitRequest) GetTotals() []*QuotaTotal {
	if m != nil {
		return m.Totals
	}
	return nil
}

func (m *RateLimitInitRequest) GetSlideCount() uint32 {
	if m != nil {
		return m.SlideCount
	}
	return 0
}

func (m *RateLimitInitRequest) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_ADAPTIVE
}

// 初始化应答
type RateLimitInitResponse struct {
	// 应答错误码
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 限流目标对象，回传给客户端
	Target *LimitTarget `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// 客户端的标识，与clientId对应，一个server全局唯一，上报时候带入
	ClientKey uint32 `protobuf:"varint,3,opt,name=clientKey,proto3" json:"clientKey,omitempty"`
	// 计数器的标识
	Counters []*QuotaCounter `protobuf:"bytes,5,rep,name=counters,proto3" json:"counters,omitempty"`
	// 实际滑窗个数
	SlideCount uint32 `protobuf:"varint,6,opt,name=slideCount,proto3" json:"slideCount,omitempty"`
	// 限流server绝对时间，单位ms
	Timestamp            int64    `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitInitResponse) Reset()         { *m = RateLimitInitResponse{} }
func (m *RateLimitInitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitInitResponse) ProtoMessage()    {}
func (*RateLimitInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{3}
}
func (m *RateLimitInitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitInitResponse.Unmarshal(m, b)
}
func (m *RateLimitInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitInitResponse.Marshal(b, m, deterministic)
}
func (dst *RateLimitInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitInitResponse.Merge(dst, src)
}
func (m *RateLimitInitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitInitResponse.Size(m)
}
func (m *RateLimitInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitInitResponse proto.InternalMessageInfo

func (m *RateLimitInitResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RateLimitInitResponse) GetTarget() *LimitTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *RateLimitInitResponse) GetClientKey() uint32 {
	if m != nil {
		return m.ClientKey
	}
	return 0
}

func (m *RateLimitInitResponse) GetCounters() []*QuotaCounter {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *RateLimitInitResponse) GetSlideCount() uint32 {
	if m != nil {
		return m.SlideCount
	}
	return 0
}

func (m *RateLimitInitResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// 限流上报请求
type RateLimitReportRequest struct {
	// 客户端标识
	ClientKey uint32 `protobuf:"varint,1,opt,name=clientKey,proto3" json:"clientKey,omitempty"`
	// 已使用的配额数
	QuotaUses []*QuotaSum `protobuf:"bytes,2,rep,name=quotaUses,proto3" json:"quotaUses,omitempty"`
	// 配额发生的时间，单位ms
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitReportRequest) Reset()         { *m = RateLimitReportRequest{} }
func (m *RateLimitReportRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitReportRequest) ProtoMessage()    {}
func (*RateLimitReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{4}
}
func (m *RateLimitReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitReportRequest.Unmarshal(m, b)
}
func (m *RateLimitReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitReportRequest.Marshal(b, m, deterministic)
}
func (dst *RateLimitReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitReportRequest.Merge(dst, src)
}
func (m *RateLimitReportRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitReportRequest.Size(m)
}
func (m *RateLimitReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitReportRequest proto.InternalMessageInfo

func (m *RateLimitReportRequest) GetClientKey() uint32 {
	if m != nil {
		return m.ClientKey
	}
	return 0
}

func (m *RateLimitReportRequest) GetQuotaUses() []*QuotaSum {
	if m != nil {
		return m.QuotaUses
	}
	return nil
}

func (m *RateLimitReportRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// 限流上报应答
type RateLimitReportResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 剩余配额数
	QuotaLefts []*QuotaLeft `protobuf:"bytes,2,rep,name=quotaLefts,proto3" json:"quotaLefts,omitempty"`
	// 限流server绝对时间，单位ms
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitReportResponse) Reset()         { *m = RateLimitReportResponse{} }
func (m *RateLimitReportResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitReportResponse) ProtoMessage()    {}
func (*RateLimitReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{5}
}
func (m *RateLimitReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitReportResponse.Unmarshal(m, b)
}
func (m *RateLimitReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitReportResponse.Marshal(b, m, deterministic)
}
func (dst *RateLimitReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitReportResponse.Merge(dst, src)
}
func (m *RateLimitReportResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitReportResponse.Size(m)
}
func (m *RateLimitReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitReportResponse proto.InternalMessageInfo

func (m *RateLimitReportResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RateLimitReportResponse) GetQuotaLefts() []*QuotaLeft {
	if m != nil {
		return m.QuotaLefts
	}
	return nil
}

func (m *RateLimitReportResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// 限流目标，针对哪部分数据进行限流
type LimitTarget struct {
	// 命名空间
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 服务名
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// 自定义标签
	Labels               string   `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LimitTarget) Reset()         { *m = LimitTarget{} }
func (m *LimitTarget) String() string { return proto.CompactTextString(m) }
func (*LimitTarget) ProtoMessage()    {}
func (*LimitTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{6}
}
func (m *LimitTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitTarget.Unmarshal(m, b)
}
func (m *LimitTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitTarget.Marshal(b, m, deterministic)
}
func (dst *LimitTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitTarget.Merge(dst, src)
}
func (m *LimitTarget) XXX_Size() int {
	return xxx_messageInfo_LimitTarget.Size(m)
}
func (m *LimitTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitTarget.DiscardUnknown(m)
}

var xxx_messageInfo_LimitTarget proto.InternalMessageInfo

func (m *LimitTarget) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LimitTarget) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *LimitTarget) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

// 阈值配置的值
type QuotaTotal struct {
	// 阈值模式
	Mode QuotaMode `protobuf:"varint,1,opt,name=mode,proto3,enum=polaris.metric.v2.QuotaMode" json:"mode,omitempty"`
	// 单位秒
	Duration uint32 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// 限流阈值
	MaxAmount            uint32   `protobuf:"varint,3,opt,name=maxAmount,proto3" json:"maxAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaTotal) Reset()         { *m = QuotaTotal{} }
func (m *QuotaTotal) String() string { return proto.CompactTextString(m) }
func (*QuotaTotal) ProtoMessage()    {}
func (*QuotaTotal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{7}
}
func (m *QuotaTotal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaTotal.Unmarshal(m, b)
}
func (m *QuotaTotal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaTotal.Marshal(b, m, deterministic)
}
func (dst *QuotaTotal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaTotal.Merge(dst, src)
}
func (m *QuotaTotal) XXX_Size() int {
	return xxx_messageInfo_QuotaTotal.Size(m)
}
func (m *QuotaTotal) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaTotal.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaTotal proto.InternalMessageInfo

func (m *QuotaTotal) GetMode() QuotaMode {
	if m != nil {
		return m.Mode
	}
	return QuotaMode_WHOLE
}

func (m *QuotaTotal) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *QuotaTotal) GetMaxAmount() uint32 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

// 限流计数器
type QuotaCounter struct {
	// 单位秒
	Duration uint32 `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	// bucket的标识，上报时候带入
	CounterKey uint32 `protobuf:"varint,2,opt,name=counterKey,proto3" json:"counterKey,omitempty"`
	// 剩余配额数，应答返回，允许为负数
	Left int64 `protobuf:"varint,3,opt,name=left,proto3" json:"left,omitempty"`
	// 实际限流模式
	Mode Mode `protobuf:"varint,4,opt,name=mode,proto3,enum=polaris.metric.v2.Mode" json:"mode,omitempty"`
	// 接入的客户端数量
	ClientCount          uint32   `protobuf:"varint,5,opt,name=clientCount,proto3" json:"clientCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaCounter) Reset()         { *m = QuotaCounter{} }
func (m *QuotaCounter) String() string { return proto.CompactTextString(m) }
func (*QuotaCounter) ProtoMessage()    {}
func (*QuotaCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{8}
}
func (m *QuotaCounter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaCounter.Unmarshal(m, b)
}
func (m *QuotaCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaCounter.Marshal(b, m, deterministic)
}
func (dst *QuotaCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaCounter.Merge(dst, src)
}
func (m *QuotaCounter) XXX_Size() int {
	return xxx_messageInfo_QuotaCounter.Size(m)
}
func (m *QuotaCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaCounter.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaCounter proto.InternalMessageInfo

func (m *QuotaCounter) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *QuotaCounter) GetCounterKey() uint32 {
	if m != nil {
		return m.CounterKey
	}
	return 0
}

func (m *QuotaCounter) GetLeft() int64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *QuotaCounter) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_ADAPTIVE
}

func (m *QuotaCounter) GetClientCount() uint32 {
	if m != nil {
		return m.ClientCount
	}
	return 0
}

// 客户端阈值使用统计
type QuotaSum struct {
	// 计数器的标识，一个server全局唯一，上报时候带入
	CounterKey uint32 `protobuf:"varint,1,opt,name=counterKey,proto3" json:"counterKey,omitempty"`
	// 已使用的配额数，上报时候带入
	Used uint32 `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	// 被限流数，上报时候带入
	Limited              uint32   `protobuf:"varint,3,opt,name=limited,proto3" json:"limited,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaSum) Reset()         { *m = QuotaSum{} }
func (m *QuotaSum) String() string { return proto.CompactTextString(m) }
func (*QuotaSum) ProtoMessage()    {}
func (*QuotaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{9}
}
func (m *QuotaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaSum.Unmarshal(m, b)
}
func (m *QuotaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaSum.Marshal(b, m, deterministic)
}
func (dst *QuotaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaSum.Merge(dst, src)
}
func (m *QuotaSum) XXX_Size() int {
	return xxx_messageInfo_QuotaSum.Size(m)
}
func (m *QuotaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaSum.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaSum proto.InternalMessageInfo

func (m *QuotaSum) GetCounterKey() uint32 {
	if m != nil {
		return m.CounterKey
	}
	return 0
}

func (m *QuotaSum) GetUsed() uint32 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *QuotaSum) GetLimited() uint32 {
	if m != nil {
		return m.Limited
	}
	return 0
}

// 客户端阈值使用统计，由服务端返回
type QuotaLeft struct {
	// 计数器的标识，一个server全局唯一，上报时候带入
	CounterKey uint32 `protobuf:"varint,1,opt,name=counterKey,proto3" json:"counterKey,omitempty"`
	// 剩余配额数，应答返回，允许为负数
	Left int64 `protobuf:"varint,2,opt,name=left,proto3" json:"left,omitempty"`
	// 当前限流模式
	Mode Mode `protobuf:"varint,3,opt,name=mode,proto3,enum=polaris.metric.v2.Mode" json:"mode,omitempty"`
	// 接入的客户端数量
	ClientCount          uint32   `protobuf:"varint,4,opt,name=clientCount,proto3" json:"clientCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaLeft) Reset()         { *m = QuotaLeft{} }
func (m *QuotaLeft) String() string { return proto.CompactTextString(m) }
func (*QuotaLeft) ProtoMessage()    {}
func (*QuotaLeft) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{10}
}
func (m *QuotaLeft) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaLeft.Unmarshal(m, b)
}
func (m *QuotaLeft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaLeft.Marshal(b, m, deterministic)
}
func (dst *QuotaLeft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaLeft.Merge(dst, src)
}
func (m *QuotaLeft) XXX_Size() int {
	return xxx_messageInfo_QuotaLeft.Size(m)
}
func (m *QuotaLeft) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaLeft.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaLeft proto.InternalMessageInfo

func (m *QuotaLeft) GetCounterKey() uint32 {
	if m != nil {
		return m.CounterKey
	}
	return 0
}

func (m *QuotaLeft) GetLeft() int64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *QuotaLeft) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_ADAPTIVE
}

func (m *QuotaLeft) GetClientCount() uint32 {
	if m != nil {
		return m.ClientCount
	}
	return 0
}

// 时间点对齐的请求
type TimeAdjustRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeAdjustRequest) Reset()         { *m = TimeAdjustRequest{} }
func (m *TimeAdjustRequest) String() string { return proto.CompactTextString(m) }
func (*TimeAdjustRequest) ProtoMessage()    {}
func (*TimeAdjustRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{11}
}
func (m *TimeAdjustRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeAdjustRequest.Unmarshal(m, b)
}
func (m *TimeAdjustRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeAdjustRequest.Marshal(b, m, deterministic)
}
func (dst *TimeAdjustRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeAdjustRequest.Merge(dst, src)
}
func (m *TimeAdjustRequest) XXX_Size() int {
	return xxx_messageInfo_TimeAdjustRequest.Size(m)
}
func (m *TimeAdjustRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeAdjustRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeAdjustRequest proto.InternalMessageInfo

// 时间点对齐的应答
type TimeAdjustResponse struct {
	// 服务器时间点，毫秒
	ServerTimestamp      int64    `protobuf:"varint,1,opt,name=serverTimestamp,proto3" json:"serverTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeAdjustResponse) Reset()         { *m = TimeAdjustResponse{} }
func (m *TimeAdjustResponse) String() string { return proto.CompactTextString(m) }
func (*TimeAdjustResponse) ProtoMessage()    {}
func (*TimeAdjustResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38, []int{12}
}
func (m *TimeAdjustResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeAdjustResponse.Unmarshal(m, b)
}
func (m *TimeAdjustResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeAdjustResponse.Marshal(b, m, deterministic)
}
func (dst *TimeAdjustResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeAdjustResponse.Merge(dst, src)
}
func (m *TimeAdjustResponse) XXX_Size() int {
	return xxx_messageInfo_TimeAdjustResponse.Size(m)
}
func (m *TimeAdjustResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeAdjustResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimeAdjustResponse proto.InternalMessageInfo

func (m *TimeAdjustResponse) GetServerTimestamp() int64 {
	if m != nil {
		return m.ServerTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*RateLimitRequest)(nil), "polaris.metric.v2.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "polaris.metric.v2.RateLimitResponse")
	proto.RegisterType((*RateLimitInitRequest)(nil), "polaris.metric.v2.RateLimitInitRequest")
	proto.RegisterType((*RateLimitInitResponse)(nil), "polaris.metric.v2.RateLimitInitResponse")
	proto.RegisterType((*RateLimitReportRequest)(nil), "polaris.metric.v2.RateLimitReportRequest")
	proto.RegisterType((*RateLimitReportResponse)(nil), "polaris.metric.v2.RateLimitReportResponse")
	proto.RegisterType((*LimitTarget)(nil), "polaris.metric.v2.LimitTarget")
	proto.RegisterType((*QuotaTotal)(nil), "polaris.metric.v2.QuotaTotal")
	proto.RegisterType((*QuotaCounter)(nil), "polaris.metric.v2.QuotaCounter")
	proto.RegisterType((*QuotaSum)(nil), "polaris.metric.v2.QuotaSum")
	proto.RegisterType((*QuotaLeft)(nil), "polaris.metric.v2.QuotaLeft")
	proto.RegisterType((*TimeAdjustRequest)(nil), "polaris.metric.v2.TimeAdjustRequest")
	proto.RegisterType((*TimeAdjustResponse)(nil), "polaris.metric.v2.TimeAdjustResponse")
	proto.RegisterEnum("polaris.metric.v2.RateLimitCmd", RateLimitCmd_name, RateLimitCmd_value)
	proto.RegisterEnum("polaris.metric.v2.Mode", Mode_name, Mode_value)
	proto.RegisterEnum("polaris.metric.v2.QuotaMode", QuotaMode_name, QuotaMode_value)
}

func init() { proto.RegisterFile("ratelimit_v2.proto", fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38) }

var fileDescriptor_ratelimit_v2_3c36bdb8b12a8d38 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xed, 0x4e, 0x13, 0x41,
	0x14, 0x65, 0xb6, 0xdb, 0xd2, 0xde, 0x16, 0x59, 0x46, 0x3e, 0x36, 0x8a, 0xd0, 0x6c, 0x62, 0xac,
	0x35, 0x69, 0xb4, 0x46, 0x8d, 0xd1, 0x98, 0x94, 0xd2, 0x84, 0x8d, 0x28, 0x30, 0x14, 0xfc, 0x8a,
	0x92, 0xa5, 0x3b, 0x98, 0x35, 0xdd, 0x6e, 0xd9, 0x9d, 0x36, 0xf8, 0x00, 0xfe, 0x30, 0x31, 0xfa,
	0x24, 0xbe, 0x9a, 0x4f, 0xe0, 0x0f, 0x33, 0xb3, 0xd3, 0xed, 0xb6, 0xdd, 0x02, 0xf2, 0x6f, 0xe7,
	0x72, 0xef, 0xb9, 0xe7, 0x9e, 0x7b, 0x86, 0x29, 0x60, 0xdf, 0x62, 0xb4, 0xed, 0xb8, 0x0e, 0x3b,
	0xea, 0x57, 0x2b, 0x5d, 0xdf, 0x63, 0x1e, 0x5e, 0xe8, 0x7a, 0x6d, 0xcb, 0x77, 0x82, 0x8a, 0x4b,
	0x99, 0xef, 0xb4, 0x2a, 0xfd, 0xaa, 0xf1, 0x4d, 0x01, 0x8d, 0x58, 0x8c, 0x6e, 0xf3, 0x4c, 0x42,
	0x4f, 0x7b, 0x34, 0x60, 0xf8, 0x01, 0xa4, 0x5a, 0xae, 0xad, 0xa3, 0x22, 0x2a, 0x5d, 0xab, 0xae,
	0x57, 0x26, 0xaa, 0x2a, 0x51, 0x45, 0xdd, 0xb5, 0x09, 0xcf, 0xc5, 0x1f, 0x60, 0xd1, 0x1f, 0x04,
	0xcd, 0x4e, 0x04, 0xa5, 0x2b, 0x45, 0x54, 0xca, 0x57, 0xef, 0x9c, 0x87, 0x11, 0x4b, 0x27, 0x89,
	0x20, 0xd8, 0x82, 0x65, 0x7f, 0xc8, 0xb1, 0xeb, 0xf9, 0x11, 0x7c, 0x4a, 0xc0, 0xdf, 0x3d, 0x0f,
	0x7e, 0xa4, 0x80, 0x4c, 0x01, 0x32, 0x7e, 0x28, 0xb0, 0x10, 0x2b, 0x09, 0xba, 0x5e, 0x27, 0xa0,
	0x57, 0x11, 0xe2, 0x13, 0x2c, 0x8d, 0xcd, 0x10, 0x62, 0x49, 0x25, 0x4a, 0x17, 0x2b, 0x11, 0xe6,
	0x93, 0x64, 0x18, 0x6c, 0xc3, 0xca, 0xc4, 0x08, 0xb2, 0x43, 0x28, 0x46, 0xf9, 0x32, 0x62, 0xc8,
	0x1e, 0xd3, 0xa0, 0x8c, 0x3f, 0x08, 0x16, 0x93, 0x16, 0x84, 0x1f, 0x43, 0x86, 0x59, 0xfe, 0x67,
	0xca, 0x84, 0x28, 0xf9, 0xea, 0x5a, 0x42, 0x37, 0x51, 0xd4, 0x14, 0x59, 0x44, 0x66, 0xe3, 0x1b,
	0x90, 0x6d, 0xb5, 0x1d, 0xda, 0x61, 0xa6, 0x2d, 0x94, 0xc8, 0x91, 0xe8, 0x8c, 0x1f, 0x41, 0x86,
	0x79, 0xcc, 0x6a, 0x07, 0x7a, 0xaa, 0x98, 0x2a, 0xe5, 0xab, 0xb7, 0x12, 0x30, 0xf7, 0x7a, 0x1e,
	0xb3, 0x9a, 0x3c, 0x8b, 0xc8, 0x64, 0xbc, 0x06, 0x10, 0xb4, 0x1d, 0x9b, 0xd6, 0xbd, 0x5e, 0x87,
	0xe9, 0x6a, 0x11, 0x95, 0xe6, 0x48, 0x2c, 0x82, 0xef, 0x81, 0xea, 0x7a, 0x36, 0xd5, 0xd3, 0x62,
	0x7b, 0x2b, 0x09, 0xa0, 0xaf, 0x3c, 0x9b, 0x12, 0x91, 0x64, 0xfc, 0x45, 0xb0, 0x94, 0xb8, 0x07,
	0x8c, 0x41, 0x6d, 0x71, 0x18, 0x24, 0x1a, 0x88, 0xef, 0x98, 0x0a, 0xca, 0x7f, 0xa9, 0xb0, 0x0a,
	0xb9, 0x70, 0xea, 0x97, 0xf4, 0xab, 0x58, 0xd7, 0x1c, 0x19, 0x06, 0xf0, 0x33, 0xc8, 0xb6, 0x38,
	0x73, 0xea, 0x07, 0x7a, 0x5a, 0x28, 0xb1, 0x3e, 0x4d, 0x89, 0x7a, 0x98, 0x47, 0xa2, 0x82, 0x31,
	0x35, 0x32, 0x13, 0x6a, 0xac, 0x42, 0x8e, 0x39, 0x2e, 0x0d, 0x98, 0xe5, 0x76, 0xf5, 0xd9, 0x22,
	0x2a, 0xa5, 0xc8, 0x30, 0x60, 0xfc, 0x42, 0xb0, 0x9c, 0x7c, 0x63, 0x46, 0x39, 0xa3, 0x71, 0xce,
	0x4f, 0x21, 0x77, 0xca, 0x09, 0x1d, 0x04, 0x34, 0xd0, 0x15, 0x41, 0xfa, 0xe6, 0x34, 0xd2, 0xfb,
	0x3d, 0x97, 0x0c, 0xb3, 0x47, 0x19, 0xa5, 0xc6, 0x19, 0x7d, 0x47, 0xb0, 0x32, 0xc5, 0xb6, 0x89,
	0x2b, 0x79, 0x0e, 0x20, 0xa0, 0xb7, 0xe9, 0x09, 0x1b, 0x30, 0x59, 0x9d, 0xc6, 0x84, 0x27, 0x91,
	0x58, 0xfe, 0x05, 0x5c, 0x3e, 0x42, 0x3e, 0xb6, 0x4d, 0x9e, 0xdc, 0xb1, 0x5c, 0x1a, 0x74, 0xad,
	0x56, 0xc8, 0x21, 0x47, 0x86, 0x01, 0xac, 0xc3, 0x6c, 0x40, 0xfd, 0xbe, 0xd3, 0xa2, 0xd2, 0xe8,
	0x83, 0x23, 0x5e, 0x86, 0x4c, 0xdb, 0x3a, 0xa6, 0xc2, 0xe7, 0xfc, 0x0f, 0xf2, 0x64, 0x9c, 0x01,
	0x0c, 0xed, 0x8d, 0xef, 0x4b, 0xdb, 0x86, 0xff, 0x74, 0xa6, 0x8e, 0x30, 0xf4, 0x2e, 0xbf, 0x5b,
	0x76, 0xcf, 0xb7, 0x98, 0xe3, 0x75, 0x44, 0xcb, 0x39, 0x12, 0x9d, 0x39, 0x57, 0xd7, 0x3a, 0xab,
	0xb9, 0xc2, 0x15, 0xd2, 0x71, 0x51, 0xc0, 0xf8, 0x8d, 0xa0, 0x10, 0xf7, 0xd3, 0x08, 0x14, 0x1a,
	0x83, 0x5a, 0x03, 0x90, 0x6e, 0xe3, 0x4e, 0x08, 0x1b, 0xc5, 0x22, 0x7c, 0x2b, 0x6d, 0x7a, 0xc2,
	0xa4, 0x7c, 0xe2, 0x3b, 0xba, 0x83, 0xea, 0x25, 0xee, 0x20, 0x2e, 0x42, 0x3e, 0x34, 0x56, 0xe8,
	0xe1, 0xb4, 0xe8, 0x10, 0x0f, 0x19, 0x6f, 0x21, 0x3b, 0x70, 0xd2, 0x18, 0x1d, 0x94, 0x44, 0xa7,
	0x17, 0x50, 0x5b, 0x12, 0x15, 0xdf, 0x7c, 0x37, 0xe2, 0x49, 0xa4, 0xb6, 0xd4, 0x62, 0x70, 0x34,
	0x7e, 0x22, 0xc8, 0x45, 0xd6, 0xb8, 0x0c, 0xb6, 0x18, 0x55, 0x49, 0x18, 0x35, 0x75, 0x85, 0x51,
	0xd5, 0xc9, 0x51, 0xaf, 0xc3, 0x42, 0xd3, 0x71, 0x69, 0xcd, 0xfe, 0xd2, 0x0b, 0xa2, 0x57, 0xea,
	0x05, 0xe0, 0x78, 0x50, 0x5e, 0x87, 0x12, 0xcc, 0x73, 0x8b, 0x51, 0xbf, 0x19, 0x59, 0x18, 0x09,
	0x62, 0xe3, 0xe1, 0xf2, 0x6d, 0x28, 0xc4, 0x5f, 0x2c, 0x9c, 0x05, 0xd5, 0x7c, 0x6d, 0x36, 0xb5,
	0x19, 0x9c, 0x87, 0xd9, 0x5a, 0x7d, 0xef, 0xc0, 0x24, 0x0d, 0x0d, 0x95, 0x9f, 0x80, 0xca, 0xb9,
	0xe2, 0x02, 0x64, 0x6b, 0x9b, 0xb5, 0xdd, 0xa6, 0x79, 0xd8, 0xd0, 0x66, 0xb0, 0x06, 0x85, 0x8d,
	0x5a, 0xb3, 0xbe, 0x75, 0xb4, 0x53, 0xaf, 0x1f, 0xec, 0xbe, 0xd3, 0x10, 0x9e, 0x87, 0x7c, 0x18,
	0xd9, 0xdf, 0xaa, 0x91, 0x86, 0xa6, 0x94, 0x0d, 0x29, 0xa2, 0xa8, 0xce, 0x41, 0xfa, 0xcd, 0xd6,
	0xce, 0x36, 0x2f, 0x05, 0xc8, 0x6c, 0x9a, 0x87, 0xe6, 0x66, 0x43, 0x43, 0x1b, 0xea, 0x7b, 0xa5,
	0x5f, 0x3d, 0xce, 0x88, 0x5f, 0x24, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x6d, 0x0c,
	0x31, 0xa7, 0x08, 0x00, 0x00,
}
