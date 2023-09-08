package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)

	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)

	router.POST("/userhome", userHomeHandler)

	router.POST("/api", apiHandler)

	router.GET("/videos/:vidd-id", proxyVideoHandler)

	router.POST("/upload/:vid-id", proxyUploadHandler)

	//router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	router.ServeFiles("/statics/img/*filepath", http.Dir("./templates/img"))

	router.ServeFiles("/videos/*filepath", http.Dir("../streamserver/videos"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
