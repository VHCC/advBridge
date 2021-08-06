package apiForms

/**
 * @apiDefine VMSServerTestDataValidate
 * @apiParam {String} account account <a style="color:red">[required]</a>.
 * @apiParam {String} password password <a style="color:red">[required]</a>.
 * @apiParam {String} protocol protocol <a style="color:red">[required]</a>.
 * @apiParam {String} host host <a style="color:red">[required]</a>.
* @apiParamExample {json} Request-Example:
{
	"account":"Admin",
	"password":"Aa123456*",
	"protocol":"http",
	"host":"192.11.9.121:80",
}
*/

type VMSServerTestDataValidate struct {
	AccountID string `json:"account" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Protocol  string `json:"protocol" binding:"required"`
	Host      string `json:"host" binding:"required"`
}
