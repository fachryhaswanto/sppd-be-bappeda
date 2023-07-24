package request

type CreateSptRequest struct {
	Jenis             string `json:"jenis" binding:"required"`
	Template          string `json:"template" binding:"required"`
	SubKegiatanId     int    `json:"subKegiatanId" binding:"required"`
	Nomor_Spt         string `json:"nomor_spt" binding:"required"`
	Tanggal_Spt       string `json:"tanggal_spt" binding:"required"`
	RekeningId        int    `json:"rekeningId" binding:"required"`
	Keperluan         string `json:"keperluan" binding:"required"`
	Tanggal_Berangkat string `json:"tanggal_berangkat" binding:"required"`
	Tanggal_Kembali   string `json:"tanggal_kembali" binding:"required"`
	Lama_Perjalanan   string `json:"lama_perjalanan" binding:"required"`
	Pejabat_Pemberi   string `json:"pejabat_pemberi" binding:"required"`
	Status            string `json:"status"`
	File_Surat_Tugas  string `json:"file_surat_tugas"`
	UserId            int    `json:"userId"`
}

type UpdateSptRequest struct {
	Jenis             string `json:"jenis"`
	Template          string `json:"string"`
	Nomor_Spt         string `json:"nomor_spt"`
	SubKegiatanId     int    `json:"subKegiatanId"`
	Tanggal_Spt       string `json:"tanggal_spt"`
	RekeningId        int    `json:"rekeningId"`
	Keperluan         string `json:"keperluan"`
	Tanggal_Berangkat string `json:"tanggal_berangkat"`
	Tanggal_Kembali   string `json:"tanggal_kembali"`
	Lama_Perjalanan   string `json:"lama_perjalanan"`
	Pejabat_Pemberi   string `json:"pejabat_pemberi"`
	Status            string `json:"status"`
	StatusSppd        int    `json:"statusSppd"`
	File_Surat_Tugas  string `json:"file_surat_tugas"`
	UserId            int    `json:"userId"`
}
