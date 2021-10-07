package controllers

import "net/http"

const (
	pong = "pong"
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

var (
	PingController pingControllerInterface = &pingController{}
)

type pingController struct{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader((http.StatusOK))
	w.Write([]byte(pong))
}
