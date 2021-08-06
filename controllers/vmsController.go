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

func (cc *VmsController) SyncHRDatabase() {
	msSQLModel.SyncHRDB()
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

	err := vmsServerModel.ConnectionTest(data.AccountID, data.Password, data.Protocol, data.Host)
	if err != nil {
		c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}