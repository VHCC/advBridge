/**
 * @apiDefine ServerConfigResponse_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "serverConfig": {
            "HRServer_Account": "rfiduser",
            "HRServer_DatabaseName": "RFID",
            "HRServer_Password": "rf!dus1r375",
            "HRServer_SQLServerHost": "172.20.2.85",
            "HRServer_ViewTableName": "RFID_Employee",
            "RFIDServer_MqttConnectionString": "tcp://104.215.147.159:1883",
            "RFIDServer_MqttTopic": "rfid_temp",
            "RFIDServer_Password": "1JFoR3YbyGaGfNGPGg19Flqzy",
            "RFIDServer_Username": "ec1aceb8-88aa-4b60-8cff-4e8e1cae9e5f:e325b491-edc1-4019-a4e8-675b7c80852c",
            "VMSServer_Account": "Admin",
            "VMSServer_Host": "172.22.24.64:7090",
            "VMSServer_Account": "Admin",
            "VMSServer_Password": "Aa123456*",
            "lastModifiedUnixTimeStamp": 1628134214
        }
 *     }
 */

/**
 * @apiDefine ServerConfigResponse_Success_List_Mac
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
        "macAddress": [
            "ac:bc:32:97:ad:55",
            "82:17:04:61:b3:00",
            "82:17:04:61:b3:01",
            "82:17:04:61:b3:00",
            "0e:bc:32:97:ad:55",
            "42:37:06:42:c6:4b",
            "42:37:06:42:c6:4b"
        ],
        "macAddressWithIP": [
            "ac:bc:32:97:ad:55 / IP:192.168.1.105",
            "82:17:04:61:b3:00 / IP:192.168.1.105",
            "82:17:04:61:b3:01 / IP:192.168.1.105",
            "82:17:04:61:b3:00 / IP:192.168.1.105",
            "0e:bc:32:97:ad:55 / IP:192.168.1.105",
            "42:37:06:42:c6:4b / IP:192.168.1.105",
            "42:37:06:42:c6:4b / IP:192.168.1.105"
        ],
        "message": "SUCCESS"
 *     }
 */

/**
 * @apiDefine ServerConfigResponse_GetEnrollUserFlag_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "user_registration": true,
 *     }
 */
