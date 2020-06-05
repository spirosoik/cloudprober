// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: github.com/google/cloudprober/rds/proto/rds.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IPConfig_IPType int32

const (
	// Default IP of the resource.
	//  - Private IP for instance resource
	//  - Forwarding rule IP for forwarding rule.
	IPConfig_DEFAULT IPConfig_IPType = 0
	// Instance's external IP.
	IPConfig_PUBLIC IPConfig_IPType = 1
	// First IP address from the first Alias IP range. For example, for
	// alias IP range "192.168.12.0/24", 192.168.12.0 will be returned.
	// Supported only on GCE.
	IPConfig_ALIAS IPConfig_IPType = 2
)

// Enum value maps for IPConfig_IPType.
var (
	IPConfig_IPType_name = map[int32]string{
		0: "DEFAULT",
		1: "PUBLIC",
		2: "ALIAS",
	}
	IPConfig_IPType_value = map[string]int32{
		"DEFAULT": 0,
		"PUBLIC":  1,
		"ALIAS":   2,
	}
)

func (x IPConfig_IPType) Enum() *IPConfig_IPType {
	p := new(IPConfig_IPType)
	*p = x
	return p
}

func (x IPConfig_IPType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IPConfig_IPType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_google_cloudprober_rds_proto_rds_proto_enumTypes[0].Descriptor()
}

func (IPConfig_IPType) Type() protoreflect.EnumType {
	return &file_github_com_google_cloudprober_rds_proto_rds_proto_enumTypes[0]
}

func (x IPConfig_IPType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *IPConfig_IPType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = IPConfig_IPType(num)
	return nil
}

// Deprecated: Use IPConfig_IPType.Descriptor instead.
func (IPConfig_IPType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescGZIP(), []int{2, 0}
}

type ListResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provider is the resource list provider, for example: "gcp", "aws", etc.
	Provider *string `protobuf:"bytes,1,req,name=provider" json:"provider,omitempty"`
	// Provider specific resource path. For example: for GCP, it could be
	// "gce_instances/<project>", "regional_forwarding_rules/<project>", etc.
	ResourcePath *string `protobuf:"bytes,2,opt,name=resource_path,json=resourcePath" json:"resource_path,omitempty"`
	// Filters for the resources list. Filters are ANDed: all filters should
	// succeed for an item to included in the result list.
	Filter []*Filter `protobuf:"bytes,3,rep,name=filter" json:"filter,omitempty"`
	// Optional. If resource has an IP (and a NIC) address, following
	// fields determine which IP address will be included in the results.
	IpConfig *IPConfig `protobuf:"bytes,4,opt,name=ip_config,json=ipConfig" json:"ip_config,omitempty"`
}

func (x *ListResourcesRequest) Reset() {
	*x = ListResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcesRequest) ProtoMessage() {}

func (x *ListResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcesRequest.ProtoReflect.Descriptor instead.
func (*ListResourcesRequest) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescGZIP(), []int{0}
}

func (x *ListResourcesRequest) GetProvider() string {
	if x != nil && x.Provider != nil {
		return *x.Provider
	}
	return ""
}

func (x *ListResourcesRequest) GetResourcePath() string {
	if x != nil && x.ResourcePath != nil {
		return *x.ResourcePath
	}
	return ""
}

func (x *ListResourcesRequest) GetFilter() []*Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListResourcesRequest) GetIpConfig() *IPConfig {
	if x != nil {
		return x.IpConfig
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *Filter) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type IPConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NIC index
	NicIndex *int32           `protobuf:"varint,1,opt,name=nic_index,json=nicIndex,def=0" json:"nic_index,omitempty"`
	IpType   *IPConfig_IPType `protobuf:"varint,3,opt,name=ip_type,json=ipType,enum=cloudprober.rds.IPConfig_IPType" json:"ip_type,omitempty"`
}

// Default values for IPConfig fields.
const (
	Default_IPConfig_NicIndex = int32(0)
)

func (x *IPConfig) Reset() {
	*x = IPConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPConfig) ProtoMessage() {}

func (x *IPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPConfig.ProtoReflect.Descriptor instead.
func (*IPConfig) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescGZIP(), []int{2}
}

func (x *IPConfig) GetNicIndex() int32 {
	if x != nil && x.NicIndex != nil {
		return *x.NicIndex
	}
	return Default_IPConfig_NicIndex
}

func (x *IPConfig) GetIpType() IPConfig_IPType {
	if x != nil && x.IpType != nil {
		return *x.IpType
	}
	return IPConfig_DEFAULT
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Resource's IP address, selected based on the request's ip_config.
	Ip *string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	// Resource's port, if any.
	Port *int32 `protobuf:"varint,5,opt,name=port" json:"port,omitempty"`
	// Resource's labels, if any.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Id associated with the resource, if any.
	Id *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	// Optional info associated with the resource. Some resource type may make use
	// of it.
	Info []byte `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescGZIP(), []int{3}
}

func (x *Resource) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Resource) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *Resource) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *Resource) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Resource) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Resource) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

type ListResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*Resource `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
}

func (x *ListResourcesResponse) Reset() {
	*x = ListResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcesResponse) ProtoMessage() {}

func (x *ListResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcesResponse.ProtoReflect.Descriptor instead.
func (*ListResourcesResponse) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescGZIP(), []int{4}
}

func (x *ListResourcesResponse) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_github_com_google_cloudprober_rds_proto_rds_proto protoreflect.FileDescriptor

var file_github_com_google_cloudprober_rds_proto_rds_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x72, 0x64, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2e, 0x72, 0x64, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2f,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x36, 0x0a, 0x09, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2e, 0x72, 0x64, 0x73, 0x2e, 0x49, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x69,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x30, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x08, 0x49, 0x50,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x39, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x49, 0x50, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x49, 0x50, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x69, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x2c, 0x0a, 0x06, 0x49, 0x50, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x49, 0x41, 0x53, 0x10, 0x02, 0x22,
	0xe0, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x50, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x32, 0x75, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x60, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x72, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescOnce sync.Once
	file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescData = file_github_com_google_cloudprober_rds_proto_rds_proto_rawDesc
)

func file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescGZIP() []byte {
	file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescOnce.Do(func() {
		file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescData)
	})
	return file_github_com_google_cloudprober_rds_proto_rds_proto_rawDescData
}

var file_github_com_google_cloudprober_rds_proto_rds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_google_cloudprober_rds_proto_rds_proto_goTypes = []interface{}{
	(IPConfig_IPType)(0),          // 0: cloudprober.rds.IPConfig.IPType
	(*ListResourcesRequest)(nil),  // 1: cloudprober.rds.ListResourcesRequest
	(*Filter)(nil),                // 2: cloudprober.rds.Filter
	(*IPConfig)(nil),              // 3: cloudprober.rds.IPConfig
	(*Resource)(nil),              // 4: cloudprober.rds.Resource
	(*ListResourcesResponse)(nil), // 5: cloudprober.rds.ListResourcesResponse
	nil,                           // 6: cloudprober.rds.Resource.LabelsEntry
}
var file_github_com_google_cloudprober_rds_proto_rds_proto_depIdxs = []int32{
	2, // 0: cloudprober.rds.ListResourcesRequest.filter:type_name -> cloudprober.rds.Filter
	3, // 1: cloudprober.rds.ListResourcesRequest.ip_config:type_name -> cloudprober.rds.IPConfig
	0, // 2: cloudprober.rds.IPConfig.ip_type:type_name -> cloudprober.rds.IPConfig.IPType
	6, // 3: cloudprober.rds.Resource.labels:type_name -> cloudprober.rds.Resource.LabelsEntry
	4, // 4: cloudprober.rds.ListResourcesResponse.resources:type_name -> cloudprober.rds.Resource
	1, // 5: cloudprober.rds.ResourceDiscovery.ListResources:input_type -> cloudprober.rds.ListResourcesRequest
	5, // 6: cloudprober.rds.ResourceDiscovery.ListResources:output_type -> cloudprober.rds.ListResourcesResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_google_cloudprober_rds_proto_rds_proto_init() }
func file_github_com_google_cloudprober_rds_proto_rds_proto_init() {
	if File_github_com_google_cloudprober_rds_proto_rds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcesRequest); i {
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
		file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPConfig); i {
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
		file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcesResponse); i {
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
			RawDescriptor: file_github_com_google_cloudprober_rds_proto_rds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_google_cloudprober_rds_proto_rds_proto_goTypes,
		DependencyIndexes: file_github_com_google_cloudprober_rds_proto_rds_proto_depIdxs,
		EnumInfos:         file_github_com_google_cloudprober_rds_proto_rds_proto_enumTypes,
		MessageInfos:      file_github_com_google_cloudprober_rds_proto_rds_proto_msgTypes,
	}.Build()
	File_github_com_google_cloudprober_rds_proto_rds_proto = out.File
	file_github_com_google_cloudprober_rds_proto_rds_proto_rawDesc = nil
	file_github_com_google_cloudprober_rds_proto_rds_proto_goTypes = nil
	file_github_com_google_cloudprober_rds_proto_rds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResourceDiscoveryClient is the client API for ResourceDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceDiscoveryClient interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
}

type resourceDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceDiscoveryClient(cc grpc.ClientConnInterface) ResourceDiscoveryClient {
	return &resourceDiscoveryClient{cc}
}

func (c *resourceDiscoveryClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.rds.ResourceDiscovery/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceDiscoveryServer is the server API for ResourceDiscovery service.
type ResourceDiscoveryServer interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
}

// UnimplementedResourceDiscoveryServer can be embedded to have forward compatible implementations.
type UnimplementedResourceDiscoveryServer struct {
}

func (*UnimplementedResourceDiscoveryServer) ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}

func RegisterResourceDiscoveryServer(s *grpc.Server, srv ResourceDiscoveryServer) {
	s.RegisterService(&_ResourceDiscovery_serviceDesc, srv)
}

func _ResourceDiscovery_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.rds.ResourceDiscovery/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceDiscovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprober.rds.ResourceDiscovery",
	HandlerType: (*ResourceDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListResources",
			Handler:    _ResourceDiscovery_ListResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/google/cloudprober/rds/proto/rds.proto",
}