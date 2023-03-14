package service

import (
	"context"
	"github.com/dannyhankk/cognitive/client"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
	"os"
)

const (
	default_path = "/http_default"
	http_default = "https://m.dannyhkk.cn/"
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
	if in.Head.Id == "" {
		return c.doResponseExp(-1, "id empty", res)
	}
	filedir := default_path + "/" + in.Head.Id
	if _, err := os.Stat(filedir); err != nil && os.IsNotExist(err) {
		err = os.Mkdir(filedir, os.ModePerm)
		util.Logger.Infof("make dir : %v", err)
	}
	go func() {
		client.CreateVoice(in.Text, in.Lang, filedir+"/"+in.Text[6:32]+".wav")
	}()

	return c.doResponse(res)
}

func (c *MyChat2Speech) doResponse(res *pb.ChatResponse) (*pb.ChatResponse, error) {

	res.Head = &pb.RspHead{
		Code: 0,
	}
	return res, nil
}

func (c *MyChat2Speech) doResponseExp(code int32, msg string, res *pb.ChatResponse) (*pb.ChatResponse, error) {
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}
