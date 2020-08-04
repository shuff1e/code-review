package Golang

/*
import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Incr() {
	client := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey,1,time.Second*5)
	lockSuccess,err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err,"lock result: ",lockSuccess)
		return
	}

	getResp := client.Get(counterKey)
	cntValue,err := getResp.Int64()

	if err == nil || err == redis.Nil {
		cntValue ++
		resp := client.Set(counterKey,cntValue,0)
		_,err := resp.Result()
		if err != nil {
			println("set value error")
		}
	}

	delResp := client.Del(lockKey)
	unlockSuccess,err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed",err)
	}
}
*/