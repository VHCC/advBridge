package controllers

import (
	"advBridge/apiForms"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	logv "github.com/sirupsen/logrus"
	"advBridge/models"
	"advBridge/utils"
	//"strings"
)

type UserController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var userModel = new(models.UserModel)
var userUtil = new(utils.UserUtil)

func (userC *UserController) InitUser() {
	//defaultCom, err := vms2ComModel.Init()
	//if err != nil {
	//	logv.Error("Init err:> ", err)
	//}
	//err = vmsCostLogModel.InsertComPointToDB(defaultCom.ID.Hex())
	//if err != nil {
	//	logv.Error("InsertComPointToDB err:> ", err)
	//}
	//defaultDepartment, _ := vmsDepModel.Init(defaultCom)

	_ = userModel.Init()
}

/**
@api {POST} /api/v1/user/loginUser User Login
@apiDescription User Login
@apiversion 0.0.1
@apiGroup 001 User
@apiName Login User

@apiUse UserLoginStructure

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
				  3:USER_NOT_FOUND (找不到User) </br>
				  6:WRONG_PASSWORD (密碼錯誤) </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {User}  user  User Info
*
* @apiUse UserResponse_Success_login
* @apiUse UserResponse_Invalid_parameter
 @apiUse UserResponse_User_Not_Found
 @apiUse UserResponse_wrong_password
*/
func (userC *UserController) LoginUser(c *gin.Context) {
	var data apiForms.LoginUserDataValidate

	// formData validation
	if c.ShouldBind(&data) != nil {
		logv.Error("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	isRoot, root, err := userModel.IsRoot(data.AccountID, data.Password)
	logv.Info("isRoot:> ", isRoot)
	if isRoot {
		root, err = userModel.UpdateUserLogin(root.UserUUID)
		err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN, root.AccountID, "SUCCESS", nil)
		c.JSON(200, gin.H{"code": 0, "message": "SUCCESS", "user": root})
		c.Abort()
		return
	}

	//var mode = strings.ToLower(models.SERVER_MODE)
	//switch(mode) {
	//case "edge":
	//	if root.Role == 9999 {
	//		err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN_FAIL, data.AccountID, "USER_NOT_FOUND", nil)
	//		c.JSON(200, gin.H{"code": 3, "message": "USER_NOT_FOUND"})
	//		c.Abort()
	//		return
	//	}
	//	break;
	//case "cloud":
	//	if root.AccountID == "Admin" && root.Password == "Aa123456*"{
	//		err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN_FAIL, data.AccountID, "USER_NOT_FOUND", nil)
	//		c.JSON(200, gin.H{"code": 3, "message": "USER_NOT_FOUND"})
	//		c.Abort()
	//		return
	//	}
	//	break;
	//}
	if root.AccountID == "Admin" && root.Password == "Aa123456*"{
		err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN_FAIL, data.AccountID, "USER_NOT_FOUND", nil)
		c.JSON(200, gin.H{"code": 3, "message": "USER_NOT_FOUND"})
		c.Abort()
		return
	}


	user, err := userModel.IsUserAccountIDExist(data.AccountID)
	if err != nil {
		logv.Error("IsUserAccountIDExist err:> ", err)
		err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN_FAIL, data.AccountID, "USER_NOT_FOUND", nil)
		c.JSON(200, gin.H{"code": 3, "message": "USER_NOT_FOUND"})
		c.Abort()
		return
	}

	if user.Password != data.Password {
		logv.Error("LoginUser err:> ", "Wrong password")
		err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN_FAIL, user.AccountID, "WRONG_PASSWORD", nil)
		c.JSON(200, gin.H{"code": 6, "message": "WRONG_PASSWORD"})
		c.Abort()
		return
	}

	//com, err := vms2ComModel.FindComUUID(user.ComUUID)
	//
	//if err != nil {
	//	logv.Error("FindComUUID err:> ", "VMS_COMPANY_NOT_FOUND")
	//	err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN_FAIL, user.AccountID, "VMS_COMPANY_NOT_FOUND", nil)
	//	c.JSON(200, gin.H{"code": 15001, "message": "VMS_COMPANY_NOT_FOUND"})
	//	c.Abort()
	//	return
	//}

	//if com.Status == 0 {
	//	logv.Error("FindComUUID err:> ", "COMPANY_STATUS_INACTIVATE")
	//	err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN_FAIL, user.AccountID, "COMPANY_STATUS_INACTIVATE", nil)
	//	c.JSON(200, gin.H{"code": 7, "message": "COMPANY_STATUS_INACTIVATE"})
	//	c.Abort()
	//	return
	//}

	user, err = userModel.UpdateUserLogin(user.UserUUID)

	err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGIN, user.AccountID, "SUCCESS", nil)
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS", "user": user})
}

/**
@api {POST} /api/v1/user/logoutUser User Logout
@apiDescription User Logout
@apiversion 0.0.1
@apiGroup 001 User
@apiName Logout User

@apiUse UserLogoutStructure

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse UserResponse_Logout_Success
* @apiUse UserResponse_Invalid_parameter
*/
func (userC *UserController) LogoutUser(c *gin.Context) {
	var data apiForms.LogoutUserDataValidate

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
	case 1001:
		c.JSON(200, gin.H{"code": 1001, "message": "USER_TOKEN_INVALID"})
		c.Abort()
		return
	}

	user, err := userModel.UpdateUserLogout(queryUser.UserUUID)
	if err != nil {
		logv.Error("UpdateUserLogout err:> ", err)
		err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGOUT_FAIL, user.AccountID, err.Error(), nil)
	}

	err = logModel.WriteLog(models.EVENT_TYPE_USER_LOGOUT, user.AccountID, "SUCCESS", nil)
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}

