package models

import (
	"advBridge/apiForms"
	"errors"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type VmsSyncRecordsModel struct {
}

type VmsSyncRecords struct {
	ID                             bson.ObjectId `json:"_id" bson:"_id"`
	SyncVmsDataCounts              int32         `json:"syncVmsDataCounts" bson:"syncVmsDataCounts"`
	RFIDDataSendCounts             int32         `json:"RFIDDataSendCounts" bson:"RFIDDataSendCounts"`
	Status                         string        `json:"status" bson:"status"`
	VMSServerProtocol              string        `json:"VMSServer_Protocol" bson:"VMSServer_Protocol"`
	VMSServerHost                  string        `json:"VMSServer_Host" bson:"VMSServer_Host"`
	RFIDServerMqttConnectionString string        `json:"RFIDServer_MqttConnectionString" bson:"RFIDServer_MqttConnectionString"`
	RFIDServerMqttTopic            string        `json:"RFIDServer_MqttTopic" bson:"RFIDServer_MqttTopic"`
	CreateUnixTimeStamp            int64         `json:"createUnixTimeStamp" bson:"createUnixTimeStamp"`
}

var vmsSyncRecordsModel = new(VmsSyncRecordsModel)

func (m *VmsSyncRecordsModel) GenerateNewInstance() (objectIdRoot bson.ObjectId, err error) {
	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)

	protocol := globalConfig.Bundle["VMSServer_Protocol"].(string)
	host := globalConfig.Bundle["VMSServer_Host"].(string)
	mqttConnectionString := globalConfig.Bundle["RFIDServer_MqttConnectionString"].(string)
	RFIDTopic = globalConfig.Bundle["RFIDServer_MqttTopic"].(string)

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_VMS_SYNC_RECORDS)
	defer collection.Database.Session.Close()

	objectIdRoot = bson.NewObjectId()
	err = collection.Insert(bson.M{
		"_id":                             objectIdRoot,
		"syncVmsDataCounts":               0,
		"RFIDDataSendCounts":              0,
		"status":                          "Sync-Processing",
		"failReason":                      "",
		"VMSServer_Protocol":              protocol,
		"VMSServer_Host":                  host,
		"RFIDServer_MqttConnectionString": mqttConnectionString,
		"RFIDServer_MqttTopic":            RFIDTopic,
		"createUnixTimeStamp":             time.Now().Unix(),
	})

	if err != nil {
		logv.Error(err.Error())
		return objectIdRoot, err
	}
	return objectIdRoot, err
}

func (m *VmsSyncRecordsModel) UpdateStatus(recordsUUID string, status string, reason string) (err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_VMS_SYNC_RECORDS)
	defer collection.Database.Session.Close()

	err = collection.UpdateId(bson.ObjectIdHex(recordsUUID), bson.M{"$set": bson.M{
		"status": status,
		"failReason": reason,
	}})
	if err != nil {
		logv.Error(err.Error())
		return errors.New(err.Error())
	}
	return err
}

func (m *VmsSyncRecordsModel) CreateKioskLocation(data apiForms.KioskLocationCreateDataValidate) (err error) {
	collectionKioskDevice := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collectionKioskDevice.Database.Session.Close()

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_KIOSK_LOCATION)
	defer collection.Database.Session.Close()

	kioskDevice := KioskDeviceInfo{}

	if _IsObjectIdHex := bson.IsObjectIdHex(*data.DeviceUUID); !_IsObjectIdHex {
		err = errors.New("invalid input to ObjectIdHex: " + *data.DeviceUUID)
		logv.Error(err.Error())
		return errors.New(err.Error())
	}

	err = collectionKioskDevice.FindId(bson.ObjectIdHex(*data.DeviceUUID)).One(&kioskDevice)
	if err != nil {
		logv.Error(err.Error())
		return errors.New("Kiosk Device UUID is not exist:> " + *data.DeviceUUID)
	}

	err = collection.Find(bson.M{"deviceUUID": data.DeviceUUID}).One(&kioskDevice)
	if err == nil {
		logv.Error("this location is exist:> ", data.DeviceUUID)
		return errors.New("this location is exist:> " + *data.DeviceUUID)
	}

	objectIdRoot := bson.NewObjectId()
	err = collection.Insert(bson.M{
		"_id":                 objectIdRoot,
		"deviceUUID":          data.DeviceUUID,
		"location":            data.Location,
		"createUnixTimeStamp": time.Now().Unix(),
	})

	if err != nil {
		logv.Error(err.Error())
		return err
	}
	return err
}

func (m *VmsSyncRecordsModel) FetchAllKioskLocation() (results []KioskLocation, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_KIOSK_LOCATION)
	defer collection.Database.Session.Close()

	err = collection.Find(bson.M{}).All(&results)

	if err != nil {
		logv.Error(err.Error())
		return results, err
	}
	return results, err
}
