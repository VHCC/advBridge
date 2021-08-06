package apiForms

/**
 * @apiDefine MSSQLTestStructure
 * @apiParam {String} host host <a style="color:red">[required]</a>.
 * @apiParam {String} account account <a style="color:red">[required]</a>.
 * @apiParam {String} password password <a style="color:red">[required]</a>.
 * @apiParam {String} DBName DBName <a style="color:red">[required]</a>.
* @apiParamExample {json} Request-Example:
{
	"host":"172.20.2.85",
	"account":"rfiduser",
	"password":"rf!dus1r375",
	"DBName":"RFID"
}
 */

type MSSQLTestDataValidate struct {
	Host      string `json:"host" binding:"required"`
	AccountID string `json:"account" binding:"required"`
	Password  string `json:"password" binding:"required"`
	DBName    string `json:"DBName" binding:"required"`
}
