package service

import (
	"context"
	"github.com/dannyhankk/cognitive/client"
	"github.com/dannyhankk/cognitive/db"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
	"os"
	"time"
)

const (
	default_path = "/http_default/video"
	http_default = "https://m.dannyhkk.cn/video"
)

type Chat2Speech interface {
	TransSpeech(ctx context.Context, in *pb.ChatRequest) (*pb.ChatResponse, error)
}

type MyChat2Speech struct {
}

func NewMyChat2Speech() Chat2Speech {
	return &MyChat2Speech{}
}

func (c *MyChat2Speech) TransSpeech(ctx context.Context, in *pb.ChatRequest) (
	*pb.ChatResponse, error) {
	res := &pb.ChatResponse{}
	if in.Head == nil || len(in.Head.Id) == 0 {
		return c.doResponseExp(-1, "id empty", res)
	}
	userInfo, _ := checkUserGen(in.Head.Id)
	if userInfo == nil {
		userInfo = &userGenInfo{
			CurGenNum:    0,
			FirstGenTime: time.Now().Unix(),
		}
	}
	if userInfo.CurGenNum >= util.RootConfig.VoiceGenLimit {
		return c.doResponseExp(-2, "user out limit", res)
	}

	fileDir := default_path + "/" + in.Head.Id
	if _, err := os.Stat(fileDir); err != nil && os.IsNotExist(err) {
		err = os.Mkdir(fileDir, os.ModePerm)
		util.Logger.Infof("make dir : %v", err)
	}
	fileName := GenVideoFileName(in.Head.Id)
	go func() {
		client.CreateVoice(in.Text, pb.LangType(in.Lang), fileDir+"/"+fileName+".wav", func() {
			saveVoiceGen(in.Head.Id, userInfo)
			db.Save(fileName, in.Title)
		})
	}()
	userInfo.CurGenNum += 1

	return c.doResponse(res)
}

func (c *MyChat2Speech) doResponse(res *pb.ChatResponse) (*pb.ChatResponse, error) {
	res.Head = &pb.RspHead{
		Code: 0,
	}
	return res, nil
}

func (c *MyChat2Speech) doResponseExp(code int32, msg string, res *pb.ChatResponse) (*pb.ChatResponse, error) {
	res.Head = &pb.RspHead{
		Code: code,
		Msg:  msg,
	}
	return res, nil
}
