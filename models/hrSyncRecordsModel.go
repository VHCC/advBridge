package models

import (
	"advBridge/apiForms"
	"errors"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
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

// HR 同步資訊
func (m *HrSyncRecordsModel) ListDataByP(data apiForms.ListByPVmsSyncRecordsDataValidate) (resultsByPage []VmsSyncRecords,
	resultsTotal []VmsSyncRecords, err error, errcode int) {
	collectionKD := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
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

// HR 同步資訊，詳細資料
func (m *HrSyncRecordsModel) GetDetailDataByP(data apiForms.ListByPVmsSyncRecordsDetailDataValidate) (resultsByPage []KioskReportResponse,
	resultsTotal []KioskReportResponse, vmsSyncRecord VmsSyncRecords, err error, errcode int) {

	collectionSR := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
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
