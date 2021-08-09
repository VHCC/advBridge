package models

import (
	"advBridge/apiForms"
	"errors"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type KioskLocationModel struct {
}

type KioskLocation struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	DeviceUUID string        `json:"deviceUUID" bson:"deviceUUID"`
	Location   string        `json:"location" bson:"location"`
}

func (m *VmsServerModel) CreateKioskLocation(data apiForms.KioskLocationCreateDataValidate) (err error) {
	collectionKioskDevice := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collectionKioskDevice.Database.Session.Close()

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_KIOSK_LOCATION)
	defer collection.Database.Session.Close()

	kioskDevice := KioskDeviceInfo{}

	if _IsObjectIdHex := bson.IsObjectIdHex(*data.DeviceUUID); !_IsObjectIdHex {
		err = errors.New("invalid input to ObjectIdHex: " + *data.DeviceUUID)
		logv.Error(err.Error())
		return errors.New(err.Error())
	}

	err = collectionKioskDevice.FindId(bson.ObjectIdHex(*data.DeviceUUID)).One(&kioskDevice)
	if err != nil {
		logv.Error(err.Error())
		return errors.New("Kiosk Device UUID is not exist:> " + *data.DeviceUUID)
	}

	err = collection.Find(bson.M{"deviceUUID": data.DeviceUUID}).One(&kioskDevice)
	if err == nil {
		logv.Error("this location is exist:> ", data.DeviceUUID)
		return errors.New("this location is exist:> " + *data.DeviceUUID)
	}

	objectIdRoot := bson.NewObjectId()
	err = collection.Insert(bson.M{
		"_id":                 objectIdRoot,
		"deviceUUID":          data.DeviceUUID,
		"location":            data.Location,
		"createUnixTimeStamp": time.Now().Unix(),
	})

	if err != nil {
		logv.Error(err.Error())
		return err
	}
	return err
}

func (m *VmsServerModel) RemoveKioskLocation(data apiForms.KioskLocationDeleteDataValidate) (err error) {
	collectionKioskDevice := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collectionKioskDevice.Database.Session.Close()

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_KIOSK_LOCATION)
	defer collection.Database.Session.Close()

	kioskDevice := KioskDeviceInfo{}

	if _IsObjectIdHex := bson.IsObjectIdHex(*data.DeviceUUID); !_IsObjectIdHex {
		err = errors.New("invalid input to ObjectIdHex: " + *data.DeviceUUID)
		logv.Error(err.Error())
		return errors.New(err.Error())
	}

	err = collectionKioskDevice.FindId(bson.ObjectIdHex(*data.DeviceUUID)).One(&kioskDevice)
	if err != nil {
		logv.Error(err.Error())
		return errors.New("Kiosk Device UUID is not exist:> " + *data.DeviceUUID)
	}

	err = collection.Remove(bson.M{"deviceUUID": *data.DeviceUUID})

	if err != nil {
		logv.Error(err.Error())
		return err
	}
	return err
}

func (m *VmsServerModel) FetchAllKioskLocation() (results []KioskLocation, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_KIOSK_LOCATION)
	defer collection.Database.Session.Close()

	err = collection.Find(bson.M{}).All(&results)

	if err != nil {
		logv.Error(err.Error())
		return results, err
	}
	return results, err
}

func (m *VmsServerModel) UpdateKioskLocation(data apiForms.KioskLocationUpdateDataValidate) (err error) {
	collectionKioskDevice := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collectionKioskDevice.Database.Session.Close()

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_KIOSK_LOCATION)
	defer collection.Database.Session.Close()

	kioskDevice := KioskDeviceInfo{}

	if _IsObjectIdHex := bson.IsObjectIdHex(*data.DeviceUUID); !_IsObjectIdHex {
		err = errors.New("invalid input to ObjectIdHex: " + *data.DeviceUUID)
		logv.Error(err.Error())
		return errors.New(err.Error())
	}

	err = collectionKioskDevice.FindId(bson.ObjectIdHex(*data.DeviceUUID)).One(&kioskDevice)
	if err != nil {
		logv.Error(err.Error())
		return errors.New("Kiosk Device UUID is not exist:> " + *data.DeviceUUID)
	}

	err = collection.Update(bson.M{"deviceUUID": *data.DeviceUUID}, bson.M{"$set":
		bson.M{"location": data.Location}})

	if err != nil {
		logv.Error(err.Error())
		return err
	}
	return err
}

