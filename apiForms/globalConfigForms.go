package apiForms

/**
 * @apiDefine GlobalConfigGetStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 */
type GetGlobalConfigDataValidate struct {
	UserToken *string `json:"userToken,omitempty" binding:"required"`
}

/**
 * @apiDefine GlobalConfigUpdateStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
				 Only root has the permission to update server config. <br/>
 * @apiParam {jsonObject} bundle bundle <br/>
							it's flexible <KEY:string, VALUE:string>
* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
	"bundle" :{
		"VMSServer_Host": "172.22.24.64:7090",
		"VMSServer_Account": "Admin",
		"VMSServer_Password": "Aa123456*",
	}
}
*/
type UpdateGlobalConfigDataValidate struct {
	UserToken *string                `json:"userToken,omitempty" binding:"required"`
	Bundle    map[string]interface{} `json:"bundle,omitempty" bson:"bundle"`
}

/**
 * @apiDefine GlobalConfigListMacStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
*/
type ListMacAddressGlobalConfigDataValidate struct {
	UserToken *string                `json:"userToken,omitempty" binding:"required"`
}
