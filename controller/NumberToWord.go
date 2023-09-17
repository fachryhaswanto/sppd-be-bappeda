package controller

import (
	"net/http"
	"sppd/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NumberToWord(cntx *gin.Context) {
	var numberString = cntx.Query("number")
	var number, _ = strconv.ParseInt(numberString, 10, 64)

	var terbilang = helper.NumberToWords(number)

	cntx.JSON(http.StatusOK, terbilang)
}
