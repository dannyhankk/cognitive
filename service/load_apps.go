package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
)

type LoadApps interface {
	AppLoad(ctx context.Context, in *pb.LoadAppsRequest) (*pb.LoadAppsResponse, error)
}

type MyLoadApps struct {
}

func NewMyLoadApps() LoadApps {
	return &MyLoadApps{}
}

func (c *MyLoadApps) AppLoad(ctx context.Context,
	in *pb.LoadAppsRequest) (*pb.LoadAppsResponse, error) {
	res := &pb.LoadAppsResponse{}
	if in.Head == nil || len(in.Head.Id) == 0 {
		util.Logger.Errorf("load apps id empty")
		return c.doResponseExp(-1, "id empty", res)
	}
	data, err := LoadAllApps(in.Head.Id)
	if err != nil {
		util.Logger.Errorf("load all apps error, %s", err.Error())
		return c.doResponseExp(-1, "load all app failed", res)
	}
	res.Apps = data
	return c.doResponse(res)
}

func (c *MyLoadApps) doResponse(res *pb.LoadAppsResponse) (*pb.LoadAppsResponse, error) {
	res.Head = &pb.RspHead{
		Code: 0,
		Msg:  "success",
	}
	return res, nil
}

func (c *MyLoadApps) doResponseExp(code int32, msg string, res *pb.LoadAppsResponse) (*pb.LoadAppsResponse, error) {
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}
