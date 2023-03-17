package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
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

	return c.doResponse(res)
}

func (c *MyDeleteApp) doResponse(res *pb.DeleteAppResponse) (*pb.DeleteAppResponse, error) {

	return res, nil
}

func (c *MyDeleteApp) doResponseExp(code int32, msg string, res *pb.DeleteAppResponse) (*pb.DeleteAppResponse, error) {

	return res, nil
}
