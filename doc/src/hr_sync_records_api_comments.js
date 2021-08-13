

/**
 * @apiDefine HRSyncRecordsResponse_List_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "dataCounts" : 2,
 *       "hrSyncRecordsData": [
            {
                "_id": "6116219091d32fe4be12346a",
                "syncVmsPersonDataCounts": 0,
                "updateVmsPersonDataCounts": 0,
                "createVmsPersonDataCounts": 1952,
                "deleteVmsPersonDataCounts": 0,
                "syncHrServerDataCounts": 1952,
                "status": "Success",
                "failReason": "",
                "VMSServer_Protocol": "http",
                "VMSServer_Host": "localhost:7080",
                "HRServer_SQLServerHost": "172.20.2.85",
                "HRServer_DatabaseName": "RFID",
                "HRServer_ViewTableName": "RFID_Employee",
                "createUnixTimeStamp": 1628840336
            },
            {
                "_id": "6116220e91d32fe4ffceb515",
                "syncVmsPersonDataCounts": 1952,
                "updateVmsPersonDataCounts": 0,
                "createVmsPersonDataCounts": 0,
                "deleteVmsPersonDataCounts": 0,
                "syncHrServerDataCounts": 1952,
                "status": "Success",
                "failReason": "",
                "VMSServer_Protocol": "http",
                "VMSServer_Host": "localhost:7080",
                "HRServer_SQLServerHost": "172.20.2.85",
                "HRServer_DatabaseName": "RFID",
                "HRServer_ViewTableName": "RFID_Employee",
                "createUnixTimeStamp": 1628840462
            }
        ],
 *     }
 */


/**
 * @apiDefine HRSyncRecordsResponse_Detail_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "dataCounts" :  1952,
 *       "hrSyncRecord": {
            "_id": "6116219091d32fe4be12346a",
            "syncVmsPersonDataCounts": 0,
            "updateVmsPersonDataCounts": 0,
            "createVmsPersonDataCounts": 1952,
            "deleteVmsPersonDataCounts": 0,
            "syncHrServerDataCounts": 1952,
            "status": "Success",
            "failReason": "",
            "VMSServer_Protocol": "http",
            "VMSServer_Host": "localhost:7080",
            "HRServer_SQLServerHost": "172.20.2.85",
            "HRServer_DatabaseName": "RFID",
            "HRServer_ViewTableName": "RFID_Employee",
            "createUnixTimeStamp": 1628840336
        },
        "message": "SUCCESS",
        "syncVmsPersons": [
            {
                "vmsPersonUUID": "6116219191d32fe4be12346c",
                "vmsPersonSerial": "6A0999E5",
                "vmsPersonName": "張美玲",
                "vmsPersonUnit": "000000000000000000B00640",
                "vmsPersonEmail": "",
                "vmsPersonMemo": "A-0001",
                "action": "create",
                "status": "SUCCESS",
                "createUnixTimestamp": 0
            },
            {
                "vmsPersonUUID": "6116219191d32fe4be12346e",
                "vmsPersonSerial": "207332F3",
                "vmsPersonName": "劉克振",
                "vmsPersonUnit": "000000000000000000B00140",
                "vmsPersonEmail": "",
                "vmsPersonMemo": "A-0002",
                "action": "create",
                "status": "SUCCESS",
                "createUnixTimestamp": 0
            },
            {
                "vmsPersonUUID": "6116219191d32fe4be123470",
                "vmsPersonSerial": "7ABC9E55",
                "vmsPersonName": "黃麗珠",
                "vmsPersonUnit": "000000000000000000B00061",
                "vmsPersonEmail": "",
                "vmsPersonMemo": "A-0033",
                "action": "create",
                "status": "SUCCESS",
                "createUnixTimestamp": 0
            }
        ]
 *     }
 */





