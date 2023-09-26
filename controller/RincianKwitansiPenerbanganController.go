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

type rincianKwitansiPenerbanganController struct {
	rincianKwitansiPenerbanganService service.RincianKwitansiPenerbanganService
}

func NewRincianKwitansiPenerbanganController(rincianKwitansiPenerbanganService service.RincianKwitansiPenerbanganService) *rincianKwitansiPenerbanganController {
	return &rincianKwitansiPenerbanganController{rincianKwitansiPenerbanganService}
}

func (c *rincianKwitansiPenerbanganController) GetRincianKwitansiPenerbangans(cntx *gin.Context) {
	var kwitansiIdString = cntx.Query("kwitansiId")
	var kwitansiId, _ = strconv.Atoi(kwitansiIdString)

	var rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan
	var err error

	if kwitansiIdString == "" {
		rincianKwitansiPenerbangans, err = c.rincianKwitansiPenerbanganService.FindAll()
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	} else {
		rincianKwitansiPenerbangans, err = c.rincianKwitansiPenerbanganService.FindByKwitansiId(kwitansiId)
		if err != nil {
			cntx.JSON(http.StatusBadRequest, gin.H{
				"error": cntx.Error(err),
			})
			return
		}
	}

	var rincianKwitansiPenerbangansResponse []responses.RincianKwitansiPenerbanganResponse

	for _, rincianKwitansiPenerbangan := range rincianKwitansiPenerbangans {
		var tempResponse = helper.ConvertToRincianKwitansiPenerbanganResponse(rincianKwitansiPenerbangan)
		rincianKwitansiPenerbangansResponse = append(rincianKwitansiPenerbangansResponse, tempResponse)
	}

	cntx.JSON(http.StatusOK, rincianKwitansiPenerbangansResponse)
}

func (c *rincianKwitansiPenerbanganController) GetRincianKwitansiPenerbangan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var rincianKwitansiPenerbangan, err = c.rincianKwitansiPenerbanganService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var rincianKwitansiPenerbanganResponse = helper.ConvertToRincianKwitansiPenerbanganResponse(rincianKwitansiPenerbangan)

	cntx.JSON(http.StatusOK, rincianKwitansiPenerbanganResponse)
}

func (c *rincianKwitansiPenerbanganController) GetRincianKwitansiPenerbanganBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var rincianKwitansiPenerbangans, err = c.rincianKwitansiPenerbanganService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rincianKwitansiPenerbangansResponse []responses.RincianKwitansiPenerbanganResponse

	for _, rincianKwitansiPenerbangan := range rincianKwitansiPenerbangans {
		var tempResponse = helper.ConvertToRincianKwitansiPenerbanganResponse(rincianKwitansiPenerbangan)
		rincianKwitansiPenerbangansResponse = append(rincianKwitansiPenerbangansResponse, tempResponse)
	}

	cntx.JSON(http.StatusOK, rincianKwitansiPenerbangansResponse)
}

func (c *rincianKwitansiPenerbanganController) CountDataBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var count, err = c.rincianKwitansiPenerbanganService.CountDataBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, count)
}

func (c *rincianKwitansiPenerbanganController) CreateRincianKwitansiPenerbangan(cntx *gin.Context) {
	var rincianKwitansiPenerbanganRequest []request.CreateRincianKwitansiPenerbanganRequest

	var err = cntx.ShouldBindJSON(&rincianKwitansiPenerbanganRequest)
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

	rincianKwitansiPenerbangans, err := c.rincianKwitansiPenerbanganService.Create(rincianKwitansiPenerbanganRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rincianKwitansiPenerbangansResponse []responses.RincianKwitansiPenerbanganResponse

	for _, rincianKwitansiPenerbangan := range rincianKwitansiPenerbangans {
		var tempResponse = helper.ConvertToRincianKwitansiPenerbanganResponse(rincianKwitansiPenerbangan)
		rincianKwitansiPenerbangansResponse = append(rincianKwitansiPenerbangansResponse, tempResponse)
	}

	cntx.JSON(http.StatusCreated, rincianKwitansiPenerbangansResponse)
}

func (c *rincianKwitansiPenerbanganController) UpdateRincianKwitansiPenerbangan(cntx *gin.Context) {
	var rincianKwitansiPenerbanganRequest request.UpdateRincianKwitansiPenerbanganRequest

	var err = cntx.ShouldBindJSON(&rincianKwitansiPenerbanganRequest)
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

	rincianKwitansiPenerbangan, err := c.rincianKwitansiPenerbanganService.Update(id, rincianKwitansiPenerbanganRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var rincianKwitansiPenerbanganResponse = helper.ConvertToRincianKwitansiPenerbanganResponse(rincianKwitansiPenerbangan)

	cntx.JSON(http.StatusOK, rincianKwitansiPenerbanganResponse)
}

func (c *rincianKwitansiPenerbanganController) DeleteRincianKwitansiPenerbangan(cntx *gin.Context) {
	var kwitansiIdString = cntx.Param("kwitansiId")
	var kwitansiId, _ = strconv.Atoi(kwitansiIdString)

	_, err := c.rincianKwitansiPenerbanganService.Delete(kwitansiId)
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
