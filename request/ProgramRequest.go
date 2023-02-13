package request

type CreateProgramRequest struct {
	Tahun      string `json:"tahun" binding:"required"`
	Kode       string `json:"kode" binding:"required"`
	Pembebanan string `json:"pembebanan" binding:"required"`
	Program    string `json:"program" binding:"required"`
}

type UpdateProgramRequest struct {
	Tahun      string `json:"tahun"`
	Kode       string `json:"kode"`
	Pembebanan string `json:"pembebanan"`
	Program    string `json:"program"`
}
