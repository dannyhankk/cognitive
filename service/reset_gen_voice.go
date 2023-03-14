package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
)

type ResetGenVoice interface {
	Reset(ctx context.Context,
		in *pb.ResetVoiceGenRequest) (*pb.ResetVoiceGenResponse, error)
}

type MyResetGenVoice struct {
}

func NewMyResetGenVoice() ResetGenVoice {
	return &MyResetGenVoice{}
}

func (c *MyResetGenVoice) Reset(
	ctx context.Context, in *pb.ResetVoiceGenRequest) (*pb.ResetVoiceGenResponse, error) {
	res := &pb.ResetVoiceGenResponse{}
	if in.Head == nil || len(in.Head.Id) == 0 {
		util.Logger.Errorf("reset request id empty")
		return c.doResponseExp(-1, "id empty", res)
	}
	ResetVoiceGen(in.Head.Id)
	return c.doResponse(res)
}

func (c *MyResetGenVoice) doResponse(res *pb.ResetVoiceGenResponse) (*pb.ResetVoiceGenResponse, error) {

	return res, nil
}

func (c *MyResetGenVoice) doResponseExp(
	code int32, msg string, res *pb.ResetVoiceGenResponse) (*pb.ResetVoiceGenResponse, error) {
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}
