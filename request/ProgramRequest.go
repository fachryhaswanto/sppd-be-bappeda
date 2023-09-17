package request

type CreateProgramRequest struct {
	Kode       string `json:"kode" binding:"required"`
	Pembebanan string `json:"pembebanan" binding:"required"`
	Program    string `json:"program" binding:"required"`
	Tahun      string `json:"tahun" binding:"required"`
}

type UpdateProgramRequest struct {
	Kode       string `json:"kode"`
	Pembebanan string `json:"pembebanan"`
	Program    string `json:"program"`
	Tahun      string `json:"tahun"`
}
