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

type dataPengikutController struct {
	dataPengikutService service.DataPengikutService
}

func NewDataPengikutController(dataPengikutService service.DataPengikutService) *dataPengikutController {
	return &dataPengikutController{dataPengikutService}
}

func (c *dataPengikutController) GetDataPengikuts(cntx *gin.Context) {
	var dataPengikuts, err = c.dataPengikutService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var dataPengikutsResponse []responses.DataPengikutResponse

	for _, dataPengikut := range dataPengikuts {
		var dataPengikutResponse = helper.ConvertToDataPengikutResponse(dataPengikut)
		dataPengikutsResponse = append(dataPengikutsResponse, dataPengikutResponse)
	}

	cntx.JSON(http.StatusOK, dataPengikutsResponse)
}

func (c *dataPengikutController) GetDataPengikut(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var dataPengikut, err = c.dataPengikutService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var dataPengikutResponse = helper.ConvertToDataPengikutResponse(dataPengikut)

	cntx.JSON(http.StatusOK, dataPengikutResponse)
}

func (c *dataPengikutController) GetDataPengikutBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var dataPengikuts, err = c.dataPengikutService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var dataPengikutsResponse []responses.DataPengikutResponse

	for _, dataPengikut := range dataPengikuts {
		var dataPengikutResponse = helper.ConvertToDataPengikutResponse(dataPengikut)
		dataPengikutsResponse = append(dataPengikutsResponse, dataPengikutResponse)
	}

	cntx.JSON(http.StatusOK, dataPengikutsResponse)
}

func (c *dataPengikutController) CreateDataPengikut(cntx *gin.Context) {
	var dataPengikutsRequest []request.CreateDataPengikutRequest

	var err = cntx.ShouldBindJSON(&dataPengikutsRequest)
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

	dataPengikuts, err := c.dataPengikutService.Create(dataPengikutsRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var dataPengikutsResponse []responses.DataPengikutResponse

	for _, dataPengikut := range dataPengikuts {
		var dataPengikutResponse = helper.ConvertToDataPengikutResponse(dataPengikut)
		dataPengikutsResponse = append(dataPengikutsResponse, dataPengikutResponse)
	}

	cntx.JSON(http.StatusCreated, dataPengikutsResponse)
}

func (c *dataPengikutController) DeleteDataPengikut(cntx *gin.Context) {
	var sptIdString = cntx.Param("sptid")
	var sptId, _ = strconv.Atoi(sptIdString)

	_, err := c.dataPengikutService.Delete(sptId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}
