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

type programController struct {
	programService service.ProgramService
}

func NewProgramController(programService service.ProgramService) *programController {
	return &programController{programService}
}

func (c *programController) GetPrograms(cntx *gin.Context) {
	var programs, err = c.programService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var programsResponse []responses.ProgramResponse

	for _, program := range programs {
		var programResponse = helper.ConvertToProgramResponse(program)
		programsResponse = append(programsResponse, programResponse)
	}

	cntx.JSON(http.StatusOK, programsResponse)
}

func (c *programController) GetProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var bidang, err = c.programService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var programResponse = helper.ConvertToProgramResponse(bidang)

	cntx.JSON(http.StatusOK, programResponse)
}

func (c *programController) CreateProgram(cntx *gin.Context) {
	var programRequest request.CreateProgramRequest

	var err = cntx.ShouldBindJSON(&programRequest)
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

	program, err := c.programService.Create(programRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusCreated, gin.H{
		"data": program,
	})
}

func (c *programController) UpdateProgram(cntx *gin.Context) {
	var programRequest request.UpdateProgramRequest

	var err = cntx.ShouldBindJSON(&programRequest)
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

	program, err := c.programService.Update(id, programRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"data": program,
	})
}

func (c *programController) DeleteProgram(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.programService.Delete(id)
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
