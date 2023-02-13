package request

type CreateInstansiRequest struct {
	Instansi      string `json:"instansi" binding:"required"`
	Kode_Instansi string `json:"kode_instansi"`
	Singkatan     string `json:"singkatan" binding:"required"`
}

type UpdateInstansiRequest struct {
	Instansi      string `json:"instansi"`
	Kode_Instansi string `json:"kode_instansi"`
	Singkatan     string `json:"singkatan"`
}
