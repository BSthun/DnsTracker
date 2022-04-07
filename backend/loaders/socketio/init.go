package socketio

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/sirupsen/logrus"

	"backend/utils/logger"
)

var Server *socketio.Server

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func Init() {
	Server = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	Server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("SOCKET CONNECTION:", s.ID())
		s.Emit("reply", "aaaa")
		return nil
	})

	go func() {
		logger.Log(logrus.Info, "STARTING SOCKETIO SERVER")

		if err := Server.Serve(); err != nil {
			logger.Log(logrus.Fatal, "FAILED TO SERVE SOCKETIO")
		}
	}()
}
