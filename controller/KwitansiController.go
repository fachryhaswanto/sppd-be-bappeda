package controller

import (
	"fmt"
	"net/http"
	"sppd/helper"
	"sppd/request"
	"sppd/responses"
	"sppd/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type kwitansiController struct {
	kwitansiService                   service.KwitansiService
	rincianKwitansiService            service.RincianKwitansiService
	rincianKwitansiPenerbanganService service.RincianKwitansiPenerbanganService
}

func NewKwitansiController(kwitansiService service.KwitansiService, rincianKwitansiService service.RincianKwitansiService, rincianKwitansiPenerbanganService service.RincianKwitansiPenerbanganService) *kwitansiController {
	return &kwitansiController{kwitansiService, rincianKwitansiService, rincianKwitansiPenerbanganService}
}

func (c *kwitansiController) GetKwitansis(cntx *gin.Context) {
	var kwitansis, err = c.kwitansiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kwitansisResponse []responses.KwitansiResponse

	for _, kwitansi := range kwitansis {
		var kwitansiResponse = helper.ConvertToKwitansiResponse(kwitansi)
		kwitansisResponse = append(kwitansisResponse, kwitansiResponse)
	}

	cntx.JSON(http.StatusOK, kwitansisResponse)
}

func (c *kwitansiController) GetKwitansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var kwitansi, err = c.kwitansiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var kwitansiResponse = helper.ConvertToKwitansiResponse(kwitansi)

	cntx.JSON(http.StatusOK, kwitansiResponse)
}

func (c *kwitansiController) GetKwitansiBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var kwitansis, err = c.kwitansiService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kwitansisResponses []responses.KwitansiResponse

	for _, kwitansi := range kwitansis {
		var kwitansiResponse = helper.ConvertToKwitansiResponse(kwitansi)
		kwitansisResponses = append(kwitansisResponses, kwitansiResponse)
	}

	cntx.JSON(http.StatusOK, kwitansisResponses)
}

func (c *kwitansiController) CountDataBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var count, err = c.kwitansiService.CountDataBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, count)
}

func (c *kwitansiController) SumTotalBayar(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var realisasi = c.kwitansiService.SumTotalBayar(whereClauseInterface)

	cntx.JSON(http.StatusOK, realisasi)
}

func (c *kwitansiController) CreateKwitansi(cntx *gin.Context) {
	var kwitansiRequest request.CreateKwitansiRequest

	var err = cntx.ShouldBindJSON(&kwitansiRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	kwitansi, err := c.kwitansiService.Create(kwitansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kwitansiResponse = helper.ConvertToKwitansiResponse(kwitansi)

	cntx.JSON(http.StatusCreated, kwitansiResponse)
}

func (c *kwitansiController) UpdateKwitansi(cntx *gin.Context) {
	var kwitansiRequest request.UpdateKwitansiRequest

	var err = cntx.ShouldBindJSON(&kwitansiRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	kwitansi, err := c.kwitansiService.Update(id, kwitansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var kwitansiResponse = helper.ConvertToKwitansiResponse(kwitansi)

	cntx.JSON(http.StatusOK, kwitansiResponse)
}

func (c *kwitansiController) UpdateTotalBayarKwitansi(cntx *gin.Context) {
	var kwitansiRequest request.UpdateTotalBayarRequest

	var err = cntx.ShouldBindJSON(&kwitansiRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	kwitansi, err := c.kwitansiService.UpdateTotalBayar(id, kwitansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var kwitansiResponse = helper.ConvertToKwitansiResponse(kwitansi)

	cntx.JSON(http.StatusOK, kwitansiResponse)
}

func (c *kwitansiController) DeleteKwitansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.kwitansiService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "data berhasil dihapus",
	})
}

func (c *kwitansiController) GenerateLaporan(cntx *gin.Context) {
	var pegawaiIdString = cntx.Query("pegawaiId")
	var pegawaiId, _ = strconv.Atoi(pegawaiIdString)

	if pegawaiIdString == "" {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing parameter pegawaiId",
		})
		return
	}

	var whereClause = map[string]interface{}{}
	whereClause["pegawaiId"] = pegawaiId

	var kwitansis, err = c.kwitansiService.FindBySearch(whereClause)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var sumTotalBayar = c.kwitansiService.SumTotalBayar(whereClause)

	for i := 0; i < len(kwitansis); i++ {
		kwitansis[i].Sppd.Spt.Tanggal_Spt = helper.FormatDate(kwitansis[i].Sppd.Spt.Tanggal_Spt)
		kwitansis[i].Sppd.Tanggal_Sppd = helper.FormatDate(kwitansis[i].Sppd.Tanggal_Sppd)
		kwitansis[i].Sppd.Spt.Tanggal_Berangkat = helper.FormatDate(kwitansis[i].Sppd.Spt.Tanggal_Berangkat)
		kwitansis[i].Sppd.Spt.Tanggal_Kembali = helper.FormatDate(kwitansis[i].Sppd.Spt.Tanggal_Kembali)
	}

	var laporanResponses []responses.LaporanResponse

	for _, kwitansi := range kwitansis {
		var tempResponse = helper.ConvertToLaporanResponse(kwitansi)
		laporanResponses = append(laporanResponses, tempResponse)
	}

	for i := 0; i < len(laporanResponses); i++ {
		laporanResponses[i].SumTotalBayar = helper.AddSeparatorInt64(sumTotalBayar)
		laporanResponses[i].TotalBayar = helper.AddSeparator(kwitansis[i].TotalBayar)
		var nomor = strconv.Itoa(i + 1)
		laporanResponses[i].Nomor = nomor

		var whereClause = make(map[string]interface{})

		//uang harian
		whereClause["kwitansiId"] = kwitansis[i].Id
		whereClause["jenis"] = "Uang Harian"
		var rincianKwitansisUangHarian, err = c.rincianKwitansiService.FindBySearch(whereClause)
		laporanResponses[i].RincianKwitansiUangHarian = make([]responses.RincianKwitansi, len(rincianKwitansisUangHarian))

		//uang representatif
		whereClause["jenis"] = "Uang Representatif"
		rincianKwitansiUangRepresentatif, err := c.rincianKwitansiService.FindBySearch(whereClause)
		laporanResponses[i].RincianKwitansiUangRepresentatif = make([]responses.RincianKwitansi, len(rincianKwitansiUangRepresentatif))

		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		//biaya hotel
		whereClause["jenis"] = "Biaya Hotel"
		rincianKwitansiBiayaHotel, err := c.rincianKwitansiService.FindBySearch(whereClause)
		laporanResponses[i].RincianKwitansiBiayaHotel = make([]responses.RincianKwitansi, len(rincianKwitansiBiayaHotel))

		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		//biaya tiket
		whereClause["jenis"] = "Biaya Tiket"
		rincianKwitansiBiayaTiket, err := c.rincianKwitansiService.FindBySearch(whereClause)
		laporanResponses[i].RincianKwitansiBiayaTiket = make([]responses.RincianKwitansi, len(rincianKwitansiBiayaTiket))

		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		//transport bandara
		whereClause["jenis"] = "Transport Bandara"
		rincianKwitansiTransportBandara, err := c.rincianKwitansiService.FindBySearch(whereClause)
		laporanResponses[i].RincianKwitansiTransportBandara = make([]responses.RincianKwitansi, len(rincianKwitansiTransportBandara))

		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		for j := 0; j < len(rincianKwitansisUangHarian); j++ {
			laporanResponses[i].RincianKwitansiUangHarian[j].Id = rincianKwitansisUangHarian[j].Id
			laporanResponses[i].RincianKwitansiUangHarian[j].KwitansiId = rincianKwitansisUangHarian[j].KwitansiId
			laporanResponses[i].RincianKwitansiUangHarian[j].Jenis = rincianKwitansisUangHarian[j].Jenis
			laporanResponses[i].RincianKwitansiUangHarian[j].HasilBayar = helper.AddSeparator(rincianKwitansisUangHarian[j].HasilBayar)
		}

		for j := 0; j < len(rincianKwitansiUangRepresentatif); j++ {
			laporanResponses[i].RincianKwitansiUangRepresentatif[j].Id = rincianKwitansiUangRepresentatif[j].Id
			laporanResponses[i].RincianKwitansiUangRepresentatif[j].KwitansiId = rincianKwitansiUangRepresentatif[j].KwitansiId
			laporanResponses[i].RincianKwitansiUangRepresentatif[j].Jenis = rincianKwitansiUangRepresentatif[j].Jenis
			laporanResponses[i].RincianKwitansiUangRepresentatif[j].HasilBayar = helper.AddSeparator(rincianKwitansiUangRepresentatif[j].HasilBayar)
		}

		for j := 0; j < len(rincianKwitansiBiayaHotel); j++ {
			laporanResponses[i].RincianKwitansiBiayaHotel[j].Id = rincianKwitansiBiayaHotel[j].Id
			laporanResponses[i].RincianKwitansiBiayaHotel[j].KwitansiId = rincianKwitansiBiayaHotel[j].KwitansiId
			laporanResponses[i].RincianKwitansiBiayaHotel[j].Jenis = rincianKwitansiBiayaHotel[j].Jenis
			laporanResponses[i].RincianKwitansiBiayaHotel[j].HasilBayar = helper.AddSeparator(rincianKwitansiBiayaHotel[j].HasilBayar)
		}

		for j := 0; j < len(rincianKwitansiBiayaTiket); j++ {
			laporanResponses[i].RincianKwitansiBiayaTiket[j].Id = rincianKwitansiBiayaTiket[j].Id
			laporanResponses[i].RincianKwitansiBiayaTiket[j].KwitansiId = rincianKwitansiBiayaTiket[j].KwitansiId
			laporanResponses[i].RincianKwitansiBiayaTiket[j].Jenis = rincianKwitansiBiayaTiket[j].Jenis
			laporanResponses[i].RincianKwitansiBiayaTiket[j].HasilBayar = helper.AddSeparator(rincianKwitansiBiayaTiket[j].HasilBayar)
		}

		for j := 0; j < len(rincianKwitansiTransportBandara); j++ {
			laporanResponses[i].RincianKwitansiTransportBandara[j].Id = rincianKwitansiTransportBandara[j].Id
			laporanResponses[i].RincianKwitansiTransportBandara[j].KwitansiId = rincianKwitansiTransportBandara[j].KwitansiId
			laporanResponses[i].RincianKwitansiTransportBandara[j].Jenis = rincianKwitansiTransportBandara[j].Jenis
			laporanResponses[i].RincianKwitansiTransportBandara[j].HasilBayar = helper.AddSeparator(rincianKwitansiTransportBandara[j].HasilBayar)
		}

		//rincian kwitansi penerbangan
		rincianKwitansiPenerbangan, err := c.rincianKwitansiPenerbanganService.FindByKwitansiId(kwitansis[i].Id)
		laporanResponses[i].RincianKwitansiPenerbangan = make([]responses.RincianKwitansiPenerbangan, len(rincianKwitansiPenerbangan))
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}

		for j := 0; j < len(rincianKwitansiPenerbangan); j++ {
			laporanResponses[i].RincianKwitansiPenerbangan[j].Id = rincianKwitansiPenerbangan[j].Id
			laporanResponses[i].RincianKwitansiPenerbangan[j].KwitansiId = rincianKwitansiPenerbangan[j].KwitansiId
			laporanResponses[i].RincianKwitansiPenerbangan[j].NamaMaskapai = rincianKwitansiPenerbangan[j].NamaMaskapai
			laporanResponses[i].RincianKwitansiPenerbangan[j].NomorTiket = rincianKwitansiPenerbangan[j].NomorTiket
			laporanResponses[i].RincianKwitansiPenerbangan[j].KodeBooking = rincianKwitansiPenerbangan[j].KodeBooking
		}
	}

	cntx.JSON(http.StatusOK, laporanResponses)

	// cntx.JSON(http.StatusOK, gin.H{
	// 	"dataLaporan": laporanResponses,
	// })
}
