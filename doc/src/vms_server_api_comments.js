/**
 * @apiDefine VMSServerResponse_Success_Kiosk_Reports
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "kioskReports": [
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
                }
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
                }
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
                }
            }
        ],
 *       "message": "SUCCESS",
 *     }
 */

/**
 * @apiDefine VMSServerResponse_Success_Kiosk_Devices
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "kioskDevices": [
            {
                "uuid": "60dbcc5c91d32fd90b06ba9e",
                "deviceName": "WFH_Ichen_Kiosk",
                "bridge_Location": "1F Lobby",
                "appVersion": "1.0.8",
                "androidID": "b9b0afa831eb09a1"
            }
        ],
 *       "message": "SUCCESS",
 *     }
 */

