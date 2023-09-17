package request

type CreateKegiatanRequest struct {
	ProgramId    int    `json:"programId" binding:"required"`
	KodeKegiatan string `json:"kodeKegiatan" binding:"required"`
	NamaKegiatan string `json:"namaKegiatan" binding:"required"`
	Tahun        string `json:"tahun" binding:"required"`
}

type UpdateKegiatanRequest struct {
	ProgramId    int    `json:"programId"`
	KodeKegiatan string `json:"kodeKegiatan"`
	NamaKegiatan string `json:"namaKegiatan"`
	Tahun        string `json:"tahun"`
}
