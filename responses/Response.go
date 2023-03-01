package responses

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type KabupatenResponse struct {
	Id   int    `json:"id"`
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type InstansiResponse struct {
	Id            int    `json:"id"`
	Instansi      string `json:"instansi"`
	Kode_Instansi string `json:"kode_instansi"`
	Singkatan     string `json:"singkatan"`
}

type BidangResponse struct {
	Id          int    `json:"id"`
	Nama_Bidang string `json:"nama_bidang"`
	Singkatan   string `json:"singkatan"`
}

type PejabatResponse struct {
	Id       int    `json:"id"`
	Nip      string `json:"nip"`
	Nama     string `json:"nama"`
	Pangkat  string `json:"pangkat"`
	Golongan string `json:"golongan"`
	Eselon   string `json:"eselon"`
	Jabatan  string `json:"jabatan"`
}

type PegawaiResponse struct {
	Id            int    `json:"id"`
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

type PegawaiByNameResponse struct {
	Nip      string `json:"nip"`
	Nama     string `json:"nama"`
	Golongan string `json:"golongan"`
	Eselon   string `json:"eselon"`
	Pangkat  string `json:"pangkat"`
	Jabatan  string `json:"jabatan"`
}

type ProgramResponse struct {
	Id         int    `json:"id"`
	Kode       string `json:"kode"`
	Pembebanan string `json:"pembebanan"`
	Program    string `json:"program"`
}

type KegiatanResponse struct {
	Id           int             `json:"id"`
	ProgramId    int             `json:"programId"`
	Program      ProgramResponse `json:"program"`
	KodeKegiatan string          `json:"kodeKegiatan"`
	NamaKegiatan string          `json:"namaKegiatan"`
}

type SubKegiatanResponse struct {
	Id           int    `json:"id"`
	Tahun        string `json:"tahun"`
	Nama_Program string `json:"nama_program"`
	Kode         string `json:"kode"`
	Kegiatan     string `json:"kegiatan"`
}

type SptResponse struct {
	Id                int    `json:"id"`
	Template          string `json:"template"`
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
	File_Surat_Tugas  string `json:"file_surat_tugas"`
	Status            string `json:"status"`
}

type SppdResponse struct {
	Id            int    `json:"id"`
	Template_Sppd string `json:"template_sppd"`
	Nomor_Sppd    string `json:"nomor_sppd"`
	Tanggal_Sppd  string `json:"tanggal_sppd"`
	Tingkat_Biaya string `json:"tingkat_biaya"`
	Instansi      string `json:"instansi"`
	Tanda_Tangan  string `json:"tanda_tangan"`
	Idspt         int    `json:"idspt"`
}

type JoinResponse struct {
	Id                int    `json:"id"`
	Template_Sppd     string `json:"template_sppd"`
	Ditugaskan        string `json:"ditugaskan"`
	Nomor_Sppd        string `json:"nomor_sppd"`
	Tanggal_Sppd      string `json:"tanggal_sppd"`
	Tingkat_Biaya     string `json:"tingkat_biaya"`
	Instansi          string `json:"instansi"`
	Tanda_Tangan      string `json:"tanda_tangan"`
	Nomor_Spt         string `json:"nomor_spt"`
	Tanggal_Spt       string `json:"tanggal_spt"`
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
}

type SemuaNamaDitugaskan struct {
	Id              int    `json:"id"`
	Tanggal_Kembali string `json:"tanggal_kembali"`
	Ditugaskan      string `json:"ditugaskan"`
	Status          string `json:"status"`
}
