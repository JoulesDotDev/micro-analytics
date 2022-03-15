package handler

import (
	"sync"
)

type Analytics struct {
	lock sync.Mutex
}

func New() *Analytics {
	return &Analytics{}
}
