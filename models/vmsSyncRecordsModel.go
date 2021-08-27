package models

import (
	"advBridge/apiForms"
	"errors"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type VmsSyncRecordsModel struct {
}

// DataStructure
type VmsSyncRecords struct {
	ID                             bson.ObjectId `json:"_id" bson:"_id"`
	SyncVmsDataCounts              int32         `json:"syncVmsDataCounts" bson:"syncVmsDataCounts"`
	RFIDDataSendCounts             int32         `json:"RFIDDataSendCounts" bson:"RFIDDataSendCounts"`
	Status                         string        `json:"status" bson:"status"`
	FailReason                     string        `json:"failReason" bson:"failReason"`
	VMSServerProtocol              string        `json:"VMSServer_Protocol" bson:"VMSServer_Protocol"`
	VMSServerHost                  string        `json:"VMSServer_Host" bson:"VMSServer_Host"`
	RFIDServerMqttConnectionString string        `json:"RFIDServer_MqttConnectionString" bson:"RFIDServer_MqttConnectionString"`
	RFIDServerMqttTopic            string        `json:"RFIDServer_MqttTopic" bson:"RFIDServer_MqttTopic"`
	CreateUnixTimeStamp            int64         `json:"createUnixTimeStamp" bson:"createUnixTimeStamp"`
}

//type VmsSyncRecordsDetail struct {
//	ID                             bson.ObjectId `json:"_id" bson:"_id"`
//	SyncVmsDataCounts              int32         `json:"syncVmsDataCounts" bson:"syncVmsDataCounts"`
//	RFIDDataSendCounts             int32         `json:"RFIDDataSendCounts" bson:"RFIDDataSendCounts"`
//	Status                         string        `json:"status" bson:"status"`
//	FailReason                     string        `json:"failReason" bson:"failReason"`
//	VMSServerProtocol              string        `json:"VMSServer_Protocol" bson:"VMSServer_Protocol"`
//	VMSServerHost                  string        `json:"VMSServer_Host" bson:"VMSServer_Host"`
//	RFIDServerMqttConnectionString string        `json:"RFIDServer_MqttConnectionString" bson:"RFIDServer_MqttConnectionString"`
//	RFIDServerMqttTopic            string        `json:"RFIDServer_MqttTopic" bson:"RFIDServer_MqttTopic"`
//	CreateUnixTimeStamp            int64         `json:"createUnixTimeStamp" bson:"createUnixTimeStamp"`
//
//}

//

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
		"status":     status,
		"failReason": reason,
	}})
	if err != nil {
		logv.Error(err.Error())
		return errors.New(err.Error())
	}
	return err
}

// VMS 同步資訊
func (m *VmsSyncRecordsModel) ListDataByP(data apiForms.ListByPVmsSyncRecordsDataValidate) (resultsByPage []VmsSyncRecords,
	resultsTotal []VmsSyncRecords, err error, errcode int) {
	collectionKD := dbConnect.UseTable(DB_Name, DB_Table_ADV_VMS_SYNC_RECORDS)
	defer collectionKD.Database.Session.Close()

	isDESC := 1
	if *data.Desc {
		isDESC = 1
	} else {
		isDESC = -1
	}

	match_stage := bson.M{}

	pipeline := []bson.M{
		{
			"$match": match_stage,
		},
		{
			"$sort": bson.M{
				*data.SortBy: isDESC,
			},
		},
	}

	pipe := collectionKD.Pipe(pipeline)
	err = pipe.All(&resultsTotal)
	if err != nil {
		logv.Error("Find Response FindId err:> ", err)
	}

	if *data.Count != -1 {
		for i := 0; i < *data.Count; i++ {
			if *data.StartIndex+i >= len(resultsTotal) {
				break
			}
			resultsByPage = append(resultsByPage, resultsTotal[*data.StartIndex+i])
		}
	} else {
		return resultsTotal, resultsTotal, err, 0
	}

	return resultsByPage, resultsTotal, err, 0
}

// VMS 同步資訊，詳細資料
func (m *VmsSyncRecordsModel) GetDetailDataByP(data apiForms.ListByPVmsSyncRecordsDetailDataValidate) (resultsByPage []KioskReportResponse,
	resultsTotal []KioskReportResponse, vmsSyncRecord VmsSyncRecords, err error, errcode int) {

	collectionSR := dbConnect.UseTable(DB_Name, DB_Table_ADV_VMS_SYNC_RECORDS)
	defer collectionSR.Database.Session.Close()

	err = collectionSR.FindId(bson.ObjectIdHex(*data.RecordUUID)).One(&vmsSyncRecord)
	if err != nil {
		logv.Error(err.Error())
		return resultsTotal, resultsTotal, vmsSyncRecord, err, 0
	}
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_REPORTS)
	defer collection.Database.Session.Close()

	isDESC := 1
	if *data.Desc {
		isDESC = 1
	} else {
		isDESC = -1
	}

	match_stage := bson.M{
		"recordsUUID": data.RecordUUID,
	}

	pipeline := []bson.M{
		{
			"$match": match_stage,
		},
		{
			"$sort": bson.M{
				*data.SortBy: isDESC,
			},
		},
	}

	pipe := collection.Pipe(pipeline)
	err = pipe.All(&resultsTotal)
	if err != nil {
		logv.Error("Find Response FindId err:> ", err)
	}

	if *data.Count != -1 {
		for i := 0; i < *data.Count; i++ {
			if *data.StartIndex+i >= len(resultsTotal) {
				break
			}
			resultsByPage = append(resultsByPage, resultsTotal[*data.StartIndex+i])
		}
	} else {
		return resultsTotal, resultsTotal, vmsSyncRecord, err, 0
	}

	return resultsByPage, resultsTotal, vmsSyncRecord, err, 0
}

func (m *VmsSyncRecordsModel) CheckVMSRecordsRetention() (info *mgo.ChangeInfo, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_VMS_SYNC_RECORDS)
	defer collection.Database.Session.Close()

	//collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	//defer collectionConfig.Database.Session.Close()
	//
	//var globalConfig GlobalConfig
	//err = collectionConfig.Find(bson.M{}).One(&globalConfig)
	//log_retention := globalConfig.Bundle["log_retention"]
	log_retention := "30"

	log_retention_target, _ := strconv.ParseInt(log_retention, 10, 64)

	//logv.Info(log_retention_target)
	//logv.Info("log_retention:> ", log_retention)
	//logv.Info(snapshot_retention)

	timestamp := time.Now().Unix()

	info, err = collection.RemoveAll(bson.M{"createUnixTimeStamp": bson.M{"$lte": timestamp - 24*60*60* int64(log_retention_target)}})
	if err != nil {
		logv.Error("Update CheckVMSRecordsRetention warn:> ", err)
	}
	return info, err
}
