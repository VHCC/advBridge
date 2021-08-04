package apiForms

/**
 * @apiDefine Vms2LogListByPStructure
 * @apiParam {String} userToken userToken <a style="color:red">[required]</a>. <br/>
 * @apiParam {String} sortBy sortBy <a style="color:red">[required]</a>. <br/>
				contains as bellow </br>
				1. accountID, </br>
				2. logType, </br>
 * @apiParam {Boolean} desc desc <a style="color:red">[required]</a>. <br/>
 * @apiParam {int} startIndex startIndex <a style="color:red">[required]</a>. <br/>
 * @apiParam {int} count count <a style="color:red">[required]</a>. <br/>

 * @apiParam {String} keyWords keyWords <a style="color:blue">[optional]</a>. <br/>
					it will fuzzy search columns as below</br>
					1. accountID
 * @apiParam {StringArray} logTypes logTypes <a style="color:blue">[optional]</a>. <br/>
						logTypes includes as below </br>
						"LOGIN", </br>
						"LOGIN-FAIL", </br>
						"LOGOUT", </br>
						"LOGOUT-FAIL", </br>
						"COMPANY-EDIT", </br>
						"TEMPLATE-CREATE", </br>
						"TEMPLATE-EDIT", </br>
						"TEMPLATE-DELETE", </br>
						"KIOSK-DEVICE-ADD", </br>
						"KIOSK-DEVICE-EDIT", </br>
						"KIOSK-DEVICE-REMOVE", </br>
						"PERSON-CREATE", </br>
						"PERSON-EDIT", </br>
						"PERSON-DELETE", </br>
						"CHECKIN-REPORTS-READ", </br>
						"ATTENDANCE-READ", </br>
						"PERSON-IMPORT", </br>
						"USER-CREATE", </br>
						"USER-EDIT", </br>
						"USER-DELETE", </br>
						"SMTP-TEST", </br>
						"RETENTION-UPDATE", </br>
						"SMTP-UPDATE", </br>
						"LICENSE-REGISTER", </br>
						"LOG-EXPORT", </br>
 * @apiParam {Integer} startTimestamp startTimestamp <a style="color:blue">[optional]</a>. <br/>
 * @apiParam {Integer} endTimestamp endTimestamp <a style="color:blue">[optional]</a>. <br/>

* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
	"sortBy": "vmsPersonSerial",
	"desc": true,
	"startIndex": 14,
	"count": 7,
	"keyWords": "Ron"
	"logTypes": [""],
	"startTimestamp": 1603347601,
	"endTimestamp": 1603347605,
}
*/
type ListByPVms2LogDataValidate struct {
	UserToken  *string   `json:"userToken" binding:"required"`
	SortBy     *string   `json:"sortBy" binding:"required"`
	Desc       *bool     `json:"desc" binding:"required"`
	StartIndex *int      `json:"startIndex" binding:"required"`
	Count      *int      `json:"count" binding:"required"`
	KeyWords   *string   `json:"keyWords,omitempty"`
	LogTypes   *[]string `json:"logTypes,omitempty"`
	StartTimestamp *int64  `json:"startTimestamp,omitempty"`
	EndTimestamp   *int64  `json:"endTimestamp,omitempty"`
}
