package main

import (
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/robfig/cron"
	"log"
)

func main() {
	m := linkedhashmap.New()
	m.Put(1,2)

	set := hashset.New()   // empty
	set.Add(1)

	c := cron.New()

	c.AddFunc("@every 1s",func() {
		log.Println("hello world")
	})
	c.Start()
	select {
	}
}