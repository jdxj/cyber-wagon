package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	gateway "github.com/jdxj/cyber-wagon/internal/gateway/proto"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/transport", nil)
	if err != nil {
		logrus.Fatalf("dial err: %s", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					return
				}

				logrus.Errorf("read err: %s", err)
				return
			}

			sMsg := &gateway.ServerMsg{}
			_ = proto.Unmarshal(message, sMsg)
			logrus.Infof("recv: %s", sMsg.GetTimestamp())
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			cMsg := &gateway.ClientMsg{
				Uuid:       "abc",
				Device:     "def",
				Timestamp:  timestamppb.Now(),
				AppVersion: "ghi",
				Data:       nil,
			}
			d, _ := proto.Marshal(cMsg)
			err := c.WriteMessage(websocket.BinaryMessage, d)
			if err != nil {
				logrus.Infof("write err: %s", err)
				return
			}
		case <-interrupt:
			logrus.Infof("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logrus.Errorf("write close err: %s", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
