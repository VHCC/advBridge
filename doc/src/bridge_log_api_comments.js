

/**
 * @apiDefine Vms2LogResponse_List_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "dataCounts" : 8,
 *      "bridgeLogs": [
                {
                    "_id": "6119d34791d32f0c66e0e903",
                    "accountID": "SYSTEM",
                    "logType": "BRIDGE_LOG-CHECK",
                    "message": "SUCCESS",
                    "detail": {
                        "removed": 0
                    },
                    "createUnixTimestamp": 1629082439
                },
                {
                    "_id": "611a04ac91d32f1815741391",
                    "accountID": "aaa",
                    "logType": "VMS-SERVER-UPDATE",
                    "message": "SUCCESS",
                    "detail": {
                        "VMSServer_Account": "Admin",
                        "VMSServer_Host": "localhost:7080",
                        "VMSServer_Password": "Aa123456*"
                    },
                    "createUnixTimestamp": 1629095084
                },
                {
                    "_id": "611a051691d32f183e26385b",
                    "accountID": "SYSTEM",
                    "logType": "VMS-KIOSK-REPORTS-SYNC-START",
                    "message": "SUCCESS",
                    "detail": {},
                    "createUnixTimestamp": 1629095190
                },
                {
                    "_id": "611a051691d32f183e26385c",
                    "accountID": "SYSTEM",
                    "logType": "VMS-KIOSK-REPORTS-SYNC-DONE",
                    "message": "SUCCESS",
                    "detail": {},
                    "createUnixTimestamp": 1629095190
                },
                {
                    "_id": "611a051691d32f183e26385d",
                    "accountID": "SYSTEM",
                    "logType": "VMS-KIOSK-DEVICE-SYNC-SUCCESS",
                    "message": "SUCCESS",
                    "detail": {},
                    "createUnixTimestamp": 1629095190
                }
            ],
 *     }
 */





