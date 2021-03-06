package models

import (
	"bytes"
	"encoding/json"
	"errors"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type VmsServerModel struct {
	protocol  string
	host      string
	userToken string
}

var MainProtocal string
var MainHost string
var MainUserToken string

// Constants
const API_login = "/api/v1/user/loginUser"
const API_listKRByPData = "/api/v2/vmsKioskReports/listKioskReportsByParameter"
const API_listKioskByPData = "/api/v2/vmsKioskDevice/listKioskDevicesByParameter"
const API_listPersonByPData = "/api/v2/vmsPerson/listVmsPersonByParameter"

const API_updatePersonByPData = "/api/v2/vmsPerson/updateVmsPerson"
const API_deletePersonByPData = "/api/v2/vmsPerson/deleteVmsPerson"
const API_createPersonByPData = "/api/v2/vmsPerson/createVmsPerson"

//  ======= login ======
type VmsLoginBody struct {
	AccountID string `json:"accountID"`
	Password  string `json:"password"`
}

type VmsLoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

// ========= listKRByPData ========
type VmsListKRByPBody struct {
	UserToken       string   `json:"userToken"`
	SortBy          string   `json:"sortBy"`
	Desc            bool     `json:"desc"`
	StartIndex      int      `json:"startIndex"`
	Count           int      `json:"count"`
	Avalo_interface []string `json:"avalo_interface"`
	StartTimestamp  int64    `json:"startTimestamp"`
	EndTimestamp    int64    `json:"endTimestamp"`
}

type VmsListKRByPResponse struct {
	Code         int           `json:"code"`
	Message      string        `json:"message"`
	KioskReports []KioskReport `json:"kioskReports"`
	DataCounts   int           `json:"dataCounts"`
}

// ============ VmsListKioskByPBody
type VmsListKioskByPBody struct {
	UserToken  string `json:"userToken"`
	SortBy     string `json:"sortBy"`
	Desc       bool   `json:"desc"`
	StartIndex int    `json:"startIndex"`
	Count      int    `json:"count"`
}

type VmsListKioskByPResponse struct {
	Code         int               `json:"code"`
	Message      string            `json:"message"`
	KioskDevices []KioskDeviceInfo `json:"kioskDevices"`
	DataCounts   int               `json:"dataCounts"`
}

// ======== VmsListPersonByPBody
type VmsListPersonByPBody struct {
	UserToken  string `json:"userToken"`
	SortBy     string `json:"sortBy"`
	Desc       bool   `json:"desc"`
	StartIndex int    `json:"startIndex"`
	Count      int    `json:"count"`
}

type VmsListPersonByPResponse struct {
	Code       int          `json:"code"`
	Message    string       `json:"message"`
	Vms2Person []Vms2Person `json:"vmsPersons"`
	DataCounts int          `json:"dataCounts"`
}

// ========= VmsUpdatePersonBody
type VmsUpdatePersonBody struct {
	UserToken     string `json:"userToken"`
	VmsPersonUUID string `json:"vmsPersonUUID"`
	VmsPersonName string `json:"vmsPersonName"`
	VmsPersonUnit string `json:"vmsPersonUnit"`
}

type VmsUpdatePersonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// =========== VmsCreatePersonBody
type VmsCreatePersonBody struct {
	UserToken       string `json:"userToken"`
	VmsPersonUUID   string `json:"vmsPersonUUID"`
	VmsPersonName   string `json:"vmsPersonName"`
	VmsPersonUnit   string `json:"vmsPersonUnit"`
	VmsPersonSerial string `json:"vmsPersonSerial"`
	VmsPersonMemo   string `json:"vmsPersonMemo"`
	VmsPersonEmail  string `json:"vmsPersonEmail"`
}

type VmsCreatePersonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// =========== VmsDeletePersonBody
type VmsDeletePersonBody struct {
	UserToken     string `json:"userToken"`
	VmsPersonUUID string `json:"vmsPersonUUID"`
}

type VmsDeletePersonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (m *VmsServerModel) LoginVMS() (err error, errCode int) {
	collectionConfig := dbConnect.UseTable(DB_Name, DB_Table_Global_Config)
	defer collectionConfig.Database.Session.Close()

	var globalConfig GlobalConfig

	err = collectionConfig.Find(bson.M{}).One(&globalConfig)

	protocol := globalConfig.Bundle["VMSServer_Protocol"].(string)
	host := globalConfig.Bundle["VMSServer_Host"].(string)
	VMSServerAccount := globalConfig.Bundle["VMSServer_Account"].(string)
	VMSServerPassword := globalConfig.Bundle["VMSServer_Password"].(string)

	loginData := VmsLoginBody{
		VMSServerAccount,
		VMSServerPassword,
	}

	loginData_json, _ := json.Marshal(loginData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", protocol+"://"+host+API_login, bytes.NewBuffer(loginData_json))
	if err != nil {
		logv.Error(err.Error())
		return errors.New(err.Error()), 101
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logv.Error(err.Error())
		return errors.New(err.Error()), 101
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return errors.New(err.Error()), 101
	}
	respBody := string(content)
	//fmt.Printf("Post request with json result: %s\n", respBody)
	vmsLoginResponse := &VmsLoginResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsLoginResponse)
	_ = errq

	defer res.Body.Close()
	if vmsLoginResponse.Code != 0 {
		logv.Error(errors.New(vmsLoginResponse.Message))
		return errors.New(vmsLoginResponse.Message), 104
	}

	logv.Info(" === Login Success, USER === ", vmsLoginResponse.User.AccountID)

	m.protocol = protocol
	MainProtocal = protocol
	m.host = host
	MainHost = host
	m.userToken = vmsLoginResponse.User.UserToken
	MainUserToken = vmsLoginResponse.User.UserToken
	return err, 0
}

func (m *VmsServerModel) ConnectionVMSTest(
	account string, pwd string, protocol string, host string) (err error) {
	resp, err := http.Get(protocol + "://" + host + "/ping")
	if err != nil {
		logv.Error(err)
		return err
		// handle error
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logv.Error(err)
		return err
		// handle error
	}

	//logv.Println(string(body))

	vmsLoginResponse := &VmsLoginResponse{}

	loginData := VmsLoginBody{
		account,
		pwd,
	}
	ba, _ := json.Marshal(loginData)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", protocol+"://"+host+"/api/v1/user/loginUser", bytes.NewBuffer(ba))
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)
	content, err := ioutil.ReadAll(res.Body)
	respBody := string(content)
	//fmt.Printf("Post request with json result: %s\n", respBody)
	err = json.Unmarshal([]byte(respBody), vmsLoginResponse)

	defer res.Body.Close()
	if vmsLoginResponse.Code != 0 {
		return errors.New(vmsLoginResponse.Message)
	}
	return err
}

func (m *VmsServerModel) SyncVMSReportData(objectID bson.ObjectId) {
	listKRByPData := VmsListKRByPBody{
		m.userToken,
		"avalo_utc_timestamp",
		false,
		0,
		500,
		[]string{"rfid"},
		(time.Now().Unix() - 60*60*24*7) * 1000,
		time.Now().Unix() * 1000,
	}
	listKRByPDataJson, _ := json.Marshal(listKRByPData)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", m.protocol+"://"+m.host+API_listKRByPData, bytes.NewBuffer(listKRByPDataJson))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		WriteLog(EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_FAIL, "SYSTEM", err.Error(), nil)

		return
	}
	respBody := string(content)

	vmsListKRByPResponse := &VmsListKRByPResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsListKRByPResponse)
	_ = errq
	defer res.Body.Close()
	if vmsListKRByPResponse.Code != 0 {
		logv.Error(errors.New(vmsListKRByPResponse.Message))
		WriteLog(EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_FAIL, "SYSTEM", vmsListKRByPResponse.Message, nil)
	}

	logv.Info(" === ListKRByP Success, Response === Counts:> ", vmsListKRByPResponse.DataCounts)

	for i := 0; i < vmsListKRByPResponse.DataCounts; i++ {
		saveReportsToBridgeDatabase(objectID.Hex(), vmsListKRByPResponse.KioskReports[i])
	}
	//vmsSyncRecordsModel.UpdateStatus(objectID.Hex(), "Success", "")
}

func (m *VmsServerModel) SyncVMSKioskDeviceData() (err error) {
	listKioskByPData := VmsListKioskByPBody{
		m.userToken,
		"deviceName",
		true,
		0,
		-1,
	}
	listKioskByPDataJson, _ := json.Marshal(listKioskByPData)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", m.protocol+"://"+m.host+API_listKioskByPData, bytes.NewBuffer(listKioskByPDataJson))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logv.Error(err.Error())
		return err
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return err
	}
	respBody := string(content)

	vmsListKioskByPResponse := &VmsListKioskByPResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsListKioskByPResponse)
	_ = errq
	defer res.Body.Close()
	if vmsListKioskByPResponse.Code != 0 {
		logv.Error(errors.New(vmsListKioskByPResponse.Message))
		WriteLog(EVENT_TYPE_VMS_KIOSK_DEVICE_SYNC_FAIL, "SYSTEM", vmsListKioskByPResponse.Message, nil)
		return errors.New(vmsListKioskByPResponse.Message)
	}
	logv.Info(" === ListKioskDeviceByP Success, Response === Counts:> ", vmsListKioskByPResponse.DataCounts)

	collectionKL := dbConnect.UseTable(DB_Name, DB_Table_ADV_KIOSK_LOCATION)
	defer collectionKL.Database.Session.Close()

	kioskLocations := []KioskLocation{}

	err = collectionKL.Find(bson.M{}).All(&kioskLocations)
	if err != nil {
		logv.Error(err.Error())
		return err
	}

	kioskDeviceUUIDArray := []string{}
	kioskLocationUUIDArray := []string{}

	for _, v := range kioskLocations {
		kioskDeviceUUIDArray = append(kioskDeviceUUIDArray, v.DeviceUUID)
		kioskLocationUUIDArray = append(kioskLocationUUIDArray, v.ID.Hex())
	}

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collection.Database.Session.Close()

	collection.DropCollection()

	for i := 0; i < vmsListKioskByPResponse.DataCounts; i++ {
		_, kioskDeviceUUIDArray, kioskLocationUUIDArray = isContainsKioskDeviceUUID(kioskDeviceUUIDArray, kioskLocationUUIDArray, vmsListKioskByPResponse.KioskDevices[i].ID.Hex())
		saveKioskDeviceToBridgeDatabase(vmsListKioskByPResponse.KioskDevices[i])
	}
	//logv.Info(kioskDeviceUUIDArray)
	for _, v := range kioskLocationUUIDArray {
		collectionKL.RemoveId(bson.ObjectIdHex(v))
	}
	return nil
}

func isContainsKioskDeviceUUID(kioskDeviceUUIDArray []string, kioskLocationUUIDArray []string, kioskDeviceUUID string, ) (
	isContained bool, newKioskDeviceArray []string, newKioskLocationArray []string) {
	for index, v := range kioskDeviceUUIDArray {
		if v == kioskDeviceUUID {
			newKioskDeviceArray = append(kioskDeviceUUIDArray[:index], kioskDeviceUUIDArray[index+1:]...)
			newKioskLocationArray = append(kioskLocationUUIDArray[:index], kioskLocationUUIDArray[index+1:]...)
			return true, newKioskDeviceArray, newKioskLocationArray
		}
	}
	return false, kioskDeviceUUIDArray, kioskLocationUUIDArray
}

func (m *VmsServerModel) SyncVMSPersonData() (syncDataCounts int) {
	listPersonByPData := VmsListPersonByPBody{
		m.userToken,
		"vmsPersonSerial",
		true,
		0,
		-1,
	}
	listPersonByPDataJson, _ := json.Marshal(listPersonByPData)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", m.protocol+"://"+m.host+API_listPersonByPData, bytes.NewBuffer(listPersonByPDataJson))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logv.Error(err.Error())
		return 0
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return 0
	}
	respBody := string(content)

	vmsListPersonByPResponse := &VmsListPersonByPResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsListPersonByPResponse)
	_ = errq
	defer res.Body.Close()
	if vmsListPersonByPResponse.Code != 0 {
		logv.Error(errors.New(vmsListPersonByPResponse.Message))
	}
	logv.Info(" === ListPersonByP Success, Response === Counts:> ", vmsListPersonByPResponse.DataCounts)
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_PERSON)
	defer collection.Database.Session.Close()
	collection.DropCollection()

	for i := 0; i < vmsListPersonByPResponse.DataCounts; i++ {
		savePersonToBridgeDatabase(vmsListPersonByPResponse.Vms2Person[i])
	}
	return vmsListPersonByPResponse.DataCounts
}

func saveReportsToBridgeDatabase(recordsUUID string, KRData KioskReport) () {
	collectionSyncRecords := dbConnect.UseTable(DB_Name, DB_Table_ADV_VMS_SYNC_RECORDS)
	defer collectionSyncRecords.Database.Session.Close()

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_REPORTS)
	defer collection.Database.Session.Close()

	kr := KioskReport{}

	err := collection.FindId(bson.ObjectIdHex(KRData.ID.Hex())).One(&kr)

	if err == nil && kr.SyncStatus {
		//logv.Error("KioskReports UUID:> ", KRData.ID.Hex(), " already exist !")
		return
	}

	sr := VmsSyncRecords{}

	if err = collectionSyncRecords.FindId(bson.ObjectIdHex(recordsUUID)).One(&sr); err != nil {
		logv.Error(err.Error())
		WriteLog(EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_FAIL, "SYSTEM", err.Error()+" :> "+recordsUUID, nil)
		return
	}

	if err = collectionSyncRecords.UpdateId(bson.ObjectIdHex(recordsUUID), bson.M{"$set": bson.M{"syncVmsDataCounts": sr.SyncVmsDataCounts + 1}}); err != nil {
		logv.Error(err.Error())
		WriteLog(EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_FAIL, "SYSTEM", err.Error()+" :> "+recordsUUID, nil)
		return
	}

	logv.Info(KRData.VmsPerson)

	if len(KRData.VmsPerson) > 0 {
		logv.Info("sync SUCCESS")
		//rfidMQTTModel.PublishToRFIDServerTest()
		if err = rfidMQTTModel.PublishToRFIDServer(KRData.VmsPerson[0].VmsPersonSerial,
			KRData.AvaloDeviceUuid,
			strconv.FormatFloat(float64(KRData.AvaloTemperature), 'f', 1, 64)); err != nil {
			logv.Error(err)
		}

		logv.Info("ADD KioskReports UUID:> ", KRData.ID.Hex())

		err = collectionSyncRecords.UpdateId(bson.ObjectIdHex(recordsUUID), bson.M{"$set": bson.M{"RFIDDataSendCounts": sr.RFIDDataSendCounts + 1}})
		if err != nil {
			logv.Error(err.Error())
			WriteLog(EVENT_TYPE_VMS_KIOSK_REPORTS_SYNC_FAIL, "SYSTEM", err.Error()+" :> "+recordsUUID, nil)
			return
		}

		err = collection.Insert(bson.M{
			"_id":                         KRData.ID,
			"recordsUUID":                 recordsUUID,
			"mappingPersonUUID":           KRData.MappingPersonUUID,
			"avalo_device":                KRData.AvaloDevice,
			"avalo_device_uuid":           KRData.AvaloDeviceUuid,
			"avalo_device_group":          KRData.AvaloDeviceGroup,
			"avalo_interface":             KRData.AvaloInterface,
			"avalo_snapshot":              KRData.AvaloSnapshot,
			"avalo_status":                KRData.AvaloStatus,
			"avalo_exception":             KRData.AvaloException,
			"avalo_serial":                KRData.AvaloSerial,
			"avalo_name":                  KRData.AvaloName,
			"avalo_visitor":               KRData.AvaloVisitor,
			"avalo_email":                 KRData.AvaloEmail,
			"avalo_mode":                  KRData.AvaloMode,
			"avalo_department":            KRData.AvaloDepartment,
			"avalo_enable_temperature":    KRData.AvaloEnableTemperature,
			"avalo_temperature":           KRData.AvaloTemperature,
			"avalo_temperature_threshold": KRData.AvaloTemperatureThreshold,
			"avalo_temperature_adjust":    KRData.AvaloTemperatureAdjust,
			"avalo_temperature_unit":      KRData.AvaloTemperatureUnit,
			"avalo_enable_mask":           KRData.AvaloEnableMask,
			"avalo_mask":                  KRData.AvaloMask,
			"avalo_utc_timestamp":         KRData.AvaloUtcTimestamp,
			"avalo_passports":             KRData.AvaloPassports,
			"report_templateUUID":         KRData.ReportTemplateUUID,
			"checkInUuid":                 KRData.CheckInUuid,
			"vmsPerson":                   KRData.VmsPerson[0],
			"syncStatus":                  true,
		})
	} else {
		logv.Info("sync FAIL")
		var vmsPerson = VmsPerson{}
		vmsPerson.ID = bson.NewObjectId()
		vmsPerson.VmsPersonSerial = "NOT EXIST"
		vmsPerson.VmsPersonName = "NOT EXIST"
		vmsPerson.VmsPersonEmail = "NOT EXIST"
		vmsPerson.VmsPersonMemo = "NOT EXIST"
		vmsPerson.VmsPersonUnit = "NOT EXIST"
		err = collection.Insert(bson.M{
			"_id":                         bson.NewObjectId(),
			"recordsUUID":                 recordsUUID,
			"mappingPersonUUID":           KRData.MappingPersonUUID,
			"avalo_device":                KRData.AvaloDevice,
			"avalo_device_uuid":           KRData.AvaloDeviceUuid,
			"avalo_device_group":          KRData.AvaloDeviceGroup,
			"avalo_interface":             KRData.AvaloInterface,
			"avalo_snapshot":              KRData.AvaloSnapshot,
			"avalo_status":                KRData.AvaloStatus,
			"avalo_exception":             KRData.AvaloException,
			"avalo_serial":                KRData.AvaloSerial,
			"avalo_name":                  KRData.AvaloName,
			"avalo_visitor":               KRData.AvaloVisitor,
			"avalo_email":                 KRData.AvaloEmail,
			"avalo_mode":                  KRData.AvaloMode,
			"avalo_department":            KRData.AvaloDepartment,
			"avalo_enable_temperature":    KRData.AvaloEnableTemperature,
			"avalo_temperature":           KRData.AvaloTemperature,
			"avalo_temperature_threshold": KRData.AvaloTemperatureThreshold,
			"avalo_temperature_adjust":    KRData.AvaloTemperatureAdjust,
			"avalo_temperature_unit":      KRData.AvaloTemperatureUnit,
			"avalo_enable_mask":           KRData.AvaloEnableMask,
			"avalo_mask":                  KRData.AvaloMask,
			"avalo_utc_timestamp":         KRData.AvaloUtcTimestamp,
			"avalo_passports":             KRData.AvaloPassports,
			"report_templateUUID":         KRData.ReportTemplateUUID,
			"checkInUuid":                 KRData.CheckInUuid,
			"vmsPerson":                   vmsPerson,
			"syncStatus":                  false,
		})
	}

}

func saveKioskDeviceToBridgeDatabase(KioskData KioskDeviceInfo) () {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collection.Database.Session.Close()

	//kiosk := KioskDeviceInfo{}

	//err := collection.FindId(bson.ObjectIdHex(KioskData.ID.Hex())).One(&kiosk)
	//if err == nil {
	//	logv.Info("Update Kiosk:> ", kiosk.ID.Hex())
	//	err = collection.RemoveId(kiosk.ID.Hex())
	//	if err == nil {
	//		logv.Error(err.Error())
	//		WriteLog(EVENT_TYPE_VMS_KIOSK_DEVICE_SYNC_FAIL, "SYSTEM", err.Error() + " :> " + kiosk.ID.Hex(), nil)
	//	}
	//}
	//logv.Info("ADD KioskDevice UUID:> ", KioskData.ID.Hex())

	err = collection.Insert(bson.M{
		"_id":                     KioskData.ID,
		"deviceName":              KioskData.DeviceName,
		"videoType":               KioskData.VideoType,
		"mode":                    KioskData.Mode,
		"memo":                    KioskData.Memo,
		"screenTimeout":           KioskData.ScreenTimeout,
		"avaloDeviceHost":         KioskData.AvaloDeviceHost,
		"avaloAlertTemp":          KioskData.AvaloAlertTemp,
		"avaloTempCompensation":   KioskData.AvaloTempCompensation,
		"avaloTempUnit":           KioskData.AvaloTempUnit,
		"isEnableTemp":            KioskData.IsEnableTemp,
		"isEnableMask":            KioskData.IsEnableMask,
		"visitorTemplateUUID":     KioskData.VisitorTemplateUUID,
		"templateUUID":            KioskData.TemplateUUID,
		"tEPEnable":               KioskData.TEPEnable,
		"tEPHost":                 KioskData.TEPHost,
		"tEPPort":                 KioskData.TEPPort,
		"tEPEnableSSL":            KioskData.TEPEnableSSL,
		"tEPAccount":              KioskData.TEPAccount,
		"tEPPassword":             KioskData.TEPPassword,
		"isRFID":                  KioskData.IsRFID,
		"isBarCodeReader":         KioskData.IsBarCodeReader,
		"isCardReader":            KioskData.IsCardReader,
		"status":                  KioskData.Status,
		"comUUID":                 KioskData.ComUUID,
		"depUUID":                 KioskData.DepUUID,
		"appUUID":                 KioskData.AppUUID,
		"appVersion":              KioskData.AppVersion,
		"androidID":               KioskData.AndroidID,
		"settingPassword":         KioskData.SettingPassword,
		"connectTimeStamp":        KioskData.ConnectTimeStamp,
		"lastHeartBeatsTimeStamp": KioskData.LastHeartBeatsTimeStamp,
		"lastSyncTimeStamp":       KioskData.LastSyncTimeStamp,
	})
}

func savePersonToBridgeDatabase(personData Vms2Person) () {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_PERSON)
	defer collection.Database.Session.Close()

	person := KioskDeviceInfo{}

	err := collection.FindId(bson.ObjectIdHex(personData.ID.Hex())).One(&person)
	if err == nil {
		logv.Info("Update person:> ", person.ID.Hex())
		err = collection.RemoveId(bson.ObjectIdHex(person.ID.Hex()))
		if err != nil {
			logv.Error(err.Error())
		}
	}
	//logv.Info(personData)
	//logv.Info(personData.VMSPersonMemo)
	//logv.Info(personData.VMSPersonSerial)
	err = collection.Insert(bson.M{
		"_id":                 personData.ID,
		"vmsPersonSerial":     personData.VMSPersonSerial,
		"vmsPersonName":       personData.VMSPersonName,
		"vmsPersonUnit":       personData.VMSPersonUnit,
		"vmsPersonEmail":      personData.VMSPersonEmail,
		"vmsPersonMemo":       personData.VMSPersonMemo,
		"isRealName":          true,
		"createUnixTimestamp": personData.CreateUnixTimestamp,
	})
}

func (m *VmsServerModel) GetAllKioskReports() (results []KioskReportResponse, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_REPORTS)
	defer collection.Database.Session.Close()

	err = collection.Find(bson.M{}).All(&results)
	if err != nil {
		logv.Error(err.Error())
		return results, err
	}
	return results, err
}

func (m *VmsServerModel) GetAllKioskDevices() (results []KioskDeviceInfoResponse, err error) {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collection.Database.Session.Close()

	err = collection.Find(bson.M{}).All(&results)
	if err != nil {
		logv.Error(err.Error())
		return results, err
	}
	return results, err
}

// ====== manipulate to VMS Person
func UpdateVMSPersonData(
	personUUID string,
	vmsPersonName string,
	vmsPersonUnit string,
) (errCode int) {
	updatePersonData := VmsUpdatePersonBody{
		MainUserToken,
		personUUID,
		vmsPersonName,
		vmsPersonUnit,
	}
	updatePersonDataJson, _ := json.Marshal(updatePersonData)

	//logv.Info(updatePersonData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", MainProtocal+"://"+MainHost+API_updatePersonByPData, bytes.NewBuffer(updatePersonDataJson))
	if err != nil {
		logv.Error(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	respBody := string(content)

	vmsUpdatePersonResponse := &VmsUpdatePersonResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsUpdatePersonResponse)
	_ = errq
	defer res.Body.Close()
	if vmsUpdatePersonResponse.Code != 0 {
		code := strconv.Itoa(vmsUpdatePersonResponse.Code)
		logv.Error(errors.New(vmsUpdatePersonResponse.Message + ", Code:> " + code))
		if vmsUpdatePersonResponse.Code == 22001 {
			return 22001
		}
		return 0
	}
	logv.Info(" === Update Person Success ===, UUID:> ", personUUID+", :> "+vmsPersonName)
	return 0
}

func CreateVMSPersonData(
	personUUID string,
	vmsPersonName string,
	vmsPersonUnit string,
	vmsPersonSerial string,
	vmsPersonMemo string,
) {
	createPersonData := VmsCreatePersonBody{
		MainUserToken,
		personUUID,
		vmsPersonName,
		vmsPersonUnit,
		vmsPersonSerial,
		vmsPersonMemo,
		"bridgeTest",
	}
	createPersonDataJson, _ := json.Marshal(createPersonData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", MainProtocal+"://"+MainHost+API_createPersonByPData, bytes.NewBuffer(createPersonDataJson))
	if err != nil {
		logv.Error(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	respBody := string(content)

	vmsCreatePersonResponse := &VmsCreatePersonResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsCreatePersonResponse)
	_ = errq
	defer res.Body.Close()
	if vmsCreatePersonResponse.Code != 0 {
		logv.Error(errors.New(vmsCreatePersonResponse.Message))
		return
	}
	logv.Info(" === Create Person Success === :> " + vmsPersonName + ", :> " + vmsPersonUnit)
}

func DeleteVMSPersonData(
	personUUID string,
) {
	deletePersonData := VmsDeletePersonBody{
		MainUserToken,
		personUUID,
	}
	deletePersonDataJson, _ := json.Marshal(deletePersonData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", MainProtocal+"://"+MainHost+API_deletePersonByPData, bytes.NewBuffer(deletePersonDataJson))
	if err != nil {
		logv.Error(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	respBody := string(content)

	vmsDeletePersonResponse := &VmsDeletePersonResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsDeletePersonResponse)
	_ = errq
	defer res.Body.Close()
	if vmsDeletePersonResponse.Code != 0 {
		logv.Error(errors.New(vmsDeletePersonResponse.Message))
		return
	}
	logv.Info(" === Delete Person Success ===")
}
