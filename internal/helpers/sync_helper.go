package helpers

import (
	"sync"
)

type MutexHelper interface {
	Lock()
	Unlock()
}

type MutexHelperImpl struct {
	mutex *sync.Mutex
}

func NewMutexHelper() MutexHelper {
	return &MutexHelperImpl{
		mutex: &sync.Mutex{},
	}
}

func (helper *MutexHelperImpl) Lock() {
	helper.mutex.Lock()
}

func (helper *MutexHelperImpl) Unlock() {
	helper.mutex.Unlock()
}
