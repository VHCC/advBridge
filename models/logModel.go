package models

import (
	"advBridge/apiForms"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type BridgeLog struct {
	ID                  bson.ObjectId          `json:"_id" bson:"_id"`
	AccountID           string                 `json:"accountID" bson:"accountID"`
	LogType             string                 `json:"logType" bson:"logType"`
	Message             string                 `json:"message" bson:"message"`
	Detail              map[string]interface{} `json:"detail" bson:"detail"`
	IsDevice            bool                   `json:"-" bson:"isDevice"`
	DeviceUUID          string                 `json:"-" bson:"deviceUUID"`
	CreateUnixTimestamp int64                  `json:"createUnixTimestamp" bson:"createUnixTimestamp"`
	FetchTimestamp      int64                  `json:"-" bson:"fetchTimestamp"`
}

type VmsLogModel struct{}

func (m *VmsLogModel) WriteLog(eventType int, account string, message string, detail map[string]interface{}) (err error) {
	return WriteLog(eventType, account, message, detail )
}


func WriteLog(eventType int, account string, message string, detail map[string]interface{}) (err error) {
	collection := dbConnect.UseTable(DB_Name, DB_ADV_BRIDGE_Log)
	defer collection.Database.Session.Close()

	objectId := bson.NewObjectId()

	eventString := ""

	switch eventType {
	case EVENT_TYPE_USER_LOGIN:
		eventString = EVENT_TYPE_USER_LOGIN_TYPE
	case EVENT_TYPE_USER_LOGIN_FAIL:
		eventString = EVENT_TYPE_USER_LOGIN_FAIL_TYPE
	case EVENT_TYPE_USER_LOGOUT:
		eventString = EVENT_TYPE_USER_LOGOUT_TYPE
	case EVENT_TYPE_USER_LOGOUT_FAIL:
		eventString = EVENT_TYPE_USER_LOGOUT_FAIL_TYPE
	case EVENT_TYPE_COMPANY_EDIT:
		eventString = EVENT_TYPE_COMPANY_EDIT_TYPE
	case EVENT_TYPE_TEMPLATE_CREATE:
		eventString = EVENT_TYPE_TEMPLATE_CREATE_TYPE
	case EVENT_TYPE_TEMPLATE_EDIT:
		eventString = EVENT_TYPE_TEMPLATE_EDIT_TYPE
	case EVENT_TYPE_TEMPLATE_DELETE:
		eventString = EVENT_TYPE_TEMPLATE_DELETE_TYPE
	case EVENT_TYPE_KIOSK_DEVICE_CONNECT:
		eventString = EVENT_TYPE_KIOSK_DEVICE_CONNECT_TYPE
	case EVENT_TYPE_KIOSK_DEVICE_EDIT:
		eventString = EVENT_TYPE_KIOSK_DEVICE_EDIT_TYPE
	case EVENT_TYPE_KIOSK_DEVICE_REMOVE:
		eventString = EVENT_TYPE_KIOSK_DEVICE_REMOVE_TYPE
	case EVENT_TYPE_PERSON_CREATE:
		eventString = EVENT_TYPE_PERSON_CREATE_TYPE
	case EVENT_TYPE_PERSON_EDIT:
		eventString = EVENT_TYPE_PERSON_EDIT_TYPE
	case EVENT_TYPE_PERSON_DELETE:
		eventString = EVENT_TYPE_PERSON_DELETE_TYPE
	case EVENT_TYPE_CHECK_IN_REPORTS_READ:
		eventString = EVENT_TYPE_CHECK_IN_REPORTS_READ_TYPE
	case EVENT_TYPE_ATTENDANCE_READ:
		eventString = EVENT_TYPE_ATTENDANCE_READ_TYPE
	case EVENT_TYPE_PERSON_IMPORT_BATCH:
		eventString = EVENT_TYPE_PERSON_IMPORT_BATCH_TYPE
	case EVENT_TYPE_USER_CREATE:
		eventString = EVENT_TYPE_USER_CREATE_TYPE
	case EVENT_TYPE_USER_EDIT:
		eventString = EVENT_TYPE_USER_EDIT_TYPE
	case EVENT_TYPE_USER_DELETE:
		eventString = EVENT_TYPE_USER_DELETE_TYPE
	case EVENT_TYPE_SMTP_TEST:
		eventString = EVENT_TYPE_SMTP_TEST_TYPE
	case EVENT_TYPE_RETENTION_UPDATE:
		eventString = EVENT_TYPE_RETENTION_UPDATE_TYPE
	case EVENT_TYPE_SMTP_UPDATE:
		eventString = EVENT_TYPE_SMTP_UPDATE_TYPE
	case EVENT_TYPE_LICENSE_REGISTER:
		eventString = EVENT_TYPE_LICENSE_REGISTER_TYPE
	case EVENT_TYPE_LOG_EXPORT:
		eventString = EVENT_TYPE_LOG_EXPORT_TYPE
	}

	// bridge
	switch eventType {
	case EVENT_TYPE_VMS_SERVER_UPDATE: eventString = EVENT_TYPE_VMS_SERVER_UPDATE_TYPE
	case EVENT_TYPE_HR_SERVER_UPDATE: eventString = EVENT_TYPE_HR_SERVER_UPDATE_TYPE
	case EVENT_TYPE_RFID_SERVER_UPDATE: eventString = EVENT_TYPE_RFID_SERVER_UPDATE_TYPE
	case EVENT_TYPE_KIOSK_LOCATION_CREATE: eventString = EVENT_TYPE_KIOSK_LOCATION_CREATE_TYPE
	case EVENT_TYPE_KIOSK_LOCATION_UPDATE: eventString = EVENT_TYPE_KIOSK_LOCATION_UPDATE_TYPE
	case EVENT_TYPE_KIOSK_LOCATION_DELETE: eventString = EVENT_TYPE_KIOSK_LOCATION_DELETE_TYPE
	case EVENT_TYPE_BRIDGE_LOG_RETENTION_CHECK: eventString = EVENT_TYPE_BRIDGE_LOG_RETENTION_CHECK_TYPE
	case EVENT_TYPE_RFID_SERVER_CONNECT_SUCCESS: eventString = EVENT_TYPE_RFID_SERVER_CONNECT_SUCCESS_TYPE
	case EVENT_TYPE_VMS_SERVER_CONNECT_SUCCESS: eventString = EVENT_TYPE_VMS_SERVER_CONNECT_SUCCESS_TYPE
	case EVENT_TYPE_HR_SERVER_CONNECT_SUCCESS: eventString = EVENT_TYPE_HR_SERVER_CONNECT_SUCCESS_TYPE
	case EVENT_TYPE_RFID_SERVER_CONNECT_FAIL: eventString = EVENT_TYPE_RFID_SERVER_CONNECT_FAIL_TYPE
	case EVENT_TYPE_VMS_SERVER_CONNECT_FAIL: eventString = EVENT_TYPE_VMS_SERVER_CONNECT_FAIL_TYPE
	case EVENT_TYPE_HR_SERVER_CONNECT_FAIL: eventString = EVENT_TYPE_HR_SERVER_CONNECT_FAIL_TYPE
	case EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_START: eventString = EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_START_TYPE
	case EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_DONE: eventString = EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_DONE_TYPE
	case EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_FAIL: eventString = EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_FAIL_TYPE
	case EVENT_TYPE_VMS_KIOSK_DEVICE_SYNC_SUCCESS: eventString = EVENT_TYPE_VMS_KIOSK_DEVICE_SYNC_SUCCESS_TYPE
	case EVENT_TYPE_VMS_KIOSK_DEVICE_SYNC_FAIL: eventString = EVENT_TYPE_VMS_KIOSK_DEVICE_SYNC_FAIL_TYPE
	case EVENT_TYPE_HR_SERVER_SYNC_FAIL: eventString = EVENT_TYPE_HR_SERVER_SYNC_FAIL_TYPE
	case EVENT_TYPE_HR_SERVER_SYNC_SUCCESS: eventString = EVENT_TYPE_HR_SERVER_SYNC_SUCCESS_TYPE
	case EVENT_TYPE_HR_RECORDS_RETENTION_CHECK: eventString = EVENT_TYPE_HR_RECORDS_RETENTION_CHECK_TYPE
	case EVENT_TYPE_VMS_RECORDS_RETENTION_CHECK: eventString = EVENT_TYPE_VMS_RECORDS_RETENTION_CHECK_TYPE


	}


	err = collection.Insert(bson.M{
		"_id":                 objectId,
		"logType":             eventString,
		"accountID":           account,
		"message":             message,
		"detail":              detail,
		"isDevice":            false,
		"createUnixTimestamp": time.Now().Unix(),
	})
	if err != nil {
		logv.Error("Write Log Insert err:> ", err)
	}
	return err
}

	func (m *VmsLogModel) ListLogByP(data apiForms.ListByPBridgeLogDataValidate) (vmsLog []BridgeLog, vmsLogTotal []BridgeLog,
	err error, errcode int) {
	collectionKD := dbConnect.UseTable(DB_Name, DB_ADV_BRIDGE_Log)
	defer collectionKD.Database.Session.Close()

	isDESC := 1
	if *data.Desc {
		isDESC = 1
	} else {
		isDESC = -1
	}

	match_stage := bson.M{}

	if data.KeyWords == nil {
		match_stage = bson.M{
			"isDevice": false,
		}
	} else {
		match_stage = bson.M{
			"isDevice": false,
			"$or": []bson.M{
				bson.M{
					"accountID": bson.RegEx{*data.KeyWords, ""},
				},
			},
		}
	}

	if data.LogTypes != nil {
		//match_stage["avalo_visitor"] = *data.TemplateUUID
		if match_stage["$or"] == nil {
			match_stage["$or"] = []bson.M{}
		}

		logTypes := *data.LogTypes

		for i := 0; i < len(*data.LogTypes); i++ {
			match_stage["$or"] = append(match_stage["$or"].([]bson.M), bson.M{"logType": logTypes[i]})
		}
	}

	if data.StartTimestamp != nil && data.EndTimestamp != nil {
		match_stage["createUnixTimestamp"] = bson.M{"$gte": *data.StartTimestamp, "$lte": *data.EndTimestamp}
	} else if data.StartTimestamp != nil {
		match_stage["createUnixTimestamp"] = bson.M{"$gte": *data.StartTimestamp}
	} else if data.EndTimestamp != nil {
		match_stage["createUnixTimestamp"] = bson.M{"$lte": *data.EndTimestamp}
	}

	logv.Info(match_stage)

	//logv.Info(data)
	//logv.Info(match_stage)
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
	err = pipe.All(&vmsLogTotal)
	if err != nil {
		logv.Error("Find Response FindId err:> ", err)
	}

	if *data.Count != -1 {
		vmsLog = []BridgeLog{}
		for i := 0; i < *data.Count; i++ {
			if *data.StartIndex+i >= len(vmsLogTotal) {
				break
			}
			vmsLog = append(vmsLog, vmsLogTotal[*data.StartIndex+i])
		}
		return vmsLog, vmsLogTotal, err, 0
	} else {
		return vmsLogTotal, vmsLogTotal, err, 0
	}
}

//// ============ DEVICE LOG ================
//func (m *VmsLogModel) WriteDeviceLog(eventType string, account string, message string, deviceTime string, deviceUUID string, detail map[string]interface{}) (err error) {
//
//	collection := dbConnect.UseTable(DB_Name, DB_VMS_Log)
//	defer collection.Database.Session.Close()
//
//	objectId := bson.NewObjectId()
//
//	eventString := eventType
//	logv.Info(deviceTime)
//	timestamp, err := strconv.ParseInt(deviceTime, 10, 64)
//	logv.Info(timestamp)
//	err = collection.Insert(bson.M{
//		"_id":                 objectId,
//		"logType":             eventString,
//		"accountID":           account,
//		"message":             message,
//		"detail":              detail,
//		"isDevice":            true,
//		"deviceUUID":          deviceUUID,
//		"createUnixTimestamp": timestamp,
//		"fetchTimestamp":      time.Now().Unix(),
//	})
//	if err != nil {
//		logv.Error("Write Log Insert err:> ", err)
//	}
//	return err
//}

func (m *VmsLogModel) CheckVmsLogRetention() (info *mgo.ChangeInfo, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_ADV_BRIDGE_Log)
	defer collection.Database.Session.Close()

	//collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	//defer collectionConfig.Database.Session.Close()
	//
	//var globalConfig GlobalConfig
	//err = collectionConfig.Find(bson.M{}).One(&globalConfig)
	//log_retention := globalConfig.Bundle["log_retention"]
	log_retention := "365"

	log_retention_target, _ := strconv.ParseInt(log_retention, 10, 64)

	//logv.Info(log_retention_target)
	logv.Info("log_retention:> ", log_retention)
	//logv.Info(snapshot_retention)

	timestamp := time.Now().Unix()

	info, err = collection.RemoveAll(bson.M{"createUnixTimestamp": bson.M{"$lte": timestamp - 24*60*60 * int64(log_retention_target)}})
	if err != nil {
		logv.Error("Update CheckVmsLogRetention warn:> ", err)
	}
	return info, err
}

