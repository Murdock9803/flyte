// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: flyteidl/service/dataproxy.proto

package service

import (
	core "github.com/flyteorg/flyte/gen/pb-go/flyteidl/core"
	_ "github.com/flyteorg/flyte/gen/pb-go/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ArtifactType
type ArtifactType int32

const (
	// ARTIFACT_TYPE_UNDEFINED is the default, often invalid, value for the enum.
	ArtifactType_ARTIFACT_TYPE_UNDEFINED ArtifactType = 0
	// ARTIFACT_TYPE_DECK refers to the deck html file optionally generated after a task, a workflow or a launch plan
	// finishes executing.
	ArtifactType_ARTIFACT_TYPE_DECK ArtifactType = 1
)

// Enum value maps for ArtifactType.
var (
	ArtifactType_name = map[int32]string{
		0: "ARTIFACT_TYPE_UNDEFINED",
		1: "ARTIFACT_TYPE_DECK",
	}
	ArtifactType_value = map[string]int32{
		"ARTIFACT_TYPE_UNDEFINED": 0,
		"ARTIFACT_TYPE_DECK":      1,
	}
)

func (x ArtifactType) Enum() *ArtifactType {
	p := new(ArtifactType)
	*p = x
	return p
}

func (x ArtifactType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArtifactType) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_service_dataproxy_proto_enumTypes[0].Descriptor()
}

func (ArtifactType) Type() protoreflect.EnumType {
	return &file_flyteidl_service_dataproxy_proto_enumTypes[0]
}

func (x ArtifactType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArtifactType.Descriptor instead.
func (ArtifactType) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_service_dataproxy_proto_rawDescGZIP(), []int{0}
}

type CreateUploadLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SignedUrl specifies the url to use to upload content to (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
	SignedUrl string `protobuf:"bytes,1,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
	// NativeUrl specifies the url in the format of the configured storage provider (e.g. s3://my-bucket/randomstring/suffix.tar)
	NativeUrl string `protobuf:"bytes,2,opt,name=native_url,json=nativeUrl,proto3" json:"native_url,omitempty"`
	// ExpiresAt defines when will the signed URL expires.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *CreateUploadLocationResponse) Reset() {
	*x = CreateUploadLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_dataproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUploadLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUploadLocationResponse) ProtoMessage() {}

func (x *CreateUploadLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_dataproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUploadLocationResponse.ProtoReflect.Descriptor instead.
func (*CreateUploadLocationResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_dataproxy_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUploadLocationResponse) GetSignedUrl() string {
	if x != nil {
		return x.SignedUrl
	}
	return ""
}

func (x *CreateUploadLocationResponse) GetNativeUrl() string {
	if x != nil {
		return x.NativeUrl
	}
	return ""
}

func (x *CreateUploadLocationResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

// CreateUploadLocationRequest specified request for the CreateUploadLocation API.
type CreateUploadLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project to create the upload location for
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Domain to create the upload location for.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Filename specifies a desired suffix for the generated location. E.g. `file.py` or `pre/fix/file.zip`.
	// +optional. By default, the service will generate a consistent name based on the provided parameters.
	Filename string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	// ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this
	// exceeds the platform allowed max.
	// +optional. The default value comes from a global config.
	ExpiresIn *durationpb.Duration `protobuf:"bytes,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	// ContentMD5 restricts the upload location to the specific MD5 provided. The ContentMD5 will also appear in the
	// generated path.
	// +required
	ContentMd5 []byte `protobuf:"bytes,5,opt,name=content_md5,json=contentMd5,proto3" json:"content_md5,omitempty"`
}

func (x *CreateUploadLocationRequest) Reset() {
	*x = CreateUploadLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_dataproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUploadLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUploadLocationRequest) ProtoMessage() {}

func (x *CreateUploadLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_dataproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUploadLocationRequest.ProtoReflect.Descriptor instead.
func (*CreateUploadLocationRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_dataproxy_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUploadLocationRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *CreateUploadLocationRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateUploadLocationRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *CreateUploadLocationRequest) GetExpiresIn() *durationpb.Duration {
	if x != nil {
		return x.ExpiresIn
	}
	return nil
}

func (x *CreateUploadLocationRequest) GetContentMd5() []byte {
	if x != nil {
		return x.ContentMd5
	}
	return nil
}

// CreateDownloadLocationRequest specified request for the CreateDownloadLocation API.
//
// Deprecated: Do not use.
type CreateDownloadLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NativeUrl specifies the url in the format of the configured storage provider (e.g. s3://my-bucket/randomstring/suffix.tar)
	NativeUrl string `protobuf:"bytes,1,opt,name=native_url,json=nativeUrl,proto3" json:"native_url,omitempty"`
	// ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this
	// exceeds the platform allowed max.
	// +optional. The default value comes from a global config.
	ExpiresIn *durationpb.Duration `protobuf:"bytes,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *CreateDownloadLocationRequest) Reset() {
	*x = CreateDownloadLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_dataproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDownloadLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadLocationRequest) ProtoMessage() {}

func (x *CreateDownloadLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_dataproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadLocationRequest.ProtoReflect.Descriptor instead.
func (*CreateDownloadLocationRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_dataproxy_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDownloadLocationRequest) GetNativeUrl() string {
	if x != nil {
		return x.NativeUrl
	}
	return ""
}

func (x *CreateDownloadLocationRequest) GetExpiresIn() *durationpb.Duration {
	if x != nil {
		return x.ExpiresIn
	}
	return nil
}

// Deprecated: Do not use.
type CreateDownloadLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SignedUrl specifies the url to use to download content from (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
	SignedUrl string `protobuf:"bytes,1,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
	// ExpiresAt defines when will the signed URL expires.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *CreateDownloadLocationResponse) Reset() {
	*x = CreateDownloadLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_dataproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDownloadLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadLocationResponse) ProtoMessage() {}

func (x *CreateDownloadLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_dataproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadLocationResponse.ProtoReflect.Descriptor instead.
func (*CreateDownloadLocationResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_dataproxy_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDownloadLocationResponse) GetSignedUrl() string {
	if x != nil {
		return x.SignedUrl
	}
	return ""
}

func (x *CreateDownloadLocationResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

// CreateDownloadLinkRequest defines the request parameters to create a download link (signed url)
type CreateDownloadLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ArtifactType of the artifact requested.
	ArtifactType ArtifactType `protobuf:"varint,1,opt,name=artifact_type,json=artifactType,proto3,enum=flyteidl.service.ArtifactType" json:"artifact_type,omitempty"`
	// ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this
	// exceeds the platform allowed max.
	// +optional. The default value comes from a global config.
	ExpiresIn *durationpb.Duration `protobuf:"bytes,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	// Types that are assignable to Source:
	//
	//	*CreateDownloadLinkRequest_NodeExecutionId
	Source isCreateDownloadLinkRequest_Source `protobuf_oneof:"source"`
}

func (x *CreateDownloadLinkRequest) Reset() {
	*x = CreateDownloadLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_dataproxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDownloadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadLinkRequest) ProtoMessage() {}

func (x *CreateDownloadLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_dataproxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateDownloadLinkRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_dataproxy_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDownloadLinkRequest) GetArtifactType() ArtifactType {
	if x != nil {
		return x.ArtifactType
	}
	return ArtifactType_ARTIFACT_TYPE_UNDEFINED
}

func (x *CreateDownloadLinkRequest) GetExpiresIn() *durationpb.Duration {
	if x != nil {
		return x.ExpiresIn
	}
	return nil
}

func (m *CreateDownloadLinkRequest) GetSource() isCreateDownloadLinkRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *CreateDownloadLinkRequest) GetNodeExecutionId() *core.NodeExecutionIdentifier {
	if x, ok := x.GetSource().(*CreateDownloadLinkRequest_NodeExecutionId); ok {
		return x.NodeExecutionId
	}
	return nil
}

type isCreateDownloadLinkRequest_Source interface {
	isCreateDownloadLinkRequest_Source()
}

type CreateDownloadLinkRequest_NodeExecutionId struct {
	// NodeId is the unique identifier for the node execution. For a task node, this will retrieve the output of the
	// most recent attempt of the task.
	NodeExecutionId *core.NodeExecutionIdentifier `protobuf:"bytes,3,opt,name=node_execution_id,json=nodeExecutionId,proto3,oneof"`
}

func (*CreateDownloadLinkRequest_NodeExecutionId) isCreateDownloadLinkRequest_Source() {}

// CreateDownloadLinkResponse defines the response for the generated links
type CreateDownloadLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SignedUrl specifies the url to use to download content from (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
	SignedUrl []string `protobuf:"bytes,1,rep,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
	// ExpiresAt defines when will the signed URL expire.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *CreateDownloadLinkResponse) Reset() {
	*x = CreateDownloadLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_dataproxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDownloadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDownloadLinkResponse) ProtoMessage() {}

func (x *CreateDownloadLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_dataproxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDownloadLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateDownloadLinkResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_dataproxy_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDownloadLinkResponse) GetSignedUrl() []string {
	if x != nil {
		return x.SignedUrl
	}
	return nil
}

func (x *CreateDownloadLinkResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

var File_flyteidl_service_dataproxy_proto protoreflect.FileDescriptor

var file_flyteidl_service_dataproxy_proto_rawDesc = []byte{
	0x0a, 0x20, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x97, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x49, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x64, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4d, 0x64, 0x35, 0x22, 0x7c, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x3a, 0x02,
	0x18, 0x01, 0x22, 0x7e, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x3a, 0x02,
	0x18, 0x01, 0x22, 0xfa, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x43, 0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12,
	0x54, 0x0a, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x76, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x2a, 0x43, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x52, 0x54, 0x49, 0x46,
	0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x52, 0x54, 0x49, 0x46, 0x41, 0x43, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4b, 0x10, 0x01, 0x32, 0x9e, 0x06, 0x0a,
	0x10, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xf0, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x92, 0x41, 0x4d, 0x1a, 0x4b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2d,
	0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x68, 0x74, 0x74, 0x70, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x69, 0x73, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x20,
	0x61, 0x74, 0x20, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x6e, 0x12, 0xa9, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xab, 0x01, 0x88, 0x02, 0x01, 0x92, 0x41, 0x7f, 0x1a, 0x7d, 0x44, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x20, 0x50, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x65, 0x61, 0x64, 0x2e, 0x20,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x72, 0x65, 0x61, 0x64, 0x2d, 0x6f,
	0x6e, 0x6c, 0x79, 0x20, 0x68, 0x74, 0x74, 0x70, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x69, 0x73, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x20, 0x61,
	0x74, 0x20, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6e,
	0x12, 0xea, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x79, 0x92, 0x41, 0x4c, 0x1a, 0x4a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73,
	0x20, 0x61, 0x20, 0x72, 0x65, 0x61, 0x64, 0x2d, 0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x68, 0x74, 0x74,
	0x70, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20,
	0x69, 0x73, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x20, 0x61, 0x74, 0x20, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0xbf, 0x01,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0xa2, 0x02, 0x03, 0x46, 0x53, 0x58, 0xaa, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x1c, 0x46,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x46, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_service_dataproxy_proto_rawDescOnce sync.Once
	file_flyteidl_service_dataproxy_proto_rawDescData = file_flyteidl_service_dataproxy_proto_rawDesc
)

func file_flyteidl_service_dataproxy_proto_rawDescGZIP() []byte {
	file_flyteidl_service_dataproxy_proto_rawDescOnce.Do(func() {
		file_flyteidl_service_dataproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_service_dataproxy_proto_rawDescData)
	})
	return file_flyteidl_service_dataproxy_proto_rawDescData
}

var file_flyteidl_service_dataproxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flyteidl_service_dataproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_flyteidl_service_dataproxy_proto_goTypes = []interface{}{
	(ArtifactType)(0),                      // 0: flyteidl.service.ArtifactType
	(*CreateUploadLocationResponse)(nil),   // 1: flyteidl.service.CreateUploadLocationResponse
	(*CreateUploadLocationRequest)(nil),    // 2: flyteidl.service.CreateUploadLocationRequest
	(*CreateDownloadLocationRequest)(nil),  // 3: flyteidl.service.CreateDownloadLocationRequest
	(*CreateDownloadLocationResponse)(nil), // 4: flyteidl.service.CreateDownloadLocationResponse
	(*CreateDownloadLinkRequest)(nil),      // 5: flyteidl.service.CreateDownloadLinkRequest
	(*CreateDownloadLinkResponse)(nil),     // 6: flyteidl.service.CreateDownloadLinkResponse
	(*timestamppb.Timestamp)(nil),          // 7: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),            // 8: google.protobuf.Duration
	(*core.NodeExecutionIdentifier)(nil),   // 9: flyteidl.core.NodeExecutionIdentifier
}
var file_flyteidl_service_dataproxy_proto_depIdxs = []int32{
	7,  // 0: flyteidl.service.CreateUploadLocationResponse.expires_at:type_name -> google.protobuf.Timestamp
	8,  // 1: flyteidl.service.CreateUploadLocationRequest.expires_in:type_name -> google.protobuf.Duration
	8,  // 2: flyteidl.service.CreateDownloadLocationRequest.expires_in:type_name -> google.protobuf.Duration
	7,  // 3: flyteidl.service.CreateDownloadLocationResponse.expires_at:type_name -> google.protobuf.Timestamp
	0,  // 4: flyteidl.service.CreateDownloadLinkRequest.artifact_type:type_name -> flyteidl.service.ArtifactType
	8,  // 5: flyteidl.service.CreateDownloadLinkRequest.expires_in:type_name -> google.protobuf.Duration
	9,  // 6: flyteidl.service.CreateDownloadLinkRequest.node_execution_id:type_name -> flyteidl.core.NodeExecutionIdentifier
	7,  // 7: flyteidl.service.CreateDownloadLinkResponse.expires_at:type_name -> google.protobuf.Timestamp
	2,  // 8: flyteidl.service.DataProxyService.CreateUploadLocation:input_type -> flyteidl.service.CreateUploadLocationRequest
	3,  // 9: flyteidl.service.DataProxyService.CreateDownloadLocation:input_type -> flyteidl.service.CreateDownloadLocationRequest
	5,  // 10: flyteidl.service.DataProxyService.CreateDownloadLink:input_type -> flyteidl.service.CreateDownloadLinkRequest
	1,  // 11: flyteidl.service.DataProxyService.CreateUploadLocation:output_type -> flyteidl.service.CreateUploadLocationResponse
	4,  // 12: flyteidl.service.DataProxyService.CreateDownloadLocation:output_type -> flyteidl.service.CreateDownloadLocationResponse
	6,  // 13: flyteidl.service.DataProxyService.CreateDownloadLink:output_type -> flyteidl.service.CreateDownloadLinkResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_flyteidl_service_dataproxy_proto_init() }
func file_flyteidl_service_dataproxy_proto_init() {
	if File_flyteidl_service_dataproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_service_dataproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUploadLocationResponse); i {
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
		file_flyteidl_service_dataproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUploadLocationRequest); i {
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
		file_flyteidl_service_dataproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDownloadLocationRequest); i {
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
		file_flyteidl_service_dataproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDownloadLocationResponse); i {
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
		file_flyteidl_service_dataproxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDownloadLinkRequest); i {
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
		file_flyteidl_service_dataproxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDownloadLinkResponse); i {
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
	file_flyteidl_service_dataproxy_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*CreateDownloadLinkRequest_NodeExecutionId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_service_dataproxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flyteidl_service_dataproxy_proto_goTypes,
		DependencyIndexes: file_flyteidl_service_dataproxy_proto_depIdxs,
		EnumInfos:         file_flyteidl_service_dataproxy_proto_enumTypes,
		MessageInfos:      file_flyteidl_service_dataproxy_proto_msgTypes,
	}.Build()
	File_flyteidl_service_dataproxy_proto = out.File
	file_flyteidl_service_dataproxy_proto_rawDesc = nil
	file_flyteidl_service_dataproxy_proto_goTypes = nil
	file_flyteidl_service_dataproxy_proto_depIdxs = nil
}