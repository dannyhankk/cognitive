package main

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/service"
)

func (s *MyServer) Chat2VoiceRequest(
	ctx context.Context, in *pb.ChatRequest) (
	out *pb.ChatResponse, err error) {
	handler := service.NewMyChat2Speech()
	return handler.TransSpeech(ctx, in)
}
