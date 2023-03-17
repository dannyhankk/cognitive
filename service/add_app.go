package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
)

type AddApp interface {
	Add(ctx context.Context, in *pb.AddAppRequest) (*pb.AddAppResponse, error)
}

type MyAddApp struct {
}

func NewMyAddApp() AddApp {
	return &MyAddApp{}
}

func (c *MyAddApp) Add(ctx context.Context,
	in *pb.AddAppRequest) (*pb.AddAppResponse, error) {
	res := &pb.AddAppResponse{}
	if in.Head == nil || len(in.Head.Id) == 0 {
		util.Logger.Errorf("add app id empty")
		return c.doResponseExp(-1, "id empty", res)
	}
	if in.App == nil {
		util.Logger.Errorf("add app error, app empty")
		return c.doResponseExp(-1, "app empty", res)
	}
	err := AddApps(in.Head.Id, in.App)
	if err != nil {
		util.Logger.Errorf("add apps error, %s", err.Error())
		return c.doResponseExp(-1, "add app error", res)
	}
	return c.doResponse(res)
}

func (c *MyAddApp) doResponse(res *pb.AddAppResponse) (*pb.AddAppResponse, error) {
	res.Head = &pb.RspHead{
		Code: 0,
		Msg:  "success",
	}
	return res, nil
}

func (c *MyAddApp) doResponseExp(code int32, msg string, res *pb.AddAppResponse) (*pb.AddAppResponse, error) {
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}
