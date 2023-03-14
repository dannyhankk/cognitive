package service

import (
	"encoding/json"
	"fmt"
	"github.com/dannyhankk/cognitive/db"
	"github.com/dannyhankk/cognitive/util"
	"github.com/go-redis/redis"
	"time"
)

const (
	DaySeconds  = 24 * 3600
	HourSeconds = 3600
)

type userGenInfo struct {
	CurGenNum    int8  `json:"cur_gen_num"`
	FirstGenTime int64 `json:"first_gen_time"`
	ResetGenTime int64 `json:"reset_gen_time,omitempty"`
}

func checkUserGen(id string) (*userGenInfo, error) {
	key := id + "_voice"
	result, err := db.Get(key)
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("check user gen info error, %s", err.Error())
	}
	userInfo := &userGenInfo{}
	err = json.Unmarshal([]byte(result), userInfo)
	if err != nil {
		db.Del(key)
		return nil, nil
	}

	timeNow := time.Now().Unix()
	if !checkSameDay(timeNow, userInfo.FirstGenTime) {
		userInfo.FirstGenTime = timeNow
		userInfo.CurGenNum = 0
		return userInfo, nil
	}
	return userInfo, nil
}
func checkSameDay(new, old int64) bool {
	if new-old >= DaySeconds {
		return false
	}
	newHour := time.Unix(new, 0).Hour()
	oldHour := time.Unix(old, 0).Hour()
	if newHour < oldHour {
		return false
	}
	return true
}

func saveVoiceGen(id string, info *userGenInfo) {
	key := id + "_voice"
	data, err := json.Marshal(info)
	if err != nil {
		util.Logger.Errorf("save user voice gen error, %s, %s", id, err.Error())
		return
	}
	db.Save(key, string(data))
}
func ResetVoiceGen(id string) error {
	key := id + "_voice"
	result, err := db.Get(key)
	if err == redis.Nil {
		return fmt.Errorf("no need to reset")
	}
	if err != nil {
		return fmt.Errorf("get user gen info error, %s", err.Error())
	}
	userInfo := &userGenInfo{}
	err = json.Unmarshal([]byte(result), userInfo)
	if err != nil {
		db.Del(key)
		return fmt.Errorf("delete user gen info, no need to reset")
	}
	if err != nil {
		return fmt.Errorf("checkUserGen error, %s", err.Error())
	}
	timeNow := time.Now().Unix()
	if checkSameDay(timeNow, userInfo.ResetGenTime) {
		return fmt.Errorf("already reset")
	}
	userInfo.ResetGenTime = timeNow
	userInfo.CurGenNum = 0

	saveData, err := json.Marshal(userInfo)
	if err != nil {
		return fmt.Errorf("reset user gen failed, %s", err.Error())
	}
	err = db.Save(key, string(saveData))
	if err != nil {
		return fmt.Errorf("reset user gen failed, %s", err.Error())
	}
	return nil
}
