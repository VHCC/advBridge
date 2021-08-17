package controllers

import (
	"advBridge/apiForms"
	"advBridge/models"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	logv "github.com/sirupsen/logrus"
)

type MsSQLController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var msSQLModel = new(models.MsSQLModel)


func (cc *MsSQLController) SyncHRDatabase() (){
	objectID, err := hrSyncRecordsModel.GenerateNewInstance()
	err, errCode := vmsServerModel.LoginVMS()
	if err != nil {
		logv.Error(err.Error() + ", code:> ", errCode)
		switch errCode {
		case 101:
			hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "Vms Server 連線失敗")
			logModel.WriteLog(models.EVENT_TYPE_VMS_SERVER_CONNECT_FAIL, "SYSTEM", "CONNECT_ERROR, " + err.Error(), nil)
			return
		case 104:
			hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "Vms Server 登入失敗")
			logModel.WriteLog(models.EVENT_TYPE_VMS_SERVER_CONNECT_FAIL, "SYSTEM", "CONNECT_ERROR, " + err.Error(), nil)
			return
		}
		return
	}

	conn, err := msSQLModel.ConnectBySystem()
	if err != nil {
		hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "HR Server 連線失敗, " + err.Error())
		logModel.WriteLog(models.EVENT_TYPE_HR_SERVER_CONNECT_FAIL, "SYSTEM", "CONNECT_ERROR, " + err.Error(), nil)
		return
	}
	err = msSQLModel.SyncHRDB(conn, objectID)
	if err != nil {
		hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "HR Server 撈取資料失敗, " + err.Error())
		logModel.WriteLog(models.EVENT_TYPE_HR_SERVER_SYNC_FAIL, "SYSTEM", "OPERATION_FAIL, " + err.Error(), nil)
		return
	}
	hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Success", "")
}




/**
@api {POST} /api/v1/hrSyncRecords/requestSyncWithHR Request Sync With HR Server
@apiDescription Request Sync With HR
@apiversion 0.0.1
@apiGroup 007 HR Server Sync Records
@apiName Request Sync With HR Server

@apiUse RequestSyncWithHRDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*                 2001:CONNECT_ERROR (參數缺少或錯誤) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse HRServerResponse_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse HRServerResponse_Connect_Err
*/
func (cc *MsSQLController) RequestSyncHRDatabase(c *gin.Context){
	var data apiForms.RequestSyncWithHRDataValidate

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

	objectID, err := hrSyncRecordsModel.GenerateNewInstance()
	err, errCode := vmsServerModel.LoginVMS()
	if err != nil {
		logv.Error(err.Error() + ", code:> ", errCode)
		switch errCode {
		case 101:
			hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "Vms Server 連線失敗")
		case 104:
			hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "Vms Server 登入失敗")
		}
		logModel.WriteLog(models.EVENT_TYPE_VMS_SERVER_CONNECT_FAIL, queryUser.AccountID, "CONNECT_ERROR, " + err.Error(), nil)
		c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
		c.Abort()
		return
	}

	conn, err := msSQLModel.ConnectBySystem()
	if err != nil {
		hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "HR Server 連線失敗, " + err.Error())
		c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
		logModel.WriteLog(models.EVENT_TYPE_HR_SERVER_CONNECT_FAIL, queryUser.AccountID, "CONNECT_ERROR, " + err.Error(), nil)
		c.Abort()
		return
	}

	syncVmsDataCounts := vmsServerModel.SyncVMSPersonData()
	hrSyncRecordsModel.UpdateVMSSync(objectID.Hex(), syncVmsDataCounts)

	err = msSQLModel.SyncHRDB(conn, objectID)
	if err != nil {
		hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Fail", "HR Server 撈取資料失敗, " + err.Error())
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		logModel.WriteLog(models.EVENT_TYPE_HR_SERVER_SYNC_FAIL, queryUser.AccountID, "OPERATION_FAIL, " + err.Error(), nil)
		c.Abort()
		return
	}
	hrSyncRecordsModel.UpdateStatus(objectID.Hex(), "Success", "")
	logModel.WriteLog(models.EVENT_TYPE_HR_SERVER_SYNC_SUCCESS, queryUser.AccountID, "SUCCESS", nil)
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}

/**
@api {POST} /api/v1/hrServer/connectTest HR Server Connection Test
@apiDescription HR Server Connection Test
@apiversion 0.0.1
@apiGroup 002 HR Server
@apiName HR Server Connection Test

@apiUse MSSQLTestStructure

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 2001:CONNECT_ERROR (參數缺少或錯誤) </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse HRServerResponse_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse HRServerResponse_Connect_Err
* @apiUse UserResponse_user_token_invalid
*/
func (cc *MsSQLController) MSSQLConnectionTest(c *gin.Context) {
	var data apiForms.MSSQLTestDataValidate

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

	_, err := msSQLModel.ConnectionTest(data.Host, data.AccountID, data.Password, data.DBName)
	if err != nil {
		c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
		logModel.WriteLog(models.EVENT_TYPE_HR_SERVER_CONNECT_FAIL, queryUser.AccountID, "CONNECT_ERROR, " + err.Error(), nil)
		c.Abort()
		return
	}
	logModel.WriteLog(models.EVENT_TYPE_HR_SERVER_CONNECT_SUCCESS, queryUser.AccountID, "SUCCESS", nil)
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}