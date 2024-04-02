package subscribe

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
)

// 定义接收到的消息结构
type ParkingData struct {
	SlotID        int    `json:"slot_id"`
	IsOccupied    bool   `json:"is_occupied"`
	OccupiedSince string `json:"occupied_since"`
	VehicleID     string `json:"vehicle_id"`
}

var db *sql.DB

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\n", msg.Topic())

	var data ParkingData
	if err := json.Unmarshal(msg.Payload(), &data); err != nil {
		log.Fatalf("Error unmarshalling message: %v", err)
		return
	}

	// 写入数据库
	_, err := db.Exec("INSERT INTO sensor_data (slot_id, is_occupied, occupied_since, vehicle_id) VALUES (?, ?, ?, ?)",
		data.SlotID, data.IsOccupied, data.OccupiedSince, data.VehicleID)
	if err != nil {
		log.Fatalf("Error inserting data into database: %v", err)
	}
	fmt.Println("Data inserted into database successfully.")
}

func Subscribe() {
	var err error
	// 连接数据库
	db, err = sql.Open("mysql", "test:123456@tcp(localhost:3306)/smart_parking")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// 配置MQTT客户端
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client_subscriber")
	opts.SetDefaultPublishHandler(messageHandler)

	// 创建并连接MQTT客户端
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}
	defer client.Disconnect(250)

	// 订阅主题
	if token := client.Subscribe("smart_parking/sensor_data", 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic: %v", token.Error())
	}

	// 保持程序运行
	for {
		time.Sleep(1 * time.Second)
	}
}
