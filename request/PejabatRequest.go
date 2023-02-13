package request

type CreatePejabatRequest struct {
	Nip      string `json:"nip" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Pangkat  string `json:"pangkat" binding:"required"`
	Golongan string `json:"golongan" binding:"required"`
	Eselon   string `json:"eselon" binding:"required"`
	Jabatan  string `json:"jabatan" binding:"required"`
}

type UpdatePejabatRequest struct {
	Nip      string `json:"nip"`
	Nama     string `json:"nama"`
	Pangkat  string `json:"pangkat"`
	Golongan string `json:"golongan"`
	Eselon   string `json:"eselon"`
	Jabatan  string `json:"jabatan"`
}
