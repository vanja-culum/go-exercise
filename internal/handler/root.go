package handler

import (
	"net/http"
)

type RootHandler struct {

}

func NewRootHandler() RootHandler {
	return RootHandler{}
}



func (h RootHandler) HandleGetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}