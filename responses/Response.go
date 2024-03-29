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

type JumlahPerjalananPegawaiResponse struct {
	Nip              string `json:"nip"`
	Nama             string `json:"nama"`
	Bidang           string `json:"bidang"`
	JumlahPerjalanan int64  `json:"jumlahPerjalanan"`
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
	Tahun      string `json:"tahun"`
}

type KegiatanResponse struct {
	Id           int             `json:"id"`
	ProgramId    int             `json:"programId"`
	Program      ProgramResponse `json:"program"`
	KodeKegiatan string          `json:"kodeKegiatan"`
	NamaKegiatan string          `json:"namaKegiatan"`
	Tahun        string          `json:"tahun"`
}

type SubKegiatanResponse struct {
	Id              int              `json:"id"`
	KegiatanId      int              `json:"kegiatanId"`
	Kegiatan        KegiatanResponse `json:"kegiatan"`
	KodeSubKegiatan string           `json:"kodeSubKegiatan"`
	NamaSubKegiatan string           `json:"namaSubKegiatan"`
	PejabatId       int              `json:"pejabatId"`
	Pejabat         PejabatResponse  `json:"pejabat"`
	BidangId        int              `json:"bidangId"`
	Bidang          BidangResponse   `json:"bidang"`
	Tahun           string           `json:"tahun"`
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
	Tahun             string              `json:"tahun"`
	PejabatId         int                 `json:"pejabatId"`
	Pejabat           PejabatResponse     `json:"pejabat"`
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
	Tahun            string          `json:"tahun"`
	Alat_Angkutan    string          `json:"alat_angkutan"`
	Instansi         string          `json:"instansi"`
	PejabatId        int             `json:"pejabatId"`
	Pejabat          PejabatResponse `json:"pejabat"`
	StatusKwitansi   int             `json:"statusKwitansi"`
	SptId            int             `json:"sptId"`
	Spt              SptResponse     `json:"spt"`
	UserId           int             `json:"userId"`
}

type DataDitugaskanResponse struct {
	Id             int             `json:"id"`
	SptId          int             `json:"sptId"`
	Spt            SptResponse     `json:"spt"`
	PegawaiId      int             `json:"pegawaiId"`
	Pegawai        PegawaiResponse `json:"pegawai"`
	NamaPegawai    string          `json:"namaPegawai"`
	StatusSppd     int             `json:"statusSppd"`
	StatusKwitansi int             `json:"statusKwitansi"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

type DataDitugaskanSptResponse struct {
	Id              int       `json:"id"`
	SptId           int       `json:"sptId"`
	PegawaiId       int       `json:"pegawaiId"`
	NamaPegawai     string    `json:"namaPegawai"`
	NipPegawai      string    `json:"nipPegawai"`
	PangkatPegawai  string    `json:"pangkatPegawai"`
	GolonganPegawai string    `json:"golonganPegawai"`
	JabatanPegawai  string    `json:"jabatanPegawai"`
	StatusSppd      int       `json:"statusSppd"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
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

type KwitansiResponse struct {
	Id            int             `json:"id"`
	SppdId        int             `json:"sppdId"`
	Sppd          SppdResponse    `json:"sppd"`
	PegawaiId     int             `json:"pegawaiId"`
	Pegawai       PegawaiResponse `json:"pegawai"`
	NomorKwitansi string          `json:"nomorKwitansi"`
	TanggalBayar  string          `json:"tanggalBayar"`
	Keperluan     string          `json:"keperluan"`
	TotalBayar    int             `json:"totalBayar"`
	Tahun         string          `json:"tahun"`
	UserId        int             `json:"userId"`
}

type RincianKwitansiResponse struct {
	Id          string           `json:"id"`
	KwitansiId  int              `json:"kwitansiId"`
	Kwitansi    KwitansiResponse `json:"kwitansi"`
	Jenis       string           `json:"jenis"`
	NamaRincian string           `json:"namaRincian"`
	JumlahBayar int              `json:"jumlahBayar"`
	Banyaknya   int              `json:"banyaknya"`
	HasilBayar  int              `json:"hasilBayar"`
}

type RincianKwitansiPenerbanganResponse struct {
	Id           string           `json:"id"`
	KwitansiId   int              `json:"kwitansiId"`
	Kwitansi     KwitansiResponse `json:"kwitansi"`
	NamaMaskapai string           `json:"namaMaskapai"`
	NomorTiket   string           `json:"nomorTiket"`
	KodeBooking  string           `json:"kodeBooking"`
}

// ========================= LAPORAN RESPONSE ===================================

type RincianKwitansiPenerbangan struct {
	Id           string `json:"id"`
	KwitansiId   int    `json:"kwitansiId"`
	NamaMaskapai string `json:"namaMaskapai"`
	NomorTiket   string `json:"nomorTiket"`
	KodeBooking  string `json:"kodeBooking"`
}

type RincianKwitansi struct {
	Id         string `json:"id"`
	KwitansiId int    `json:"kwitansiId"`
	Jenis      string `json:"jenis"`
	HasilBayar string `json:"hasilBayar"`
}

type LaporanResponse struct {
	Id                               int                          `json:"id"`
	Nomor                            string                       `json:"nomor"`
	Nip                              string                       `json:"nip"`
	NamaPegawai                      string                       `json:"namaPegawai"`
	NamaKegiatan                     string                       `json:"namaKegiatan"`
	NomorSpt                         string                       `json:"nomorSpt"`
	TanggalSpt                       string                       `json:"tanggalSpt"`
	NomorSppd                        string                       `json:"nomorSppd"`
	TanggalSppd                      string                       `json:"tanggalSppd"`
	KeperluanSpt                     string                       `json:"keperluanSpt"`
	TempatTujuan                     string                       `json:"tempatTujuan"`
	LamaPerjalanan                   string                       `json:"lamaPerjalanan"`
	TanggalBerangkat                 string                       `json:"tanggalBerangkat"`
	TanggalKembali                   string                       `json:"tanggalKembali"`
	NomorKwitansi                    string                       `json:"nomorKwitansi"`
	TanggalBayar                     string                       `json:"tanggalBayar"`
	Keperluan                        string                       `json:"keperluan"`
	TotalBayar                       string                       `json:"totalBayar"`
	Tahun                            string                       `json:"tahun"`
	SumTotalBayar                    string                       `json:"sumTotalBayar"`
	RincianKwitansiUangHarian        []RincianKwitansi            `json:"rincianKwitansiUangHarian"`
	RincianKwitansiUangRepresentatif []RincianKwitansi            `json:"rincianKwitansiUangRepresentatif"`
	RincianKwitansiBiayaHotel        []RincianKwitansi            `json:"rincianKwitansiBiayaHotel"`
	RincianKwitansiBiayaTiket        []RincianKwitansi            `json:"rincianKwitansiBiayaTiket"`
	RincianKwitansiTransportBandara  []RincianKwitansi            `json:"rincianKwitansiTransportBandara"`
	RincianKwitansiPenerbangan       []RincianKwitansiPenerbangan `json:"rincianKwitansiPenerbangan"`
}
