package request

type CreateBidangRequest struct {
	Nama_Bidang string `json:"nama_bidang" binding:"required"`
	Singkatan   string `json:"singkatan" binding:"required"`
}

type UpdateBidangRequest struct {
	Nama_Bidang string `json:"nama_bidang"`
	Singkatan   string `json:"singkatan"`
}
