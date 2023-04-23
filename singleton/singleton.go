package singleton

import (
	"sync"
)

type Singleton struct {
}

var mu sync.Mutex

func GetInstance(instance *Singleton) *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &Singleton{}
		return instance
	} else {
		return instance
	}
}
