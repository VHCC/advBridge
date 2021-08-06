package models

import (
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"net"
	"os"
	"time"
)

type GlobalConfig struct {
	ID                        bson.ObjectId          `json:"-" bson:"_id"`
	Bundle                    map[string]interface{} `json:"bundle" bson:"bundle"`
	LastModifiedUnixTimeStamp int64                  `json:"lastModifiedUnixTimeStamp" bson:"lastModifiedUnixTimeStamp"`
}

type GlobalConfigModel struct{}

var serverIp = ""

func initGlobalConfigModel() (err error) {
	logv.Info("GlobalConfigModel init()")
	globalConfig, err := FindFromDB()
	_ = globalConfig
	if err != nil {
		logv.Error("GlobalConfigModel init() err:> ", err)
		globalConfig, err = InsertToDB()
	}
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			logv.Info("IPv4: ", ipv4)
			serverIp = ipv4.String()
		}
	}
	logv.Info(getMacAddr())
	return err
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func getMacAddrWithIp() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a+" / IP:"+serverIp)
		}
	}
	return as, nil
}

func (m *GlobalConfigModel) ListMacAddress(isWithIP bool) (response []string, err error) {
	if isWithIP {
		return getMacAddrWithIp()
	} else {
		return getMacAddr()
	}
}

func InsertToDB() (response map[string]interface{}, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collection.Database.Session.Close()

	objectId := bson.NewObjectId()
	err = collection.Insert(bson.M{
		"_id": objectId,
		"bundle": bson.M{
			//"serverinternetaddress": "https://vms.ichenprocin.dsmynas.com",
			//"serverinternetaddress": "172.22.28.1", //spec
			//"smtp":                  "",
			//"port":                  "587",
			//"user_registration":     true,
			//"checkin_retention":     "60",
			//"log_retention":         "365",
			//"snapshot_retention":    "60",

			"VMSServer_Protocol": "http",
			"VMSServer_Host": "vms:80",
			"VMSServer_Account": "",
			"VMSServer_Password": "",

			"RFIDServer_MqttConnectionString": "tcp://104.215.147.159:1883",
			"RFIDServer_MqttTopic": "rfid_temp",
			"RFIDServer_Username": "ec1aceb8-88aa-4b60-8cff-4e8e1cae9e5f:e325b491-edc1-4019-a4e8-675b7c80852c",
			"RFIDServer_Password": "1JFoR3YbyGaGfNGPGg19Flqzy",

			"HRServer_SQLServerHost": "172.20.2.85",
			"HRServer_Account":       "rfiduser",
			"HRServer_Password":      "rf!dus1r375",
			"HRServer_DatabaseName":  "RFID",
			"HRServer_ViewTableName": "RFID_Employee",
		},
		"lastModifiedUnixTimeStamp": time.Now().Unix(),
	})
	var globalConfig GlobalConfig
	if err != nil {
		logv.Error("InsertToDB Response Insert err:> ", err)
		return response, err
		//panic(err)
	}

	err = collection.FindId(objectId).One(&globalConfig)
	if err != nil {
		logv.Error("InsertToDB Response FindId err:> ", err)
		//panic(err)
	}
	logv.Info(globalConfig)

	response = make(map[string]interface{})
	if len(globalConfig.Bundle) != 0 {
		response = globalConfig.Bundle
	}
	response["lastModifiedUnixTimeStamp"] = globalConfig.LastModifiedUnixTimeStamp
	return response, err
}

// find the only one Server Config
func (m *GlobalConfigModel) FindFromDB() (response map[string]interface{}, err error) {
	response, err = FindFromDB()
	return response, err
}

func FindFromDB() (response map[string]interface{}, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collection.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collection.Find(bson.M{}).One(&globalConfig)
	if err != nil {
		logv.Error("FindFromDB Response FindId err:> ", err)
		return response, err
	}

	response = make(map[string]interface{})
	if len(globalConfig.Bundle) != 0 {
		response = globalConfig.Bundle
	}
	response["lastModifiedUnixTimeStamp"] = globalConfig.LastModifiedUnixTimeStamp

	return response, err
}
