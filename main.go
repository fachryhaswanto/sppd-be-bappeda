package main

import (
	"log"
	"os"
	"sppd/config"
	"sppd/controller"
	"sppd/middleware"
	"sppd/repository"
	"sppd/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var appConfig = config.AppConfig{}

	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Terdapat kesalahan memuat file .env")
	}

	appConfig.DatabaseHost = os.Getenv("APP_DATABASE_HOST")
	// appConfig.DatabaseURL = os.Getenv("DATABASE_URL")
	appConfig.DatabasePassword = os.Getenv("APP_DATABASE_PASSWORD")
	appConfig.DatabasePort = os.Getenv("APP_DATABASE_PORT")
	appConfig.DatabaseName = os.Getenv("APP_DATABASE_NAME")
	appConfig.AllowOrigins = os.Getenv("APP_ALLOW_ORIGIN")
	appConfig.AppPort = os.Getenv("APP_PORT")

	config.ConnectDB(appConfig)

	var userRepository = repository.NewUserRepository(config.DB)
	var userService = service.NewUserService(userRepository)
	var userController = controller.NewUserController(userService)

	var kabupatenRepository = repository.NewKabupatenRepository(config.DB)
	var kabupatenService = service.NewKabupatenService(kabupatenRepository)
	var kabupatenController = controller.NewKabupatenController(kabupatenService)

	var instansiRepository = repository.NewInstansiRepository(config.DB)
	var instansiService = service.NewInstansiService(instansiRepository)
	var instansiController = controller.NewInstansiController(instansiService)

	var bidangRepository = repository.NewBidangRepository(config.DB)
	var bidangService = service.NewBidangService(bidangRepository)
	var bidangController = controller.NewBidangController(bidangService)

	var pejabatRepository = repository.NewPejabatRepository(config.DB)
	var pejabatService = service.NewPejabatService(pejabatRepository)
	var pejabatController = controller.NewPejabatController(pejabatService)

	var pegawaiRepository = repository.NewPegawaiRepository(config.DB)
	var pegawaiService = service.NewPegawaiService(pegawaiRepository)
	var pegawaiController = controller.NewPegawaiController(pegawaiService)

	var programRepository = repository.NewProgramRepository(config.DB)
	var programService = service.NewProgramService(programRepository)
	var programController = controller.NewProgramController(programService)

	var kegiatanRepository = repository.NewKegiatanRepository(config.DB)
	var kegiatanService = service.NewKegiatanService(kegiatanRepository)
	var kegiatanController = controller.NewKegiatanController(kegiatanService)

	var subKegiatanRepository = repository.NewSubKegiatanRepository(config.DB)
	var subKegiatanService = service.NewSubKegiatanService(subKegiatanRepository)
	var subKegiatanController = controller.NewSubKegiatanController(subKegiatanService)

	var rekeningRepository = repository.NewRekeningRepository(config.DB)
	var rekeningService = service.NewRekeningService(rekeningRepository)
	var rekeningController = controller.NewRekeningController(rekeningService)

	var sptRepository = repository.NewSptRepository(config.DB)
	var sptService = service.NewSptService(sptRepository)
	var sptController = controller.NewSptController(sptService)

	var sppdRepository = repository.NewSppdRepository(config.DB)
	var sppdService = service.NewSppdService(sppdRepository)
	var sppdController = controller.NewSppdController(sppdService, sptService)

	var dataDitugaskanRepository = repository.NewDataDitugaskanRepository(config.DB)
	var dataDitugaskanService = service.NewDataDitugaskanService(dataDitugaskanRepository, sptRepository)
	var dataDitugaskanController = controller.NewDataDitugaskanController(dataDitugaskanService)

	var dataPengikutRepository = repository.NewDataPengikutRepository(config.DB)
	var dataPengikutService = service.NewDataPengikutService(dataPengikutRepository)
	var dataPengikutController = controller.NewDataPengikutController(dataPengikutService)

	var server = gin.Default()

	server.MaxMultipartMemory = 5 << 20

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "PUT", "DELETE", "GET", "PATCH"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "Accept", "Origin", "cookie-sppd", "Access-Control-Allow-Credentials"},
	}))

	server.GET("/auth/session", middleware.CheckAuth, userController.Validate)
	server.POST("/auth/login", userController.LoginUser)
	server.POST("/auth/logout", userController.LogoutUser)

	server.GET("/kabupaten", middleware.CheckAuth, kabupatenController.GetKabupatens)
	server.GET("/kabupaten/:id", middleware.CheckAuth, kabupatenController.GetKabupaten)
	server.POST("/kabupaten", middleware.CheckAuth, kabupatenController.CreateKabupaten)
	server.PATCH("/kabupaten/:id", middleware.CheckAuth, kabupatenController.UpdateKabupaten)
	server.DELETE("/kabupaten/:id", middleware.CheckAuth, kabupatenController.DeleteKabupaten)

	server.GET("/instansi", middleware.CheckAuth, instansiController.GetInstansis)
	server.GET("/instansi/:id", middleware.CheckAuth, instansiController.GetInstansi)
	server.POST("/instansi", middleware.CheckAuth, instansiController.CreateInstansi)
	server.PATCH("/instansi/:id", middleware.CheckAuth, instansiController.UpdateInstansi)
	server.DELETE("/instansi/:id", middleware.CheckAuth, instansiController.DeleteInstansi)

	server.GET("/bidang", middleware.CheckAuth, bidangController.GetBidangs)
	server.GET("/bidang/:id", middleware.CheckAuth, bidangController.GetBidang)
	server.POST("/bidang", middleware.CheckAuth, bidangController.CreateBidang)
	server.PATCH("/bidang/:id", middleware.CheckAuth, bidangController.UpdateBidang)
	server.DELETE("/bidang/:id", middleware.CheckAuth, bidangController.DeleteBidang)

	server.GET("/pejabat", middleware.CheckAuth, pejabatController.GetPejabats)
	server.GET("/pejabat/:id", middleware.CheckAuth, pejabatController.GetPejabat)
	server.GET("/pejabat/nama/:name", middleware.CheckAuth, pejabatController.GetPejabatByName)
	server.GET("/pejabat/search", middleware.CheckAuth, pejabatController.GetPejabatBySearch)
	server.POST("/pejabat", middleware.CheckAuth, pejabatController.CreatePejabat)
	server.PATCH("/pejabat/:id", middleware.CheckAuth, pejabatController.UpdatePejabat)
	server.DELETE("/pejabat/:id", middleware.CheckAuth, pejabatController.DeletePejabat)

	server.GET("/pegawai", middleware.CheckAuth, pegawaiController.GetPegawais)
	server.GET("/pegawai/:id", middleware.CheckAuth, pegawaiController.GetPegawai)
	server.GET("/pegawai/nama/:name", middleware.CheckAuth, pegawaiController.GetPegawaiByName)
	server.GET("/pegawai/search", middleware.CheckAuth, pegawaiController.GetPegawaisBySearch)
	server.POST("/pegawai", middleware.CheckAuth, pegawaiController.CreatePegawai)
	server.PATCH("/pegawai/:id", middleware.CheckAuth, pegawaiController.UpdatePegawai)
	server.PATCH("/pegawai/batch", middleware.CheckAuth, pegawaiController.UpdateBatchPegawaiStatusPerjalanan)
	server.DELETE("/pegawai/:id", middleware.CheckAuth, pegawaiController.DeletePegawai)

	server.GET("/program", middleware.CheckAuth, programController.GetPrograms)
	server.GET("/program/:id", middleware.CheckAuth, programController.GetProgram)
	server.POST("/program", middleware.CheckAuth, programController.CreateProgram)
	server.PATCH("/program/:id", middleware.CheckAuth, programController.UpdateProgram)
	server.DELETE("/program/:id", middleware.CheckAuth, programController.DeleteProgram)

	server.GET("/kegiatan", middleware.CheckAuth, kegiatanController.GetKegiatans)
	server.GET("/kegiatan/:id", middleware.CheckAuth, kegiatanController.GetKegiatan)
	server.POST("/kegiatan", middleware.CheckAuth, kegiatanController.CreateKegiatan)
	server.PATCH("/kegiatan/:id", middleware.CheckAuth, kegiatanController.UpdateKegiatan)
	server.DELETE("/kegiatan/:id", middleware.CheckAuth, kegiatanController.DeleteKegiatan)

	server.GET("/subkegiatan", middleware.CheckAuth, subKegiatanController.GetSubKegiatans)
	server.GET("/subkegiatan/:id", middleware.CheckAuth, subKegiatanController.GetSubKegiatan)
	server.POST("/subkegiatan", middleware.CheckAuth, subKegiatanController.CreateSubKegiatan)
	server.PATCH("/subkegiatan/:id", middleware.CheckAuth, subKegiatanController.UpdateSubKegiatan)
	server.DELETE("/subkegiatan/:id", middleware.CheckAuth, subKegiatanController.DeleteSubKegiatan)

	server.GET("/rekening", middleware.CheckAuth, rekeningController.GetRekenings)
	server.GET("/rekening/:id", middleware.CheckAuth, rekeningController.GetRekening)
	server.POST("/rekening", middleware.CheckAuth, rekeningController.CreateRekening)
	server.PATCH("/rekening/:id", middleware.CheckAuth, rekeningController.UpdateRekening)
	server.DELETE("/rekening/:id", middleware.CheckAuth, rekeningController.DeleteRekening)

	server.GET("/spt", middleware.CheckAuth, sptController.GetSpts)
	server.GET("/spt/:id", middleware.CheckAuth, sptController.GetSpt)
	server.GET("/spt/ditugaskan", middleware.CheckAuth, sptController.GetAllDitugaskan)
	server.GET("/spt/search", middleware.CheckAuth, sptController.GetSptsBySearch)
	server.POST("/spt", middleware.CheckAuth, sptController.CreateSpt)
	server.PATCH("/spt/:id", middleware.CheckAuth, sptController.UpdateSpt)
	server.PATCH("/spt/updatestatus/:id/:value", middleware.CheckAuth, sptController.UpdateStatusSppd)
	server.DELETE("/spt/:id", middleware.CheckAuth, sptController.DeleteSpt)

	server.GET("/sppd", middleware.CheckAuth, sppdController.GetSppds)
	server.GET("/sppd/:id", middleware.CheckAuth, sppdController.GetSppd)
	server.GET("/sppd/search", middleware.CheckAuth, sppdController.GetSppdBySearch)
	server.GET("/sppd/count", middleware.CheckAuth, sppdController.CountDataBySptId)
	server.POST("/sppd", middleware.CheckAuth, sppdController.CreateSppd)
	server.PATCH("/sppd/:id", middleware.CheckAuth, sppdController.UpdateSppd)
	server.PATCH("/sppd/updatebysptid/:sptid", middleware.CheckAuth, sppdController.UpdateSppdBySptId)
	server.DELETE("/sppd/:id", middleware.CheckAuth, sppdController.DeleteSppd)
	server.DELETE("/sppd/sptid/:sptid", middleware.CheckAuth, sppdController.DeleteSppdBySptId)

	server.GET("/dataditugaskan", middleware.CheckAuth, dataDitugaskanController.GetDataDitugaskans)
	server.GET("/dataditugaskan/:id", middleware.CheckAuth, dataDitugaskanController.GetDataDitugaskan)
	server.GET("/dataditugaskan/search", middleware.CheckAuth, dataDitugaskanController.GetDataDitugaskansBySearch)
	server.GET("/dataditugaskan/count", middleware.CheckAuth, dataDitugaskanController.CountDataBySptId)
	server.POST("/dataditugaskan", middleware.CheckAuth, dataDitugaskanController.CreateDataDitugaskan)
	server.PATCH("/dataditugaskan/updatestatus/:sptid/:value", middleware.CheckAuth, dataDitugaskanController.UpdateStatusSppd)
	server.DELETE("/dataditugaskan/:sptid", middleware.CheckAuth, dataDitugaskanController.DeleteDataDitugaskan)

	server.GET("/datapengikut", middleware.CheckAuth, dataPengikutController.GetDataPengikuts)
	server.GET("/datapengikut/:id", middleware.CheckAuth, dataPengikutController.GetDataPengikut)
	server.GET("/datapengikut/search", middleware.CheckAuth, dataPengikutController.GetDataPengikutBySearch)
	server.POST("/datapengikut", middleware.CheckAuth, dataPengikutController.CreateDataPengikut)
	server.DELETE("/datapengikut/:sptid", middleware.CheckAuth, dataPengikutController.DeleteDataPengikut)

	server.POST("/upload/", controller.UploadFile)

	server.GET("/getfile/:filename", controller.GetFile)
	// http.HandleFunc("/getfile/", controller.GetFile)

	server.Run(":" + appConfig.AppPort)

}
