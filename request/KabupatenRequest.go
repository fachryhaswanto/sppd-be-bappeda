package request

type CreateKabupatenRequest struct {
	Kode string `json:"kode" binding:"required"`
	Nama string `json:"nama" binding:"required"`
}

type UpdateKabupatenRequest struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}
