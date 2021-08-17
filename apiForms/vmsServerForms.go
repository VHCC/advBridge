package apiForms

/**
 * @apiDefine VMSServerTestDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} account account <a style="color:red">[required]</a>.
 * @apiParam {String} password password <a style="color:red">[required]</a>.
 * @apiParam {String} protocol protocol <a style="color:red">[required]</a>.
 * @apiParam {String} host host <a style="color:red">[required]</a>.
* @apiParamExample {json} Request-Example:
{
	"userToken": "5on_WOzj-08nSxTfgkaz12HYwswk8b9fRV4Ej9hyTMs=",
	"account":"Admin",
	"password":"Aa123456*",
	"protocol":"http",
	"host":"192.11.9.121:80",
}
*/
type VMSServerTestDataValidate struct {
	UserToken *string `json:"userToken" binding:"required"`
	AccountID string  `json:"account" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Protocol  string  `json:"protocol" binding:"required"`
	Host      string  `json:"host" binding:"required"`
}

/**
 * @apiDefine VMSServerKioskReportsFetchDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM="
}
*/
type VMSServerKioskReportsFetchDataValidate struct {
	UserToken *string `json:"userToken" binding:"required"`
}

/**
 * @apiDefine VMSServerKioskDevicesFetchDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM="
}
*/
type VMSServerKioskDevicesFetchDataValidate struct {
	UserToken *string `json:"userToken" binding:"required"`
}

/**
 * @apiDefine VMSServerKioskDeviceUpdateDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} deviceUUID deviceUUID <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} location location <a style="color:red">[required]</a>. <br/>
* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
	"deviceUUID": "60dbcbf391d32fd90b06ba92",
	"location": "1F Lobby"
}
*/
type VMSServerKioskDeviceUpdateDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
	DeviceUUID *string `json:"deviceUUID" binding:"required"`
	Location   *string `json:"location" binding:"required"`
}
