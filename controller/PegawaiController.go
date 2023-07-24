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

type pegawaiController struct {
	pegawaiService service.PegawaiService
}

func NewPegawaiController(pegawaiService service.PegawaiService) *pegawaiController {
	return &pegawaiController{pegawaiService}
}

func (c *pegawaiController) GetPegawais(cntx *gin.Context) {
	var pegawais, err = c.pegawaiService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var pegawaisResponse []responses.PegawaiResponse

	for _, pegawai := range pegawais {
		var pegawaiResponse = helper.ConvertToPegawaiResponse(pegawai)
		pegawaisResponse = append(pegawaisResponse, pegawaiResponse)
	}

	cntx.JSON(http.StatusOK, pegawaisResponse)
}

func (c *pegawaiController) GetPegawai(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var pegawai, err = c.pegawaiService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var pegawaiResponse = helper.ConvertToPegawaiResponse(pegawai)

	cntx.JSON(http.StatusOK, pegawaiResponse)
}

func (c *pegawaiController) GetPegawaiByName(cntx *gin.Context) {
	var name = cntx.Param("name")

	var pegawai, err = c.pegawaiService.FindByName(name)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var pegawaiResponse = helper.ConvertToPegawaiByNameResponse(pegawai)

	cntx.JSON(http.StatusOK, pegawaiResponse)
}

func (c *pegawaiController) GetPegawaisBySearch(cntx *gin.Context) {
	var whereClauseString = cntx.Request.URL.Query()
	var whereClauseInterface = map[string]interface{}{}

	for k, v := range whereClauseString {
		interfaceKey := k
		interfaceVal := v

		whereClauseInterface[interfaceKey] = interfaceVal
	}

	var pegawais, err = c.pegawaiService.FindBySearch(whereClauseInterface)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var pegawaisResponse []responses.PegawaiResponse

	for _, pegawai := range pegawais {
		var pegawaiResponse = helper.ConvertToPegawaiResponse(pegawai)
		pegawaisResponse = append(pegawaisResponse, pegawaiResponse)
	}

	cntx.JSON(http.StatusOK, pegawaisResponse)

}

func (c *pegawaiController) CreatePegawai(cntx *gin.Context) {
	var pegawaiRequest request.CreatePegawaiRequest

	var err = cntx.ShouldBindJSON(&pegawaiRequest)
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

	pegawai, err := c.pegawaiService.Create(pegawaiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"data": pegawai,
	})
}

func (c *pegawaiController) UpdatePegawai(cntx *gin.Context) {
	var pegawaiRequest request.UpdatePegawaiRequest

	var err = cntx.ShouldBindJSON(&pegawaiRequest)
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

	pegawai, err := c.pegawaiService.Update(id, pegawaiRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": pegawai,
	})
}

func (c *pegawaiController) UpdateBatchPegawaiStatusPerjalanan(cntx *gin.Context) {
	var updatePegawaiBatchRequest request.UpdatePegawaiBatchStatusPerjalanan

	var err = cntx.ShouldBindJSON(&updatePegawaiBatchRequest)
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

	err = c.pegawaiService.UpdateBatchStatusPerjalanan(updatePegawaiBatchRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data status perjalanan pegawai berhasil diubah",
	})
}

func (c *pegawaiController) DeletePegawai(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.pegawaiService.Delete(id)
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
