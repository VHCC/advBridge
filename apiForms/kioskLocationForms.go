package apiForms

/**
 * @apiDefine KioskLocationCreateDataValidate
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
type KioskLocationCreateDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
	DeviceUUID *string `json:"deviceUUID" binding:"required"`
	Location   *string `json:"location" binding:"required"`
}

/**
 * @apiDefine KioskLocationDeleteDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} deviceUUID deviceUUID <a style="color:red">[required]</a>. <br/>
* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
	"deviceUUID": "60dbcbf391d32fd90b06ba92",
}
*/
type KioskLocationDeleteDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
	DeviceUUID *string `json:"deviceUUID" binding:"required"`
}

/**
 * @apiDefine KioskLocationFetchAllDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
}
*/
type KioskLocationFetchAllDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
}

/**
 * @apiDefine KioskLocationUpdateDataValidate
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
type KioskLocationUpdateDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
	DeviceUUID *string `json:"deviceUUID" binding:"required"`
	Location   *string `json:"location" binding:"required"`
}
