package controllers

import (
	"advBridge/apiForms"
	"advBridge/models"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	logv "github.com/sirupsen/logrus"
)

type VmsSyncRecordsController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var vmsSyncRecordsModel = new(models.VmsSyncRecordsModel)

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
* @apiSuccess     {Integer}  dataCounts  counts of VMS 同步紀錄
* @apiSuccess     {JsonArray} vmsSyncRecordsData VMS 同步紀錄
* @apiSuccess     {String} vmsSyncRecordsData-status VMS 同步狀態。<br>
								1. Success <br>
								2. Fail
* @apiSuccess     {Integer} dataCounts VMS 同步紀錄總筆數
*
* @apiUse VmsSyncRecordsResponse_List_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *VmsSyncRecordsController) ListVmsSyncRecordsByParameter(c *gin.Context) {
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
* @apiSuccess     {Integer}  dataCounts  counts of 簽到紀錄
* @apiSuccess     {JsonObject} vmsSyncRecord VMS 同步紀錄
* @apiSuccess     {JsonArray} syncKioskReports 簽到紀錄同步清單
* @apiSuccess     {String} syncKioskReports-avalo_status 簽到紀錄狀況。<br>
							1. filling-out <br>
							2. check-in <br>
							3. authorized <br>
							4. exception
* @apiSuccess     {String} syncKioskReports-avalo_exception 簽到紀錄例外情形。<br>
							1. high-fever <br>
							2. no-mask <br>
							3. invalid-code <br>
							4 .authorization-fail <br>
							5. reject
* @apiSuccess     {Integer} dataCounts 人員同步清單總筆數
*
* @apiUse VmsSyncRecordsResponse_Detail_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *VmsSyncRecordsController) ListVmsSyncRecordsDetailByParameter(c *gin.Context) {
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

var vmsServerController = new(VmsController)


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
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse HRServerResponse_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse HRServerResponse_Connect_Err
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *VmsSyncRecordsController) RequestSyncWithVMS(c *gin.Context) {
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
		c.JSON(200, gin.H{"code": 2001, "message": "OPERATION_FAIL, " + err.Error()})
		logModel.WriteLog(models.EVENT_TYPE_VMS_SERVER_SYNC_FAIL, queryUser.AccountID, "OPERATION_FAIL, " + err.Error(), nil)
		c.Abort()
		return
	}
	logModel.WriteLog(models.EVENT_TYPE_VMS_SERVER_SYNC_SUCCESS, queryUser.AccountID, "SUCCESS", nil)
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}

func (cc *VmsSyncRecordsController) CheckVMSRecordsRetentions() {
	info, err := vmsSyncRecordsModel.CheckVMSRecordsRetention()
	if err != nil {
		logv.Error("CheckVMSRecordsRetentions warn:> ", err)
	}
	detailJson := map[string]interface{}{"removed": info.Removed}
	logModel.WriteLog(models.EVENT_TYPE_VMS_RECORDS_RETENTION_CHECK, "SYSTEM", "SUCCESS", detailJson)
}
