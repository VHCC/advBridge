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

type HrSyncRecordsModel struct {
}

// DataStructure
type HrSyncRecords struct {
	ID                        bson.ObjectId `json:"_id" bson:"_id"`
	SyncVmsPersonDataCounts   int32         `json:"syncVmsPersonDataCounts" bson:"syncVmsPersonDataCounts"`
	UpdateVmsPersonDataCounts int32         `json:"updateVmsPersonDataCounts" bson:"updateVmsPersonDataCounts"`
	CreateVmsPersonDataCounts int32         `json:"createVmsPersonDataCounts" bson:"createVmsPersonDataCounts"`
	DeleteVmsPersonDataCounts int32         `json:"deleteVmsPersonDataCounts" bson:"deleteVmsPersonDataCounts"`
	SyncHrServerDataCounts    int32         `json:"syncHrServerDataCounts" bson:"syncHrServerDataCounts"`
	Status                    string        `json:"status" bson:"status"`
	FailReason                string        `json:"failReason" bson:"failReason"`
	VMSServerProtocol         string        `json:"VMSServer_Protocol" bson:"VMSServer_Protocol"`
	VMSServerHost             string        `json:"VMSServer_Host" bson:"VMSServer_Host"`
	HRServerSQLServerHost     string        `json:"HRServer_SQLServerHost" bson:"HRServer_SQLServerHost"`
	HRServerDatabaseName      string        `json:"HRServer_DatabaseName" bson:"HRServer_DatabaseName"`
	HRServerViewTableName     string        `json:"HRServer_ViewTableName" bson:"HRServer_ViewTableName"`
	CreateUnixTimeStamp       int64         `json:"createUnixTimeStamp" bson:"createUnixTimeStamp"`
}

var hrSyncRecordsModel = new(HrSyncRecordsModel)

func (m *HrSyncRecordsModel) GenerateNewInstance() (objectIdRoot bson.ObjectId, err error) {
	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)

	protocol := globalConfig.Bundle["VMSServer_Protocol"].(string)
	host := globalConfig.Bundle["VMSServer_Host"].(string)
	HRServerSQLServerHost := globalConfig.Bundle["HRServer_SQLServerHost"].(string)
	HRServerDatabaseName := globalConfig.Bundle["HRServer_DatabaseName"].(string)
	HRServerViewTableName := globalConfig.Bundle["HRServer_ViewTableName"].(string)

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
	defer collection.Database.Session.Close()

	objectIdRoot = bson.NewObjectId()
	err = collection.Insert(bson.M{
		"_id":                       objectIdRoot,
		"syncVmsPersonDataCounts":   0,
		"updateVmsPersonDataCounts": 0,
		"createVmsPersonDataCounts": 0,
		"deleteVmsPersonDataCounts": 0,
		"syncHrServerDataCounts":    0,
		"status":                    "Sync-Processing",
		"failReason":                "",
		"VMSServer_Protocol":        protocol,
		"VMSServer_Host":            host,
		"HRServer_SQLServerHost":    HRServerSQLServerHost,
		"HRServer_DatabaseName":     HRServerDatabaseName,
		"HRServer_ViewTableName":    HRServerViewTableName,
		"createUnixTimeStamp":       time.Now().Unix(),
	})

	if err != nil {
		logv.Error(err.Error())
		return objectIdRoot, err
	}
	return objectIdRoot, err
}

func (m *HrSyncRecordsModel) UpdateStatus(recordsUUID string, status string, reason string) (err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
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

func (m *HrSyncRecordsModel) UpdateVMSSync(recordsUUID string, syncDataCounts int) (err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
	defer collection.Database.Session.Close()

	err = collection.UpdateId(bson.ObjectIdHex(recordsUUID), bson.M{"$set": bson.M{
		"syncVmsPersonDataCounts":     syncDataCounts,
	}})
	if err != nil {
		logv.Error(err.Error())
		return errors.New(err.Error())
	}
	return err
}

// HR 同步資訊
func (m *HrSyncRecordsModel) ListDataByP(data apiForms.ListByPHrSyncRecordsDataValidate) (resultsByPage []HrSyncRecords,
	resultsTotal []HrSyncRecords, err error, errcode int) {
	collectionHSR := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
	defer collectionHSR.Database.Session.Close()

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

	pipe := collectionHSR.Pipe(pipeline)
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

// HR 同步資訊，詳細資料
func (m *HrSyncRecordsModel) GetDetailDataByP(data apiForms.ListByPHrSyncRecordsDetailDataValidate) (resultsByPage []SyncVms2PersonResponse,
	resultsTotal []SyncVms2PersonResponse, hrSyncRecord HrSyncRecords, err error, errcode int) {

	collectionSR := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
	defer collectionSR.Database.Session.Close()

	err = collectionSR.FindId(bson.ObjectIdHex(*data.RecordUUID)).One(&hrSyncRecord)
	if err != nil {
		logv.Error(err.Error())
		return resultsTotal, resultsTotal, hrSyncRecord, err, 0
	}
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS_PERSON)
	defer collection.Database.Session.Close()

	isDESC := 1
	if *data.Desc {
		isDESC = 1
	} else {
		isDESC = -1
	}

	match_stage := bson.M{
		"hrSyncRecordsUUID": data.RecordUUID,
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
		return resultsTotal, resultsTotal, hrSyncRecord, err, 0
	}

	return resultsByPage, resultsTotal, hrSyncRecord, err, 0
}

func (m *HrSyncRecordsModel) CheckHrRecordsRetention() (info *mgo.ChangeInfo, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
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

	info, err = collection.RemoveAll(bson.M{"createUnixTimestamp": bson.M{"$lte": timestamp - 24*60*60* int64(log_retention_target)}})
	if err != nil {
		logv.Error("Update CheckHrRecordsRetention warn:> ", err)
	}
	return info, err
}
