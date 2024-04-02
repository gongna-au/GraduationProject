import paho.mqtt.client as mqtt
import time

broker_address = "emqx"
client = mqtt.Client("PythonClient")
client.connect(broker_address)

while True:
    message = "Hello EMQX from Python"
    client.publish("testtopic/1", message)
    print(f"Sent message: {message}")
    time.sleep(5)  # Sends a message every 5 seconds
