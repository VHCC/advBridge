package models

import (
	"database/sql"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/mgo.v2/bson"
)

type DBConnection struct {
	session *mgo.Session
}

func NewConnection(host string, dbConfig map[string]interface{}) (conn *DBConnection) {

	logv.Info("dbUtil, MongoDB :> ", host)
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	logv.Info(dbConfig["authSource"].(string))
	myDB := session.DB(dbConfig["authSource"].(string)) //这里的关键是连接mongodb后，选择admin数据库，然后登录，确保账号密码无误之后，该连接就一直能用了
	//出现server returned error on SASL authentication step: Authentication failed. 这个错也是因为没有在admin数据库下登录
	err = myDB.Login(dbConfig["account"].(string), dbConfig["password"].(string))
	if err != nil {
		logv.Error("DB Login-error:> ", err)
	}

	conn = &DBConnection{session}
	return conn
}

func NewMSSQLConnection() (conn *sql.DB) {
	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)
	SQLServerHost := globalConfig.Bundle["HRServer_SQLServerHost"].(string)
	SQLServerAccount := globalConfig.Bundle["HRServer_Account"].(string)
	SQLServerPassword := globalConfig.Bundle["HRServer_Password"].(string)
	SQLServerDBName := globalConfig.Bundle["HRServer_DatabaseName"].(string)

	conn, err := sql.Open("mssql",
		//"server=" + SQLServerHost +
		"server=" + SQLServerHost +
		";port=1433" +
		";user id=" + SQLServerAccount +
		";password=" + SQLServerPassword +
		";database=" + SQLServerDBName)
	if err != nil {
		logv.Error("Connecting Error:> ", err)
		return conn
	}
	logv.Info("dbUtil, MSSQL :> ", SQLServerHost + ":1433")
	return conn
}

func (conn *DBConnection) UseTable(dbName string, tableName string) (collection *mgo.Collection) {
	//logv.Info("UseTable:> " , dbName)
	conn.session.Refresh()
	s := conn.session.Copy()
	return s.DB(dbName).C(tableName)
}
