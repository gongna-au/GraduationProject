package publish

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish() {
	// MQTT客户端配置
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client")

	// 创建MQTT客户端
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)
	for {
		publishSensorData(client)
		fmt.Println("Data published. Waiting for next round...")
		time.Sleep(5 * time.Second) // 等待5秒
	}
	// 发布消息

}

// 模拟数据发送
func publishSensorData(client mqtt.Client) {
	topicPrefix := "smart_parking/"

	// 1. 停车场传感器数据
	parkingData := map[string]interface{}{
		"slot_id":        101,
		"is_occupied":    true,
		"occupied_since": time.Now().Format(time.RFC3339),
		"vehicle_id":     "ABC123",
	}
	publishMessage(client, topicPrefix+"sensor_data", parkingData)

	// 2. 车辆入口和出口记录
	vehicleEntryExitData := map[string]interface{}{
		"vehicle_id":       "ABC123",
		"timestamp":        time.Now().Format(time.RFC3339),
		"action":           "entry", // 或 "exit"
		"entry_exit_point": 1,
	}
	publishMessage(client, topicPrefix+"vehicle_entry_exit", vehicleEntryExitData)

	// 3. 用户交互数据
	userInteractionData := map[string]interface{}{
		"reservation_id":    456,
		"user_id":           789,
		"slot_id":           101,
		"reserved_from":     time.Now().Format(time.RFC3339),
		"reserved_until":    time.Now().Add(time.Hour * 1).Format(time.RFC3339),
		"payment_timestamp": time.Now().Format(time.RFC3339),
		"amount":            10.0,
		"payment_method":    "credit_card",
	}
	publishMessage(client, topicPrefix+"user_interaction", userInteractionData)

	// 4. 设备运行状态
	deviceStatusData := map[string]interface{}{
		"device_id":          "sensor_101",
		"device_type":        "sensor",
		"is_online":          true,
		"last_online":        time.Now().Format(time.RFC3339),
		"malfunction_report": "None",
	}
	publishMessage(client, topicPrefix+"device_status", deviceStatusData)

	sensorLogData := map[string]interface{}{
		"sensor_id": "sensor_101",
		"status":    "occupied",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	publishMessage(client, topicPrefix+"sensor_logs", sensorLogData)
}

// 发布消息到指定主题
func publishMessage(client mqtt.Client, topic string, payload interface{}) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshalling payload: %v\n", err)
		return
	}

	token := client.Publish(topic, 0, false, jsonPayload)
	token.Wait()
}
