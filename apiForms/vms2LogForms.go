package apiForms

/**
 * @apiDefine ListByPBridgeLogDataValidate
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
						1. "VMS-SERVER-UPDATE", </br>
						2. "HR-SERVER-UPDATE", </br>
						3. "RFID-SERVER-UPDATE", </br>
						4. "KIOSK-LOCATION-CREATE", </br>
						5. "KIOSK-LOCATION-UPDATE", </br>
						6. "KIOSK-LOCATION-DELETE", </br>
						7. "BRIDGE_LOG-CHECK", </br>
						8. "RFID-CONNECT-SUCCESS", </br>
						9. "VMS-CONNECT-SUCCESS", </br>
						10. "HR-CONNECT-SUCCESS", </br>
						11. "RFID-CONNECT-FAIL", </br>
						12. "VMS-CONNECT-FAIL", </br>
						13. "HR-CONNECT-FAIL", </br>
						14. "VMS-KIOSK-REPORTS-SYNC-START", </br>
						15. "VMS-KIOSK-REPORTS-SYNC-DONE", </br>
						16. "VMS-KIOSK-REPORTS-SYNC-FAIL", </br>
						17. "VMS-KIOSK-DEVICE-SYNC-SUCCESS", </br>
						18. "VMS-KIOSK-DEVICE-SYNC-FAIL", </br>
						19. "HR-SERVER-SYNC-FAIL", </br>
						20. "HR-SERVER-SYNC-SUCCESS", </br>
						21. "VMS-SERVER-SYNC-SUCCESS", </br>
						22. "VNS-SERVER-SYNC-FAIL", </br>
						22. "VNS-SERVER-SYNC-FAIL", </br>
						23. "HR-RECORDS-CHECK", </br>
						24. "VMS-RECORDS-CHECK", </br>
 * @apiParam {Integer} startTimestamp startTimestamp <a style="color:blue">[optional]</a>. <br/>
 * @apiParam {Integer} endTimestamp endTimestamp <a style="color:blue">[optional]</a>. <br/>

* @apiParamExample {json} Request-Example:
{
	"userToken": "dO3Hi3AsYKrW4KmH_5rWo1uM6vpJRloCV3trWtuD1XM=",
	"sortBy": "vmsPersonSerial",
	"desc": true,
	"startIndex": 14,
	"count": 7,
	"keyWords": "Ruby"
	"logTypes": [""],
	"startTimestamp": 1603347601,
	"endTimestamp": 1603347605,
}
*/
type ListByPBridgeLogDataValidate struct {
	UserToken      *string   `json:"userToken" binding:"required"`
	SortBy         *string   `json:"sortBy" binding:"required"`
	Desc           *bool     `json:"desc" binding:"required"`
	StartIndex     *int      `json:"startIndex" binding:"required"`
	Count          *int      `json:"count" binding:"required"`
	KeyWords       *string   `json:"keyWords,omitempty"`
	LogTypes       *[]string `json:"logTypes,omitempty"`
	StartTimestamp *int64    `json:"startTimestamp,omitempty"`
	EndTimestamp   *int64    `json:"endTimestamp,omitempty"`
}
