package util

import (
	"fmt"
	"sync"
	"time"
)

var (
	IdLock sync.Mutex
)

func GenerateID(id string) string {
	return fmt.Sprintf("%s_%d", id, time.Now().UnixMilli())
}
