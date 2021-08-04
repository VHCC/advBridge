package apiForms

/**
 * @apiDefine UserCreateStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} accountID The user's accountID <a style="color:red">[required]</a> <a style="color:green">[unique]</a>.
 * @apiParam {String} email The user's email <a style="color:red">[required]</a> <a style="color:green">[unique]</a>.
 * @apiParam {String} password The user's password <a style="color:red">[required]</a>.
 * @apiParam {String} depUUID belong to Department UUID <a style="color:red">[required]</a>.
 * @apiParam {String} userMemo memo for user <a style="color:blue">[optional]</a>.
 * @apiParam {Integer} role role of the user. <a style="color:red">[required]</a> </br>
						9999: 原廠使用者 root </br>
						5000: 單一公司管理者 Company manager </br>
						1000: 單一部門使用者 user </br>

 * @apiParam {Array} permission user's permission list <a style="color:blue">[optional]</a>. <br/>
						101 as 數據分析 permission. <br/>
						201 as 註冊表單 permission. <br/>
						202 as 名單審核 permission. <br/>
						203 as 統計報表 permission. <br/>
						301 as 註冊管理 permission. <br/>
						302 as 報到流程管理 permission. <br/>
						303 as 報導數字統計 permission. <br/>
						401 as 帳號管理 permission. <br/>
						402 as 部門管理 permission. <br/>
 @apiParamExample {json} Request-Example:
 	    {
			"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
			"accountID": "User005",
			"email": "user005@gmail.com",
			"password": "123456",
			"role": 1000,
			"depUUID": "5f840c8b019ba9193ec99684",
			"permission": [101, 201, 303]
		}
*/
type CreateUserDataValidate struct {
	UserToken  *string `json:"userToken,omitempty" binding:"required"`
	AccountID  string  `json:"accountID" binding:"required"`
	Email      string  `json:"email" binding:"required"`
	Password   string  `json:"password" binding:"required"`
	DepUUID    *string `json:"depUUID" binding:"required"`
	UserMemo   *string `json:"userMemo"`
	Role       int32   `json:"role" binding:"required"`
	Permission *[]int  `json:"permission"`
	Memo       string  `json:"memo"`
}

/**
 * @apiDefine UserListByProjectStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} projectDID projectDID <a style="color:blue">[optional]</a>. <br/>
 * @apiParam {IntArray} role role <a style="color:blue">[optional]</a>. <br/>
 */
type ListUserByProjectDataValidate struct {
	UserToken  *string `json:"userToken,omitempty" binding:"required"`
	ProjectDID string  `json:"projectDID,omitempty"`
	Role       []int32 `json:"role,omitempty"`
}

/**
 * @apiDefine UserListByGroupStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} groupDID groupDID <a style="color:red">[required]</a>.
 */
type ListUserByGroupDataValidate struct {
	UserToken *string `json:"userToken,omitempty" binding:"required"`
	GroupDID  string  `json:"groupDID" binding:"required"`
	Role      []int32 `json:"-"`
}

/**
 * @apiDefine UserUpdateStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} userUUID The user's UUID <a style="color:red">[required]</a>.
 * @apiParam {String} accountID The user's accountID <a style="color:blue">[optional]</a>.
 * @apiParam {String} password The user's password <a style="color:blue">[optional]</a>.
 * @apiParam {String} currentDepUUID current Department view UUID <a style="color:blue">[optional]</a>.
 * @apiParam {String} depUUID belong to Department UUID <a style="color:blue">[optional]</a>.
 * @apiParam {String} userMemo memo for user <a style="color:blue">[optional]</a>.
 * @apiParam {Array} permission user's permission list. <a style="color:blue">[optional]</a> <br/>
						101 as 數據分析 permission. <br/>
						201 as 註冊表單 permission. <br/>
						202 as 名單審核 permission. <br/>
						203 as 統計報表 permission. <br/>
						301 as 註冊管理 permission. <br/>
						302 as 報到流程管理 permission. <br/>
						303 as 報導數字統計 permission. <br/>
						401 as 帳號管理 permission. <br/>
						402 as 部門管理 permission. <br/>
*/
type UpdateUserDataValidate struct {
	UserToken      *string `json:"userToken,omitempty" binding:"required"`
	UserUUID       *string `json:"userUUID,omitempty" binding:"required"`
	AccountID      string  `json:"accountID,omitempty"`
	Password       string  `json:"password,omitempty"`
	UserMemo       *string `json:"userMemo,omitempty"`
	ComUUID        *string `json:"comUUID,omitempty"`
	CurrentDepUUID *string `json:"currentDepUUID,omitempty"`
	DepUUID        *string `json:"depUUID,omitempty"`
	Permission     *[]int  `json:"permission,omitempty"`
}

/**
 * @apiDefine UserRemoveStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} userUUID userUUID <a style="color:red">[required]</a>.
 */
type DeleteUserDataValidate struct {
	UserToken *string `json:"userToken,omitempty" binding:"required"`
	UserUUID  *string `json:"userUUID,omitempty" binding:"required"`
}

/**
 * @apiDefine FetchUserInfoStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} userUUID userUUID <a style="color:red">[required]</a>.
 */
type FetchUserInfoDataValidate struct {
	UserToken *string `json:"userToken,omitempty" binding:"required"`
	UserUUID  *string `json:"userUUID,omitempty" binding:"required"`
}

/**
 * @apiDefine UserLoginStructure
 * @apiParam {String} accountID accountID <a style="color:red">[required]</a>.
 * @apiParam {String} password password email <a style="color:red">[required]</a>.
 */
type LoginUserDataValidate struct {
	AccountID string `json:"accountID,omitempty" binding:"required"`
	Password  string `json:"password,omitempty" binding:"required"`
}

/**
 * @apiDefine UserLogoutStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 */
type LogoutUserDataValidate struct {
	UserToken *string `json:"userToken,omitempty" binding:"required"`
}

/**
 * @apiDefine UserListAllStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 */
type ListUserAllDataValidate struct {
	UserToken *string `json:"userToken,omitempty" binding:"required"`
}

/**
 * @apiDefine UserListStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 */
type ListUsersDataValidate struct {
	UserToken *string `json:"userToken,omitempty" binding:"required"`
}

// ============== V2 ====================
/**
 * @apiDefine UserCreateV2Structure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} comUUID comUUID <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} accountID The user's accountID <a style="color:red">[required]</a> <a style="color:green">[unique]</a>.
 * @apiParam {String} email The user's email <a style="color:red">[required]</a> <a style="color:green">[unique]</a>.
 * @apiParam {String} password The user's password <a style="color:red">[required]</a>.
 * @apiParam {String} userMemo memo for user <a style="color:blue">[optional]</a>.
 * @apiParam {Integer} role role of the user. <a style="color:red">[required]</a> </br>
						5000: amdin of Company </br>
						1000: normal user </br>

 * @apiParam {Array} permission user's permission list <a style="color:blue">[optional]</a>. <br/>
						101 as 數據分析 permission. <br/>
						201 as 註冊表單 permission. <br/>
						202 as 名單審核 permission. <br/>
						203 as 統計報表 permission. <br/>
						301 as 註冊管理 permission. <br/>
						302 as 報到流程管理 permission. <br/>
						303 as 報導數字統計 permission. <br/>
						401 as 帳號管理 permission. <br/>
						402 as 部門管理 permission. <br/>
 @apiParamExample {json} Request-Example:
 	    {
			"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
			"comUUID": "5f17d3d791d32f2045749cb1",
			"accountID": "User005",
			"email": "user005@gmail.com",
			"password": "123456",
			"userMemo": "test memo",
			"role": 1000,
			"permission": [101, 201, 303]
		}
*/
type CreateUserV2DataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
	ComUUID    *string `json:"comUUID" binding:"required"`
	AccountID  string  `json:"accountID" binding:"required"`
	Email      string  `json:"email" binding:"required"`
	Password   string  `json:"password" binding:"required"`
	UserMemo   *string `json:"userMemo,omitempty"`
	Role       int32   `json:"role" binding:"required"`
	Permission *[]int  `json:"permission,omitempty"`
}

/**
 * @apiDefine UserEnrollV2Structure
 * @apiParam {String} accountID The user's accountID <a style="color:red">[required]</a> <a style="color:green">[unique]</a>.
 * @apiParam {String} email The user's email <a style="color:red">[required]</a> <a style="color:green">[unique]</a>.
 * @apiParam {String} password The user's password <a style="color:red">[required]</a>.
 * @apiParam {String} comName comName  <a style="color:red">[required]</a> <a style="color:green">[unique]</a>.
 * @apiParam {String} comMemo comMemo  <a style="color:blue">[optional]</a>
 @apiParamExample {json} Request-Example:
 	    {
			"accountID": "Ichen3",
			"email":"r99521323@gmail.com",
			"password": "123456",
			"comName": "AiCS3",
			"comMemo": "33333"
		}
*/
type EnrollUserNComV2DataValidate struct {
	AccountID *string `json:"accountID" binding:"required"`
	Email     *string `json:"email" binding:"required"`
	Password  *string `json:"password" binding:"required"`
	ComName   *string `json:"comName" binding:"required"`
	ComMemo   *string `json:"comMemo,omitempty"`
}

/**
 * @apiDefine UserUpdateV2Structure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} userUUID The user's UUID <a style="color:red">[required]</a>.
 * @apiParam {String} comUUID comUUID <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} accountID The user's accountID <a style="color:blue">[optional]</a>.
 * @apiParam {String} email The user's email <a style="color:blue">[optional]</a>.
 * @apiParam {String} password The user's password <a style="color:blue">[optional]</a>.
 * @apiParam {String} role The user's role <a style="color:blue">[optional]</a>.
						1000 : normal user
						5000 : admin of company
 * @apiParam {String} userMemo memo for user <a style="color:blue">[optional]</a>.
 * @apiParam {Array} permission user's permission list. <a style="color:blue">[optional]</a> <br/>
						101 as 數據分析 permission. <br/>
						201 as 註冊表單 permission. <br/>
						202 as 名單審核 permission. <br/>
						203 as 統計報表 permission. <br/>
						301 as 註冊管理 permission. <br/>
						302 as 報到流程管理 permission. <br/>
						303 as 報導數字統計 permission. <br/>
						401 as 帳號管理 permission. <br/>
						402 as 部門管理 permission. <br/>
* @apiParam {Boolean} allowReviewNonVisitorData Allow to review non-visitor check-in data <a style="color:blue">[optional]</a>.
*/
type UpdateUserV2DataValidate struct {
	UserToken                 *string `json:"userToken,omitempty" binding:"required"`
	UserUUID                  *string `json:"userUUID,omitempty" binding:"required"`
	ComUUID                   *string `json:"comUUID,omitempty" binding:"required"`
	AccountID                 string  `json:"accountID,omitempty"`
	Email                     string  `json:"email,omitempty"`
	Password                  string  `json:"password,omitempty"`
	UserMemo                  *string `json:"userMemo,omitempty"`
	Role                      *int32  `json:"role,omitempty"`
	Permission                *[]int  `json:"permission,omitempty"`
	AllowReviewNonVisitorData *bool   `json:"allowReviewNonVisitorData,omitempty" `
}
