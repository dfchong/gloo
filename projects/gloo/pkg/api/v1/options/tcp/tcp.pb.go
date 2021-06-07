// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tcp/tcp.proto

package tcp

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
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

// Contains various settings for Envoy's tcp proxy filter.
// See here for more information: https://www.envoyproxy.io/docs/envoy/v1.10.0/api-v2/config/filter/network/tcp_proxy/v2/tcp_proxy.proto#envoy-api-msg-config-filter-network-tcp-proxy-v2-tcpproxy
type TcpProxySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxConnectAttempts *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	IdleTimeout        *duration.Duration    `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// If set, this configures tunneling, e.g. configuration options to tunnel multiple TCP
	// payloads over a shared HTTP tunnel. If this message is absent, the payload
	// will be proxied upstream as per usual.
	TunnelingConfig *TcpProxySettings_TunnelingConfig `protobuf:"bytes,12,opt,name=tunneling_config,json=tunnelingConfig,proto3" json:"tunneling_config,omitempty"`
}

func (x *TcpProxySettings) Reset() {
	*x = TcpProxySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpProxySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpProxySettings) ProtoMessage() {}

func (x *TcpProxySettings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpProxySettings.ProtoReflect.Descriptor instead.
func (*TcpProxySettings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescGZIP(), []int{0}
}

func (x *TcpProxySettings) GetMaxConnectAttempts() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxConnectAttempts
	}
	return nil
}

func (x *TcpProxySettings) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *TcpProxySettings) GetTunnelingConfig() *TcpProxySettings_TunnelingConfig {
	if x != nil {
		return x.TunnelingConfig
	}
	return nil
}

// Configuration for tunneling TCP over other transports or application layers.
type TcpProxySettings_TunnelingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hostname to send in the synthesized CONNECT headers to the upstream proxy.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *TcpProxySettings_TunnelingConfig) Reset() {
	*x = TcpProxySettings_TunnelingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpProxySettings_TunnelingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpProxySettings_TunnelingConfig) ProtoMessage() {}

func (x *TcpProxySettings_TunnelingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpProxySettings_TunnelingConfig.ProtoReflect.Descriptor instead.
func (*TcpProxySettings_TunnelingConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TcpProxySettings_TunnelingConfig) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x63, 0x70, 0x2f, 0x74, 0x63, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x63, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x10, 0x54, 0x63,
	0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4e,
	0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x3c,
	0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x65, 0x0a, 0x10,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x74, 0x63, 0x70, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x54, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x2d, 0x0a, 0x0f, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x46, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74,
	0x63, 0x70, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_goTypes = []interface{}{
	(*TcpProxySettings)(nil),                 // 0: tcp.options.gloo.solo.io.TcpProxySettings
	(*TcpProxySettings_TunnelingConfig)(nil), // 1: tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig
	(*wrappers.UInt32Value)(nil),             // 2: google.protobuf.UInt32Value
	(*duration.Duration)(nil),                // 3: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_depIdxs = []int32{
	2, // 0: tcp.options.gloo.solo.io.TcpProxySettings.max_connect_attempts:type_name -> google.protobuf.UInt32Value
	3, // 1: tcp.options.gloo.solo.io.TcpProxySettings.idle_timeout:type_name -> google.protobuf.Duration
	1, // 2: tcp.options.gloo.solo.io.TcpProxySettings.tunneling_config:type_name -> tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpProxySettings); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpProxySettings_TunnelingConfig); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tcp_tcp_proto_depIdxs = nil
}
