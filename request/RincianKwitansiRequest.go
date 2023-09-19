package request

type CreateRincianKwitansiRequest struct {
	Id          string `json:"id" binding:"required"`
	KwitansiId  int    `json:"kwitansiId" binding:"required"`
	Jenis       string `json:"jenis" binding:"required"`
	NamaRincian string `json:"namaRincian" binding:"required"`
	JumlahBayar int    `json:"jumlahBayar" binding:"required"`
	Banyaknya   int    `json:"banyaknya" binding:"required"`
	HasilBayar  int    `json:"hasilBayar" binding:"required"`
}

type UpdateRincianKwitansiRequest struct {
	Id          string `json:"id"`
	KwitansiId  int    `json:"kwitansiId"`
	Jenis       string `json:"jenis"`
	NamaRincian string `json:"namaRincian"`
	JumlahBayar int    `json:"jumlahBayar"`
	Banyaknya   int    `json:"banyaknya"`
	HasilBayar  int    `json:"hasilBayar"`
}
