package request

type CreateSubKegiatanRequest struct {
	KegiatanId      int    `json:"kegiatanId" binding:"required"`
	KodeSubKegiatan string `json:"kodeSubKegiatan" binding:"required"`
	NamaSubKegiatan string `json:"namaSubKegiatan" binding:"required"`
	PejabatId       int    `json:"pejabatId" binding:"required"`
	Tahun           string `json:"tahun" binding:"required"`
}

type UpdateSubKegiatanRequest struct {
	KegiatanId      int    `json:"kegiatanId"`
	KodeSubKegiatan string `json:"kodeSubKegiatan"`
	NamaSubKegiatan string `json:"namaSubKegiatan"`
	PejabatId       int    `json:"pejabatId"`
	Tahun           string `json:"tahun"`
}
