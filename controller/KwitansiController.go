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
	kwitansiService service.KwitansiService
}

func NewKwitansiController(kwitansiService service.KwitansiService) *kwitansiController {
	return &kwitansiController{kwitansiService}
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
