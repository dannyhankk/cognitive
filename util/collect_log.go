package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/dannyhankk/cognitive/db"
	"github.com/go-redis/redis"
	"io"
	"os"
	"strings"
	"time"
)

var (
	DefaultAppsClickKey = "UserCLick"
)

type AppClick struct {
	Click map[string]int64 `json:"click"`
}

type AppReport struct {
	Id string `json:"id"`
}

func LoadReportLog(click *AppClick) (*AppClick, error) {
	yesterday := time.Unix(time.Now().Unix()-24*3600, 0).Format("2006/01/02")
	Logger.Infof("read day: %s", yesterday)
	file, err := os.OpenFile("/root/working/app_report.log", os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open log file error, %s", err.Error())
	}
	tmpClick := click
	r := bufio.NewReader(file)
	for {
		data, rErr := r.ReadBytes('\n')
		if rErr == io.EOF {
			Logger.Infof("read end")
			break
		}
		if rErr != nil {
			Logger.Errorf("read error, %s", rErr.Error())
			break
		}
		if !strings.Contains(string(data), yesterday) {
			continue
		}
		strs := strings.Split(string(data), " ")
		Logger.Errorf("%+v", strs)
		if len(strs) < 2 {
			Logger.Errorf("wrong log format, %s", string(data))
			continue
		}
		reportSt := &AppReport{}
		err := json.Unmarshal([]byte(strs[1]), reportSt)
		if err != nil {
			Logger.Errorf("wrong log format, %s", err.Error())
			continue
		}
		if _, ok := tmpClick.Click[reportSt.Id]; ok {
			tmpClick.Click[reportSt.Id] += 1
			continue
		}
		tmpClick.Click[reportSt.Id] = 1

	}
	Logger.Infof("click: %+v", tmpClick)
	return tmpClick, nil
}

func LoadUserClick() (*AppClick, error) {
	appClick, err := db.Get(DefaultAppsClickKey)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("load user apps error, %s", err)
	}
	if err == redis.Nil {
		return nil, nil
	}
	click := &AppClick{}
	err = json.Unmarshal([]byte(appClick), &click)
	if err != nil {
		return nil, fmt.Errorf("unmarshal user click error, %s", err.Error())
	}
	return click, nil
}

func SaveUserClick(appclick *AppClick) error {
	data, err := json.Marshal(appclick)
	if err != nil {
		return fmt.Errorf("marshal app click  error, %s", err.Error())
	}
	return db.Save(DefaultAppsClickKey, string(data))
}
