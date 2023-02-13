package request

type CreateSubKegiatanRequest struct {
	Tahun        string `json:"tahun" binding:"required"`
	Nama_Program string `json:"nama_program" binding:"required"`
	Kode         string `json:"kode" binding:"required"`
	Kegiatan     string `json:"kegiatan" binding:"required"`
}

type UpdateSubKegiatanRequest struct {
	Tahun        string `json:"tahun"`
	Nama_Program string `json:"nama_program"`
	Kode         string `json:"kode"`
	Kegiatan     string `json:"kegiatan"`
}
