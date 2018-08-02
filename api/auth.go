package main

import (
	"net/http"
	"github.com/Zereker/video_server/api/session"
)

var HeaderFiledSession = "X-Session-Id"
var HeaderFiledUsername = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sessionId := r.Header.Get(HeaderFiledSession)
	if len(sessionId) == 0 {
		return false
	}
	username, ok := session.IsSessionExpired(sessionId)
	if ok {
		return false
	}
	r.Header.Add(HeaderFiledSession, username)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HeaderFiledUsername)
	if len(username) == 0 {
		return false
	}
	return true
}
