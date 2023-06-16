package qcsocket

import (
	"admin-server/internal/consts"
	"admin-server/pkg/goredis"
	"admin-server/pkg/logger"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type SubData struct {
	Type int64       `json:"type"`
	Data interface{} `json:"data"`
}

func Ws(c context.Context, connect *websocket.Conn) {

	SubChannel := make(chan SubData)

	pubsub := goredis.GetRedisDBWithContext(c).Subscribe(c, consts.RedisWebSocketWord)
	// Close the subscription when we are done.
	defer pubsub.Close()

	log := logger.GetLoggerByCtx(c).With(
		zap.String("log_from", "websocket"),
	)

	ch := pubsub.Channel()
	go func() {
		for msg := range ch {
			temp := SubData{}

			if err := json.Unmarshal([]byte(msg.Payload), &temp); err != nil {
				log.Error("解析Unmarshal错误:" + err.Error())
			}
			SubChannel <- temp
		}
	}()

	for {
		data := <-SubChannel
		dataByte, err := json.Marshal(data)
		if err != nil {
			log.Error("解析Unmarshal错误:" + err.Error())
		}
		log.Info("websocket发送数据", zap.String("data", string(dataByte)))
		err = connect.WriteMessage(websocket.TextMessage, dataByte)
		if err != nil {
			log.Error(err.Error(), zap.String("data", string(dataByte)))
			break
		}
	}

}
