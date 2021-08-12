package apiForms

/**
 * @apiDefine ListByPVmsSyncRecordsDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} sortBy  排序欄位, 請參考 vmsSyncRecordsData 欄位 {JsonObject} <a style="color:red">[required]</a>. <br/>
 * @apiParam {Boolean} desc  升冪/降冪 <a style="color:red">[required]</a>. <br/>
 * @apiParam {int} startIndex  請求的開始筆數 <a style="color:red">[required]</a>. <br/>
 * @apiParam {int} count  請求的筆數, 設置 -1 則返回Server全部資料 <a style="color:red">[required]</a>. <br/>

* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
	"sortBy": "syncVmsDataCounts",
	"desc": true,
	"startIndex": 0,
	"count": 10
}
*/
type ListByPHrSyncRecordsDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
	SortBy     *string `json:"sortBy" binding:"required"`
	Desc       *bool   `json:"desc" binding:"required"`
	StartIndex *int    `json:"startIndex" binding:"required"`
	Count      *int    `json:"count" binding:"required"`
}

/**
 * @apiDefine ListByPVmsSyncRecordsDetailDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} recordUUID VMS 同步資訊 UUID <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} sortBy  排序欄位, 請參考 vmsSyncRecordsData 欄位 {JsonObject} <a style="color:red">[required]</a>. <br/>
 * @apiParam {Boolean} desc  升冪/降冪 <a style="color:red">[required]</a>. <br/>
 * @apiParam {int} startIndex  請求的開始筆數 <a style="color:red">[required]</a>. <br/>
 * @apiParam {int} count  請求的筆數, 設置 -1 則返回Server全部資料 <a style="color:red">[required]</a>. <br/>

* @apiParamExample {json} Request-Example:
{
	"userToken": "5on_WOzj-08nSxTfgkaz12HYwswk8b9fRV4Ej9hyTMs=",
    "recordUUID" : "611338fb91d32f6386d4ff3f",
    "sortBy": "avalo_temperature",
	"desc": false,
	"startIndex": 0,
	"count": 10
}
*/
type ListByPHrSyncRecordsDetailDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
	RecordUUID *string `json:"recordUUID" binding:"required"`
	SortBy     *string `json:"sortBy" binding:"required"`
	Desc       *bool   `json:"desc" binding:"required"`
	StartIndex *int    `json:"startIndex" binding:"required"`
	Count      *int    `json:"count" binding:"required"`
}

/**
 * @apiDefine RequestSyncWithHRDataValidate
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>

* @apiParamExample {json} Request-Example:
{
	"userToken": "5on_WOzj-08nSxTfgkaz12HYwswk8b9fRV4Ej9hyTMs=",
}
*/
type RequestSyncWithHRDataValidate struct {
	UserToken  *string `json:"userToken" binding:"required"`
}
