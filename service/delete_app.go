package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
)

type DeleteApp interface {
	Delete(ctx context.Context, in *pb.DeleteAppRequest) (*pb.DeleteAppResponse, error)
}

type MyDeleteApp struct {
}

func NewMyDeleteApp() DeleteApp {
	return &MyDeleteApp{}
}

func (c *MyDeleteApp) Delete(ctx context.Context,
	in *pb.DeleteAppRequest) (*pb.DeleteAppResponse, error) {
	res := &pb.DeleteAppResponse{}
	if in.Head == nil || len(in.Head.Id) == 0 {
		util.Logger.Errorf("add app id empty")
		return c.doResponseExp(-1, "id empty", res)
	}
	err := DeleteApps(in.Head.Id, in.Id)
	return c.doResponse(res)
}

func (c *MyDeleteApp) doResponse(res *pb.DeleteAppResponse) (*pb.DeleteAppResponse, error) {

	return res, nil
}

func (c *MyDeleteApp) doResponseExp(code int32, msg string, res *pb.DeleteAppResponse) (*pb.DeleteAppResponse, error) {

	return res, nil
}
