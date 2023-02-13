package request

type CreateSptRequest struct {
	Template          string `json:"template" binding:"required"`
	Nomor_Spt         string `json:"nomor_spt" binding:"required"`
	Tanggal_Spt       string `json:"tanggal_spt" binding:"required"`
	Ditugaskan        string `json:"ditugaskan" binding:"required"`
	Jenis_Perjalanan  string `json:"jenis_perjalanan" binding:"required"`
	Keperluan         string `json:"keperluan" binding:"required"`
	Alat_Angkutan     string `json:"alat_angkutan" binding:"required"`
	Tempat_Berangkat  string `json:"tempat_berangkat" binding:"required"`
	Tempat_Tujuan     string `json:"tempat_tujuan" binding:"required"`
	Tanggal_Berangkat string `json:"tanggal_berangkat" binding:"required"`
	Tanggal_Kembali   string `json:"tanggal_kembali" binding:"required"`
	Lama_Perjalanan   string `json:"lama_perjalanan" binding:"required"`
	Pejabat_Pemberi   string `json:"pejabat_pemberi" binding:"required"`
	Status            string `json:"status"`
	File_Surat_Tugas  string `json:"file_surat_tugas"`
}

type UpdateSptRequest struct {
	Template          string `json:"string"`
	Nomor_Spt         string `json:"nomor_spt"`
	Tanggal_Spt       string `json:"tanggal_spt"`
	Ditugaskan        string `json:"ditugaskan"`
	Jenis_Perjalanan  string `json:"jenis_perjalanan"`
	Keperluan         string `json:"keperluan"`
	Alat_Angkutan     string `json:"alat_angkutan"`
	Tempat_Berangkat  string `json:"tempat_berangkat"`
	Tempat_Tujuan     string `json:"tempat_tujuan"`
	Tanggal_Berangkat string `json:"tanggal_berangkat"`
	Tanggal_Kembali   string `json:"tanggal_kembali"`
	Lama_Perjalanan   string `json:"lama_perjalanan"`
	Pejabat_Pemberi   string `json:"pejabat_pemberi"`
	Status            string `json:"status"`
	File_Surat_Tugas  string `json:"file_surat_tugas"`
}
