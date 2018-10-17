// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/dns/proto/config.proto

package proto

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

// DNS query types from https://en.wikipedia.org/wiki/List_of_DNS_record_types
type QueryType int32

const (
	QueryType_NONE       QueryType = 0
	QueryType_A          QueryType = 1
	QueryType_NS         QueryType = 2
	QueryType_CNAME      QueryType = 5
	QueryType_SOA        QueryType = 6
	QueryType_PTR        QueryType = 12
	QueryType_MX         QueryType = 15
	QueryType_TXT        QueryType = 16
	QueryType_RP         QueryType = 17
	QueryType_AFSDB      QueryType = 18
	QueryType_SIG        QueryType = 24
	QueryType_KEY        QueryType = 25
	QueryType_AAAA       QueryType = 28
	QueryType_LOC        QueryType = 29
	QueryType_SRV        QueryType = 33
	QueryType_NAPTR      QueryType = 35
	QueryType_KX         QueryType = 36
	QueryType_CERT       QueryType = 37
	QueryType_DNAME      QueryType = 39
	QueryType_APL        QueryType = 42
	QueryType_DS         QueryType = 43
	QueryType_SSHFP      QueryType = 44
	QueryType_IPSECKEY   QueryType = 45
	QueryType_RRSIG      QueryType = 46
	QueryType_NSEC       QueryType = 47
	QueryType_DNSKEY     QueryType = 48
	QueryType_DHCID      QueryType = 49
	QueryType_NSEC3      QueryType = 50
	QueryType_NSEC3PARAM QueryType = 51
	QueryType_TLSA       QueryType = 52
	QueryType_HIP        QueryType = 55
	QueryType_CDS        QueryType = 59
	QueryType_CDNSKEY    QueryType = 60
	QueryType_OPENPGPKEY QueryType = 61
	QueryType_TKEY       QueryType = 249
	QueryType_TSIG       QueryType = 250
	QueryType_URI        QueryType = 256
	QueryType_CAA        QueryType = 257
	QueryType_TA         QueryType = 32768
	QueryType_DLV        QueryType = 32769
)

var QueryType_name = map[int32]string{
	0:     "NONE",
	1:     "A",
	2:     "NS",
	5:     "CNAME",
	6:     "SOA",
	12:    "PTR",
	15:    "MX",
	16:    "TXT",
	17:    "RP",
	18:    "AFSDB",
	24:    "SIG",
	25:    "KEY",
	28:    "AAAA",
	29:    "LOC",
	33:    "SRV",
	35:    "NAPTR",
	36:    "KX",
	37:    "CERT",
	39:    "DNAME",
	42:    "APL",
	43:    "DS",
	44:    "SSHFP",
	45:    "IPSECKEY",
	46:    "RRSIG",
	47:    "NSEC",
	48:    "DNSKEY",
	49:    "DHCID",
	50:    "NSEC3",
	51:    "NSEC3PARAM",
	52:    "TLSA",
	55:    "HIP",
	59:    "CDS",
	60:    "CDNSKEY",
	61:    "OPENPGPKEY",
	249:   "TKEY",
	250:   "TSIG",
	256:   "URI",
	257:   "CAA",
	32768: "TA",
	32769: "DLV",
}
var QueryType_value = map[string]int32{
	"NONE":       0,
	"A":          1,
	"NS":         2,
	"CNAME":      5,
	"SOA":        6,
	"PTR":        12,
	"MX":         15,
	"TXT":        16,
	"RP":         17,
	"AFSDB":      18,
	"SIG":        24,
	"KEY":        25,
	"AAAA":       28,
	"LOC":        29,
	"SRV":        33,
	"NAPTR":      35,
	"KX":         36,
	"CERT":       37,
	"DNAME":      39,
	"APL":        42,
	"DS":         43,
	"SSHFP":      44,
	"IPSECKEY":   45,
	"RRSIG":      46,
	"NSEC":       47,
	"DNSKEY":     48,
	"DHCID":      49,
	"NSEC3":      50,
	"NSEC3PARAM": 51,
	"TLSA":       52,
	"HIP":        55,
	"CDS":        59,
	"CDNSKEY":    60,
	"OPENPGPKEY": 61,
	"TKEY":       249,
	"TSIG":       250,
	"URI":        256,
	"CAA":        257,
	"TA":         32768,
	"DLV":        32769,
}

func (x QueryType) Enum() *QueryType {
	p := new(QueryType)
	*p = x
	return p
}
func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}
func (x *QueryType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QueryType_value, data, "QueryType")
	if err != nil {
		return err
	}
	*x = QueryType(value)
	return nil
}
func (QueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_d89b08fc73ea537a, []int{0}
}

type ProbeConf struct {
	// Domain to use when making DNS queries
	ResolvedDomain *string `protobuf:"bytes,1,opt,name=resolved_domain,json=resolvedDomain,def=www.google.com." json:"resolved_domain,omitempty"`
	// Export stats after these many milliseconds
	StatsExportIntervalMsec *int32 `protobuf:"varint,2,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	// DNS Query Type
	QueryType *QueryType `protobuf:"varint,3,opt,name=query_type,json=queryType,enum=cloudprober.probes.dns.QueryType,def=15" json:"query_type,omitempty"`
	// Minimum number of answers expected. Default behavior is to return success
	// if DNS response status is NOERROR.
	MinAnswers           *uint32  `protobuf:"varint,4,opt,name=min_answers,json=minAnswers,def=0" json:"min_answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeConf) Reset()         { *m = ProbeConf{} }
func (m *ProbeConf) String() string { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()    {}
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_d89b08fc73ea537a, []int{0}
}
func (m *ProbeConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeConf.Unmarshal(m, b)
}
func (m *ProbeConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeConf.Marshal(b, m, deterministic)
}
func (dst *ProbeConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeConf.Merge(dst, src)
}
func (m *ProbeConf) XXX_Size() int {
	return xxx_messageInfo_ProbeConf.Size(m)
}
func (m *ProbeConf) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeConf.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeConf proto.InternalMessageInfo

const Default_ProbeConf_ResolvedDomain string = "www.google.com."
const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000
const Default_ProbeConf_QueryType QueryType = QueryType_MX
const Default_ProbeConf_MinAnswers uint32 = 0

func (m *ProbeConf) GetResolvedDomain() string {
	if m != nil && m.ResolvedDomain != nil {
		return *m.ResolvedDomain
	}
	return Default_ProbeConf_ResolvedDomain
}

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

func (m *ProbeConf) GetQueryType() QueryType {
	if m != nil && m.QueryType != nil {
		return *m.QueryType
	}
	return Default_ProbeConf_QueryType
}

func (m *ProbeConf) GetMinAnswers() uint32 {
	if m != nil && m.MinAnswers != nil {
		return *m.MinAnswers
	}
	return Default_ProbeConf_MinAnswers
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.dns.ProbeConf")
	proto.RegisterEnum("cloudprober.probes.dns.QueryType", QueryType_name, QueryType_value)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/dns/proto/config.proto", fileDescriptor_config_d89b08fc73ea537a)
}

var fileDescriptor_config_d89b08fc73ea537a = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x53, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0x4a, 0x81, 0x2c, 0x08, 0xcf, 0x3d, 0x68, 0x75, 0x74, 0xa6, 0xa0, 0x8e, 0x1d,
	0xd4, 0x24, 0x80, 0x33, 0x3a, 0x45, 0x0f, 0x4b, 0x12, 0x20, 0x43, 0x9b, 0xc6, 0xdd, 0xc8, 0xd4,
	0x53, 0xa6, 0xb4, 0x4b, 0xcd, 0x4c, 0x9b, 0x94, 0x24, 0x50, 0x7b, 0x2b, 0xff, 0x8b, 0xff, 0xa1,
	0x27, 0x3d, 0x39, 0x6f, 0x0b, 0x8e, 0x07, 0x4f, 0xf9, 0xde, 0x7b, 0xbf, 0xef, 0x7d, 0x2f, 0x4b,
	0x0e, 0x86, 0x49, 0xf9, 0xed, 0xea, 0xdc, 0xec, 0x67, 0x63, 0x6b, 0x98, 0x65, 0xc3, 0x91, 0xb4,
	0xfa, 0xa3, 0xec, 0x6a, 0x30, 0xc9, 0xb3, 0x73, 0x99, 0x5b, 0xea, 0x53, 0x58, 0x83, 0xb4, 0x40,
	0x59, 0x66, 0x56, 0x3f, 0x4b, 0x2f, 0x92, 0xa1, 0xa9, 0x0a, 0xfa, 0xf0, 0x1f, 0xd4, 0x5c, 0xa0,
	0xe6, 0x20, 0x2d, 0xb6, 0x7f, 0x6a, 0xc4, 0x08, 0xb1, 0x74, 0xb2, 0xf4, 0x82, 0x7e, 0x20, 0x9b,
	0xb9, 0x2c, 0xb2, 0xd1, 0xb5, 0x1c, 0xc4, 0x83, 0x6c, 0xdc, 0x4b, 0xd2, 0x9a, 0x56, 0xd7, 0x1a,
	0x46, 0x73, 0x73, 0x3a, 0x9d, 0x9a, 0x8b, 0x50, 0xcc, 0x37, 0xf9, 0xc6, 0x1d, 0xe7, 0x2a, 0x8c,
	0x1e, 0x92, 0x27, 0x45, 0xd9, 0x2b, 0x8b, 0x58, 0x7e, 0x9f, 0x64, 0x79, 0x19, 0x27, 0x69, 0x29,
	0xf3, 0xeb, 0xde, 0x28, 0x1e, 0x17, 0xb2, 0x5f, 0xd3, 0xeb, 0x5a, 0xa3, 0xda, 0xac, 0xee, 0xda,
	0xb6, 0x6d, 0xf3, 0x47, 0x0a, 0xf4, 0x14, 0xe7, 0xdf, 0x62, 0xed, 0x42, 0xf6, 0xa9, 0x4b, 0xc8,
	0xe5, 0x95, 0xcc, 0x67, 0x71, 0x39, 0x9b, 0xc8, 0x5a, 0xa5, 0xae, 0x35, 0x36, 0xf6, 0xb6, 0xcc,
	0xff, 0x1f, 0x6e, 0x7e, 0x46, 0x32, 0x9a, 0x4d, 0x64, 0x53, 0x6f, 0x77, 0xb9, 0x71, 0x79, 0x57,
	0xd2, 0x6d, 0xb2, 0x36, 0x4e, 0xd2, 0xb8, 0x97, 0x16, 0x53, 0x99, 0x17, 0xb5, 0xa5, 0xba, 0xd6,
	0xb8, 0xdf, 0xd4, 0x6c, 0x4e, 0xc6, 0x49, 0xca, 0x16, 0xcd, 0x9d, 0x1f, 0x15, 0x62, 0xfc, 0x5d,
	0x40, 0x57, 0xc9, 0x52, 0xd0, 0x09, 0x3c, 0xb8, 0x47, 0xab, 0x44, 0x63, 0xa0, 0xd1, 0x65, 0xa2,
	0x07, 0x02, 0x74, 0x6a, 0x90, 0xaa, 0x13, 0xb0, 0xb6, 0x07, 0x55, 0xba, 0x42, 0x2a, 0xa2, 0xc3,
	0x60, 0x19, 0x45, 0x18, 0x71, 0x58, 0x47, 0xa8, 0xdd, 0x85, 0x4d, 0x6c, 0x44, 0xdd, 0x08, 0x00,
	0x1b, 0x3c, 0x84, 0x07, 0xe8, 0x62, 0x47, 0xc2, 0x3d, 0x04, 0xaa, 0x5c, 0xfe, 0x31, 0xd4, 0x50,
	0x9c, 0x7a, 0x5f, 0xe1, 0x31, 0x66, 0x31, 0xc6, 0x18, 0x3c, 0xc5, 0x56, 0xab, 0xe3, 0xc0, 0x33,
	0x05, 0xf1, 0x33, 0xd8, 0x42, 0x63, 0xc0, 0x70, 0xf9, 0x73, 0xdc, 0x75, 0xda, 0x85, 0x17, 0x88,
	0x3b, 0x1e, 0x8f, 0xe0, 0x25, 0x0e, 0x5d, 0x75, 0xcb, 0x2b, 0x34, 0xb0, 0xb0, 0x05, 0x3b, 0x48,
	0xb9, 0x02, 0x5e, 0xe3, 0x4c, 0x88, 0x93, 0xa3, 0x10, 0xde, 0xd0, 0x75, 0xb2, 0xea, 0x87, 0xc2,
	0x73, 0x30, 0xed, 0x2d, 0x0e, 0x38, 0xc7, 0x0b, 0x4c, 0xf5, 0x93, 0xc2, 0x73, 0xc0, 0xa2, 0x84,
	0x2c, 0xbb, 0x81, 0x40, 0xc0, 0x56, 0x5b, 0x4f, 0x1c, 0xdf, 0x85, 0x5d, 0x95, 0x2e, 0x3c, 0x67,
	0x1f, 0xf6, 0xe8, 0x06, 0x21, 0x4a, 0x86, 0x8c, 0xb3, 0x36, 0xec, 0xa3, 0x37, 0x6a, 0x09, 0x06,
	0xef, 0x30, 0xfa, 0xc4, 0x0f, 0xe1, 0x3d, 0x0a, 0xc7, 0x15, 0x70, 0x40, 0xd7, 0xc8, 0x8a, 0x73,
	0xbb, 0xee, 0x23, 0x1a, 0x3b, 0xa1, 0x17, 0x84, 0xc7, 0x21, 0xd6, 0x9f, 0xa8, 0x41, 0x96, 0x22,
	0x54, 0xbf, 0x34, 0x25, 0xf1, 0x92, 0xdf, 0x1a, 0x5d, 0x25, 0x95, 0x2f, 0xdc, 0x87, 0xb9, 0x8e,
	0xca, 0x61, 0x0c, 0x6e, 0x50, 0xe9, 0x11, 0x83, 0xf9, 0x1c, 0x1f, 0xbd, 0xe2, 0xb6, 0xce, 0xe0,
	0x66, 0xae, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x00, 0x3c, 0x6c, 0xf2, 0x02, 0x00, 0x00,
}
