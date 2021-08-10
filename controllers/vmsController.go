package controllers

import (
	"advBridge/apiForms"
	"advBridge/models"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	logv "github.com/sirupsen/logrus"
)

type VmsController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var vmsServerModel = new(models.VmsServerModel)

func (cc *VmsController) SyncVMSKioskReportsData() {
	err := vmsServerModel.LoginVMS()
	if err != nil {
		logv.Error(err.Error())
		return
	}
	vmsServerModel.SyncVMSReportData()
	vmsServerModel.SyncVMSKioskDeviceData()
}

/**
@api {POST} /api/v1/vmsServer/connectTest VMS Server Connection Test
@apiDescription VMS Server Connection Test
@apiversion 0.0.1
@apiGroup 003 VMS Server
@apiName VMS Server Connection Test

@apiUse VMSServerTestDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*                 2001:CONNECT_ERROR </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse HRServerResponse_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse HRServerResponse_Connect_Err
*/
func (cc *VmsController) VmsServerConnectionTest(c *gin.Context) {
	var data apiForms.VMSServerTestDataValidate

	// formData validation
	if c.ShouldBind(&data) != nil {
		logv.Error("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	err := vmsServerModel.ConnectionVMSTest(data.AccountID, data.Password, data.Protocol, data.Host)
	if err != nil {
		c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}

/**
@api {POST} /api/v1/vmsServer/fetchVMSKioskReports List VMS Kiosk Reports
@apiDescription List VMS Kiosk Kiosk
@apiversion 0.0.1
@apiGroup 003 VMS Server
@apiName List VMS Kiosk Kiosk

@apiUse VMSServerKioskReportsFetchDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonArray}  kioskReports kioskReports
*
* @apiUse VMSServerResponse_Success_Kiosk_Reports
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *VmsController) FetchVmsKioskReports(c *gin.Context) {
	var data apiForms.VMSServerKioskReportsFetchDataValidate

	// formData validation
	if c.ShouldBind(&data) != nil {
		logv.Error("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	checkResult, queryUser := userModel.UserTokenCheck(data.UserToken)
	_ = queryUser
	switch checkResult {
	case 1:
	case 2:
	case 1001:
		c.JSON(200, gin.H{"code": 1001, "message": "USER_TOKEN_INVALID"})
		c.Abort()
		return
	}

	kioskReports, err := vmsServerModel.GetAllKioskReports()
	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS", "kioskReports": kioskReports})
}

/**
@api {POST} /api/v1/vmsServer/fetchVMSKioskDevices List VMS Kiosk Devices
@apiDescription List VMS Kiosk Devices
@apiversion 0.0.1
@apiGroup 003 VMS Server
@apiName List VMS Kiosk Devices

@apiUse VMSServerKioskReportsFetchDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonArray}  kioskDevices kioskDevices
*
* @apiUse VMSServerResponse_Success_Kiosk_Devices
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *VmsController) FetchVmsKioskDevices(c *gin.Context) {
	var data apiForms.VMSServerKioskDevicesFetchDataValidate

	// formData validation
	if c.ShouldBind(&data) != nil {
		logv.Error("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	checkResult, queryUser := userModel.UserTokenCheck(data.UserToken)
	_ = queryUser
	switch checkResult {
	case 1:
	case 2:
	case 1001:
		c.JSON(200, gin.H{"code": 1001, "message": "USER_TOKEN_INVALID"})
		c.Abort()
		return
	}

	kioskDevices, err := vmsServerModel.GetAllKioskDevices()
	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS", "kioskDevices": kioskDevices})
}