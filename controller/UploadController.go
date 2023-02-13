package controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(cntx *gin.Context) {

	var file, err = cntx.FormFile("file")
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var fileName = filepath.Base(file.Filename)
	if err = cntx.SaveUploadedFile(file, "./documents/"+fileName+".pdf"); err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
	})
}
