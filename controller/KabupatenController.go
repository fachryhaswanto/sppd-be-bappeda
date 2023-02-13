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

type kabupatenController struct {
	kabupatenService service.KabupatenService
}

func NewKabupatenController(kabupatenService service.KabupatenService) *kabupatenController {
	return &kabupatenController{kabupatenService}
}

func (c *kabupatenController) GetKabupatens(cntx *gin.Context) {
	var kabupatens, err = c.kabupatenService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var kabupatensResponse []responses.KabupatenResponse

	for _, kabupaten := range kabupatens {
		var kabupatenResponse = helper.ConvertToKabupatenResponse(kabupaten)

		kabupatensResponse = append(kabupatensResponse, kabupatenResponse)
	}

	cntx.JSON(http.StatusOK, kabupatensResponse)
}

func (c *kabupatenController) GetKabupaten(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var kabupaten, err = c.kabupatenService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}

	var kabupatenResponse = helper.ConvertToKabupatenResponse(kabupaten)

	cntx.JSON(http.StatusOK, kabupatenResponse)
}

func (c *kabupatenController) CreateKabupaten(cntx *gin.Context) {
	var kabupatenRequest request.CreateKabupatenRequest

	var err = cntx.ShouldBindJSON(&kabupatenRequest)
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

	kabupaten, err := c.kabupatenService.Create(kabupatenRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"data": kabupaten,
	})
}

func (c *kabupatenController) UpdateKabupaten(cntx *gin.Context) {
	var kabupatenRequest request.UpdateKabupatenRequest

	var err = cntx.ShouldBindJSON(&kabupatenRequest)
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

	kabupaten, err := c.kabupatenService.Update(id, kabupatenRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": kabupaten,
	})
}

func (c *kabupatenController) DeleteKabupaten(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.kabupatenService.Delete(id)
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
