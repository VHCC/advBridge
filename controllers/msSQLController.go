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

func (cc *MsSQLController) SyncHRDatabase() {
	msSQLModel.SyncHRDB()
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
*                 2001:CONNECT_ERROR (參數缺少或錯誤) </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse HRServerResponse_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse HRServerResponse_Connect_Err
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

	_, err := msSQLModel.ConnectionTest(data.Host, data.AccountID, data.Password, data.DBName)
	if err != nil {
		c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}