package request

type CreatePegawaiRequest struct {
	Nip           string `json:"nip" binding:"required"`
	Nama          string `json:"nama" binding:"required"`
	Jenis_Kelamin string `json:"jenis_kelamin" binding:"required"`
	Status        string `json:"status" binding:"required"`
	Tempat_Lahir  string `json:"tempat_lahir" binding:"required"`
	Tanggal_Lahir string `json:"tanggal_lahir" binding:"required"`
	Instansi      string `json:"instansi" binding:"required"`
	Bidang        string `json:"bidang" binding:"required"`
	Golongan      string `json:"golongan" binding:"required"`
	Eselon        string `json:"eselon" binding:"required"`
	Pangkat       string `json:"pangkat" binding:"required"`
	Jabatan       string `json:"jabatan" binding:"required"`
}

type UpdatePegawaiRequest struct {
	Nip           string `json:"nip"`
	Nama          string `json:"nama"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Status        string `json:"status"`
	Tempat_Lahir  string `json:"tempat_lahir"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
	Instansi      string `json:"instansi"`
	Bidang        string `json:"bidang"`
	Golongan      string `json:"golongan"`
	Eselon        string `json:"eselon"`
	Pangkat       string `json:"pangkat"`
	Jabatan       string `json:"jabatan"`
}
