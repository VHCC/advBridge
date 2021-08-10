package controllers

import (
	"advBridge/apiForms"
	"advBridge/models"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	logv "github.com/sirupsen/logrus"
)

type KioskLocationController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var kioskLocationModel = new(models.VmsServerModel)

func (cc *KioskLocationController) SyncVMSKioskReportsData() {
}


/**
@api {POST} /api/v1/kioskLocation/create Create Kiosk Location
@apiDescription Create Kiosk Location
@apiversion 0.0.1
@apiGroup 004 Kiosk Location
@apiName Create Kiosk Location

@apiUse KioskLocationCreateDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse KioskLocationResponse_Success_Create_Remove_Edit
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *KioskLocationController) CreateLocation(c *gin.Context) {
	var data apiForms.KioskLocationCreateDataValidate

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

	err := vmsServerModel.CreateKioskLocation(data)
	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}

/**
@api {POST} /api/v1/kioskLocation/remove Delete Kiosk Location
@apiDescription Delete Kiosk Location
@apiversion 0.0.1
@apiGroup 004 Kiosk Location
@apiName Delete Kiosk Location

@apiUse KioskLocationDeleteDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse KioskLocationResponse_Success_Create_Remove_Edit
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *KioskLocationController) RemoveLocation(c *gin.Context) {
	var data apiForms.KioskLocationDeleteDataValidate

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

	err := vmsServerModel.RemoveKioskLocation(data)
	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}

/**
@api {POST} /api/v1/kioskLocation/fetchAll Fetch All Of Kiosk Locations
@apiDescription Fetch All Of Kiosk Locations
@apiversion 0.0.1
@apiGroup 004 Kiosk Location
@apiName Fetch All Of Kiosk Locations

@apiUse KioskLocationFetchAllDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse KioskLocationResponse_Success_FetchAll
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *KioskLocationController) FetchAllLocation(c *gin.Context) {
	var data apiForms.KioskLocationFetchAllDataValidate

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

	kioskLocations, err := vmsServerModel.FetchAllKioskLocation()
	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}

	if len(kioskLocations) == 0 {
		c.JSON(200, gin.H{"code": 0, "message": "SUCCESS",
			"kioskLocations": []string{}})
	} else {

		c.JSON(200, gin.H{"code": 0, "message": "SUCCESS",
			"kioskLocations": kioskLocations})
	}
}

/**
@api {POST} /api/v1/kioskLocation/edit Edit Kiosk Location
@apiDescription Edit Kiosk Location
@apiversion 0.0.1
@apiGroup 004 Kiosk Location
@apiName Edit Kiosk Location

@apiUse KioskLocationUpdateDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse KioskLocationResponse_Success_Create_Remove_Edit
* @apiUse UserResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (cc *KioskLocationController) EditLocation(c *gin.Context) {
	var data apiForms.KioskLocationUpdateDataValidate

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

	err := vmsServerModel.UpdateKioskLocation(data)
	if err != nil {
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}