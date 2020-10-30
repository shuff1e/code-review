package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	runtime.LockOSThread()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("worker:", unix.Getpid())
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:           ":18282",
		Handler:        router,
		ReadTimeout:    600 * time.Second,
		WriteTimeout:   600 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("listen:", unix.Getppid(),unix.Getegid())
	s.ListenAndServe()
}