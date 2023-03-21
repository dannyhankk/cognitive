package util

import (
	"time"
)

var (
	Ticker *time.Ticker
)

func InitTicker() error {
	Ticker = time.NewTicker(time.Second * 1800)
	go func() {
		ExecuteTask()
		for {
			select {
			case <-Ticker.C:
				if !check2am() {
					continue
				}
				ExecuteTask()
			}
		}
	}()
	return nil
}

func check2am() bool {
	tm := time.Unix(time.Now().Unix(), 0)
	if tm.Hour() < 2 || tm.Hour() > 2 {
		return false
	}
	if tm.Minute() < 30 {
		return false
	}
	return true
}

func ExecuteTask() {
	click, err := LoadUserClick()
	if err != nil {
		Logger.Errorf("load user click error, %s", err.Error())
		return
	}
	if click == nil {
		click = &AppClick{
			Click: make(map[string]int64),
		}
	}
	afterClick, err := LoadReportLog(click)
	if err != nil {
		Logger.Errorf("load report log error, %s", err.Error())
		return
	}
	err = SaveUserClick(afterClick)
	if err != nil {
		Logger.Errorf("save user click error, %s", err.Error())
	}
}
