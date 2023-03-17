package service

import (
	"context"
	"fmt"
	"github.com/dannyhankk/cognitive/db"
	pb "github.com/dannyhankk/cognitive/proto"
	"os"
	"path/filepath"
	"strings"
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

	var files []string
	root := default_voice_path + "/" + in.Head.Id
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Size() == 0 {
			return nil
		}
		if !strings.Contains(info.Name(), "wav") {
			return nil
		}
		files = append(files, info.Name())
		return nil
	})
	if err != nil {
		return c.doResponseExp(-1, fmt.Sprintf("scan user files error, %s", err.Error()), res)
	}
	res.VideoList = []*pb.Video{}
	for _, file := range files {
		fileNameWithOutEx := file[0 : len(file)-4]
		title, _ := db.Get(fileNameWithOutEx)
		res.VideoList = append(res.VideoList, &pb.Video{
			Src:    http_voice_default + "/" + in.Head.Id + "/" + file,
			Author: "Kevin",
			Title:  title,
		})
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
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	res.VideoList = []*pb.Video{}
	return res, nil
}
