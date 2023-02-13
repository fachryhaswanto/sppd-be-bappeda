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

type bidangController struct {
	bidangService service.BidangService
}

func NewBidangController(bidangService service.BidangService) *bidangController {
	return &bidangController{bidangService}
}

func (c *bidangController) GetBidangs(cntx *gin.Context) {
	var bidangs, err = c.bidangService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var bidangsResponse []responses.BidangResponse

	for _, bidang := range bidangs {
		var bidangResponse = helper.ConvertToBidangResponse(bidang)
		bidangsResponse = append(bidangsResponse, bidangResponse)
	}

	cntx.JSON(http.StatusOK, bidangsResponse)
}

func (c *bidangController) GetBidang(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var bidang, err = c.bidangService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var bidangResponse = helper.ConvertToBidangResponse(bidang)

	cntx.JSON(http.StatusOK, bidangResponse)
}

func (c *bidangController) CreateBidang(cntx *gin.Context) {
	var bidangRequest request.CreateBidangRequest

	var err = cntx.ShouldBindJSON(&bidangRequest)
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

	bidang, err := c.bidangService.Create(bidangRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"data": bidang,
	})
}

func (c *bidangController) UpdateBidang(cntx *gin.Context) {
	var bidangRequest request.UpdateBidangRequest

	var err = cntx.ShouldBindJSON(&bidangRequest)
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

	bidang, err := c.bidangService.Update(id, bidangRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": bidang,
	})
}

func (c *bidangController) DeleteBidang(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.bidangService.Delete(id)
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
