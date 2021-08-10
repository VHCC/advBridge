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
)

type VmsServerModel struct {
	protocol  string
	host      string
	userToken string
}

// Constants
const API_login = "/api/v1/user/loginUser"
const API_listKRByPData = "/api/v2/vmsKioskReports/listKioskReportsByParameter"
const API_listKioskByPData = "/api/v2/vmsKioskDevice/listKioskDevicesByParameter"
const API_listPersonByPData = "/api/v2/vmsPerson/listVmsPersonByParameter"

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
	m.host = host
	m.userToken = vmsLoginResponse.User.UserToken
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
	errq := json.Unmarshal([]byte(respBody), vmsLoginResponse)
	_ = errq

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
		true,
		0,
		-1,
		[]string{"rfid"},
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
		return
	}
	respBody := string(content)

	vmsListKRByPResponse := &VmsListKRByPResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsListKRByPResponse)
	_ = errq
	defer res.Body.Close()
	if vmsListKRByPResponse.Code != 0 {
		logv.Error(errors.New(vmsListKRByPResponse.Message))
	}

	logv.Info(" === ListKRByP Success, Response === Counts:> ", vmsListKRByPResponse.DataCounts)

	for i := 0; i < vmsListKRByPResponse.DataCounts; i++ {
		saveReportsToBridgeDatabase(objectID.Hex(), vmsListKRByPResponse.KioskReports[i])
	}
	vmsSyncRecordsModel.UpdateStatus(objectID.Hex(), "Success", "")
}

func (m *VmsServerModel) SyncVMSKioskDeviceData() {
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
		return
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return
	}
	respBody := string(content)

	vmsListKioskByPResponse := &VmsListKioskByPResponse{}
	errq := json.Unmarshal([]byte(respBody), vmsListKioskByPResponse)
	_ = errq
	defer res.Body.Close()
	if vmsListKioskByPResponse.Code != 0 {
		logv.Error(errors.New(vmsListKioskByPResponse.Message))
	}
	logv.Info(" === ListKioskDeviceByP Success, Response === Counts:> ", vmsListKioskByPResponse.DataCounts)
	for i := 0; i < vmsListKioskByPResponse.DataCounts; i++ {
		saveKioskDeviceToBridgeDatabase(vmsListKioskByPResponse.KioskDevices[i])
	}
}

func (m *VmsServerModel) SyncVMSPersonData() {
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
		return
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logv.Error(err.Error())
		return
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
	for i := 0; i < vmsListPersonByPResponse.DataCounts; i++ {
		savePersonToBridgeDatabase(vmsListPersonByPResponse.Vms2Person[i])
	}
}

func saveReportsToBridgeDatabase(recordsUUID string, KRData KioskReport) () {
	collectionSyncRecords := dbConnect.UseTable(DB_Name, DB_Table_ADV_VMS_SYNC_RECORDS)
	defer collectionSyncRecords.Database.Session.Close()

	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_REPORTS)
	defer collection.Database.Session.Close()

	kr := KioskReport{}

	err := collection.FindId(bson.ObjectIdHex(KRData.ID.Hex())).One(&kr)
	if err == nil {
		//logv.Error("KioskReports UUID:> ", KRData.ID.Hex(), " already exist !")
		return
	}

	sr := VmsSyncRecords{}

	err = collectionSyncRecords.FindId(bson.ObjectIdHex(recordsUUID)).One(&sr)
	if err != nil {
		logv.Error(err.Error())
		return
	}

	err = collectionSyncRecords.UpdateId(bson.ObjectIdHex(recordsUUID), bson.M{"$set": bson.M{"syncVmsDataCounts": sr.SyncVmsDataCounts + 1}})
	if err != nil {
		logv.Error(err.Error())
		return
	}

	rfidMQTTModel.PublishToRFIDServer(KRData.VmsPerson[0].VmsPersonSerial,
		KRData.AvaloDeviceUuid,
		strconv.FormatFloat(float64(KRData.AvaloTemperature), 'f', 1, 64))

	logv.Info("ADD KioskReports UUID:> ", KRData.ID.Hex())

	err = collectionSyncRecords.UpdateId(bson.ObjectIdHex(recordsUUID), bson.M{"$set": bson.M{"RFIDDataSendCounts": sr.RFIDDataSendCounts + 1}})
	if err != nil {
		logv.Error(err.Error())
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
	})
}

func saveKioskDeviceToBridgeDatabase(KioskData KioskDeviceInfo) () {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES)
	defer collection.Database.Session.Close()

	kiosk := KioskDeviceInfo{}

	err := collection.FindId(bson.ObjectIdHex(KioskData.ID.Hex())).One(&kiosk)
	if err == nil {
		logv.Info("Update Kiosk:> ", kiosk.ID.Hex())
		err = collection.RemoveId(kiosk.ID.Hex())
		if err == nil {
			logv.Error(err.Error())
		}
	}
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

func savePersonToBridgeDatabase(PersonData Vms2Person) () {
	collection := dbConnect.UseTable(DB_Name, DB_Table_ADV_SYNC_VMS_PERSON)
	defer collection.Database.Session.Close()

	person := KioskDeviceInfo{}

	err := collection.FindId(bson.ObjectIdHex(PersonData.ID.Hex())).One(&person)
	if err == nil {
		logv.Info("Update person:> ", person.ID.Hex())
		err = collection.RemoveId(person.ID.Hex())
		if err == nil {
			logv.Error(err.Error())
		}
	}

	err = collection.Insert(bson.M{
		"_id":                 PersonData.ID,
		"vmsPersonSerial":     PersonData.VMSPersonSerial,
		"vmsPersonName":       PersonData.VMSPersonName,
		"vmsPersonUnit":       PersonData.VMSPersonUnit,
		"vmsPersonEmail":      PersonData.VMSPersonEmail,
		"vmsPersonMemo":       PersonData.VMSPersonMemo,
		"isRealName":          PersonData.IsRealName,
		"createUnixTimestamp": PersonData.CreateUnixTimestamp,
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
