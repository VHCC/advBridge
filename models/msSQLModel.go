package models

import (
	"database/sql"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
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

type HrSyncRecords struct {
}

type MsSQLModel struct{}

func (m *MsSQLModel) ConnectionTest(
	host string, account string, pwd string, DBName string) (conn *sql.DB, err error) {
	//logv.Info(host)
	//logv.Info(account)
	//logv.Info(pwd)
	//logv.Info(DBName)
	//defer conn.Close()
	conn, err = sql.Open("mssql",
		"server=" + host +
			";port=1433" +
			";user id=" + account +
			";password=" + pwd +
			";database=" + DBName)
	if err != nil {
		logv.Error("Connecting Error:> ", err)
		return conn, err
	}
	err = conn.Ping()
	if err != nil {
		logv.Error("Connecting Error:> ", err)
		return conn, err
	}
	logv.Info("ConnectionVMSTest, MSSQL :> ", host + ":1433")
	return conn, err
}


func (m *MsSQLModel) SyncHRDB() (err error) {
	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)

	SQLServerTableName := globalConfig.Bundle["HRServer_ViewTableName"].(string)


	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_User)
	collection.DropCollection()
	defer collection.Database.Session.Close()

	stmt, err := msSQLConnect.Prepare("select * from " + SQLServerTableName)
	if err != nil {
		logv.Println("Query Error", err)
		return
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		logv.Println("Query Error", err)
		return
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
			err = collection.Insert(bson.M{
				"_id":                 objectIdRoot,
				"Memo":                EMNO,
				"Name":                NAME,
				"Serial":                 EPC,
				"MEB_CardNo":          MEB_CardNo,
				"createUnixTimeStamp": time.Now().Unix(),
			})
			if err != nil {
				logv.Println("Mongodb Insert Error:> ", err)
			}
		}
	}
	logv.Info(" === SyncHRDB Done !!! === ")
	return err
}
