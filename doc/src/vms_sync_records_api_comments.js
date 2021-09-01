

/**
 * @apiDefine VmsSyncRecordsResponse_List_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "dataCounts" : 4,
 *       "vmsSyncRecordsData": [
            {
                "_id": "611338fb91d32f6386d4ff3f",
                "syncVmsDataCounts": 3,
                "RFIDDataSendCounts": 3,
                "status": "Success",
                "failReason": "",
                "VMSServer_Protocol": "http",
                "VMSServer_Host": "localhost:7080",
                "RFIDServer_MqttConnectionString": "tcp://104.215.147.159:1883",
                "RFIDServer_MqttTopic": "rfid_temp",
                "createUnixTimeStamp": 1628649723
            },
            {
                "_id": "611338e791d32f6386d4ff3b",
                "syncVmsDataCounts": 0,
                "RFIDDataSendCounts": 0,
                "status": "Success",
                "failReason": "",
                "VMSServer_Protocol": "http",
                "VMSServer_Host": "localhost:7080",
                "RFIDServer_MqttConnectionString": "tcp://104.215.147.159:1883",
                "RFIDServer_MqttTopic": "rfid_temp",
                "createUnixTimeStamp": 1628649703
            },
            {
                "_id": "611338f191d32f6386d4ff3d",
                "syncVmsDataCounts": 0,
                "RFIDDataSendCounts": 0,
                "status": "Success",
                "failReason": "",
                "VMSServer_Protocol": "http",
                "VMSServer_Host": "localhost:7080",
                "RFIDServer_MqttConnectionString": "tcp://104.215.147.159:1883",
                "RFIDServer_MqttTopic": "rfid_temp",
                "createUnixTimeStamp": 1628649713
            },
            {
                "_id": "6113390591d32f6386d4ff41",
                "syncVmsDataCounts": 0,
                "RFIDDataSendCounts": 0,
                "status": "Success",
                "failReason": "",
                "VMSServer_Protocol": "http",
                "VMSServer_Host": "localhost:7080",
                "RFIDServer_MqttConnectionString": "tcp://104.215.147.159:1883",
                "RFIDServer_MqttTopic": "rfid_temp",
                "createUnixTimeStamp": 1628649733
            }
        ]
 *     }
 */


/**
 * @apiDefine VmsSyncRecordsResponse_Detail_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "dataCounts" : 3,
 *       "syncKioskReports": [
        {
            "uuid": "610cffac91d32f9c0389c643",
            "avalo_device": "WFH_Ichen_Kiosk",
            "avalo_device_uuid": "60dbcc5c91d32fd90b06ba9e",
            "avalo_mode": "advanced",
            "avalo_interface": "rfid",
            "avalo_snapshot": "ADFADF",
            "avalo_status": "exception",
            "avalo_exception": "high-fever",
            "avalo_temperature": 36.5,
            "avalo_temperature_threshold": 38,
            "avalo_temperature_adjust": 0.5,
            "avalo_temperature_unit": "C",
            "avalo_mask": true,
            "avalo_utc_timestamp": 1628241836287,
            "vmsPerson": {
                "_id": "60e59ced91d32fb556a1d3f5",
                "vmsPersonEmail": "temp person",
                "vmsPersonMemo": "temp memo",
                "vmsPersonName": "temp name",
                "vmsPersonSerial": "temp serialNumber",
                "vmsPersonUnit": "temp unit"
            },
            "syncStatus": false
        },
        {
            "uuid": "610cffcd91d32f9c0389c645",
            "avalo_device": "WFH_Ichen_Kiosk",
            "avalo_device_uuid": "60dbcc5c91d32fd90b06ba9e",
            "avalo_mode": "advanced",
            "avalo_interface": "rfid",
            "avalo_snapshot": "ADFADF",
            "avalo_status": "exception",
            "avalo_exception": "high-fever",
            "avalo_temperature": 36.5,
            "avalo_temperature_threshold": 38,
            "avalo_temperature_adjust": 0.5,
            "avalo_temperature_unit": "C",
            "avalo_mask": true,
            "avalo_utc_timestamp": 1628241869100,
            "vmsPerson": {
                "_id": "60e59ced91d32fb556a1d3f5",
                "vmsPersonEmail": "temp person",
                "vmsPersonMemo": "temp memo",
                "vmsPersonName": "temp name",
                "vmsPersonSerial": "temp serialNumber",
                "vmsPersonUnit": "temp unit"
            },
            "syncStatus": true
        },
        {
            "uuid": "6110c1bd91d32fa798177c16",
            "avalo_device": "WFH_Ichen_Kiosk",
            "avalo_device_uuid": "60dbcc5c91d32fd90b06ba9e",
            "avalo_mode": "advanced",
            "avalo_interface": "rfid",
            "avalo_snapshot": "ADFADF",
            "avalo_status": "exception",
            "avalo_exception": "high-fever",
            "avalo_temperature": 36.5,
            "avalo_temperature_threshold": 38,
            "avalo_temperature_adjust": 0.5,
            "avalo_temperature_unit": "C",
            "avalo_mask": true,
            "avalo_utc_timestamp": 1628488125103,
            "vmsPerson": {
                "_id": "60e59ced91d32fb556a1d3f5",
                "vmsPersonEmail": "temp person",
                "vmsPersonMemo": "temp memo",
                "vmsPersonName": "temp name",
                "vmsPersonSerial": "temp serialNumber",
                "vmsPersonUnit": "temp unit"
            },
            "syncStatus": true
        }
    ],
    "vmsSyncRecord": {
        "_id": "611338fb91d32f6386d4ff3f",
        "syncVmsDataCounts": 3,
        "RFIDDataSendCounts": 3,
        "status": "Success",
        "failReason": "",
        "VMSServer_Protocol": "http",
        "VMSServer_Host": "localhost:7080",
        "RFIDServer_MqttConnectionString": "tcp://104.215.147.159:1883",
        "RFIDServer_MqttTopic": "rfid_temp",
        "createUnixTimeStamp": 1628649723
    }
 *     }
 */





