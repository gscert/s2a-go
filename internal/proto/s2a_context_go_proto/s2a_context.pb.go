// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: internal/proto/s2a_context/s2a_context.proto

package s2a_context_go_proto

import (
	common_go_proto "github.com/google/s2a/internal/proto/common_go_proto"
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

type S2AContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The application protocol negotiated for this connection, e.g., 'grpc'.
	ApplicationProtocol string `protobuf:"bytes,1,opt,name=application_protocol,json=applicationProtocol,proto3" json:"application_protocol,omitempty"`
	// The TLS version number that the S2A's handshaker module used to set up the
	// session.
	TlsVersion common_go_proto.TLSVersion `protobuf:"varint,2,opt,name=tls_version,json=tlsVersion,proto3,enum=s2a.proto.TLSVersion" json:"tls_version,omitempty"`
	// The TLS ciphersuite negotiated by the S2A's handshaker module.
	Ciphersuite common_go_proto.Ciphersuite `protobuf:"varint,3,opt,name=ciphersuite,proto3,enum=s2a.proto.Ciphersuite" json:"ciphersuite,omitempty"`
	// The authenticated identity of the peer.
	PeerIdentity *common_go_proto.Identity `protobuf:"bytes,4,opt,name=peer_identity,json=peerIdentity,proto3" json:"peer_identity,omitempty"`
	// The local identity used during session setup. This could be:
	//   - The local identity that the client specifies in ClientSessionStartReq.
	//   - One of the local identities that the server specifies in
	//     ServerSessionStartReq.
	//   - If neither client or server specifies local identities, the S2A picks the
	//     default one. In this case, this field will contain that identity.
	LocalIdentity *common_go_proto.Identity `protobuf:"bytes,5,opt,name=local_identity,json=localIdentity,proto3" json:"local_identity,omitempty"`
	// The SHA256 hash of the peer certificate used in the handshake.
	PeerCertFingerprint []byte `protobuf:"bytes,6,opt,name=peer_cert_fingerprint,json=peerCertFingerprint,proto3" json:"peer_cert_fingerprint,omitempty"`
	// The SHA256 hash of the local certificate used in the handshake.
	LocalCertFingerprint []byte `protobuf:"bytes,7,opt,name=local_cert_fingerprint,json=localCertFingerprint,proto3" json:"local_cert_fingerprint,omitempty"`
	// Set to true if a cached session was reused to resume the handshake.
	IsHandshakeResumed bool `protobuf:"varint,8,opt,name=is_handshake_resumed,json=isHandshakeResumed,proto3" json:"is_handshake_resumed,omitempty"`
}

func (x *S2AContext) Reset() {
	*x = S2AContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_s2a_context_s2a_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2AContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2AContext) ProtoMessage() {}

func (x *S2AContext) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_s2a_context_s2a_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2AContext.ProtoReflect.Descriptor instead.
func (*S2AContext) Descriptor() ([]byte, []int) {
	return file_internal_proto_s2a_context_s2a_context_proto_rawDescGZIP(), []int{0}
}

func (x *S2AContext) GetApplicationProtocol() string {
	if x != nil {
		return x.ApplicationProtocol
	}
	return ""
}

func (x *S2AContext) GetTlsVersion() common_go_proto.TLSVersion {
	if x != nil {
		return x.TlsVersion
	}
	return common_go_proto.TLSVersion(0)
}

func (x *S2AContext) GetCiphersuite() common_go_proto.Ciphersuite {
	if x != nil {
		return x.Ciphersuite
	}
	return common_go_proto.Ciphersuite(0)
}

func (x *S2AContext) GetPeerIdentity() *common_go_proto.Identity {
	if x != nil {
		return x.PeerIdentity
	}
	return nil
}

func (x *S2AContext) GetLocalIdentity() *common_go_proto.Identity {
	if x != nil {
		return x.LocalIdentity
	}
	return nil
}

func (x *S2AContext) GetPeerCertFingerprint() []byte {
	if x != nil {
		return x.PeerCertFingerprint
	}
	return nil
}

func (x *S2AContext) GetLocalCertFingerprint() []byte {
	if x != nil {
		return x.LocalCertFingerprint
	}
	return nil
}

func (x *S2AContext) GetIsHandshakeResumed() bool {
	if x != nil {
		return x.IsHandshakeResumed
	}
	return false
}

var File_internal_proto_s2a_context_s2a_context_proto protoreflect.FileDescriptor

var file_internal_proto_s2a_context_s2a_context_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x32, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x73, 0x32, 0x61,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x32, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x03,
	0x0a, 0x0a, 0x53, 0x32, 0x41, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x31, 0x0a, 0x14,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x36, 0x0a, 0x0b, 0x74, 0x6c, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x32, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x4c, 0x53, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x6c, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x73, 0x75, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73,
	0x32, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x73,
	0x75, 0x69, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x73, 0x75, 0x69, 0x74,
	0x65, 0x12, 0x38, 0x0a, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x32, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x70,
	0x65, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x32, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x70, 0x65, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74,
	0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x69, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x32, 0x61, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x32, 0x61, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_s2a_context_s2a_context_proto_rawDescOnce sync.Once
	file_internal_proto_s2a_context_s2a_context_proto_rawDescData = file_internal_proto_s2a_context_s2a_context_proto_rawDesc
)

func file_internal_proto_s2a_context_s2a_context_proto_rawDescGZIP() []byte {
	file_internal_proto_s2a_context_s2a_context_proto_rawDescOnce.Do(func() {
		file_internal_proto_s2a_context_s2a_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_s2a_context_s2a_context_proto_rawDescData)
	})
	return file_internal_proto_s2a_context_s2a_context_proto_rawDescData
}

var file_internal_proto_s2a_context_s2a_context_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_proto_s2a_context_s2a_context_proto_goTypes = []interface{}{
	(*S2AContext)(nil),               // 0: s2a.proto.S2AContext
	(common_go_proto.TLSVersion)(0),  // 1: s2a.proto.TLSVersion
	(common_go_proto.Ciphersuite)(0), // 2: s2a.proto.Ciphersuite
	(*common_go_proto.Identity)(nil), // 3: s2a.proto.Identity
}
var file_internal_proto_s2a_context_s2a_context_proto_depIdxs = []int32{
	1, // 0: s2a.proto.S2AContext.tls_version:type_name -> s2a.proto.TLSVersion
	2, // 1: s2a.proto.S2AContext.ciphersuite:type_name -> s2a.proto.Ciphersuite
	3, // 2: s2a.proto.S2AContext.peer_identity:type_name -> s2a.proto.Identity
	3, // 3: s2a.proto.S2AContext.local_identity:type_name -> s2a.proto.Identity
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_proto_s2a_context_s2a_context_proto_init() }
func file_internal_proto_s2a_context_s2a_context_proto_init() {
	if File_internal_proto_s2a_context_s2a_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_s2a_context_s2a_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2AContext); i {
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
			RawDescriptor: file_internal_proto_s2a_context_s2a_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_s2a_context_s2a_context_proto_goTypes,
		DependencyIndexes: file_internal_proto_s2a_context_s2a_context_proto_depIdxs,
		MessageInfos:      file_internal_proto_s2a_context_s2a_context_proto_msgTypes,
	}.Build()
	File_internal_proto_s2a_context_s2a_context_proto = out.File
	file_internal_proto_s2a_context_s2a_context_proto_rawDesc = nil
	file_internal_proto_s2a_context_s2a_context_proto_goTypes = nil
	file_internal_proto_s2a_context_s2a_context_proto_depIdxs = nil
}
