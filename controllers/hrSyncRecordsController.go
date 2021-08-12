package controllers

import (
	"advBridge/apiForms"
	"advBridge/models"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	logv "github.com/sirupsen/logrus"
)

type HrSyncRecordsController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var hrSyncRecordsModel = new(models.HrSyncRecordsModel)

/**
@api {POST} /api/v1/vmsSyncRecords/listVmsSyncRecords List Vms Sync Records
@apiDescription List Vms Sync Records
@apiversion 0.0.1
@apiGroup 006 VMS Sync Records
@apiName List Vms Sync Records

@apiUse ListByPVmsSyncRecordsDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonArray} vmsSyncRecordsData VMS 同步紀錄
* @apiSuccess     {Integer} dataCounts VMS 同步紀錄總筆數
*
* @apiUse VmsSyncRecordsResponse_List_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *HrSyncRecordsController) ListHRSyncRecordsByParameter(c *gin.Context) {
	var data apiForms.ListByPVmsSyncRecordsDataValidate

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

	vmsSyncRecordsData, vmsSyncRecordsDataTotal, err, errCode := vmsSyncRecordsModel.ListDataByP(data)

	if err != nil {
		switch errCode {
		case 999:
		}
	}

	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}

	if len(vmsSyncRecordsData) == 0 {
		vmsSyncRecordsData = []models.VmsSyncRecords{}
	}

	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS",
		"vmsSyncRecordsData": vmsSyncRecordsData, "dataCounts": len(vmsSyncRecordsDataTotal)})
}




/**
@api {POST} /api/v1/vmsSyncRecords/getVmsSyncRecordsDetail Detail of Vms Sync Record
@apiDescription Detail of Vms Sync Record
@apiversion 0.0.1
@apiGroup 006 VMS Sync Records
@apiName Detail of Vms Sync Record

@apiUse ListByPVmsSyncRecordsDetailDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonObject} vmsSyncRecord VMS 同步紀錄
* @apiSuccess     {JsonArray} syncKioskReports 人員同步清單
* @apiSuccess     {Integer} dataCounts 人員同步清單總筆數
*
* @apiUse VmsSyncRecordsResponse_Detail_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *HrSyncRecordsController) ListHRSyncRecordsDetailByParameter(c *gin.Context) {
	var data apiForms.ListByPVmsSyncRecordsDetailDataValidate

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

	syncKioskReports, syncKioskReportsTotal, vmsSyncRecord, err, errCode := vmsSyncRecordsModel.GetDetailDataByP(data)

	if err != nil {
		switch errCode {
		case 999:
		}
	}

	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}

	if len(syncKioskReports) == 0 {
		syncKioskReports = []models.KioskReportResponse{}
	}

	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS",
		"vmsSyncRecord": vmsSyncRecord,
		"syncKioskReports": syncKioskReports,
		"dataCounts": len(syncKioskReportsTotal)})
}



/**
@api {POST} /api/v1/vmsSyncRecords/requestSyncWithVMS Request Sync With Vms
@apiDescription Request Sync With Vms
@apiversion 0.0.1
@apiGroup 006 VMS Sync Records
@apiName Request Sync With Vms

@apiUse RequestSyncWithVMSDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 2001:CONNECT_ERROR </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse HRServerResponse_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse HRServerResponse_Connect_Err
* @apiUse UserResponse_user_token_invalid
*/
func (cc *HrSyncRecordsController) RequestSyncWithHR(c *gin.Context) {
	var data apiForms.RequestSyncWithVMSDataValidate

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


	err := vmsServerController.SyncVMSKioskReportsData()
	if err != nil {
		if err != nil {
			c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
			c.Abort()
			return
		}
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}
