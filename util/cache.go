package util

import (
	"errors"
	"sync"
	"time"
)

type chatInfo struct {
	lastMessage    string
	lastUpdateTime int64
}

var (
	chatHistory sync.Map
	gapSecond   = int64(300)
)

func ReplaceChatInfo(id string, message string) error {
	timeNow := time.Now().Unix()
	data := &chatInfo{}
	data.lastMessage += message
	data.lastUpdateTime = timeNow
	chatHistory.Store(id, data)
	return nil
}
func UpdateChatInfo(id string, message string) error {
	timeNow := time.Now().Unix()
	value, ok := chatHistory.Load(id)
	data := &chatInfo{}
	if ok {
		data = value.(*chatInfo)
		// 时间错误或者超过5分分钟，忽略之前信息
		if data.lastUpdateTime > timeNow ||
			data.lastUpdateTime+gapSecond <= timeNow {
			data.lastMessage = ""
		}
	}
	data.lastMessage += message
	data.lastUpdateTime = timeNow
	chatHistory.Store(id, data)
	return nil
}

func LoadChatInfo(id string) (string, error) {
	timeNow := time.Now().Unix()
	value, ok := chatHistory.Load(id)
	data := &chatInfo{}
	if ok {
		data = value.(*chatInfo)
		// 时间错误或者超过5分分钟，忽略之前信息
		if data.lastUpdateTime > timeNow ||
			data.lastUpdateTime+gapSecond <= timeNow {
			chatHistory.Delete(id)
			return "", errors.New("message expired")
		}
		return data.lastMessage, nil
	}
	return "", errors.New("message not exist")
}
