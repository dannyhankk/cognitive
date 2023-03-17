package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
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

	return c.doResponse(res)
}

func (c *MyLoadApps) doResponse(res *pb.LoadAppsResponse) (*pb.LoadAppsResponse, error) {

	return res, nil
}

func (c *MyLoadApps) doResponseExp(code int32, msg string, res *pb.LoadAppsResponse) (*pb.LoadAppsResponse, error) {

	return res, nil
}
