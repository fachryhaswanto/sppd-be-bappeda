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

type pejabatController struct {
	pejabatService service.PejabatService
}

func NewPejabatController(pejabatService service.PejabatService) *pejabatController {
	return &pejabatController{pejabatService}
}

func (c *pejabatController) GetPejabats(cntx *gin.Context) {
	var pejabats, err = c.pejabatService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var pejabatsResponse []responses.PejabatResponse

	for _, pejabat := range pejabats {
		var pejabatResponse = helper.ConvertToPejabatResponse(pejabat)
		pejabatsResponse = append(pejabatsResponse, pejabatResponse)
	}

	cntx.JSON(http.StatusOK, pejabatsResponse)
}

func (c *pejabatController) GetPejabat(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var pejabat, err = c.pejabatService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var pejabatResponse = helper.ConvertToPejabatResponse(pejabat)

	cntx.JSON(http.StatusOK, pejabatResponse)
}

func (c *pejabatController) GetPejabatByName(cntx *gin.Context) {
	var name = cntx.Param("name")

	var pejabat, err = c.pejabatService.FindByName(name)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var pejabatResponse = helper.ConvertToPejabatResponse(pejabat)

	cntx.JSON(http.StatusOK, pejabatResponse)
}

func (c *pejabatController) GetPejabatBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceValue := v

		whereClauseInterface[interfaceKey] = interfaceValue
	}

	var pejabats, err = c.pejabatService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var pejabatsResponse []responses.PejabatResponse

	for _, pejabat := range pejabats {
		var pejabatResponse = helper.ConvertToPejabatResponse(pejabat)
		pejabatsResponse = append(pejabatsResponse, pejabatResponse)
	}

	cntx.JSON(http.StatusOK, pejabatsResponse)
}

func (c *pejabatController) CreatePejabat(cntx *gin.Context) {
	var pejabatRequest request.CreatePejabatRequest

	var err = cntx.ShouldBindJSON(&pejabatRequest)
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

	pejabat, err := c.pejabatService.Create(pejabatRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"data": pejabat,
	})
}

func (c *pejabatController) UpdatePejabat(cntx *gin.Context) {
	var pejabatRequest request.UpdatePejabatRequest

	var err = cntx.ShouldBindJSON(&pejabatRequest)
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

	pejabat, err := c.pejabatService.Update(id, pejabatRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": pejabat,
	})
}

func (c *pejabatController) DeletePejabat(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.pejabatService.Delete(id)
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
