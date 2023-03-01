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

type sptController struct {
	sptService  service.SptService
	sppdService service.SppdService
}

func NewSptController(sptService service.SptService, sppdService service.SppdService) *sptController {
	return &sptController{
		sptService,
		sppdService,
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

	var sppd model.Sppd
	sppd.Idspt = spt.Id

	_, err = c.sppdService.Create(sppd)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"data": spt,
	})
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
