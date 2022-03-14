package handler

import (
	"sync"
)

type Lock struct {
	mu sync.Mutex
}

func CreateLock() *Lock {
	return &Lock{}
}
