package service

import (
	"context"
	pb "github.com/dannyhankk/cognitive/proto"
)

type FetchVideos interface {
	Fetch(ctx context.Context, in *pb.FetchVideoRequest) (*pb.FetchVideoResponse, error)
}

type MyFetchVideos struct {
}

func NewMyFetchVideos() FetchVideos {
	return &MyFetchVideos{}
}

func (c *MyFetchVideos) Fetch(ctx context.Context,
	in *pb.FetchVideoRequest) (*pb.FetchVideoResponse, error) {
	res := &pb.FetchVideoResponse{}
	res.VideoList = []*pb.Video{
		{
			Src:    "https://m.dannyhkk.cn/englishpod_D0091dg.mp3",
			Author: "Kevin",
			Title:  "englishpod_dg_91",
		},
		{
			Src:    "https://m.dannyhkk.cn/englishpod_D0091dg.mp3",
			Author: "Kevin",
			Title:  "englishpod_dg_90",
		},
	}
	return c.doResponse(res)
}

func (c *MyFetchVideos) doResponse(res *pb.FetchVideoResponse) (*pb.FetchVideoResponse, error) {
	res.Head = &pb.RspHead{
		Code: 0,
	}
	return res, nil
}

func (c *MyFetchVideos) doResponseExp(code int32, msg string, res *pb.FetchVideoResponse) (*pb.FetchVideoResponse, error) {
	return nil, nil
}
