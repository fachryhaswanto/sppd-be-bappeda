package controller

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// func GetFile(cntx *gin.Context) {
// 	var fileName = cntx.Param("filename")

// 	var byteFile, err = ioutil.ReadFile("./documents/" + fileName + ".pdf")
// 	if err != nil {
// 		cntx.JSON(http.StatusBadRequest, gin.H{
// 			"error": cntx.Error(err),
// 		})
// 	}

// 	cntx.Header("Content-Disposition", "attachment; filename="+fileName+".pdf")
// 	cntx.Data(http.StatusOK, "application/octet-stream", byteFile)
// }

func GetFile(cntx *gin.Context) {
	var fileName = cntx.Param("filename")

	var file, err = os.OpenFile("./documents/"+fileName+".pdf", os.O_RDWR, 0666)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		cntx.JSON(http.StatusInternalServerError, cntx.Error(err))
		return
	}

	cntx.Header("Content-Disposition", "attachment; filename="+fileName+".pdf")

	if _, err := io.Copy(cntx.Writer, file); err != nil {
		cntx.JSON(http.StatusInternalServerError, cntx.Error(err))
	}
}
