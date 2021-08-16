package models

import (
	"fmt"
	logv "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"strconv"
)

var setting_db_json = db_Init()
var isPortInputString = true

func db_Init() (dbConfig map[string]interface{}) {
	logv.Info(" $$$ database init $$$ ")
	setting_db_json, err := ioutil.ReadFile("setting_db.json")
	if err != nil {
		logv.Error(err)
	}
	err = bson.UnmarshalJSON([]byte(string(setting_db_json)), &dbConfig)
	if err != nil {
		logv.Error(err)
	}

	logv.Info(" - db:> ", dbConfig["collection"])
	logv.Info(" - ip:> ", dbConfig["ip"])
	logv.Info(" - port:> ", dbConfig["port"])
	logv.Info(" - auth account:> ", dbConfig["account"])
	//logv.Info(dbConfig["password"])
	logv.Info(" - authSource:> ", dbConfig["authSource"])
	//logv.Info("mode:> ", dbConfig["mode"])
	var port = dbConfig["port"]
	switch v := port.(type) {
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		fmt.Printf("Integer: %v\n", v)
	case float64:
		isPortInputString = false
		// v is a float64 here, so e.g. v + 1.0 is possible.
		fmt.Printf("Float64: %v\n", v)
	case string:
		isPortInputString = true
		// v is a string here, so e.g. v + " Yeah!" is possible.
		fmt.Printf("String: %v\n", v)
	default:
		// And here I'm feeling dumb. ;)
		fmt.Printf("I don't know, ask stackoverflow.")
	}

	//_, err = os.Open("edge.mode.env")
	//if err == nil {
	//	//logv.Info("edge.mode.env Exists")
	//	SERVER_MODE = "edge"
	//}
	//_, err = os.Open("cloud.mode.env")
	//if err == nil {
	//	//logv.Error(err)
	//	//logv.Error("edge.mode.env Exists")
	//	SERVER_MODE = "cloud"
	//}
	//
	//logv.Info("MODE:> " + SERVER_MODE)
	return dbConfig
}

var db_host = setting_db_json["ip"].(string)

var db_port = FloatToString(setting_db_json["port"].(float64))

var server = db_host + ":" + db_port
var dbConnect = NewConnection(server, setting_db_json)
//var msSQLConnect = NewMSSQLConnection()
var err = initGlobalConfigModel()
var SERVER_MODE = "edge"

var DB_Name = setting_db_json["collection"].(string)

// ====== 中介 ======
var DB_Table_ADV_HR_User = "bridge_adv_users"
var DB_Table_ADV_SYNC_VMS_KIOSK_REPORTS = "bridge_adv_sync_vms_kiosk_reports"
var DB_Table_ADV_SYNC_VMS_KIOSK_DEVICES = "bridge_adv_sync_vms_kiosk_devices"
var DB_Table_ADV_SYNC_VMS_PERSON = "bridge_adv_sync_vms_person"

var DB_Table_ADV_KIOSK_LOCATION = "bridge_adv_kiosk_location"
var DB_Table_ADV_VMS_SYNC_RECORDS = "bridge_adv_vms_sync_records"
var DB_Table_ADV_HR_SYNC_RECORDS = "bridge_adv_hr_sync_records"
var DB_Table_ADV_HR_SYNC_RECORDS_PERSON = "bridge_adv_hr_sync_records_person"

var DB_Table_User = "users"
var DB_Table_Project = "projects"
var DB_Table_Group = "groups"
var DB_Table_Device = "devices"
var DB_Table_Global_Config = "globalConfigs"
var DB_Table_User_Preference = "userPreferences"
var DB_Table_File_Uploader_Version = "fileUploaderVersion"

var DB_VMS_Table_Form = "vmsForms"
var DB_VMS_Table_Form_Review_Item = "vmsReviewItems"
var DB_VMS_Table_Company = "vmsCompany"
var DB_VMS_Table_Department = "vmsDepartment"
var DB_VMS_Table_Client_Layout = "vmsClientLayout"
var DB_VMS_Table_Client_Device = "vmsClientDevice"
var DB_VMS_Table_Visitor = "vmsVisitor"

var DB_VMS_Table_Cost_Log = "vmsCostLog"
var DB_VMS_Table_Company_Point = "vmsCompanyPoint"
var DB_VMS_Table_Licence_Log = "vmsLicenceLog"

var DB_VMS_2_Table_Person = "vms2Person"
var DB_VMS_2_Table_Template = "vms2Template"

var DB_VMS_2_Table_Kiosk_Device = "vms2KioskDevice"
var DB_VMS_2_Table_Kiosk_Device_Connection_Number = "vms2KioskDeviceConnectionNumber"
var DB_VMS_2_Table_Kiosk_Device_Log_File = "vms2KioskDeviceLogFile"

var DB_VMS_2_Table_Kiosk_Reports = "vms2KioskReports"
var DB_VMS_2_Table_Check_In_Form = "vms2CheckInForm"

var DB_VMS_Log = "vmsLog"
var DB_ADV_BRIDGE_Log = "bridgeLog"

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}
