package controllers

import (
	"advBridge/apiForms"
	"advBridge/models"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	logv "github.com/sirupsen/logrus"
)

type BridgeLogController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

/**
@api {POST} /api/v2/bridgeLog/listByParameter List Bridge Logs
@apiDescription bridge server Log
@apiversion 0.2.0
@apiGroup 009 BRIDGE LOG
@apiName List Bridge Log By parameter

@apiUse ListByPBridgeLogDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonArray}  bridgeLogs  Logs of Bridge Server
* @apiSuccess     {JsonObject}  bridgeLogs-detail  detail of log <br>
								such as logType = BRIDGE_LOG-CHECK, detail will show how many records were removed <br>
								logType = VMS-SERVER-UPDATE, detail will print the modified info of VMS-SERVER.
*
* @apiUse Vms2LogResponse_List_Success
* @apiUse UserResponse_user_token_invalid
*/
func (vmsPersonC *BridgeLogController) ListVmsLogByPData(c *gin.Context) {
	var data apiForms.ListByPBridgeLogDataValidate

	// formData validation
	if c.ShouldBindJSON(&data) != nil {
		logv.Error("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	checkResult, queryUser := userModel.UserTokenCheck(data.UserToken)
	_ = queryUser
	switch checkResult {
	case 1001:
		c.JSON(200, gin.H{"code": 1001, "message": "USER_TOKEN_INVALID"})
		c.Abort()
		return
	}
	bridgeLogs, bridgeLogsTotal, err, errCode := logModel.ListLogByP(data)
	if err != nil {
		switch errCode {
		case 999:
		}
	}

	if len(bridgeLogs) == 0 {
		bridgeLogs = []models.BridgeLog{}
	}

	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS", "bridgeLogs": bridgeLogs, "dataCounts": len(bridgeLogsTotal)})
}

func (vmsPersonC *BridgeLogController) CheckBridgeLogRetentions() {
	info, err := logModel.CheckVmsLogRetention()
	if err != nil {
		logv.Error("CheckBridgeLogRetentions warn:> ", err)
	}
	detailJson := map[string]interface{}{"removed": info.Removed}
	logModel.WriteLog(models.EVENT_TYPE_BRIDGE_LOG_RETENTION_CHECK, "SYSTEM", "SUCCESS", detailJson)
}