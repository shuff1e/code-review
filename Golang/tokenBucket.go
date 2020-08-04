package Golang

import (
	"fmt"
	"time"
)

var fillInterval = time.Millisecond * 10
var capacity = 100
var tokenBucket = make(chan struct{},capacity)

var FillToken = func() {
	ticker := time.NewTicker(fillInterval)
	for {
		select {
		case <- ticker.C:
			select {
			case tokenBucket <- struct{}{}:
			default:

			}
		}
		time.Sleep(time.Second)
		fmt.Println("current token cnt:",len(tokenBucket),time.Now())
	}
}

var TakeAvailable = func(block bool) bool {
	var takeResult bool
	if block {
		select {
		case <- tokenBucket:
			takeResult = true
		}
	} else {
		select {
		case <- tokenBucket:
			takeResult = true
		default:
			takeResult = false
		}
	}
	return takeResult
}
