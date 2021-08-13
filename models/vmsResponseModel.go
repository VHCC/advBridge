package models

import "gopkg.in/mgo.v2/bson"

type KioskReport struct {
	ID                        bson.ObjectId `json:"uuid" bson:"_id"`
	MappingPersonUUID         string        `json:"mappingPersonUUID" bson:"mappingPersonUUID"`
	AvaloDevice               string        `json:"avalo_device" bson:"avalo_device"`
	AvaloDeviceUuid           string        `json:"avalo_device_uuid" bson:"avalo_device_uuid"`
	AvaloDeviceGroup          string        `json:"avalo_device_group" bson:"avalo_device_group"`
	AvaloMode                 string        `json:"avalo_mode" bson:"avalo_mode"`
	AvaloInterface            string        `json:"avalo_interface" bson:"avalo_interface"`
	AvaloSnapshot             string        `json:"avalo_snapshot" bson:"avalo_snapshot"`
	AvaloStatus               string        `json:"avalo_status" bson:"avalo_status"`
	AvaloException            string        `json:"avalo_exception" bson:"avalo_exception"`
	AvaloSerial               string        `json:"avalo_serial" bson:"avalo_serial"`
	AvaloName                 string        `json:"avalo_name" bson:"avalo_name"`
	AvaloVisitor              bool          `json:"avalo_visitor" bson:"avalo_visitor"`
	AvaloEmail                string        `json:"avalo_email" bson:"avalo_email"`
	AvaloDepartment           string        `json:"avalo_department" bson:"avalo_department"`
	AvaloEnableTemperature    bool          `json:"avalo_enable_temperature" bson:"avalo_enable_temperature"`
	AvaloTemperature          float32       `json:"avalo_temperature" bson:"avalo_temperature"`
	AvaloTemperatureThreshold float32       `json:"avalo_temperature_threshold" bson:"avalo_temperature_threshold"`
	AvaloTemperatureAdjust    float32       `json:"avalo_temperature_adjust" bson:"avalo_temperature_adjust"`
	AvaloTemperatureUnit      string        `json:"avalo_temperature_unit" bson:"avalo_temperature_unit"`
	AvaloEnableMask           bool          `json:"avalo_enable_mask" bson:"avalo_enable_mask"`
	AvaloMask                 bool          `json:"avalo_mask" bson:"avalo_mask"`
	AvaloUtcTimestamp         int64         `json:"avalo_utc_timestamp" bson:"avalo_utc_timestamp"`
	AvaloPassports            string        `json:"avalo_passports" bson:"avalo_passports"`
	ReportTemplateUUID        string        `json:"report_templateUUID" bson:"report_templateUUID"`
	CheckInUuid               string        `json:"checkInUuid" bson:"checkInUuid"`
	VmsPerson                 []VmsPerson   `json:"vmsPerson" bson:"vmsPerson"`
}

type KioskReportResponse struct {
	ID                        bson.ObjectId `json:"uuid" bson:"_id"`
	MappingPersonUUID         string        `json:"-" bson:"mappingPersonUUID"`
	AvaloDevice               string        `json:"avalo_device" bson:"avalo_device"`
	AvaloDeviceUuid           string        `json:"avalo_device_uuid" bson:"avalo_device_uuid"`
	AvaloDeviceGroup          string        `json:"-" bson:"avalo_device_group"`
	AvaloMode                 string        `json:"avalo_mode" bson:"avalo_mode"`
	AvaloInterface            string        `json:"avalo_interface" bson:"avalo_interface"`
	AvaloSnapshot             string        `json:"avalo_snapshot" bson:"avalo_snapshot"`
	AvaloStatus               string        `json:"avalo_status" bson:"avalo_status"`
	AvaloException            string        `json:"avalo_exception" bson:"avalo_exception"`
	AvaloSerial               string        `json:"-" bson:"avalo_serial"`
	AvaloName                 string        `json:"-" bson:"avalo_name"`
	AvaloVisitor              bool          `json:"-" bson:"avalo_visitor"`
	AvaloEmail                string        `json:"-" bson:"avalo_email"`
	AvaloDepartment           string        `json:"-" bson:"avalo_department"`
	AvaloEnableTemperature    bool          `json:"-" bson:"avalo_enable_temperature"`
	AvaloTemperature          float32       `json:"avalo_temperature" bson:"avalo_temperature"`
	AvaloTemperatureThreshold float32       `json:"avalo_temperature_threshold" bson:"avalo_temperature_threshold"`
	AvaloTemperatureAdjust    float32       `json:"avalo_temperature_adjust" bson:"avalo_temperature_adjust"`
	AvaloTemperatureUnit      string        `json:"avalo_temperature_unit" bson:"avalo_temperature_unit"`
	AvaloEnableMask           bool          `json:"-" bson:"avalo_enable_mask"`
	AvaloMask                 bool          `json:"avalo_mask" bson:"avalo_mask"`
	AvaloUtcTimestamp         int64         `json:"avalo_utc_timestamp" bson:"avalo_utc_timestamp"`
	AvaloPassports            string        `json:"-" bson:"avalo_passports"`
	ReportTemplateUUID        string        `json:"-" bson:"report_templateUUID"`
	CheckInUuid               string        `json:"-" bson:"checkInUuid"`
	VmsPerson                 VmsPerson     `json:"vmsPerson" bson:"vmsPerson"`
}

type VmsPerson struct {
	ID              bson.ObjectId `json:"_id" bson:"_id"`
	VmsPersonEmail  string        `json:"vmsPersonEmail" bson:"vmsPersonEmail"`
	VmsPersonMemo   string        `json:"vmsPersonMemo" bson:"vmsPersonMemo"`
	VmsPersonName   string        `json:"vmsPersonName" bson:"vmsPersonName"`
	VmsPersonSerial string        `json:"vmsPersonSerial" bson:"vmsPersonSerial"`
	VmsPersonUnit   string        `json:"vmsPersonUnit" bson:"vmsPersonUnit"`
}

type KioskDeviceInfo struct {
	ID                      bson.ObjectId `json:"uuid" bson:"_id"`
	DeviceName              string        `json:"deviceName" bson:"deviceName"`
	VideoType               int           `json:"videoType" bson:"videoType"`
	Mode                    int           `json:"mode" bson:"mode"`
	Memo                    string        `json:"memo" bson:"memo"`
	ScreenTimeout           int           `json:"screenTimeout" bson:"screenTimeout"`
	AvaloDeviceHost         string        `json:"avaloDeviceHost" bson:"avaloDeviceHost"`
	AvaloAlertTemp          float32       `json:"avaloAlertTemp" bson:"avaloAlertTemp"`
	AvaloTempCompensation   float32       `json:"avaloTempCompensation" bson:"avaloTempCompensation"`
	AvaloTempUnit           string        `json:"avaloTempUnit" bson:"avaloTempUnit"`
	IsEnableTemp            bool          `json:"isEnableTemp" bson:"isEnableTemp"`
	IsEnableMask            bool          `json:"isEnableMask" bson:"isEnableMask"`
	VisitorTemplateUUID     string        `json:"visitorTemplateUUID" bson:"visitorTemplateUUID"`
	TemplateUUID            string        `json:"templateUUID" bson:"templateUUID"`
	TEPEnable               bool          `json:"tEPEnable" bson:"tEPEnable"`
	TEPHost                 string        `json:"tEPHost" bson:"tEPHost"`
	TEPPort                 string        `json:"tEPPort" bson:"tEPPort"`
	TEPEnableSSL            bool          `json:"tEPEnableSSL" bson:"tEPEnableSSL"`
	TEPAccount              string        `json:"tEPAccount" bson:"tEPAccount"`
	TEPPassword             string        `json:"tEPPassword" bson:"tEPPassword"`
	IsRFID                  bool          `json:"isRFID" bson:"isRFID"`
	IsBarCodeReader         bool          `json:"isBarCodeReader" bson:"isBarCodeReader"`
	IsCardReader            bool          `json:"isCardReader" bson:"isCardReader"`
	Status                  int           `json:"status" bson:"status"`
	ComUUID                 string        `json:"comUUID" bson:"comUUID"`
	DepUUID                 string        `json:"depUUID" bson:"depUUID"`
	AppUUID                 string        `json:"appUUID" bson:"appUUID"`
	AppVersion              string        `json:"appVersion" bson:"appVersion"`
	AndroidID               string        `json:"androidID" bson:"androidID"`
	SettingPassword         string        `json:"settingPassword" bson:"settingPassword"`
	ConnectTimeStamp        int64         `json:"connectTimeStamp" bson:"connectTimeStamp"`
	LastHeartBeatsTimeStamp int64         `json:"lastHeartBeatsTimeStamp" bson:"lastHeartBeatsTimeStamp"`
	LastSyncTimeStamp       int64         `json:"lastSyncTimeStamp" bson:"lastSyncTimeStamp"`
}

type KioskDeviceInfoResponse struct {
	ID                      bson.ObjectId `json:"uuid" bson:"_id"`
	DeviceName              string        `json:"deviceName" bson:"deviceName"`
	VideoType               int           `json:"-" bson:"videoType"`
	Mode                    int           `json:"-" bson:"mode"`
	Memo                    string        `json:"-" bson:"memo"`
	ScreenTimeout           int           `json:"-" bson:"screenTimeout"`
	AvaloDeviceHost         string        `json:"-" bson:"avaloDeviceHost"`
	AvaloAlertTemp          float32       `json:"-" bson:"avaloAlertTemp"`
	AvaloTempCompensation   float32       `json:"-" bson:"avaloTempCompensation"`
	AvaloTempUnit           string        `json:"-" bson:"avaloTempUnit"`
	IsEnableTemp            bool          `json:"-" bson:"isEnableTemp"`
	IsEnableMask            bool          `json:"-" bson:"isEnableMask"`
	VisitorTemplateUUID     string        `json:"-" bson:"visitorTemplateUUID"`
	TemplateUUID            string        `json:"-" bson:"templateUUID"`
	TEPEnable               bool          `json:"-" bson:"tEPEnable"`
	TEPHost                 string        `json:"-" bson:"tEPHost"`
	TEPPort                 string        `json:"-" bson:"tEPPort"`
	TEPEnableSSL            bool          `json:"-" bson:"tEPEnableSSL"`
	TEPAccount              string        `json:"-" bson:"tEPAccount"`
	TEPPassword             string        `json:"-" bson:"tEPPassword"`
	IsRFID                  bool          `json:"-" bson:"isRFID"`
	IsBarCodeReader         bool          `json:"-" bson:"isBarCodeReader"`
	IsCardReader            bool          `json:"-" bson:"isCardReader"`
	Status                  int           `json:"-" bson:"status"`
	ComUUID                 string        `json:"-" bson:"comUUID"`
	DepUUID                 string        `json:"-" bson:"depUUID"`
	AppUUID                 string        `json:"-" bson:"appUUID"`
	AppVersion              string        `json:"appVersion" bson:"appVersion"`
	AndroidID               string        `json:"androidID" bson:"androidID"`
	SettingPassword         string        `json:"-" bson:"settingPassword"`
	ConnectTimeStamp        int64         `json:"-" bson:"connectTimeStamp"`
	LastHeartBeatsTimeStamp int64         `json:"-" bson:"lastHeartBeatsTimeStamp"`
	LastSyncTimeStamp       int64         `json:"-" bson:"lastSyncTimeStamp"`
}

type Vms2Person struct {
	ID                  bson.ObjectId `json:"vmsPersonUUID" bson:"_id"`
	VMSPersonSerial     string        `json:"vmsPersonSerial" bson:"vmsPersonSerial"`
	VMSPersonName       string        `json:"vmsPersonName" bson:"vmsPersonName"`
	VMSPersonUnit       string        `json:"vmsPersonUnit" bson:"vmsPersonUnit"`
	VMSPersonEmail      string        `json:"vmsPersonEmail" bson:"vmsPersonEmail"`
	VMSPersonMemo       string        `json:"vmsPersonMemo" bson:"vmsPersonMemo"`
	IsRealName          bool          `json:"isRealName" bson:"isRealName"`
	CreateUnixTimestamp int64         `json:"createUnixTimestamp" bson:"createUnixTimestamp"`
}

type Vms2PersonResponse struct {
	ID                  bson.ObjectId `json:"vmsPersonUUID" bson:"_id,omitempty"`
	VMSPersonSerial     string        `json:"vmsPersonSerial" bson:"vmsPersonSerial"`
	VMSPersonName       string        `json:"vmsPersonName" bson:"vmsPersonName"`
	VMSPersonUnit       string        `json:"vmsPersonUnit" bson:"vmsPersonUnit"`
	VMSPersonEmail      string        `json:"vmsPersonEmail" bson:"vmsPersonEmail"`
	VMSPersonMemo       string        `json:"vmsPersonMemo" bson:"vmsPersonMemo"`
	IsRealName          bool          `json:"-" bson:"isRealName"`
	CreateUnixTimestamp int64         `json:"createUnixTimestamp" bson:"createUnixTimestamp"`
}

type SyncVms2PersonResponse struct {
	ID                  bson.ObjectId `json:"vmsPersonUUID" bson:"_id,omitempty"`
	VMSPersonSerial     string        `json:"vmsPersonSerial" bson:"vmsPersonSerial"`
	VMSPersonName       string        `json:"vmsPersonName" bson:"vmsPersonName"`
	VMSPersonUnit       string        `json:"vmsPersonUnit" bson:"vmsPersonUnit"`
	VMSPersonEmail      string        `json:"vmsPersonEmail" bson:"vmsPersonEmail"`
	VMSPersonMemo       string        `json:"vmsPersonMemo" bson:"vmsPersonMemo"`
	Action              string        `json:"action" bson:"action"`
	Status              string        `json:"status" bson:"status"`
	IsRealName          bool          `json:"-" bson:"isRealName"`
	CreateUnixTimestamp int64         `json:"createUnixTimestamp" bson:"createUnixTimestamp"`
}
