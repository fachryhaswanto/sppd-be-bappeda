package request

type CreateRincianKwitansiPenerbanganRequest struct {
	Id           string `json:"id" binding:"required"`
	KwitansiId   int    `json:"kwitansiId" binding:"required"`
	NamaMaskapai string `json:"namaMaskapai" binding:"required"`
	NomorTiket   string `json:"nomorTiket" binding:"required"`
	KodeBooking  string `json:"kodeBooking" binding:"required"`
}

type UpdateRincianKwitansiPenerbanganRequest struct {
	Id           string `json:"id"`
	KwitansiId   int    `json:"kwitansiId"`
	NamaMaskapai string `json:"namaMaskapai"`
	NomorTiket   string `json:"nomorTiket"`
	KodeBooking  string `json:"kodeBooking"`
}
