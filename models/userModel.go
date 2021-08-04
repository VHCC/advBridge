package models

import (
	"advBridge/utils"
	"errors"
	"fmt"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type User struct {
	UserUUID                  bson.ObjectId `json:"userUUID" bson:"_id,omitempty"`
	AccountID                 string        `json:"accountID" bson:"accountID"`
	Email                     string        `json:"email" bson:"email"`
	Password                  string        `json:"-" bson:"password"`
	Role                      int32         `json:"role" bson:"role"`
	ComUUID                   string        `json:"-" bson:"comUUID"`
	CurrentDepUUID            string        `json:"-" bson:"currentDepUUID"`
	DepUUID                   string        `json:"-" bson:"depUUID"`
	UserMemo                  string        `json:"userMemo" bson:"userMemo"`
	Permission                []int         `json:"-" bson:"permission"`
	CreateUnixTimeStamp       int64         `json:"-" bson:"createUnixTimeStamp"`
	LastModifiedUnixTimeStamp int64         `json:"-" bson:"lastModifiedUnixTimeStamp"`
	LastLoginUnixTimeStamp    int64         `json:"lastLoginUnixTimeStamp" bson:"lastLoginUnixTimeStamp"`
	AllowReviewNonVisitorData bool          `json:"-" bson:"allowReviewNonVisitorData"`
	UserToken                 string        `json:"userToken" bson:"userToken"`
}

type UserList struct {
	UserUUID                  bson.ObjectId `json:"userUUID" bson:"_id,omitempty"`
	AccountID                 string        `json:"accountID" bson:"accountID"`
	Email                     string        `json:"email" bson:"email"`
	Password                  string        `json:"-" bson:"password"`
	Role                      int32         `json:"role" bson:"role"`
	ComUUID                   string        `json:"comUUID" bson:"comUUID"`
	CurrentDepUUID            string        `json:"-" bson:"currentDepUUID"`
	DepUUID                   string        `json:"-" bson:"depUUID"`
	UserMemo                  string        `json:"userMemo" bson:"userMemo"`
	Permission                []int         `json:"permission" bson:"permission"`
	CreateUnixTimeStamp       int64         `json:"createUnixTimeStamp" bson:"createUnixTimeStamp"`
	LastModifiedUnixTimeStamp int64         `json:"lastModifiedUnixTimeStamp" bson:"lastModifiedUnixTimeStamp"`
	LastLoginUnixTimeStamp    int64         `json:"lastLoginUnixTimeStamp" bson:"lastLoginUnixTimeStamp"`
	AllowReviewNonVisitorData bool          `json:"allowReviewNonVisitorData" bson:"allowReviewNonVisitorData"`
	UserToken                 string        `json:"-" bson:"userToken"`
}

type UserModel struct{}

var cryptUtil = new(utils.CryptUtil)

func (m *UserModel) FetchUserInfoByUUID(userUUID string) (user UserList, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	if _IsObjectIdHex := bson.IsObjectIdHex(userUUID); !_IsObjectIdHex {
		err = errors.New("invalid input to ObjectIdHex: " + userUUID)
		return user, err
	}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(userUUID)}).One(&user)
	return user, err
}

func (m *UserModel) IsUserExist(
	email string,
	projectDID string,
	role int32) (user User, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	err = collection.Find(
		bson.M{"email": email,
			"projectDID": projectDID,
			"role":       role}).One(&user)
	return user, err
}

func (m *UserModel) Init() (err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	users := []UserList{}
	err = collection.Find(bson.M{}).All(&users)

	if len(users) == 0 {
		//objectIdAdmin := bson.NewObjectId()
		//err = collection.Insert(bson.M{
		//	"_id":                       objectIdAdmin,
		//	"accountID":                 "Admin",
		//	"email":                     "",
		//	"password":                  "Aa123456*",
		//	"role":                      5000,
		//	//"comUUID":                   defaultCom.ID.Hex(),
		//	//"currentDepUUID":            defaultDep.ID.Hex(),
		//	//"depUUID":                   defaultDep.ID.Hex(),
		//	"userMemo":                  "Default User Memo",
		//	"permission":                []int{},
		//	"allowReviewNonVisitorData": false,
		//	"createUnixTimeStamp":       time.Now().Unix(),
		//	"lastModifiedUnixTimeStamp": time.Now().Unix(),
		//})

		objectIdRoot := bson.NewObjectId()
		err = collection.Insert(bson.M{
			"_id":                       objectIdRoot,
			"accountID":                 "Admin",
			"email":                     "root",
			"password":                  "Az123567!",
			"role":                      9999,
			//"comUUID":                   "root",
			//"currentDepUUID":            "root",
			//"depUUID":                   "root",
			"userMemo":                  "Root Memo",
			"permission":                []int{},
			"allowReviewNonVisitorData": false,
			"createUnixTimeStamp":       time.Now().Unix(),
			"lastModifiedUnixTimeStamp": time.Now().Unix(),
		})
	}

	return err
}




func (m *UserModel) UpdateUserLogin(
	userUUID bson.ObjectId) (user User, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	userToken, err := cryptUtil.Generate()
	err = collection.UpdateId(userUUID, bson.M{"$set": bson.M{
		"userToken":              userToken,
		"lastLoginUnixTimeStamp": time.Now().Unix(),
	}})
	if err != nil {
		logv.Info("UpdateToDB Response UpdateId err:> ", err)
		return user, err
	}
	err = collection.FindId(userUUID).One(&user)
	return user, err
}

func (m *UserModel) UpdateUserLogout(
	userUUID bson.ObjectId) (user User, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	userToken, err := cryptUtil.Generate()
	err = collection.UpdateId(userUUID, bson.M{"$set": bson.M{
		"userToken":              userToken,
	}})
	if err != nil {
		logv.Info("UpdateToDB Response UpdateId err:> ", err)
		return user, err
	}
	err = collection.FindId(userUUID).One(&user)
	return user, err
}

func (m *UserModel) IsUserEmailExist(
	email string) (user User, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	err = collection.Find(
		bson.M{
			"email": email,
		}).One(&user)
	if err != nil {
		logv.Error("IsUserEmailExist Response FindId err:> ", err)
	}
	return user, err
}

func (m *UserModel) IsRoot(accountID string, password string) (isRoot bool, root User, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	err = collection.Find(
		bson.M{
			"accountID": accountID,
			"password":  password,
		}).One(&root)
	if err != nil {
		logv.Error("IsRoot Response FindId err:> ", err)
		return false, root, err
	}

	var mode = strings.ToLower(SERVER_MODE)
	logv.Info("mode:> ", mode)
	switch(mode) {
	case "edge":
		if root.Role == 5000 {
			return true, root, err
		} else {
			return false, root, err
		}
		break;
	case "cloud":
		if root.Role == 9999 {
			return true, root, err
		} else {
			return false, root, err
		}
		break;
	}
	if root.Role == 9999 {
		return true, root, err
	} else {
		return false, root, err
	}
}

func (m *UserModel) IsUserAccountIDExist(
	accountID string) (user User, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	err = collection.Find(
		bson.M{
			"accountID": accountID,
		}).One(&user)
	if err != nil {
		logv.Error("IsUserAccountIDExist Response FindId err:> ", err)
	}
	return user, err
}

func (m *UserModel) UserTokenCheck(
	userToken *string) (checkStatus int32, user UserList) {
	//logv.Info(*userToken)
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	err := collection.Find(
		bson.M{
			"userToken": *userToken,
		}).One(&user)

	checkStatus = user.Role
	if err != nil {
		fmt.Println("UserTokenCheck Response FindId err:> ", err)
		checkStatus = 1001 // userToken not Found
		//panic(err)
	}

	return checkStatus, user
}

func (m *UserModel) FindAllFromDB() (users []UserList, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_User)
	defer collection.Database.Session.Close()

	var query = bson.M{}
	err = collection.Find(query).All(&users)
	if err != nil {
		fmt.Println("FindByGroupFromDB Response Find err:> ", err)
	}
	return users, err
}


