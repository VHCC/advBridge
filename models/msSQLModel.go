package models

import (
	"database/sql"
	"fmt"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ADVUser struct {
	UUID                bson.ObjectId `json:"UUID" bson:"_id"`
	EMNO                string        `json:"Memo" bson:"Memo"`
	NAME                string        `json:"Name" bson:"Name"`
	EPC                 string        `json:"Serial" bson:"Serial"`
	MEBCARDNO           string        `json:"-" bson:"MEB_CardNo"`
	CreateUnixTimeStamp int64         `json:"createUnixTimeStamp" bson:"createUnixTimeStamp"`
}

type MsSQLModel struct{}

var wg sync.WaitGroup

func showData(txt string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("RealOutput==>", txt)
	}
	wg.Done()
}

func (m *MsSQLModel) ConnectionTest(
	host string, account string, pwd string, DBName string) (conn *sql.DB, err error) {

	connectString := "sqlserver://"+account+":"+pwd +"@"+host+":1433??database=" +DBName+"&dial+timeout=3"
	//conn, err = sql.Open("mssql",
	//	"server="+host+
	//		//";port=1433"+
	//		";user id="+account+
	//		";password="+pwd+
	//		";database="+DBName)

	logv.Info(connectString)
	//conn, err = sql.Open("sqlserver", u.String())
	conn, err = sql.Open("sqlserver", connectString)
	defer conn.Close()
	if err != nil {
		logv.Error("ConnectionHRTest, Connecting Error:> ", err)
		return conn, err
	}

	err = conn.Ping()
	if err != nil {
		logv.Error("ConnectionHRTest, Connecting Error:> ", err)
		conn.Close()
		return conn, err
	}

	logv.Info("ConnectionHRTest, MSSQL :> ", host+":1433")
	return conn, err
}

func (m *MsSQLModel) ConnectBySystem() (conn *sql.DB, err error) {
	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)
	SQLServerHost := globalConfig.Bundle["HRServer_SQLServerHost"].(string)
	SQLServerAccount := globalConfig.Bundle["HRServer_Account"].(string)
	SQLServerPassword := globalConfig.Bundle["HRServer_Password"].(string)
	SQLServerDBName := globalConfig.Bundle["HRServer_DatabaseName"].(string)

	conn, err = sql.Open("mssql",
		//"server=" + SQLServerHost +
		"server=" + SQLServerHost +
			";port=1433" +
			";user id=" + SQLServerAccount +
			";password=" + SQLServerPassword +
			";database=" + SQLServerDBName)

	//defer conn.Close()
	if err != nil {
		logv.Error("ConnectionHRTest, Connecting Error:> ", err)
		return conn, err
	}

	//var bgCtx = context.Background()
	//var ctx2SecondTimeout, cancelFunc2SecondTimeout = context.WithTimeout(bgCtx, time.Second*2)
	//defer cancelFunc2SecondTimeout()

	err = conn.Ping()
	if err != nil {
		logv.Error("ConnectBySystem, Connecting Error:> ", err)
		return conn, err
	}
	logv.Info("ConnectBySystem, MSSQL :> ", SQLServerHost+":1433")
	return conn, err
}

func isContainsRFID(personSerialArray []string, personUUIDArray []string, personUnitArray []string, rfid string) (
	isContained bool, personUUID string,
	newPersonUUIDArray []string, newPersonSerialArray []string,
	newPersonUnitArray []string, vmsPersonUnit string) {
	for index, v := range personSerialArray {
		if v == rfid {
			vmsPersonUnit = personUnitArray[index]
			personUUID = personUUIDArray[index]
			newPersonUUIDArray = append(personUUIDArray[:index], personUUIDArray[index+1:]...)
			newPersonSerialArray = append(personSerialArray[:index], personSerialArray[index+1:]...)
			newPersonUnitArray = append(personUnitArray[:index], personUnitArray[index+1:]...)
			return true, personUUID, newPersonUUIDArray, newPersonSerialArray, newPersonUnitArray, vmsPersonUnit
		}
	}
	return false, "", personUUIDArray, personSerialArray, personUnitArray, ""
}

func (m *MsSQLModel) SyncHRDB(conn *sql.DB, hrSyncObjectID bson.ObjectId) (err error) {
	defer conn.Close()

	collectionHRSyncRecordsPerson := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS_PERSON)
	defer collectionHRSyncRecordsPerson.Database.Session.Close()

	collectionHRSyncRecords := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_SYNC_RECORDS)
	defer collectionHRSyncRecords.Database.Session.Close()

	hr := HrSyncRecords{}

	err = collectionHRSyncRecords.FindId(bson.ObjectIdHex(hrSyncObjectID.Hex())).One(&hr)
	if err != nil {
		logv.Error(err.Error())
		return err
	}

	collectionVP := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_PERSON)
	defer collectionVP.Database.Session.Close()

	vmsPersons := []VmsPerson{}
	err = collectionVP.Find(bson.M{}).All(&vmsPersons)
	if err != nil {
		logv.Error(err.Error())
		return err
	}

	personUUIDArray := []string{}
	personSerialArray := []string{}
	personUnitArray := []string{}

	for _, v := range vmsPersons {
		personUUIDArray = append(personUUIDArray, v.ID.Hex())
		personSerialArray = append(personSerialArray, v.VmsPersonSerial)
		personUnitArray = append(personUnitArray, v.VmsPersonUnit)
	}

	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)

	SQLServerTableName := globalConfig.Bundle["HRServer_ViewTableName"].(string)

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_HR_User)
	collection.DropCollection()
	defer collection.Database.Session.Close()

	stmt, err := conn.Prepare("select * from " + SQLServerTableName)
	if err != nil {
		logv.Println("SyncHRDB, Query Error:> ", err)
		return err
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		logv.Println("SyncHRDB, Query Error:> ", err)
		return err
	}
	defer row.Close()
	for row.Next() {
		//logv.Info(row.Columns())
		var EMNO string
		var NAME string
		var EPC string
		var MEB_CardNo string
		if err := row.Scan(&EMNO, &NAME, &EPC, &MEB_CardNo); err == nil {
			err = collectionHRSyncRecords.FindId(bson.ObjectIdHex(hrSyncObjectID.Hex())).One(&hr)
			if err != nil {
				logv.Error(err.Error())
				return err
			}

			//logv.Info(EMNO, NAME, EPC, MEB_CardNo)
			objectIdRoot := bson.NewObjectId()

			//logv.Info(MEB_CardNo, ", isNum:> ",  IsNum(MEB_CardNo))
			if IsNum(MEB_CardNo) {
				rfidNum, _ := strconv.Atoi(MEB_CardNo)

				//logv.Info(strings.ToUpper(strconv.FormatInt(int64(rfidNum), 16)))
				var hexString = strings.ToUpper(strconv.FormatInt(int64(rfidNum), 16))
				var reverseHexString = ""
				if len(hexString) == 7 {
					//if EMNO == "I-0111" {
					//	bbb := hexString[5:7] + hexString[3:5] + hexString[1:3] + "0" + hexString[0:1]
					//	logv.Info(bbb + ", " + EMNO)
					//}
					reverseHexString = hexString[5:7] + hexString[3:5] + hexString[1:3] + "0" + hexString[0:1]
					//logv.Info(reverseHexString + ", " + EMNO)
				} else if len(hexString) == 8{
					//if EMNO == "I-0107" {
					//	bbb := hexString[6:8] + hexString[4:6] + hexString[2:4] + hexString[0:2]
					//	logv.Info(bbb + ", " + EMNO)
					//}
					reverseHexString = hexString[6:8] + hexString[4:6] + hexString[2:4] + hexString[0:2]
					//logv.Info(reverseHexString + ", " + EMNO)
				} else {
					reverseHexString = hexString[4:6] + hexString[2:4] + hexString[0:2] + "00"
					//logv.Info(reverseHexString + ", " + EMNO)
				}
				err = collection.Insert(bson.M{
					"_id":                 objectIdRoot,
					"Memo":                EMNO,
					"Name":                NAME,
					"Serial":              EPC,
					"MEB_CardNo":          MEB_CardNo,
					"RFIDCard":            reverseHexString,
					"createUnixTimeStamp": time.Now().Unix(),
				})

				var personUUID string
				var isContainsRFIDString bool
				var matchPersonUnit string

				isContainsRFIDString, personUUID, personUUIDArray, personSerialArray, personUnitArray, matchPersonUnit = isContainsRFID(personSerialArray, personUUIDArray, personUnitArray, reverseHexString)
				if  isContainsRFIDString {
					if EPC != matchPersonUnit {
						//logv.Info(EPC + " : " + matchPersonUnit + ", :> " + NAME)
						errCode := UpdateVMSPersonData(personUUID, NAME, EPC)
						if errCode == 22001 {
							CreateVMSPersonData(personUUID, NAME, EPC, reverseHexString, EMNO)
							err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"createVmsPersonDataCounts": hr.CreateVmsPersonDataCounts + 1}})
							if err != nil {
								logv.Error(err.Error())
								return err
							}
							collectionHRSyncRecordsPerson.Insert(bson.M{
								"_id": bson.ObjectIdHex(personUUID),
								"vmsPersonName": NAME,
								"vmsPersonUnit": EPC,
								"vmsPersonSerial": reverseHexString,
								"vmsPersonMemo": EMNO,
								"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
								"action": "create",
								"status": "SUCCESS",
							})
						} else {
							err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"updateVmsPersonDataCounts": hr.UpdateVmsPersonDataCounts + 1}})
							if err != nil {
								logv.Error(err.Error())
								return err
							}
							collectionHRSyncRecordsPerson.Insert(bson.M{
								"_id": bson.ObjectIdHex(personUUID),
								"vmsPersonName": NAME,
								"vmsPersonUnit": EPC,
								"vmsPersonSerial": reverseHexString,
								"vmsPersonMemo": EMNO,
								"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
								"action": "update",
								"status": "SUCCESS",
							})
						}
					} else {
						// KEEP
						collectionHRSyncRecordsPerson.Insert(bson.M{
							"_id": bson.ObjectIdHex(personUUID),
							"vmsPersonName": NAME,
							"vmsPersonUnit": EPC,
							"vmsPersonSerial": reverseHexString,
							"vmsPersonMemo": EMNO,
							"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
							"action": "keep",
							"status": "SUCCESS",
						})
					}

				} else {
					CreateVMSPersonData(personUUID, NAME, EPC, reverseHexString, EMNO)
					err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"createVmsPersonDataCounts": hr.CreateVmsPersonDataCounts + 1}})
					if err != nil {
						logv.Error(err.Error())
						return err
					}
					collectionHRSyncRecordsPerson.Insert(bson.M{
						"_id": bson.NewObjectId(),
						"vmsPersonName": NAME,
						"vmsPersonUnit": EPC,
						"vmsPersonSerial": reverseHexString,
						"vmsPersonMemo": EMNO,
						"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
						"action": "create",
						"status": "SUCCESS",
					})
				}
			} else {
				err = collection.Insert(bson.M{
					"_id":                 objectIdRoot,
					"Memo":                EMNO,
					"Name":                NAME,
					"Serial":              EPC,
					"MEB_CardNo":          MEB_CardNo,
					"RFIDCard":            MEB_CardNo,
					"createUnixTimeStamp": time.Now().Unix(),
				})

				var personUUID string
				var isContainsRFIDString bool
				var matchPersonUnit string

				isContainsRFIDString, personUUID, personUUIDArray, personSerialArray, personUnitArray, matchPersonUnit = isContainsRFID(personSerialArray, personUUIDArray, personUnitArray, MEB_CardNo)
				if isContainsRFIDString {
					if EPC != matchPersonUnit {
						//logv.Info(EPC + " : " + matchPersonUnit)
						errCode := UpdateVMSPersonData(personUUID, NAME, EPC)
						if errCode == 22001 {
							CreateVMSPersonData(personUUID, NAME, EPC, MEB_CardNo, EMNO)
							err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"createVmsPersonDataCounts": hr.CreateVmsPersonDataCounts + 1}})
							if err != nil {
								logv.Error(err.Error())
								return err
							}
							collectionHRSyncRecordsPerson.Insert(bson.M{
								"_id": bson.NewObjectId(),
								"vmsPersonName": NAME,
								"vmsPersonUnit": EPC,
								"vmsPersonSerial": MEB_CardNo,
								"vmsPersonMemo": EMNO,
								"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
								"action": "create",
								"status": "SUCCESS",
							})
						} else {
							err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"updateVmsPersonDataCounts": hr.UpdateVmsPersonDataCounts + 1}})
							if err != nil {
								logv.Error(err.Error())
								return err
							}
							collectionHRSyncRecordsPerson.Insert(bson.M{
								"_id": bson.NewObjectId(),
								"vmsPersonName": NAME,
								"vmsPersonUnit": EPC,
								"vmsPersonSerial": MEB_CardNo,
								"vmsPersonMemo": EMNO,
								"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
								"action": "update",
								"status": "SUCCESS",
							})
						}
					} else {
						// KEEP
						collectionHRSyncRecordsPerson.Insert(bson.M{
							"_id": bson.ObjectIdHex(personUUID),
							"vmsPersonName": NAME,
							"vmsPersonUnit": EPC,
							"vmsPersonSerial": MEB_CardNo,
							"vmsPersonMemo": EMNO,
							"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
							"action": "keep",
							"status": "SUCCESS",
						})
					}
				} else {
					CreateVMSPersonData(personUUID, NAME, EPC, MEB_CardNo, EMNO)
					err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"createVmsPersonDataCounts": hr.CreateVmsPersonDataCounts + 1}})
					if err != nil {
						logv.Error(err.Error())
						return err
					}
					collectionHRSyncRecordsPerson.Insert(bson.M{
						"_id": bson.NewObjectId(),
						"vmsPersonName": NAME,
						"vmsPersonUnit": EPC,
						"vmsPersonSerial": MEB_CardNo,
						"vmsPersonMemo": EMNO,
						"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
						"action": "create",
						"status": "SUCCESS",
					})
				}
			}

			if err != nil {
				logv.Println("Mongodb Insert Error:> ", err)
				return err
			}
			err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"syncHrServerDataCounts": hr.SyncHrServerDataCounts + 1}})
			if err != nil {
				logv.Error(err.Error())
				return err
			}
		}
	}
	for _, uuid := range personUUIDArray {
		DeleteVMSPersonData(uuid)
		err = collectionHRSyncRecords.UpdateId(bson.ObjectIdHex(hrSyncObjectID.Hex()), bson.M{"$set": bson.M{"deleteVmsPersonDataCounts": hr.DeleteVmsPersonDataCounts + 1}})
		if err != nil {
			logv.Error(err.Error())
			return err
		}

		vmsPerson := VmsPerson{}
		err = collectionVP.FindId(bson.ObjectIdHex(uuid)).One(&vmsPerson)
		if err != nil {
			logv.Error(err.Error())
		}

		collectionHRSyncRecordsPerson.Insert(bson.M{
			"_id": bson.NewObjectId(),
			"vmsPersonName": vmsPerson.VmsPersonName,
			"vmsPersonUnit": vmsPerson.VmsPersonUnit,
			"vmsPersonSerial": vmsPerson.VmsPersonSerial,
			"vmsPersonMemo": vmsPerson.VmsPersonMemo,
			"hrSyncRecordsUUID": hrSyncObjectID.Hex(),
			"status": "delete",
		})
	}

	logv.Info(" === SyncHRDB Done !!! === ")
	return err
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

