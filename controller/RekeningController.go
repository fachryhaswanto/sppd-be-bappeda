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

type rekeningController struct {
	rekeningService service.RekeningService
}

func NewRekeningController(rekeningService service.RekeningService) *rekeningController {
	return &rekeningController{rekeningService}
}

func (c *rekeningController) GetRekenings(cntx *gin.Context) {
	var rekenings, err = c.rekeningService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var rekeningsResponse []responses.RekeningResponse

	for _, rekening := range rekenings {
		var rekeningResponse = helper.ConvertToRekeningResponse(rekening)
		rekeningsResponse = append(rekeningsResponse, rekeningResponse)
	}

	cntx.JSON(http.StatusOK, rekeningsResponse)
}

func (c *rekeningController) GetRekening(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var rekening, err = c.rekeningService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var rekeningResponse = helper.ConvertToRekeningResponse(rekening)

	cntx.JSON(http.StatusOK, rekeningResponse)
}

func (c *rekeningController) CreateRekening(cntx *gin.Context) {
	var rekeningRequest request.CreateRekeningRequest

	var err = cntx.ShouldBindJSON(&rekeningRequest)
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

	rekening, err := c.rekeningService.Create(rekeningRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rekeningResponse = helper.ConvertToRekeningResponse(rekening)

	cntx.JSON(http.StatusCreated, rekeningResponse)
}

func (c *rekeningController) UpdateRekening(cntx *gin.Context) {
	var rekeningRequest request.UpdateRekeningRequest

	var err = cntx.ShouldBindJSON(&rekeningRequest)
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

	rekening, err := c.rekeningService.Update(id, rekeningRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var rekeningResponse = helper.ConvertToRekeningResponse(rekening)

	cntx.JSON(http.StatusOK, rekeningResponse)
}

func (c *rekeningController) DeleteRekening(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.rekeningService.Delete(id)
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
