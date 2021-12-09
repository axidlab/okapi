// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: okapi/transport/v1/transport.proto

package okapiproto

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

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  []byte         `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Key      *JsonWebKey    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	AppendTo *SignedMessage `protobuf:"bytes,3,opt,name=append_to,json=appendTo,proto3" json:"append_to,omitempty"`
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{0}
}

func (x *SignRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SignRequest) GetKey() *JsonWebKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SignRequest) GetAppendTo() *SignedMessage {
	if x != nil {
		return x.AppendTo
	}
	return nil
}

type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *SignedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{1}
}

func (x *SignResponse) GetMessage() *SignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *SignedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Key     *JsonWebKey    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyRequest) GetMessage() *SignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *VerifyRequest) GetKey() *JsonWebKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type PackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderKey      *JsonWebKey         `protobuf:"bytes,1,opt,name=sender_key,json=senderKey,proto3" json:"sender_key,omitempty"`
	ReceiverKey    *JsonWebKey         `protobuf:"bytes,2,opt,name=receiver_key,json=receiverKey,proto3" json:"receiver_key,omitempty"`
	AssociatedData []byte              `protobuf:"bytes,3,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
	Plaintext      []byte              `protobuf:"bytes,4,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	Mode           EncryptionMode      `protobuf:"varint,5,opt,name=mode,proto3,enum=pbmse.v1.EncryptionMode" json:"mode,omitempty"`
	Algorithm      EncryptionAlgorithm `protobuf:"varint,6,opt,name=algorithm,proto3,enum=pbmse.v1.EncryptionAlgorithm" json:"algorithm,omitempty"`
}

func (x *PackRequest) Reset() {
	*x = PackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackRequest) ProtoMessage() {}

func (x *PackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackRequest.ProtoReflect.Descriptor instead.
func (*PackRequest) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{4}
}

func (x *PackRequest) GetSenderKey() *JsonWebKey {
	if x != nil {
		return x.SenderKey
	}
	return nil
}

func (x *PackRequest) GetReceiverKey() *JsonWebKey {
	if x != nil {
		return x.ReceiverKey
	}
	return nil
}

func (x *PackRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

func (x *PackRequest) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

func (x *PackRequest) GetMode() EncryptionMode {
	if x != nil {
		return x.Mode
	}
	return EncryptionMode_ENCRYPTION_MODE_UNSPECIFIED
}

func (x *PackRequest) GetAlgorithm() EncryptionAlgorithm {
	if x != nil {
		return x.Algorithm
	}
	return EncryptionAlgorithm_ENCRYPTION_ALGORITHM_UNSPECIFIED
}

type PackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *EncryptedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PackResponse) Reset() {
	*x = PackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackResponse) ProtoMessage() {}

func (x *PackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackResponse.ProtoReflect.Descriptor instead.
func (*PackResponse) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{5}
}

func (x *PackResponse) GetMessage() *EncryptedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type UnpackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderKey   *JsonWebKey       `protobuf:"bytes,1,opt,name=sender_key,json=senderKey,proto3" json:"sender_key,omitempty"`
	ReceiverKey *JsonWebKey       `protobuf:"bytes,2,opt,name=receiver_key,json=receiverKey,proto3" json:"receiver_key,omitempty"`
	Message     *EncryptedMessage `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UnpackRequest) Reset() {
	*x = UnpackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnpackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnpackRequest) ProtoMessage() {}

func (x *UnpackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnpackRequest.ProtoReflect.Descriptor instead.
func (*UnpackRequest) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{6}
}

func (x *UnpackRequest) GetSenderKey() *JsonWebKey {
	if x != nil {
		return x.SenderKey
	}
	return nil
}

func (x *UnpackRequest) GetReceiverKey() *JsonWebKey {
	if x != nil {
		return x.ReceiverKey
	}
	return nil
}

func (x *UnpackRequest) GetMessage() *EncryptedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type UnpackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plaintext []byte `protobuf:"bytes,1,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
}

func (x *UnpackResponse) Reset() {
	*x = UnpackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnpackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnpackResponse) ProtoMessage() {}

func (x *UnpackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnpackResponse.ProtoReflect.Descriptor instead.
func (*UnpackResponse) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{7}
}

func (x *UnpackResponse) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

type CoreMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type    string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Body    []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	To      []string `protobuf:"bytes,4,rep,name=to,proto3" json:"to,omitempty"`
	From    string   `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Created int64    `protobuf:"varint,6,opt,name=created,json=created_time,proto3" json:"created,omitempty"`
	Expires int64    `protobuf:"varint,7,opt,name=expires,json=expires_time,proto3" json:"expires,omitempty"`
}

func (x *CoreMessage) Reset() {
	*x = CoreMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_okapi_transport_v1_transport_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreMessage) ProtoMessage() {}

func (x *CoreMessage) ProtoReflect() protoreflect.Message {
	mi := &file_okapi_transport_v1_transport_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreMessage.ProtoReflect.Descriptor instead.
func (*CoreMessage) Descriptor() ([]byte, []int) {
	return file_okapi_transport_v1_transport_proto_rawDescGZIP(), []int{8}
}

func (x *CoreMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoreMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CoreMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *CoreMessage) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *CoreMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CoreMessage) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *CoreMessage) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

var File_okapi_transport_v1_transport_proto protoreflect.FileDescriptor

var file_okapi_transport_v1_transport_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2f,
	0x6b, 0x65, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x6d,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x53, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x2b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x34, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x22, 0x41, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x6d,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x6b, 0x61, 0x70,
	0x69, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65,
	0x62, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2b, 0x0a, 0x0e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0xb7, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x6b, 0x61,
	0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57,
	0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65,
	0x79, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x22, 0x44, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x55, 0x6e, 0x70, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f,
	0x6b, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x73, 0x6f,
	0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62,
	0x4b, 0x65, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x34, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x0e, 0x55, 0x6e, 0x70, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x53, 0x0a, 0x1a, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x2e, 0x6f, 0x6b, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x69, 0x6e, 0x73,
	0x69, 0x63, 0x2d, 0x69, 0x64, 0x2f, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x12, 0x4f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_okapi_transport_v1_transport_proto_rawDescOnce sync.Once
	file_okapi_transport_v1_transport_proto_rawDescData = file_okapi_transport_v1_transport_proto_rawDesc
)

func file_okapi_transport_v1_transport_proto_rawDescGZIP() []byte {
	file_okapi_transport_v1_transport_proto_rawDescOnce.Do(func() {
		file_okapi_transport_v1_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_okapi_transport_v1_transport_proto_rawDescData)
	})
	return file_okapi_transport_v1_transport_proto_rawDescData
}

var file_okapi_transport_v1_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_okapi_transport_v1_transport_proto_goTypes = []interface{}{
	(*SignRequest)(nil),      // 0: okapi.transport.v1.SignRequest
	(*SignResponse)(nil),     // 1: okapi.transport.v1.SignResponse
	(*VerifyRequest)(nil),    // 2: okapi.transport.v1.VerifyRequest
	(*VerifyResponse)(nil),   // 3: okapi.transport.v1.VerifyResponse
	(*PackRequest)(nil),      // 4: okapi.transport.v1.PackRequest
	(*PackResponse)(nil),     // 5: okapi.transport.v1.PackResponse
	(*UnpackRequest)(nil),    // 6: okapi.transport.v1.UnpackRequest
	(*UnpackResponse)(nil),   // 7: okapi.transport.v1.UnpackResponse
	(*CoreMessage)(nil),      // 8: okapi.transport.v1.CoreMessage
	(*JsonWebKey)(nil),       // 9: okapi.keys.v1.JsonWebKey
	(*SignedMessage)(nil),    // 10: pbmse.v1.SignedMessage
	(EncryptionMode)(0),      // 11: pbmse.v1.EncryptionMode
	(EncryptionAlgorithm)(0), // 12: pbmse.v1.EncryptionAlgorithm
	(*EncryptedMessage)(nil), // 13: pbmse.v1.EncryptedMessage
}
var file_okapi_transport_v1_transport_proto_depIdxs = []int32{
	9,  // 0: okapi.transport.v1.SignRequest.key:type_name -> okapi.keys.v1.JsonWebKey
	10, // 1: okapi.transport.v1.SignRequest.append_to:type_name -> pbmse.v1.SignedMessage
	10, // 2: okapi.transport.v1.SignResponse.message:type_name -> pbmse.v1.SignedMessage
	10, // 3: okapi.transport.v1.VerifyRequest.message:type_name -> pbmse.v1.SignedMessage
	9,  // 4: okapi.transport.v1.VerifyRequest.key:type_name -> okapi.keys.v1.JsonWebKey
	9,  // 5: okapi.transport.v1.PackRequest.sender_key:type_name -> okapi.keys.v1.JsonWebKey
	9,  // 6: okapi.transport.v1.PackRequest.receiver_key:type_name -> okapi.keys.v1.JsonWebKey
	11, // 7: okapi.transport.v1.PackRequest.mode:type_name -> pbmse.v1.EncryptionMode
	12, // 8: okapi.transport.v1.PackRequest.algorithm:type_name -> pbmse.v1.EncryptionAlgorithm
	13, // 9: okapi.transport.v1.PackResponse.message:type_name -> pbmse.v1.EncryptedMessage
	9,  // 10: okapi.transport.v1.UnpackRequest.sender_key:type_name -> okapi.keys.v1.JsonWebKey
	9,  // 11: okapi.transport.v1.UnpackRequest.receiver_key:type_name -> okapi.keys.v1.JsonWebKey
	13, // 12: okapi.transport.v1.UnpackRequest.message:type_name -> pbmse.v1.EncryptedMessage
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_okapi_transport_v1_transport_proto_init() }
func file_okapi_transport_v1_transport_proto_init() {
	if File_okapi_transport_v1_transport_proto != nil {
		return
	}
	file_okapi_keys_v1_keys_proto_init()
	file_pbmse_v1_pbmse_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_okapi_transport_v1_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackRequest); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackResponse); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnpackRequest); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnpackResponse); i {
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
		file_okapi_transport_v1_transport_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreMessage); i {
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
			RawDescriptor: file_okapi_transport_v1_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_okapi_transport_v1_transport_proto_goTypes,
		DependencyIndexes: file_okapi_transport_v1_transport_proto_depIdxs,
		MessageInfos:      file_okapi_transport_v1_transport_proto_msgTypes,
	}.Build()
	File_okapi_transport_v1_transport_proto = out.File
	file_okapi_transport_v1_transport_proto_rawDesc = nil
	file_okapi_transport_v1_transport_proto_goTypes = nil
	file_okapi_transport_v1_transport_proto_depIdxs = nil
}
