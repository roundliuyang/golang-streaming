package main

import (
	"github.com/alanhou/golang-streaming/video_server/api/defs"
	"github.com/alanhou/golang-streaming/video_server/api/session"
	"net/http"
	"log"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"
// var HEADER_FIELD_SESSION = "session"
// var HEADER_FIELD_UNAME = "username"

func ValidateUserSession(r *http.Request) bool  {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}


func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	// sid := r.Header.Get(HEADER_FIELD_SESSION)
	log.Printf("GetUserInfo:ValidateUser: %n", uname)
	if len(uname) == 0 {
	// if len(sid) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}