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
}

func NewSppdController(sppdService service.SppdService) *sppdController {
	return &sppdController{sppdService}
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

func (c *sppdController) GetJoin(cntx *gin.Context) {
	var joins, err = c.sppdService.FindJoin()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var joinsResponse []responses.JoinResponse

	for _, join := range joins {
		var joinResponse = helper.ConvertToJoinResponse(join)
		joinsResponse = append(joinsResponse, joinResponse)
	}

	cntx.JSON(http.StatusOK, joinsResponse)
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

// func (c *sppdController) CreateSppd(cntx *gin.Context) {
// 	var sppdRequest request.CreateSppdRequest

// 	var err = cntx.ShouldBindJSON(&sppdRequest)
// 	if err != nil {
// 		var errorMessages = []string{}

// 		for _, e := range err.(validator.ValidationErrors) {
// 			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
// 			errorMessages = append(errorMessages, errorMessage)
// 		}

// 		cntx.JSON(http.StatusBadRequest, gin.H{
// 			"error": errorMessages,
// 		})
// 		return
// 	}

// 	sppd, err := c.sppdService.Create(sppdRequest)
// 	if err != nil {
// 		cntx.JSON(http.StatusBadRequest, gin.H{
// 			"error": cntx.Error(err),
// 		})
// 		return
// 	}

// 	cntx.JSON(http.StatusCreated, gin.H{
// 		"data": sppd,
// 	})
// }

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
