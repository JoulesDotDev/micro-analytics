package handler

type Analytics struct {
	lock *Lock
}

func New() *Analytics {
	return &Analytics{
		lock: CreateLock(),
	}
}
