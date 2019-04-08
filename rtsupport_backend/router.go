package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Router struct {
	rules map[string]Handler
}

type Handler func(*Client, interface{})

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

func NewRouter() *Router {
	return &Router {
		rules: make(map[string]Handler),
	}
}

func (r *Router) FindeHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}

func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(socket, e.FindeHandler)
	go client.Write()
	client.Read()
}