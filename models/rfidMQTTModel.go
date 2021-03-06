package models

import (
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type RFIDMQTTModel struct {
	mqttClient mqtt.Client
}

var mainMqttClient mqtt.Client

var rfidMQTTModel = new(RFIDMQTTModel)

var RFIDTopic = "rfid_temp"

func (m *RFIDMQTTModel) ConnectionToRFIDServer() (err error) {
	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)

	mqttConnectionString := globalConfig.Bundle["RFIDServer_MqttConnectionString"].(string)
	RFIDTopic = globalConfig.Bundle["RFIDServer_MqttTopic"].(string)
	rfidServerUsername := globalConfig.Bundle["RFIDServer_Username"].(string)
	rfidServerPassword := globalConfig.Bundle["RFIDServer_Password"].(string)

	//var broker = "104.215.147.159"
	//var port = 1883
	opts := mqtt.NewClientOptions()
	//opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.AddBroker(mqttConnectionString)

	opts.SetClientID(bson.NewObjectId().Hex())
	opts.SetUsername(rfidServerUsername)
	opts.SetPassword(rfidServerPassword)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logv.Error(token.Error().Error())
		return errors.New(token.Error().Error())
	}
	logv.Info(" ==== mqttClient connected SUCCESS ==== ")
	//m.mqttClient = client
	mainMqttClient = client
	return nil
}

func (m *RFIDMQTTModel) DisconnectionToRFIDServer() () {
	if mainMqttClient != nil {
		mainMqttClient.Disconnect(0)
	}
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf(" --- Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println(" --- MQTT Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf(" --- MQTT Connect lost: %v", err)
}

type RFIDDataBody struct {
	ETAG        string `json:"etag"`
	TIME        string `json:"time"`
	ANT0        string `json:"ant0"`
	ANT1        string `json:"ant1"`
	ANT2        string `json:"ant2"`
	ANT3        string `json:"ant3"`
	MACADDR     string `json:"macaddr"`
	TEMPERATURE string `json:"temperature"`
}

func (m *RFIDMQTTModel) PublishToRFIDServerTest() (){
	if mainMqttClient == nil || !mainMqttClient.IsConnected() {
		logv.Info("QQQQ")
	}
	RFIDData := RFIDDataBody{
		"000000000000000000B02514",
		time.Now().String(),
		"",
		"",
		"",
		"",
		"74:fe:48:1d:0c:11",
		"36.3",
	}
	RFIDData_json, _ := json.Marshal(RFIDData)
	err := mainMqttClient.Publish(RFIDTopic, 1, false, RFIDData_json)
	if err != nil {
	}
}

func (m *RFIDMQTTModel) PublishToRFIDServer(ETAG string, kioskUUID string, temp string) (err error){
	if mainMqttClient == nil || !mainMqttClient.IsConnected() {
		return errors.New("MQTT did not connected yet")
	}
	RFIDData := RFIDDataBody{
		ETAG,
		time.Now().String(),
		"",
		"",
		"",
		"",
		kioskUUID,
		temp,
	}
	RFIDDataJson, _ := json.Marshal(RFIDData)
	mainMqttClient.Publish(RFIDTopic, 1, false, RFIDDataJson)
	return err
}

// ======= DEBUG ========
func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish(RFIDTopic, 1, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	token := client.Subscribe(RFIDTopic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", RFIDTopic)
}

// ========== Test
func (m *RFIDMQTTModel) TestConnectionToRFIDServer(
	mqttConnectionString string,
	rfidServerUsername string,
	rfidServerPassword string,
	) (err error) {

	//var broker = "104.215.147.159"
	//var port = 1883
	opts := mqtt.NewClientOptions()
	//opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.AddBroker(mqttConnectionString)
	timeout, _ := time.ParseDuration("3s");
	opts.SetClientID(bson.NewObjectId().Hex())
	opts.SetUsername(rfidServerUsername)
	opts.SetPassword(rfidServerPassword)
	//opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetConnectTimeout(timeout)
	//opts.OnConnect = connectHandler
	//opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logv.Error(token.Error().Error())
		return errors.New(token.Error().Error())
	}
	return nil
}
