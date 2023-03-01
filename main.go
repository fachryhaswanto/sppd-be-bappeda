package main

import (
	"log"
	"os"
	"sppd/config"
	"sppd/controller"
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

	var sppdRepository = repository.NewSppdRepository(config.DB)
	var sppdService = service.NewSppdService(sppdRepository)
	var sppdController = controller.NewSppdController(sppdService)

	var sptRepository = repository.NewSptRepository(config.DB)
	var sptService = service.NewSptService(sptRepository)
	var sptController = controller.NewSptController(sptService, sppdService)

	var server = gin.Default()

	server.MaxMultipartMemory = 5 << 20

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://" + os.Getenv("APP_ALLOW_ORIGINS"), "http://localhost:3000"},
		AllowMethods: []string{"POST", "PUT", "DELETE", "GET", "PATCH"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers"},
	}))

	server.GET("/user/:username/:password", userController.GetUser)

	server.GET("/kabupaten", kabupatenController.GetKabupatens)
	server.GET("/kabupaten/:id", kabupatenController.GetKabupaten)
	server.POST("/kabupaten", kabupatenController.CreateKabupaten)
	server.PATCH("/kabupaten/:id", kabupatenController.UpdateKabupaten)
	server.DELETE("/kabupaten/:id", kabupatenController.DeleteKabupaten)

	server.GET("/instansi", instansiController.GetInstansis)
	server.GET("/instansi/:id", instansiController.GetInstansi)
	server.POST("/instansi", instansiController.CreateInstansi)
	server.PATCH("/instansi/:id", instansiController.UpdateInstansi)
	server.DELETE("/instansi/:id", instansiController.DeleteInstansi)

	server.GET("/bidang", bidangController.GetBidangs)
	server.GET("/bidang/:id", bidangController.GetBidang)
	server.POST("/bidang", bidangController.CreateBidang)
	server.PATCH("/bidang/:id", bidangController.UpdateBidang)
	server.DELETE("/bidang/:id", bidangController.DeleteBidang)

	server.GET("/pejabat", pejabatController.GetPejabats)
	server.GET("/pejabat/:id", pejabatController.GetPejabat)
	server.GET("/pejabat/nama/:name", pejabatController.GetPejabatByName)
	server.POST("/pejabat", pejabatController.CreatePejabat)
	server.PATCH("/pejabat/:id", pejabatController.UpdatePejabat)
	server.DELETE("/pejabat/:id", pejabatController.DeletePejabat)

	server.GET("/pegawai", pegawaiController.GetPegawais)
	server.GET("/pegawai/:id", pegawaiController.GetPegawai)
	server.GET("/pegawai/nama/:name", pegawaiController.GetPegawaiByName)
	server.POST("/pegawai", pegawaiController.CreatePegawai)
	server.PATCH("/pegawai/:id", pegawaiController.UpdatePegawai)
	server.DELETE("/pegawai/:id", pegawaiController.DeletePegawai)

	server.GET("/program", programController.GetPrograms)
	server.GET("/program/:id", programController.GetProgram)
	server.POST("/program", programController.CreateProgram)
	server.PATCH("/program/:id", programController.UpdateProgram)
	server.DELETE("/program/:id", programController.DeleteProgram)

	server.GET("/kegiatan", kegiatanController.GetKegiatans)
	server.GET("/kegiatan/:id", kegiatanController.GetKegiatan)
	server.POST("/kegiatan", kegiatanController.CreateKegiatan)
	server.PATCH("/kegiatan/:id", kegiatanController.UpdateKegiatan)
	server.DELETE("/kegiata/:id", kegiatanController.DeleteKegiatan)

	server.GET("/subkegiatan", subKegiatanController.GetSubKegiatans)
	server.GET("/subkegiatan/:id", subKegiatanController.GetSubKegiatan)
	server.POST("/subkegiatan", subKegiatanController.CreateSubKegiatan)
	server.PATCH("/subkegiatan/:id", subKegiatanController.UpdateSubKegiatan)
	server.DELETE("/subkegiatan/:id", subKegiatanController.DeleteSubKegiatan)

	server.GET("/spt", sptController.GetSpts)
	server.GET("/spt/:id", sptController.GetSpt)
	server.GET("/spt/ditugaskan", sptController.GetAllDitugaskan)
	server.POST("/spt", sptController.CreateSpt)
	server.PATCH("/spt/:id", sptController.UpdateSpt)
	server.DELETE("/spt/:id", sptController.DeleteSpt)

	server.GET("/sppd/", sppdController.GetSppds)
	server.GET("/join", sppdController.GetJoin)
	server.GET("/sppd/:id", sppdController.GetSppd)
	server.PATCH("/sppd/:id", sppdController.UpdateSppd)
	server.DELETE("/sppd/:id", sppdController.DeleteSppd)

	server.POST("/upload/", controller.UploadFile)

	server.GET("/getfile/:filename", controller.GetFile)
	// http.HandleFunc("/getfile/", controller.GetFile)

	server.Run(":" + appConfig.AppPort)

}
