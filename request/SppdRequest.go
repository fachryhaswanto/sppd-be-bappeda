package request

type CreateSppdRequest struct {
	Template_Sppd string `json:"template_sppd"`
	Nomor_Sppd    string `json:"nomor_sppd"`
	Tanggal_Sppd  string `json:"tanggal_sppd"`
	Tingkat_Biaya string `json:"tingkat_biaya"`
	Instansi      string `json:"instansi"`
	Tanda_tangan  string `json:"tanda_tangan"`
}

type UpdateSppdRequest struct {
	Template_Sppd string `json:"template_sppd"`
	Nomor_Sppd    string `json:"nomor_sppd"`
	Tanggal_Sppd  string `json:"tanggal_sppd"`
	Tingkat_Biaya string `json:"tingkat_biaya"`
	Instansi      string `json:"instansi"`
	Tanda_tangan  string `json:"tanda_tangan"`
	Idspt         int    `json:"idspt"`
}
