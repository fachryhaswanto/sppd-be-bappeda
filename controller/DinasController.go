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

type instansiController struct {
	instansiService service.InstansiService
}

func NewInstansiController(instansiService service.InstansiService) *instansiController {
	return &instansiController{instansiService}
}

func (c *instansiController) GetInstansis(cntx *gin.Context) {
	var instansis, err = c.instansiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var instansisResponse []responses.InstansiResponse

	for _, instansi := range instansis {
		var instansiResponse = helper.ConvertToInstansiResponse(instansi)
		instansisResponse = append(instansisResponse, instansiResponse)
	}

	cntx.JSON(http.StatusOK, instansisResponse)
}

func (c *instansiController) GetInstansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var instansi, err = c.instansiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var instansiResponse = helper.ConvertToInstansiResponse(instansi)

	cntx.JSON(http.StatusOK, instansiResponse)
}

func (c *instansiController) CreateInstansi(cntx *gin.Context) {
	var instansiRequest request.CreateInstansiRequest

	var err = cntx.ShouldBindJSON(&instansiRequest)
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

	instansi, err := c.instansiService.Create(instansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"data": instansi,
	})
}

func (c *instansiController) UpdateInstansi(cntx *gin.Context) {
	var instansiRequest request.UpdateInstansiRequest

	var err = cntx.ShouldBindJSON(&instansiRequest)
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

	instansi, err := c.instansiService.Update(id, instansiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": instansi,
	})
}

func (c *instansiController) DeleteInstansi(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.instansiService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}
