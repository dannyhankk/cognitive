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

func (c *MyChat2Speech) TransSpeech(ctx context.Context, in *pb.ChatRequest) (
	*pb.ChatResponse, error) {

	res := &pb.ChatResponse{}
	res.VideoItem = &pb.Video{
		Src:    "https://m.dannyhkk.cn/englishpod_D0091dg.mp3",
		Author: "Kevin",
		Title:  "englishpod_dg_89",
	}
	return c.doResponse(res)
}

func (c *MyChat2Speech) doResponse(res *pb.ChatResponse) (*pb.ChatResponse, error) {

	res.Head = &pb.RspHead{
		Code: 0,
	}
	return res, nil
}

func (c *MyChat2Speech) doResponseExp(code int32, msg string, res *pb.ChatResponse) (*pb.ChatResponse, error) {

	return nil, nil
}
