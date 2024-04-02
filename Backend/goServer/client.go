package main

import (
	"fmt"
	"net/http"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

var messages []string // 用于存储接收到的消息
var mu sync.Mutex     // 用于同步访问`messages`

// MQTT消息处理器
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	mu.Lock()
	messages = append(messages, string(msg.Payload())) // 存储接收到的消息
	mu.Unlock()
}

func main() {
	// MQTT客户端设置
	opts := mqtt.NewClientOptions().AddBroker("tcp://emqx:1883").SetClientID("GoClient")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("testtopic/1", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	// Gin路由设置
	router := gin.Default()
	router.GET("/messages", func(c *gin.Context) {
		mu.Lock()
		c.JSON(http.StatusOK, messages) // 通过HTTP接口展示消息
		mu.Unlock()
	})

	// 启动Gin服务
	router.Run(":8080") // 监听在8080端口
}
