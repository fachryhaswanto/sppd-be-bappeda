package request

type CreateKwitansiRequest struct {
	SppdId        int    `json:"sppdId" binding:"required"`
	NomorKwitansi string `json:"nomorKwitansi" binding:"required"`
	TanggalBayar  string `json:"tanggalBayar" binding:"required"`
	Keperluan     string `json:"keperluan" binding:"required"`
	TotalBayar    int    `json:"totalBayar"`
	Tahun         string `json:"tahun"`
	UserId        int    `json:"userId" binding:"required"`
}

type UpdateKwitansiRequest struct {
	SppdId        int    `json:"sppdId"`
	NomorKwitansi string `json:"nomorKwitansi"`
	TanggalBayar  string `json:"tanggalBayar"`
	Keperluan     string `json:"keperluan"`
	TotalBayar    int    `json:"totalBayar"`
	Tahun         string `json:"tahun"`
	UserId        int    `json:"userId"`
}

type UpdateTotalBayarRequest struct {
	TotalBayar int `json:"totalBayar"`
}
