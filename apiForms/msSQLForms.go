package apiForms

/**
 * @apiDefine MSSQLTestStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} host host <a style="color:red">[required]</a>.
 * @apiParam {String} account account <a style="color:red">[required]</a>.
 * @apiParam {String} password password <a style="color:red">[required]</a>.
 * @apiParam {String} DBName DBName <a style="color:red">[required]</a>.
* @apiParamExample {json} Request-Example:
{
	"userToken": "5on_WOzj-08nSxTfgkaz12HYwswk8b9fRV4Ej9hyTMs=",
	"host":"172.20.2.85",
	"account":"rfiduser",
	"password":"rf!dus1r375",
	"DBName":"RFID"
}
*/

type MSSQLTestDataValidate struct {
	UserToken *string `json:"userToken" binding:"required"`
	Host      string `json:"host" binding:"required"`
	AccountID string `json:"account" binding:"required"`
	Password  string `json:"password" binding:"required"`
	DBName    string `json:"DBName" binding:"required"`
}
