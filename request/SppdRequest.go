package request

type CreateSppdRequest struct {
	Template_Sppd    string `json:"template_sppd" binding:"required"`
	Jenis            string `json:"jenis" binding:"required"`
	Nomor_Sppd       string `json:"nomor_sppd" binding:"required"`
	PegawaiId        int    `json:"pegawaiId" binding:"required"`
	Tempat_Berangkat string `json:"tempat_berangkat" binding:"required"`
	Tempat_Tujuan    string `json:"tempat_tujuan" binding:"required"`
	Tanggal_Sppd     string `json:"tanggal_sppd" binding:"required"`
	Alat_Angkutan    string `json:"alat_angkutan" binding:"required"`
	Instansi         string `json:"instansi" binding:"required"`
	PejabatId        int    `json:"pejabatId" binding:"required"`
	SptId            int    `json:"sptId" binding:"required"`
	UserId           int    `json:"userId" binding:"required"`
}

type UpdateSppdRequest struct {
	Template_Sppd    string `json:"template_sppd"`
	Jenis            string `json:"jenis"`
	Nomor_Sppd       string `json:"nomor_sppd"`
	PegawaiId        int    `json:"pegawaiId"`
	Tempat_Berangkat string `json:"tempat_berangkat"`
	Tempat_Tujuan    string `json:"tempat_tujuan"`
	Tanggal_Sppd     string `json:"tanggal_sppd"`
	Alat_Angkutan    string `json:"alat_angkutan"`
	Instansi         string `json:"instansi"`
	PejabatId        int    `json:"pejabatId"`
	SptId            int    `json:"sptId"`
	UserId           int    `json:"userId"`
}
