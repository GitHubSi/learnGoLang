package client

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/golang/glog"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	//本地调试，关闭origin校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Info(err)
		return
	}

	client := Client{conn: conn}
	go client.read()
	go client.write()
}
