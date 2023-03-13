package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
)

type Chat2Speech interface {
	TransSpeech(ctx context.Context, in *pb.ChatRequest) (*pb.ChatResponse, error)
}

type MyChat2Speech struct {
}

func NewMyChat2Speech() Chat2Speech {
	return &MyChat2Speech{}
}

func (c MyChat2Speech) TransSpeech(ctx context.Context, in *pb.ChatRequest) (
	*pb.ChatResponse, error) {

	return nil, nil
}