package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	gateway "github.com/jdxj/cyber-wagon/internal/gateway/proto"
	"github.com/jdxj/cyber-wagon/internal/pkg/network/web"
)

var upgrader = websocket.Upgrader{}

func transport(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Errorf("upgrade err: %s", err)
		return
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				logrus.Infof("client closed")
				break
			}

			logrus.Errorf("read msg err: %s", err)
			break
		}

		cMsg := &gateway.ClientMsg{}
		_ = proto.Unmarshal(message, cMsg)
		logrus.Infof("recv: %s", cMsg.GetUuid())

		sMsg := &gateway.ServerMsg{
			Timestamp: timestamppb.Now(),
			Data:      nil,
		}
		d, _ := proto.Marshal(sMsg)
		err = c.WriteMessage(websocket.BinaryMessage, d)
		if err != nil {
			logrus.Errorf("write err: %s", err)
			break
		}
	}
}

func main() {
	web.Start(":8080", func(root gin.IRouter) {
		root.GET("/transport", transport)
	})
}
