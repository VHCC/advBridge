package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	logv "github.com/sirupsen/logrus"
	_ "image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"advBridge/controllers"
	"advBridge/models"
	"strings"
	"time"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jasonlvhit/gocron"
)

var exportPORT = "7090"
var buildVersion = "v0.0.1"

func main() {
	logv.Info(" === adb Bridge server start === ")
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/js", "./templates")

	// ================ v1 ===================
	ticker10sv1 := time.NewTicker(1 * 10 * time.Second)
	ticker1m1 := time.NewTicker(1 * 60 * time.Second)
	ticker5mv1 := time.NewTicker(5 * 60 * time.Second)
	v1 := router.Group("/api/v1")
	{
		userController := new(controllers.UserController)

		userController.InitUser()

		//v1.POST("/user/create", userController.CreateUser)
		//v1.POST("/user/updateUser", userController.UpdateUser)
		//v1.POST("/user/deleteUser", userController.DeleteUser)
		v1.POST("/user/loginUser", userController.LoginUser)
		v1.POST("/user/logoutUser", userController.LogoutUser)
		//v1.POST("/user/listAllUsers", userController.ListAllUsers)
		//v1.POST("/user/listUsers", userController.ListUsers)
		//v1.POST("/user/fetchUserInfo", userController.FetchUserInfo)

		// ======== 中介 ========
		msSQLController := new(controllers.MsSQLController)
		v1.POST("/hrServer/connectTest", msSQLController.MSSQLConnectionTest)


		vmsServerController := new(controllers.VmsController)
		vmsServerController.SyncVMSKioskReportsData()
		v1.POST("/vmsServer/connectTest", vmsServerController.VmsServerConnectionTest)
		v1.POST("/vmsServer/fetchVMSKioskReports", vmsServerController.FetchVmsKioskReports)
		v1.POST("/vmsServer/fetchVMSKioskDevices", vmsServerController.FetchVmsKioskDevices)

		mqttTopicController := new(controllers.TopicController)
		mqttTopicController.Init()

		kioskLocationController := new(controllers.KioskLocationController)
		v1.POST("/kioskLocation/create", kioskLocationController.CreateLocation)
		v1.POST("/kioskLocation/delete", kioskLocationController.RemoveLocation)
		v1.POST("/kioskLocation/fetchAll", kioskLocationController.FetchAllLocation)
		v1.POST("/kioskLocation/edit", kioskLocationController.EditLocation)


		//vmsFormController := new(controllers.VmsFormController)
		// ========== VMS ============
		//v1.POST("/form/createForm", vmsFormController.CreateForm)

		//vmsDepartmentController := new(controllers.VmsDepartmentController)
		//v1.POST("/dep/createDep", vmsDepartmentController.CreateDep)
		//v1.POST("/dep/getInfoDep", vmsDepartmentController.GetDepInfo)
		//v1.POST("/dep/deleteDep", vmsDepartmentController.DeleteDep)
		//v1.POST("/dep/fetchAllDeps", vmsDepartmentController.FetchAllDepartments)
		//v1.POST("/dep/updateDep", vmsDepartmentController.UpdateDep)

		//vmsCompanyController := new(controllers.VmsCompanyController)
		//v1.GET("/com/verifyCompany", vmsCompanyController.EnrollCompany)
		//v1.POST("/com/getInfoCompany", vmsCompanyController.GetCompanyInfo)

		//vmsVisitorController := new(controllers.VmsVisitorController)
		//v1.POST("/visitor/createVisitor", vmsVisitorController.CreateVisitor)
		//v1.POST("/visitor/getInfoVisitor", vmsVisitorController.GetVisitorInfo)
		//v1.POST("/visitor/getInfoVisitorByCheckCode", vmsVisitorController.GetVisitorByCheckCode)
		//v1.POST("/visitor/fetchVisitorsByPersonUUID", vmsVisitorController.FetchVisitorsByPersonUUID)
		//v1.POST("/visitor/fetchAllVisitors", vmsVisitorController.FetchAllVisitors)
		//v1.POST("/visitor/deleteVisitor", vmsVisitorController.DeleteVisitor)
		//v1.POST("/visitor/fetchFormRelatedVisitors", vmsVisitorController.FetchVisitorsByFormUUID)
		//v1.POST("/visitor/updateVisitor", vmsVisitorController.UpdateVisitor)

		//vmsClientController := new(controllers.VmsClientLayoutController)
		//v1.POST("/clientLayout/getInfoClientLayout", vmsClientController.GetClientLayoutInfo)
		//v1.POST("/clientLayout/updateClientLayout", vmsClientController.UpdateClientLayout)

		//vmsDeviceController := new(controllers.VmsClientDeviceController)
		//v1.POST("/clientDevice/deviceLogin", vmsDeviceController.DeviceLoginCheck)
		//v1.POST("/clientDevice/deviceHeartBeat", vmsDeviceController.DeviceHeartBeat)
		//v1.POST("/clientDevice/listDevices", vmsDeviceController.ListClientDevice)
		//v1.POST("/clientDevice/deviceLogout", vmsDeviceController.DeviceLogout)
		//v1.POST("/clientDevice/updateDevice", vmsDeviceController.UpdateDevice)

		//vmsCostLogController := new(controllers.VmsCostLogController)
		//v1.POST("/billing/fetchCompanyAccountingInfo", vmsCostLogController.FetchCompanyAccountingInfo)
		//v1.POST("/costLog/getCostLogReports", vmsCostLogController.GetCostLogReports)
		//v1.POST("/statistics/getDashboard", vmsCostLogController.GetDashboardInfo)
		//v1.POST("/statistics/getVisitorsInPeriods", vmsCostLogController.GetVisitorsInPeriods)

		//atLicenceController := new(controllers.AtLicenceController)
		//v1.POST("/licence/deposit", atLicenceController.DepositPoint)
		//v1.POST("/licence/activateLicence", atLicenceController.ActivateLicense)
		//v1.POST("/licence/listByParameter", atLicenceController.ListByParameter)
		//v1.POST("/licence/faceFeature", atLicenceController.ActivateFaceFeature)

		//vmsRootController := new(controllers.VmsRootController)
		//v1.POST("/root/getRootDashboardInfo", vmsRootController.GetDashboardInfo)

		go func() {
			for {
				select {
				case t5 := <-ticker5mv1.C:
					_ = t5
					mqttTopicController.SendDataToServer()
					//logv.Info("Tick 5 min at:> ", t5)
					//logv.Info(" === CheckKioskDeviceStatus === ")
					//atLicenceController.CheckKioskDeviceStatus()
					//_ = t5
					//case t30s := <-ticker60s_v2.C:
					//logv.Info("Tick 60 s at:> ", t30s)
					//vmsCostLogController.CheckClientDeviceCostAndWriteLog()
				case t10s := <-ticker10sv1.C:
					//logv.Info("Tick 10 s at:> ", t10s)
					_ = t10s
					targetTimeStamp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
						04, 59, 59, 0, time.UTC)
					nowTimeStamp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
						time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
					//logv.Info("now:> ", nowTimeStamp.Unix(), " || targetTime:> ", targetTimeStamp.Unix())
					//logv.Info()
					//
					if nowTimeStamp.Unix() > targetTimeStamp.Unix() {
					}
					case t1m := <-ticker1m1.C:
						_ = t1m
						vmsServerController.SyncVMSKioskReportsData()
				}
			}
		}()
	}

	gocron.ChangeLoc(time.UTC)
	gocron.Every(1).Days().At("05:00").Do(syncTask)
	//gocron.Every(1).Second().Do(syncTask)
	gocron.Start()

	// ====================== v2 =======================
	ticker10sv2 := time.NewTicker(1 * 10 * time.Second)
	ticker3mv2 := time.NewTicker(3 * 60 * time.Second)
	ticker60mv2 := time.NewTicker(60 * 60 * time.Second)
	ticker10mv2 := time.NewTicker(10 * 60 * time.Second)
	//v2 := router.Group("/api/v2")
	{
		//vms2PersonController := new(controllers.Vms2PersonController)
		//v2.POST("/vmsPerson/createVmsPerson", vms2PersonController.CreateVMSPerson)
		//v2.POST("/vmsPerson/getInfoVmsPerson", vms2PersonController.GetVmsPersonInfo)
		//v2.POST("/vmsPerson/deleteVmsPerson", vms2PersonController.DeleteVmsPerson)
		//v2.POST("/vmsPerson/updateVmsPerson", vms2PersonController.UpdateVmsPerson)
		//v2.POST("/vmsPerson/checkVmsPersonSerialByDevice", vms2PersonController.CheckVmsPersonSerialByDevice)
		//v2.POST("/vmsPerson/checkVmsPersonUUIDByDevice", vms2PersonController.CheckVmsPersonUUIDByDevice)
		//v2.POST("/vmsPerson/listVmsPersonByParameter", vms2PersonController.ListVmsPersonByPData)
		//v2.POST("/vmsPerson/importPersonByBatch", vms2PersonController.ImportBatchPerson)

		//vmsF2TemplateController := new(controllers.Vms2TemplateController)
		// ========== VMS ============
		//v2.POST("/template/createTemplate", vmsF2TemplateController.CreateTemplate)
		//v2.POST("/template/fetchAllTemplateList", vmsF2TemplateController.FetchAllTemplates)
		//v2.POST("/template/getTemplateInfo", vmsF2TemplateController.GetTemplateInfo)
		//v2.POST("/template/deleteTemplate", vmsF2TemplateController.DeleteTemplate)
		//v2.POST("/template/updateTemplate", vmsF2TemplateController.UpdateTemplate)
		//v2.POST("/template/duplicateTemplate", vmsF2TemplateController.DuplicateTemplate)
		//v2.POST("/template/importTemplate", vmsF2TemplateController.ImportTemplate)

		//vms2ServerConfigController := new(controllers.GlobalConfigController)
		//v2.POST("/serverConfig/getConfig", vms2ServerConfigController.GetGlobalConfig)
		//v2.GET("/serverConfig/getEnrollUserFlag", vms2ServerConfigController.GetEnrollUserFlag)
		//v2.POST("/serverConfig/updateConfig", vms2ServerConfigController.UpdateGlobalConfig)
		//v2.POST("/serverConfig/listServerMacAddress", vms2ServerConfigController.ListServerMacAddress)

		//vms2ServerActionController := new(controllers.Vms2ServerActionController)
		//v2.POST("/serverAction/sendTestMail", vms2ServerActionController.SendTestMail)

		//vms2KioskDeviceController := new(controllers.Vms2KioskDeviceController)
		//v2.POST("/vmsKioskDevice/connectionSwitch", vms2KioskDeviceController.ChangeConnectionFlag)
		//v2.POST("/vmsKioskDevice/generateMappedDeviceCode", vms2KioskDeviceController.GenerateMappedDeviceCode)
		//v2.POST("/vmsKioskDevice/pollingCheckConnectionFlag", vms2KioskDeviceController.PollingCheckConnectionFlag)
		//v2.POST("/vmsKioskDevice/deviceConnect", vms2KioskDeviceController.ConnectVMSKioskDevice)
		//v2.POST("/vmsKioskDevice/tryToConnectVMS", vms2KioskDeviceController.TryToConnectVMS) // dedicated for KioskDevice Connect
		//v2.POST("/vmsKioskDevice/getInfoVmsKioskDevice", vms2KioskDeviceController.GetKDInfo)
		//v2.POST("/vmsKioskDevice/updateVmsKioskDevice", vms2KioskDeviceController.UpdateKDInfo)
		//v2.POST("/vmsKioskDevice/fetchAllVmsKioskDevices", vms2KioskDeviceController.FetchAllKDs)
		//v2.POST("/vmsKioskDevice/listKioskDevicesByParameter", vms2KioskDeviceController.ListKDByPData)
		//v2.POST("/vmsKioskDevice/deviceDeactivate", vms2KioskDeviceController.DeactivateVmsKioskDevice)
		//v2.POST("/vmsKioskDevice/deviceHeartBeats", vms2KioskDeviceController.HeartBeats)
		//v2.POST("/vmsKioskDevice/deviceApplyUpdate", vms2KioskDeviceController.ApplyUpdate)
		//v2.POST("/vmsKioskDevice/deviceSync", vms2KioskDeviceController.SyncKDInfo)
		//v2.POST("/vmsKioskDevice/authorizeTimeCheck", vms2KioskDeviceController.AuthTimeCheck)
		//v2.POST("/vmsKioskDevice/removeDevice", vms2KioskDeviceController.RemoveDevice)
		//v2.POST("/vmsKioskDevice/removeDeviceByWeb", vms2KioskDeviceController.RemoveDeviceByWeb)
		//v2.POST("/vmsKioskDevice/uploadDeviceLogFile", vms2KioskDeviceController.UploadLogFile)

		// ========== Device Log File ===========
		//vms2KioskDeviceLogController := new(controllers.Vms2KioskDeviceLogController)
		//v2.POST("/vmsKioskDeviceLog/updateDeviceLogFileList", vms2KioskDeviceLogController.UpdateLogFileList)
		//v2.POST("/vmsKioskDeviceLog/listLogFileByKioskUUID", vms2KioskDeviceLogController.ListLogFileByKioskUUID)
		//v2.POST("/vmsKioskDeviceLog/downloadLogFile", vms2KioskDeviceLogController.DownloadFileLog)

		//vms2KioskReportsController := new(controllers.Vms2KioskReportsController)
		//v2.POST("/vmsKioskReports/uploadKioskData", vms2KioskReportsController.UploadKR)
		//v2.POST("/vmsKioskReports/listKioskReports", vms2KioskReportsController.ListKRData)
		//v2.POST("/vmsKioskReports/listKioskReportsByParameter", vms2KioskReportsController.ListKRByPData)
		//v2.POST("/vmsKioskReports/exportKioskReportsByParameter", vms2KioskReportsController.ExportKRByPData)
		//v2.POST("/vmsKioskReports/getKioskReportsDetail", vms2KioskReportsController.GetKRDetail)
		//v2.POST("/vmsKioskReports/updateKioskReportsDetail", vms2KioskReportsController.UpdateKRDetail)
		//v2.POST("/vmsKioskReports/fillFormByReportUUID", vms2KioskReportsController.FillFormByReportUUID)

		//vms2AttendanceController := new(controllers.Vms2AttendanceController)
		//v2.POST("/vmsAttendance/listByParameter", vms2AttendanceController.ListByParameterNew)
		//v2.POST("/vmsAttendance/listByParameterWithScope", vms2AttendanceController.ListByParameterNewWithScope)
		//v2.POST("/vmsAttendance/exportByParameter", vms2AttendanceController.ExportByParameter)
		//v2.POST("/vmsAttendance/exportByParameterWithScope", vms2AttendanceController.ExportByParameterNewWithScope)
		//v2.POST("/vmsAttendance/getDetailByPersonUUID", vms2AttendanceController.GetDetailByPersonUUID)

		//vms2CompanyController := new(controllers.Vms2CompanyController)
		//v2.POST("/com/createCompany", vms2CompanyController.CreateCompany)
		//v2.POST("/com/registerCompany", vms2CompanyController.RegisterCompany)
		//v2.POST("/com/listCompaniesByParameter", vms2CompanyController.ListComByPData)
		//v2.POST("/com/getCompanyDetail", vms2CompanyController.GetCompanyDetail)
		//v2.POST("/com/updateCompany", vms2CompanyController.UpdateCom)
		//v2.POST("/com/deleteCompany", vms2CompanyController.DeleteCompany)

		//userControllerV2 := new(controllers.UserController)
		//v2.POST("/user/createWithComUUID", userControllerV2.CreateUserWithComUUID)
		//v2.POST("/user/updateUserWithComUUID", userControllerV2.UpdateUserWithComUUID)
		//v2.POST("/user/enrollUserAndCompany", userControllerV2.EnrollCompanyAndUser)

		//vms2LogController := new(controllers.Vms2LogController)
		//v2.POST("/vmsLog/listByParameter", vms2LogController.ListVmsLogByPData)
		//v2.POST("/vmsLog/listDeviceLogByParameter", vms2KioskDeviceController.ListDeviceLog)
		//v2.POST("/vmsLog/exportLogByParameter", vms2LogController.ExportVmsLogByPData)

		go func() {
			for {
				select {
				case t3 := <-ticker3mv2.C:
					_ = t3
					//logv.Info("Tick 3 min at:> ", t3)
					//vms2KioskDeviceController.CheckKioskDeviceStatus()
				case t10s := <-ticker10sv2.C:
					//logv.Info("Tick 10 s at:> ", t10s)
					_ = t10s
					targetTimeStamp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
						04, 59, 59, 0, time.UTC)
					nowTimeStamp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),
						time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
					//logv.Info("now:> ", nowTimeStamp.Unix(), " || targetTime:> ", targetTimeStamp.Unix())
					//
					if nowTimeStamp.Unix() > targetTimeStamp.Unix() {
						//logv.Info(" ============= CheckKioskReportsRetentions ===========")
						// TODO Check retentionData
						//vms2KioskReportsController.CheckKioskReportsRetentions()
						//vms2LogController.CheckVmsLogRetentions()
					}
				case t60m := <-ticker60mv2.C:
					_ = t60m
					//logv.Info(" ============= CheckLogFileStatus ===========")
					//vms2KioskDeviceLogController.CheckLogFileStatus()
					break
				case t10m := <-ticker10mv2.C:
					_ = t10m
					//vms2KioskDeviceController.NotifyDeviceSyncPerson()
					break
				}
			}
		}()
	}

	router.StaticFS("/doc/html/apidoc", http.Dir("doc/html/"))

	router.NoRoute(func(c *gin.Context) { //
		c.String(http.StatusNotFound, "Not Found")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/**
	@api {Get} /server/info get server version Info
	@apiDescription vms server Info
	@apiversion 0.0.1
	@apiGroup 000 SERVER INFO
	@apiName get server version Info

	* @apiSuccess     {String}  serverInfo serverInfo
	*
	*/
	router.GET("/server/info", func(c *gin.Context) {
		c.String(200, "%s", "ADB Bridge v0.00.01")
	})

	router.Use(static.Serve("/", static.LocalFile("web/", true)))
	// STEP 4：除了有定義路由的 API 之外，其他都會到前端框架
	// https://github.com/go-ggz/ggz/blob/master/api/index.go
	router.NoRoute(func(ctx *gin.Context) {
		file, _ := ioutil.ReadFile("web/index.html")
		etag := fmt.Sprintf("%x", md5.Sum(file)) //nolint:gosec
		ctx.Header("ETag", etag)
		ctx.Header("Cache-Control", "no-cache")

		if match := ctx.GetHeader("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				ctx.Status(http.StatusNotModified)
				//這裡若沒 return 的話，會執行到 ctx.Data
				return
			}
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:> [%+v]", oscall)

		ticker5mv1.Stop()
		ticker10sv2.Stop()
		ticker3mv2.Stop()

		cancel()
	}()

	if err := serve(ctx, router); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}

	//err := s.ListenAndServe()
	//if err != nil {
	//	logv.Error(err)
	//}

	log.Println("Exit Program")
}

func serve(ctx context.Context, router http.Handler) (err error) {

	srv := &http.Server{
		Addr:           ":" + exportPORT,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logv.Fatalf("listen:%+s\n", err)
		}
	}()

	logv.Info("server started, Listening PORT:> ", exportPORT, ", ", "MODE:> "+ models.SERVER_MODE)

	<-ctx.Done()

	log.Println("server stopped!!!")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		logv.Fatalf("server Shutdown Failed:%+s ", err)
	}

	logv.Info("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func syncTask() {
	msSQLController := new(controllers.MsSQLController)
	logv.Info(" ============= msSQLController.SyncHRDatabase ===========")
	msSQLController.SyncHRDatabase()
}
