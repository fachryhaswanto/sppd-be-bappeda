package responses

import "time"

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
	Id               int            `json:"id"`
	Nip              string         `json:"nip"`
	Nama             string         `json:"nama"`
	Jenis_Kelamin    string         `json:"jenis_kelamin"`
	Status           string         `json:"status"`
	Tempat_Lahir     string         `json:"tempat_lahir"`
	Tanggal_Lahir    string         `json:"tanggal_lahir"`
	Instansi         string         `json:"instansi"`
	BidangId         int            `json:"bidangId"`
	Bidang           BidangResponse `json:"bidang"`
	Golongan         string         `json:"golongan"`
	Eselon           string         `json:"eselon"`
	Pangkat          string         `json:"pangkat"`
	Jabatan          string         `json:"jabatan"`
	StatusPerjalanan string         `json:"statusPerjalanan"`
	UserId           int            `json:"userId"`
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
	Id              int              `json:"id"`
	KegiatanId      int              `json:"kegiatanId"`
	Kegiatan        KegiatanResponse `json:"kegiatan"`
	KodeSubKegiatan string           `json:"kodeSubKegiatan"`
	NamaSubKegiatan string           `json:"namaSubKegiatan"`
	PejabatId       int              `json:"pejabatId"`
	Pejabat         PejabatResponse  `json:"pejabat"`
}

type RekeningResponse struct {
	Id           int    `json:"id"`
	KodeRekening string `json:"kodeRekening"`
	NamaRekening string `json:"namaRekening"`
}

type SptResponse struct {
	Id                int                 `json:"id"`
	Jenis             string              `json:"jenis"`
	Template          string              `json:"template"`
	SubKegiatanId     int                 `json:"subKegiatanId"`
	SubKegiatan       SubKegiatanResponse `json:"subKegiatan"`
	Nomor_Spt         string              `json:"nomor_spt"`
	Tanggal_Spt       string              `json:"tanggal_spt"`
	RekeningId        int                 `json:"rekeningId"`
	Rekening          RekeningResponse    `json:"rekening"`
	Keperluan         string              `json:"keperluan"`
	Tanggal_Berangkat string              `json:"tanggal_berangkat"`
	Tanggal_Kembali   string              `json:"tanggal_kembali"`
	Lama_Perjalanan   string              `json:"lama_perjalanan"`
	Pejabat_Pemberi   string              `json:"pejabat_pemberi"`
	Status            string              `json:"status"`
	StatusSppd        int                 `json:"statusSppd"`
	File_Surat_Tugas  string              `json:"file_surat_tugas"`
	UserId            int                 `json:"userId"`
}

type SppdResponse struct {
	Id               int             `json:"id"`
	Template_Sppd    string          `json:"template_sppd"`
	Jenis            string          `json:"jenis"`
	Nomor_Sppd       string          `json:"nomor_sppd"`
	PegawaiId        int             `json:"pegawaiId"`
	Pegawai          PegawaiResponse `json:"pegawai"`
	Tanggal_Sppd     string          `json:"tanggal_sppd"`
	Tempat_Berangkat string          `json:"tempat_berangkat"`
	Tempat_Tujuan    string          `json:"tempat_tujuan"`
	Alat_Angkutan    string          `json:"alat_angkutan"`
	Instansi         string          `json:"instansi"`
	PejabatId        int             `json:"pejabatId"`
	Pejabat          PejabatResponse `json:"pejabat"`
	SptId            int             `json:"sptId"`
	Spt              SptResponse     `json:"spt"`
	UserId           int             `json:"userId"`
}

type DataDitugaskanResponse struct {
	Id         int             `json:"id"`
	SptId      int             `json:"sptId"`
	Spt        SptResponse     `json:"spt"`
	PegawaiId  int             `json:"pegawaiId"`
	Pegawai    PegawaiResponse `json:"pegawai"`
	StatusSppd int             `json:"statusSppd"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}

type DataPengikutResponse struct {
	Id        int             `json:"id"`
	SptId     int             `json:"sptId"`
	Spt       SptResponse     `json:"spt"`
	PegawaiId int             `json:"pegawaiId"`
	Pegawai   PegawaiResponse `json:"pegawai"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

type SemuaNamaDitugaskan struct {
	Id              int    `json:"id"`
	Tanggal_Kembali string `json:"tanggal_kembali"`
	Ditugaskan      string `json:"ditugaskan"`
	Status          string `json:"status"`
}
