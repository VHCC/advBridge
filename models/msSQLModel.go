package models

import (
	"context"
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
	//logv.Info(host)
	//logv.Info(account)
	//logv.Info(pwd)
	//logv.Info(DBName)
	//defer conn.Close()

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

	//var bgCtx = context.TODO()
	//var ctx2SecondTimeout, cancelFunc2SecondTimeout = context.WithTimeout(bgCtx, time.Second*3)
	//defer cancelFunc2SecondTimeout()

	err = conn.Ping()
	if err != nil {
		logv.Error("ConnectionHRTest, Connecting Error:> ", err)
		conn.Close()
		return conn, err
	}

	//if host != "172.20.2.85" {
	//	logv.Error("ConnectionHRTest, Connecting Error:> ", errors.New("connect Timeout:> " + host))
	//	return conn, errors.New("connect Timeout:> " + host)
	//}

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

	defer conn.Close()
	if err != nil {
		logv.Error("ConnectionHRTest, Connecting Error:> ", err)
		return conn, err
	}

	var bgCtx = context.Background()
	var ctx2SecondTimeout, cancelFunc2SecondTimeout = context.WithTimeout(bgCtx, time.Second*2)
	defer cancelFunc2SecondTimeout()

	err = conn.PingContext(ctx2SecondTimeout)
	if err != nil {
		logv.Error("ConnectBySystem, Connecting Error:> ", err)
		return conn, err
	}
	logv.Info("ConnectBySystem, MSSQL :> ", SQLServerHost+":1433")
	return conn, err
}

func isContainsRFID(personSerialArray []string, s []string, rfid string) (isContained bool, personUUID string, newPersonUUIDArray []string, newPersonSerialArray []string) {
	//logv.Info("array:> ", len(s))
	for index, v := range personSerialArray {
		if v == rfid {
			//logv.Info(v, ":> ", s[index])
			newPersonUUIDArray = append(s[:index], s[index+1:]...)
			newPersonSerialArray = append(personSerialArray[:index], personSerialArray[index+1:]...)
			return true, s[index], newPersonUUIDArray, newPersonSerialArray
		}
	}
	return false, "", s, personSerialArray
}

func (m *MsSQLModel) SyncHRDB(conn *sql.DB) (err error) {
	collectionVP := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_PERSON)
	defer collectionVP.Database.Session.Close()

	vmsPersons := []VmsPerson{}
	err = collectionVP.Find(bson.M{}).All(&vmsPersons)
	if err != nil {
		logv.Error(err.Error())
	}

	personUUIDArray := []string{}
	personSerialArray := []string{}

	for _, v := range vmsPersons {
		personUUIDArray = append(personUUIDArray, v.ID.Hex())
		personSerialArray = append(personSerialArray, v.VmsPersonSerial)
	}

	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)

	SQLServerTableName := globalConfig.Bundle["HRServer_ViewTableName"].(string)

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_User)
	collection.DropCollection()
	defer collection.Database.Session.Close()

	stmt, err := conn.Prepare("select * from " + SQLServerTableName)
	if err != nil {
		logv.Println("Query Error", err)
		return err
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		logv.Println("Query Error", err)
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

				isContainsRFIDString, personUUID, personUUIDArray, personSerialArray = isContainsRFID(personSerialArray, personUUIDArray, reverseHexString)
				if  isContainsRFIDString {
					errCode := UpdateVMSPersonData(personUUID, NAME, EPC)
					if errCode == 22001 {
						CreateVMSPersonData(personUUID, NAME, EMNO, reverseHexString, EPC)
					}
				} else {
					CreateVMSPersonData(personUUID, NAME, EMNO, reverseHexString, EPC)
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

				isContainsRFIDString, personUUID, personUUIDArray, personSerialArray = isContainsRFID(personSerialArray, personUUIDArray, MEB_CardNo)
				if isContainsRFIDString {
					errCode := UpdateVMSPersonData(personUUID, NAME, EPC)
					if errCode == 22001 {
						CreateVMSPersonData(personUUID, NAME, EMNO, MEB_CardNo, EPC)
					}
				} else {
					CreateVMSPersonData(personUUID, NAME, EMNO, MEB_CardNo, EPC)
				}
			}

			if err != nil {
				logv.Println("Mongodb Insert Error:> ", err)
				return err
			}
		}
	}
	for _, uuid := range personUUIDArray {
		DeleteVMSPersonData(uuid)
	}
	logv.Info(" === SyncHRDB Done !!! === ")
	return err
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

