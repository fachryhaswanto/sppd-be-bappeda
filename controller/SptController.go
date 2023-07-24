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

type sptController struct {
	sptService service.SptService
}

func NewSptController(sptService service.SptService) *sptController {
	return &sptController{
		sptService,
	}
}

func (c *sptController) GetSpts(cntx *gin.Context) {
	var spts, err = c.sptService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var sptsResponse []responses.SptResponse

	for _, spt := range spts {
		var sptResponse = helper.ConvertToSptResponse(spt)
		sptsResponse = append(sptsResponse, sptResponse)
	}

	cntx.JSON(http.StatusOK, sptsResponse)
}

func (c *sptController) GetSpt(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var spt, err = c.sptService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var sptResponse = helper.ConvertToSptResponse(spt)

	cntx.JSON(http.StatusOK, sptResponse)
}

func (c *sptController) GetAllDitugaskan(cntx *gin.Context) {
	var ditugaskans, err = c.sptService.FindAllDitugaskan()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var ditugaskansResponse []responses.SemuaNamaDitugaskan

	for _, ditugaskan := range ditugaskans {
		var ditugaskanResponse = helper.ConvertToSemuaNamaDitugaskan(ditugaskan)
		ditugaskansResponse = append(ditugaskansResponse, ditugaskanResponse)
	}

	cntx.JSON(http.StatusOK, ditugaskansResponse)
}

func (c *sptController) GetSptsBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var spts, err = c.sptService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var sptsResponse []responses.SptResponse

	for _, spt := range spts {
		var sptResponse = helper.ConvertToSptResponse(spt)
		sptsResponse = append(sptsResponse, sptResponse)
	}

	cntx.JSON(http.StatusOK, sptsResponse)
}

func (c *sptController) CreateSpt(cntx *gin.Context) {
	var sptRequest request.CreateSptRequest

	var err = cntx.ShouldBindJSON(&sptRequest)
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

	spt, err := c.sptService.Create(sptRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var sptResponse = helper.ConvertToSptResponse(spt)

	cntx.JSON(http.StatusCreated, sptResponse)
}

func (c *sptController) UpdateSpt(cntx *gin.Context) {
	var sptRequest request.UpdateSptRequest

	var err = cntx.ShouldBindJSON(&sptRequest)
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

	spt, err := c.sptService.Update(id, sptRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": spt,
	})
}

func (c *sptController) UpdateStatusSppd(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var valueString = cntx.Param("value")
	var value, _ = strconv.Atoi(valueString)

	var err = c.sptService.UpdateStatusSppd(id, value)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, cntx.Error(err))
		return
	}

	cntx.JSON(http.StatusOK, "data spt berhasil diperbarui")
}

func (c *sptController) DeleteSpt(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.sptService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})

	// var e = os.Remove("./documents/" + idString + ".pdf")
	// if e != nil {
	// 	return
	// }
}
