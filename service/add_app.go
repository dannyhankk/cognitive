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

	return c.doResponse(res)
}

func (c *MyAddApp) doResponse(res *pb.AddAppResponse) (*pb.AddAppResponse, error) {

	return res, nil
}

func (c *MyAddApp) doResponseExp(code int32, msg string, res *pb.AddAppResponse) (*pb.AddAppResponse, error) {
	return res, nil
}
