package request

type CreateDataPengikutRequest struct {
	SptId     int `json:"sptId" binding:"required"`
	PegawaiId int `json:"pegawaiId" binding:"required"`
}
