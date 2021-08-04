package db

import (
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type DBConnection struct {
	session *mgo.Session
}

func NewConnection(host string, dbConfig map[string]interface{}) (conn *DBConnection) {

	logv.Info("dbUtil, NewConnection :> ", host)
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

func (conn *DBConnection) UseTable(dbName string, tableName string) (collection *mgo.Collection) {
	//logv.Info("UseTable:> " , dbName)
	conn.session.Refresh()
	s := conn.session.Copy()
	return s.DB(dbName).C(tableName)
}
