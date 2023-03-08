// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: cognitive.proto

package cognitive

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// Text2SpeechClient is the client API for Text2Speech service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Text2SpeechClient interface {
	Chat2VoiceRequest(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
}

type text2SpeechClient struct {
	cc grpc.ClientConnInterface
}

func NewText2SpeechClient(cc grpc.ClientConnInterface) Text2SpeechClient {
	return &text2SpeechClient{cc}
}

func (c *text2SpeechClient) Chat2VoiceRequest(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, "/cognitive.Text2Speech/Chat2VoiceRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Text2SpeechServer is the server API for Text2Speech service.
// All implementations must embed UnimplementedText2SpeechServer
// for forward compatibility
type Text2SpeechServer interface {
	Chat2VoiceRequest(context.Context, *ChatRequest) (*ChatResponse, error)
	mustEmbedUnimplementedText2SpeechServer()
}

// UnimplementedText2SpeechServer must be embedded to have forward compatible implementations.
type UnimplementedText2SpeechServer struct {
}

func (UnimplementedText2SpeechServer) Chat2VoiceRequest(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat2VoiceRequest not implemented")
}
func (UnimplementedText2SpeechServer) mustEmbedUnimplementedText2SpeechServer() {}

// UnsafeText2SpeechServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Text2SpeechServer will
// result in compilation errors.
type UnsafeText2SpeechServer interface {
	mustEmbedUnimplementedText2SpeechServer()
}

func RegisterText2SpeechServer(s grpc.ServiceRegistrar, srv Text2SpeechServer) {
	s.RegisterService(&Text2Speech_ServiceDesc, srv)
}

func _Text2Speech_Chat2VoiceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Text2SpeechServer).Chat2VoiceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cognitive.Text2Speech/Chat2VoiceRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Text2SpeechServer).Chat2VoiceRequest(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Text2Speech_ServiceDesc is the grpc.ServiceDesc for Text2Speech service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Text2Speech_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cognitive.Text2Speech",
	HandlerType: (*Text2SpeechServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chat2VoiceRequest",
			Handler:    _Text2Speech_Chat2VoiceRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cognitive.proto",
}
