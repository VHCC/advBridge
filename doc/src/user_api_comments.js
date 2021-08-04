/**
 * @apiDefine UserResponse_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "user": {
                    "userUUID": "5e687b3d91d32f3da23f24a9",
                    "displayName": "new Boy",
                    "email": "userAsia@advantech.com.tw",
                    "role": 1000,
                    "comUUID": "5f7ecb4a91d32f51ab239479",
                    "currentDepUUID": "5f7ecb5d91d32f51ab23947a",
                    "depUUID": "5f7ecb5d91d32f51ab23947a",
                    "userMemo": "test",
                    "permission": [201, 202]
                }
 *     }
 */

/**
 * @apiDefine UserResponse_Enroll_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "company": {
            "comUUID": "60b5b1d75944ba346680b0d9",
            "comName": "AiCS3",
            "comMemo": "33333",
            "status": 1,
            "devicesCount": 0,
            "userCount": 1,
            "userTotal": 0,
            "users": [
                    {
                        "userUUID": "60b5b1d75944ba346680b0db",
                        "accountID": "Ichen3",
                        "email": "r99521323@gmail.com",
                        "role": 5000,
                        "comUUID": "60b5b1d75944ba346680b0d9",
                        "userMemo": "",
                        "permission": [],
                        "createUnixTimeStamp": 1622520279,
                        "lastModifiedUnixTimeStamp": 0,
                        "lastLoginUnixTimeStamp": 0,
                        "allowReviewNonVisitorData": false
                    }
                ],
                "createUnixTimestamp": 1622520279
            },
        "message": "SUCCESS",
        "user": {
                "userUUID": "60b5b1d75944ba346680b0db",
                "accountID": "Ichen3",
                "email": "r99521323@gmail.com",
                "role": 5000,
                "comUUID": "60b5b1d75944ba346680b0d9",
                "userMemo": "",
                "permission": [],
                "createUnixTimeStamp": 1622520279,
                "lastModifiedUnixTimeStamp": 0,
                "lastLoginUnixTimeStamp": 0,
                "allowReviewNonVisitorData": false,
                "userToken": "6KTpv3HKzMRqejHYv7fr5KHbwBUi1VqdxpaaC83uP8M="
            }
 *     }
 */

/**
 * @apiDefine UserResponse_Success_login
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "user": {
                "userUUID": "610a4b4591d32f2bcb6a468a",
                "accountID": "Admin",
                "email": "root",
                "role": 9999,
                "userMemo": "Root Memo",
                "lastLoginUnixTimeStamp": 1628067081,
                "userToken": "_PSI7W1sT7W7UtTzJffrtTHNAILlhCHfwLLaNUDgW-M="
            }
 *     }
 */

/**
 * @apiDefine UserResponse_Logout_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *     }
 */

/**
 * @apiDefine UserResponse_Delete_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "userUUID": "5e68907291d32f4aee777f92"
 *     }
 */

/**
 * @apiDefine UserResponse_List_Success
 *
 *   @apiSuccessExample SUCCESS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 0,
 *       "message": "SUCCESS",
 *       "users": [
                    {
                        "userUUID": "5e6875df91d32f3c2a0a78c9",
                        "displayName": "Frank",
                        "email": "user1@advantech.com.tw",
                        "comUUID": "5f7ecb4a91d32f51ab239479",
                        "currentDepUUID": "5f7ecb5d91d32f51ab23947a",
                        "depUUID": "5f7ecb5d91d32f51ab23947a",
                        "userMemo": "AAA",
                        "role": 1000,
                        "permission": [201, 401],
                        "createUnixTimeStamp": 1586229547,
                        "lastModifiedUnixTimeStamp": 1586229547,
                        "lastLoginUnixTimeStamp": 1586229584,
                    },
                    {
                        "userUUID": "5e6878f891d32f3c2a0a78ca",
                        "displayName": "Jose",
                        "email": "user2@advantech.com.tw",
                        "comUUID": "5f7ecb4a91d32f51ab239479",
                        "currentDepUUID": "5f7ecb5d91d32f51ab23947a",
                        "depUUID": "5f7ecb5d91d32f51ab23947a",
                        "userMemo": "BBB",
                        "role": 2000,
                        "permission": [301, 302],
                        "createUnixTimeStamp": 1586229547,
                        "lastModifiedUnixTimeStamp": 1586229547,
                        "lastLoginUnixTimeStamp": 1586229584,
                    }
                ]
 *     }
 */

/**
 * @apiDefine UserResponse_Invalid_parameter
 *
 *   @apiSuccessExample INVALID_PARAMETERS:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 1,
 *       "message": "INVALID_PARAMETERS"
 *     }
 */

/**
 * @apiDefine UserResponse_User_Exist
 *
 *   @apiSuccessExample USER_EXIST:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 2,
 *       "message": "USER_EXIST"
 *     }
 */

/**
 * @apiDefine UserResponse_User_Email_Exist
 *
 *   @apiSuccessExample USER_EMAIL_EXIST:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 2001,
 *       "message": "USER_EMAIL_EXIST"
 *     }
 */

/**
 * @apiDefine UserResponse_User_Account_ID_Exist
 *
 *   @apiSuccessExample USER_ACCOUNT_ID_EXIST:
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 2002,
 *       "message": "USER_ACCOUNT_ID_EXIST"
 *     }
 */

/**
 * @apiDefine UserResponse_User_Not_Found
 *
 *   @apiSuccessExample USER_NOT_FOUND
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 3,
 *       "message": "USER_NOT_FOUND"
 *     }
 */

/**
 * @apiDefine UserResponse_Project_not_found
 *
 *   @apiSuccessExample PROJECT_NOT_FOUND
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 4,
 *       "message": "PROJECT_NOT_FOUND"
 *     }
 */

/**
 * @apiDefine UserResponse_Group_not_found
 *
 *   @apiSuccessExample GROUP_NOT_FOUND
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 5,
 *       "message": "GROUP_NOT_FOUND"
 *     }
 */

/**
 * @apiDefine UserResponse_wrong_password
 *
 *   @apiSuccessExample WRONG_PASSWORD
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 6,
 *       "message": "WRONG_PASSWORD"
 *     }
 */

/**
 * @apiDefine UserResponse_Company_Inactivate
 *
 *   @apiSuccessExample COMPANY_STATUS_INACTIVATE
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 7,
 *       "message": "COMPANY_STATUS_INACTIVATE"
 *     }
 */

/**
 * @apiDefine UserResponse_user_token_invalid
 *
 *   @apiSuccessExample USER_TOKEN_INVALID
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 1001,
 *       "message": "USER_TOKEN_INVALID"
 *     }
 */

/**
 * @apiDefine UserResponse_no_permission_to_access
 *
 *   @apiSuccessExample NO_USER_PERMISSION
 *     HTTP/1.1 200 OK
 *     {
 *       "code": 4001,
 *       "message": "NO_USER_PERMISSION"
 *     }
 */
