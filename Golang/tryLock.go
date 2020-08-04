package Golang

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{},1)
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <- l.c:
		lockResult = true
	default:
		lockResult = false
	}
	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}


