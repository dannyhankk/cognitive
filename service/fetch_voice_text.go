package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
	"os"
)

type FetchVoiceText interface {
	FetchText(ctx context.Context,
		in *pb.FetchVoiceTextRequest) (*pb.FetchVoiceTextResponse, error)
}

type MyFetchVoiceText struct {
}

func NewMyFetchVoiceText() FetchVoiceText {
	return &MyFetchVoiceText{}
}

func (c *MyFetchVoiceText) FetchText(ctx context.Context,
	in *pb.FetchVoiceTextRequest) (*pb.FetchVoiceTextResponse, error) {
	res := &pb.FetchVoiceTextResponse{}
	if in.Head == nil || len(in.Head.Id) == 0 {
		util.Logger.Errorf("FetchText request id empty")
		return c.doResponseExp(-1, "request id empty", res)
	}
	if len(in.FileUrl) == 0 {
		util.Logger.Errorf("FetchText request url empty")
		return c.doResponseExp(-2, "request file url empty", res)
	}
	data, err := os.ReadFile(in.FileUrl)
	if err != nil {
		util.Logger.Errorf("read file error, %s, %s", in.FileUrl, err.Error())
		return c.doResponseExp(-2, "read file error", res)
	}
	res.Text = string(data)
	return c.doResponse(res)
}

func (c *MyFetchVoiceText) doResponse(res *pb.FetchVoiceTextResponse) (
	*pb.FetchVoiceTextResponse, error) {
	res.Head = &pb.RspHead{
		Code: 0,
		Msg:  "success",
	}
	return res, nil

}

func (c *MyFetchVoiceText) doResponseExp(code int32, msg string, res *pb.FetchVoiceTextResponse) (
	*pb.FetchVoiceTextResponse, error) {
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}
