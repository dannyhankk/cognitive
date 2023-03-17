package service

import (
	"context"
	"fmt"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
	"os"
	"path/filepath"
	"strings"
)

type LoadIcon interface {
	Load(ctx context.Context, in *pb.LoadIconsRequest) (*pb.LoadIconsResponse, error)
}

type MyLoadIcon struct {
}

func NewMyLoadIcon() LoadIcon {
	return &MyLoadIcon{}
}

func (c *MyLoadIcon) Load(ctx context.Context,
	in *pb.LoadIconsRequest) (*pb.LoadIconsResponse, error) {
	res := &pb.LoadIconsResponse{}
	var icons []string
	root := default_path + default_icon_path
	err := filepath.Walk(root, func(
		path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Size() == 0 {
			return nil
		}
		if !strings.Contains(info.Name(), "png") {
			return nil
		}
		icons = append(icons, http_default+default_icon_path+info.Name())
		return nil
	})
	if err != nil {
		util.Logger.Errorf("scan failed, %s", err.Error())
		return c.doResponseExp(-1, fmt.Sprintf("scan icon path error"), res)
	}
	res.Icons = icons
	util.Logger.Infof("icons: %+v", res)
	return c.doResponse(res)
}

func (c *MyLoadIcon) doResponse(res *pb.LoadIconsResponse) (*pb.LoadIconsResponse,
	error) {
	res.Head = &pb.RspHead{
		Code: 0,
		Msg:  "success",
	}
	return res, nil
}

func (c *MyLoadIcon) doResponseExp(code int32, msg string, res *pb.LoadIconsResponse) (*pb.LoadIconsResponse, error) {
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}
