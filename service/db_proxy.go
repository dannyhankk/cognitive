package service

import (
	"encoding/json"
	"fmt"
	"github.com/dannyhankk/cognitive/db"
	pb "github.com/dannyhankk/cognitive/proto"
	"github.com/dannyhankk/cognitive/util"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

const (
	DaySeconds  = 24 * 3600
	HourSeconds = 3600

	DefaultAppsKey = "UserApp"
)

var appsLock sync.Mutex

type userGenInfo struct {
	CurGenNum    int8  `json:"cur_gen_num"`
	FirstGenTime int64 `json:"first_gen_time"`
	ResetGenTime int64 `json:"reset_gen_time,omitempty"`
}

type storeAppInfo struct {
	Title    string `json:"title"`
	Describe string `json:"describe"`
	Prompt   string `json:"prompt"`
	Id       string `json:"id"`
	Creator  string `json:"creator"`
}
type AppStore struct {
	Apps map[string]*storeAppInfo `json:"apps"`
}

func GenVideoFileName(id string) string {
	return fmt.Sprintf("%s-%d", id, time.Now().UnixMilli())
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

func AddApps(id string, app *pb.AppInfo) error {
	appsLock.Lock()
	defer appsLock.Unlock()
	appString, err := db.Get(DefaultAppsKey)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("load user apps error, %s", err)
	}
	apps := &AppStore{}
	if err == redis.Nil {
		apps.Apps = make(map[string]*storeAppInfo)
	} else {
		err = json.Unmarshal([]byte(appString), &apps)
		if err != nil {
			return fmt.Errorf("unmarshal user apps error, %s", err.Error())
		}
	}
	appId := util.GenerateID(id)
	newApp := &storeAppInfo{
		Title:    app.Title,
		Describe: app.Describe,
		Prompt:   app.Prompt,
		Id:       appId,
		Creator:  id,
	}
	apps.Apps[appId] = newApp
	storeString, err := json.Marshal(apps)
	if err != nil {
		return fmt.Errorf("marshal store app error, %s", err.Error())
	}
	err = db.Save(DefaultAppsKey, string(storeString))
	if err != nil {
		return fmt.Errorf("store app error")
	}
	return nil
}
func DeleteApps(id string, appId string) error {
	appsLock.Lock()
	defer appsLock.Unlock()
	appString, err := db.Get(DefaultAppsKey)
	if err != nil && err != redis.Nil {
		return fmt.Errorf("load user apps error, %s", err)
	}
	if err == redis.Nil {
		return fmt.Errorf("apps empty")
	}
	apps := &AppStore{}
	err = json.Unmarshal([]byte(appString), &apps)
	if err != nil {
		return fmt.Errorf("unmarshal user apps error, %s", err.Error())
	}
	if _, ok := apps.Apps[appId]; !ok {
		return fmt.Errorf("app not found error")
	}
	if apps.Apps[appId].creator != id {
		return fmt.Errorf("not your app error")
	}
	delete(apps.Apps, appId)
	storeString, err := json.Marshal(apps)
	if err != nil {
		return fmt.Errorf("marshal store app error, %s", err.Error())
	}
	err = db.Save(DefaultAppsKey, string(storeString))
	if err != nil {
		return fmt.Errorf("store app error")
	}
	return nil
}

func LoadAllApps(id string) ([]*pb.AppInfo, error) {
	appString, err := db.Get(DefaultAppsKey)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("load user apps error, %s", err)
	}
	if err == redis.Nil {
		return nil, fmt.Errorf("apps empty")
	}
	apps := &AppStore{}
	err = json.Unmarshal([]byte(appString), &apps)
	if err != nil {
		return nil, fmt.Errorf("unmarshal user apps error, %s", err.Error())
	}
	var res []*pb.AppInfo
	for _, val := range apps.Apps {
		res = append(res, &pb.AppInfo{
			Title:    val.title,
			Describe: val.describe,
			Prompt:   val.prompt,
			Id:       val.id,
		})
	}
	return res, nil
}
