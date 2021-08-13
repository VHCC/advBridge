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
@api {POST} /api/v1/hrSyncRecords/listHrSyncRecords List HRServer Sync Records
@apiDescription List Vms Sync Records
@apiversion 0.0.1
@apiGroup 007 HR Server Sync Records
@apiName List Vms Sync Records

@apiUse ListByPHrSyncRecordsDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonArray} hrSyncRecordsData HR Server 同步紀錄
* @apiSuccess     {Integer} dataCounts HR Server 同步紀錄總筆數
*
* @apiUse HRSyncRecordsResponse_List_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *HrSyncRecordsController) ListHRSyncRecordsByParameter(c *gin.Context) {
	var data apiForms.ListByPHrSyncRecordsDataValidate

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

	hrSyncRecordsData, hrSyncRecordsDataTotal, err, errCode := hrSyncRecordsModel.ListDataByP(data)

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

	if len(hrSyncRecordsData) == 0 {
		hrSyncRecordsData = []models.HrSyncRecords{}
	}

	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS",
		"hrSyncRecordsData": hrSyncRecordsData, "dataCounts": len(hrSyncRecordsDataTotal)})
}




/**
@api {POST} /api/v1/vmsSyncRecords/getVmsSyncRecordsDetail Detail of HR Server Sync Record
@apiDescription Detail of HR Server Sync Record
@apiversion 0.0.1
@apiGroup 007 HR Server Sync Records
@apiName Detail of HR Server Sync Record

@apiUse ListByPHrSyncRecordsDetailDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonObject} hrSyncRecord HR SERVER 同步紀錄
* @apiSuccess     {JsonArray} syncVmsPersons 人員同步清單
* @apiSuccess     {Integer} dataCounts 人員同步清單總筆數
*
* @apiUse HRSyncRecordsResponse_Detail_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *HrSyncRecordsController) ListHRSyncRecordsDetailByParameter(c *gin.Context) {
	var data apiForms.ListByPHrSyncRecordsDetailDataValidate

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

	syncVmsPersons, syncVmsPersonsTotal, hrSyncRecord, err, errCode := hrSyncRecordsModel.GetDetailDataByP(data)

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

	if len(syncVmsPersons) == 0 {
		syncVmsPersons = []models.SyncVms2PersonResponse{}
	}

	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS",
		"hrSyncRecord": hrSyncRecord,
		"syncVmsPersons": syncVmsPersons,
		"dataCounts": len(syncVmsPersonsTotal)})
}



