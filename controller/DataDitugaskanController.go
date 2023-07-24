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

type dataDitugaskanController struct {
	dataDitugaskanService service.DataDitugaskanService
}

func NewDataDitugaskanController(dataDitugaskanService service.DataDitugaskanService) *dataDitugaskanController {
	return &dataDitugaskanController{dataDitugaskanService}
}

func (c *dataDitugaskanController) GetDataDitugaskans(cntx *gin.Context) {
	var dataDitugaskans, err = c.dataDitugaskanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var dataDitugaskansResponse []responses.DataDitugaskanResponse

	for _, dataDitugaskan := range dataDitugaskans {
		var dataDitugaskanResponse = helper.ConvertToDataDitugaskanResponse(dataDitugaskan)
		dataDitugaskansResponse = append(dataDitugaskansResponse, dataDitugaskanResponse)
	}

	cntx.JSON(http.StatusOK, dataDitugaskansResponse)
}

func (c *dataDitugaskanController) GetDataDitugaskan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var dataDitugaskan, err = c.dataDitugaskanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var dataDitugaskanResponse = helper.ConvertToDataDitugaskanResponse(dataDitugaskan)

	cntx.JSON(http.StatusOK, dataDitugaskanResponse)
}

func (c *dataDitugaskanController) GetDataDitugaskansBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var dataDitugaskans, err = c.dataDitugaskanService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var dataDitugaskansResponse []responses.DataDitugaskanResponse

	for _, dataDitugaskan := range dataDitugaskans {
		var dataDitugaskanResponse = helper.ConvertToDataDitugaskanResponse(dataDitugaskan)
		dataDitugaskansResponse = append(dataDitugaskansResponse, dataDitugaskanResponse)
	}

	cntx.JSON(http.StatusOK, dataDitugaskansResponse)

}

func (c *dataDitugaskanController) CountDataBySptId(cntx *gin.Context) {
	var sptIdString = cntx.Query("sptid")
	var sptId, _ = strconv.Atoi(sptIdString)

	var jumlah, err = c.dataDitugaskanService.CountDataBySptId(sptId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, jumlah)
}

func (c *dataDitugaskanController) CreateDataDitugaskan(cntx *gin.Context) {
	var dataDitugaskanRequest []request.CreateDataDitugaskanRequest

	var err = cntx.ShouldBindJSON(&dataDitugaskanRequest)
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

	dataDitugaskans, err := c.dataDitugaskanService.Create(dataDitugaskanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, cntx.Error(err))
		return
	}

	var dataDitugaskansResponse []responses.DataDitugaskanResponse

	for _, dataDitugaskan := range dataDitugaskans {
		var dataDitugaskanResponse = helper.ConvertToDataDitugaskanResponse(dataDitugaskan)
		dataDitugaskansResponse = append(dataDitugaskansResponse, dataDitugaskanResponse)
	}

	cntx.JSON(http.StatusCreated, dataDitugaskansResponse)
}

func (c *dataDitugaskanController) UpdateStatusSppd(cntx *gin.Context) {
	var sptIdString = cntx.Param("sptid")
	var sptId, _ = strconv.Atoi(sptIdString)

	var valueString = cntx.Param("value")
	var value, _ = strconv.Atoi(valueString)

	var pegawaiIdString = cntx.Query("pegawaiid")
	var pegawaiId, _ = strconv.Atoi(pegawaiIdString)

	var err = c.dataDitugaskanService.UpdateStatus(sptId, pegawaiId, value)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, cntx.Error(err))
		return
	}

	cntx.JSON(http.StatusOK, "data ditugaskan berhasil diperbarui")
}

func (c *dataDitugaskanController) DeleteDataDitugaskan(cntx *gin.Context) {
	var sptIdString = cntx.Param("sptid")
	var sptId, _ = strconv.Atoi(sptIdString)

	_, err := c.dataDitugaskanService.Delete(sptId)
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
