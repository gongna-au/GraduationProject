const mqtt = require('mqtt')
const client = mqtt.connect('mqtt://localhost:1883')

client.on('connect', function () {
  setInterval(function () {
    const message = JSON.stringify({ temperature: 22.5, humidity: 60 })
    client.publish('sensor/data', message)
    console.log('Message Sent', message)
  }, 5000)
})
