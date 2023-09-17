package controller

import (
	"fmt"
	"net/http"
	"sppd/helper"
	"sppd/model"
	"sppd/request"
	"sppd/responses"
	"sppd/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type rincianKwitansiController struct {
	rincianKwitansiService service.RincianKwitansiService
}

func NewRincianKwitansiController(rincianKwitansiService service.RincianKwitansiService) *rincianKwitansiController {
	return &rincianKwitansiController{rincianKwitansiService}
}

func (c *rincianKwitansiController) GetRincianKwitansis(cntx *gin.Context) {
	var kwitansiIdString = cntx.Query("kwitansiId")
	var kwitansiId, _ = strconv.Atoi(kwitansiIdString)

	var rincianKwitansis []model.RincianKwitansi
	var err error

	if kwitansiIdString == "" {
		rincianKwitansis, err = c.rincianKwitansiService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	} else {
		rincianKwitansis, err = c.rincianKwitansiService.FindByKwitansiId(kwitansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	}

	var rincianKwitansisResponse []responses.RincianKwitansiResponse

	for _, rincianKwitansi := range rincianKwitansis {
		var rincianKwitansiResponse = helper.ConvertToRincianKwitansiResponse(rincianKwitansi)
		rincianKwitansisResponse = append(rincianKwitansisResponse, rincianKwitansiResponse)
	}

	cntx.JSON(http.StatusOK, rincianKwitansisResponse)
}

func (c *rincianKwitansiController) GetRincianKwitansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var rincianKwitansi, err = c.rincianKwitansiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var rincianKwitansiResponse = helper.ConvertToRincianKwitansiResponse(rincianKwitansi)

	cntx.JSON(http.StatusOK, rincianKwitansiResponse)
}

func (c *rincianKwitansiController) GetRincianKwitansiBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var rincianKwitansis, err = c.rincianKwitansiService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rincianKwitansisResponse []responses.RincianKwitansiResponse

	for _, rincianKwitansi := range rincianKwitansis {
		var rincianKwitansiResponse = helper.ConvertToRincianKwitansiResponse(rincianKwitansi)
		rincianKwitansisResponse = append(rincianKwitansisResponse, rincianKwitansiResponse)
	}

	cntx.JSON(http.StatusOK, rincianKwitansisResponse)
}

func (c *rincianKwitansiController) CountDataBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var count, err = c.rincianKwitansiService.CountDataBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, count)
}

func (c *rincianKwitansiController) CreateRincianKwitansi(cntx *gin.Context) {
	var rincianKwitansiRequest []request.CreateRincianKwitansiRequest

	var err = cntx.ShouldBindJSON(&rincianKwitansiRequest)
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

	rincianKwitansis, err := c.rincianKwitansiService.Create(rincianKwitansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rincianKwitansisResponse []responses.RincianKwitansiResponse

	for _, rincianKwitansi := range rincianKwitansis {
		var rincianKwitansiResponse = helper.ConvertToRincianKwitansiResponse(rincianKwitansi)
		rincianKwitansisResponse = append(rincianKwitansisResponse, rincianKwitansiResponse)
	}

	cntx.JSON(http.StatusCreated, rincianKwitansisResponse)
}

func (c *rincianKwitansiController) UpdateRincianKwitansi(cntx *gin.Context) {
	var rincianKwitansiRequest request.UpdateRincianKwitansiRequest

	var err = cntx.ShouldBindJSON(&rincianKwitansiRequest)
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

	rincianKwitansi, err := c.rincianKwitansiService.Update(id, rincianKwitansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var rincianKwitansiResponse = helper.ConvertToRincianKwitansiResponse(rincianKwitansi)

	cntx.JSON(http.StatusOK, rincianKwitansiResponse)
}

func (c *rincianKwitansiController) DeleteRincianKwitansi(cntx *gin.Context) {
	var kwitansiIdString = cntx.Param("kwitansiId")
	var kwitansiId, _ = strconv.Atoi(kwitansiIdString)

	_, err := c.rincianKwitansiService.Delete(kwitansiId)
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
