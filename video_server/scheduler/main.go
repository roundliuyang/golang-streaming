package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/alanhou/golang-streaming/video_server/scheduler/taskrunner"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	// 或者用for死循环
	// c :=make(chan int)
	go taskrunner.Start()
	r := RegisterHandlers()
	// <- c
	http.ListenAndServe(":9001", r)
}
