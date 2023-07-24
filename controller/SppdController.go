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

type sppdController struct {
	sppdService service.SppdService
	sptService  service.SptService
}

func NewSppdController(sppdService service.SppdService, sptService service.SptService) *sppdController {
	return &sppdController{sppdService, sptService}
}

func (c *sppdController) GetSppds(cntx *gin.Context) {
	var sppds, err = c.sppdService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var sppdsResponse []responses.SppdResponse
	for _, sppd := range sppds {
		var sppdResponse = helper.ConvertToSppdResponse(sppd)
		sppdsResponse = append(sppdsResponse, sppdResponse)
	}

	cntx.JSON(http.StatusOK, sppdsResponse)
}

func (c *sppdController) GetSppd(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var sppd, err = c.sppdService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var sppdResponse = helper.ConvertToSppdResponse(sppd)

	cntx.JSON(http.StatusOK, sppdResponse)
}

func (c *sppdController) GetSppdBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var sppds, err = c.sppdService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var sppdsResponse []responses.SppdResponse

	for _, sppd := range sppds {
		var sppdResponse = helper.ConvertToSppdResponse(sppd)
		sppdsResponse = append(sppdsResponse, sppdResponse)
	}

	cntx.JSON(http.StatusOK, sppdsResponse)
}

func (c *sppdController) CountDataBySptId(cntx *gin.Context) {
	var sptIdString = cntx.Query("sptid")
	var sptId, _ = strconv.Atoi(sptIdString)

	var jumlah, err = c.sppdService.CountDataBySptId(sptId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	cntx.JSON(http.StatusOK, jumlah)
}

func (c *sppdController) CreateSppd(cntx *gin.Context) {
	var sppdRequest request.CreateSppdRequest

	var err = cntx.ShouldBindJSON(&sppdRequest)
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

	sppd, err := c.sppdService.Create(sppdRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var sppdResponse = helper.ConvertToSppdResponse(sppd)

	cntx.JSON(http.StatusCreated, sppdResponse)
}

func (c *sppdController) UpdateSppd(cntx *gin.Context) {
	var sppdRequest request.UpdateSppdRequest

	var err = cntx.ShouldBindJSON(&sppdRequest)
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

	sppd, err := c.sppdService.Update(id, sppdRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": sppd,
	})
}

func (c *sppdController) UpdateSppdBySptId(cntx *gin.Context) {
	var sppdRequest request.UpdateSppdRequest

	var err = cntx.ShouldBindJSON(&sppdRequest)
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

	var sptIdString = cntx.Param("sptid")
	var sptId, _ = strconv.Atoi(sptIdString)

	sppd, err := c.sppdService.UpdateBySptId(sptId, sppdRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var sppdRespone = helper.ConvertToSppdResponse(sppd)

	cntx.JSON(http.StatusOK, sppdRespone)
}

func (c *sppdController) DeleteSppd(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.sppdService.Delete(id)
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

func (c *sppdController) DeleteSppdBySptId(cntx *gin.Context) {
	var sptIdString = cntx.Param("sptid")
	var sptId, _ = strconv.Atoi(sptIdString)

	err := c.sppdService.DeleteBySptId(sptId)
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
