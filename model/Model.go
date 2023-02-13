package model

type User struct {
	Id       int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Role     string `gorm:"column:role;not null"`
}

type Kabupaten struct {
	Id   int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Kode string `gorm:"column:kode;not null"`
	Nama string `gorm:"column:nama;not null"`
}

type Instansi struct {
	Id            int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Instansi      string `gorm:"column:instansi;not null"`
	Kode_Instansi string `gorm:"column:kode_instansi;"`
	Singkatan     string `gorm:"column:singkatan;not null"`
}

type Bidang struct {
	Id          int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Nama_Bidang string `gorm:"column:nama_bidang;not null"`
	Singkatan   string `gorm:"column:singkatan;not null"`
}

type Pejabat struct {
	Id       int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Nip      string `gorm:"column:nip;not null"`
	Nama     string `gorm:"column:nama;not null"`
	Pangkat  string `gorm:"column:pangkat;not null"`
	Golongan string `gorm:"column:golongan;not null"`
	Eselon   string `gorm:"column:eselon;not null"`
	Jabatan  string `gorm:"column:jabatan;not null"`
}

type Pegawai struct {
	Id            int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Nip           string `gorm:"column:nip;not null"`
	Nama          string `gorm:"column:nama;not null"`
	Jenis_Kelamin string `gorm:"column:jenis_kelamin;not null"`
	Status        string `gorm:"column:status;not null"`
	Tempat_Lahir  string `gorm:"column:tempat_lahir;not null"`
	Tanggal_Lahir string `gorm:"column:tanggal_lahir;not null"`
	Instansi      string `gorm:"column:instansi;not null"`
	Bidang        string `gorm:"column:bidang;not null"`
	Golongan      string `gorm:"column:golongan;not null"`
	Eselon        string `gorm:"column:eselon;not null"`
	Pangkat       string `gorm:"column:pangkat;not null"`
	Jabatan       string `gorm:"column:jabatan;not null"`
}

type Program struct {
	Id         int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Tahun      string `gorm:"column:tahun;not null"`
	Kode       string `gorm:"column:kode;not null"`
	Pembebanan string `gorm:"column:pembebanan;not null"`
	Program    string `gorm:"column:program;not null"`
}

type Sub_Kegiatan struct {
	Id           int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Tahun        string `gorm:"column:tahun;not null"`
	Nama_Program string `gorm:"column:nama_program;not null"`
	Kode         string `gorm:"column:kode;not null"`
	Kegiatan     string `gorm:"column:kegiatan;not null"`
}

type Spt struct {
	Id                int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Template          string `gorm:"column:template;not null"`
	Nomor_Spt         string `gorm:"column:nomor_spt;not null"`
	Tanggal_Spt       string `gorm:"column:tanggal_spt;not null"`
	Ditugaskan        string `gorm:"ditugaskan;not null"`
	Jenis_Perjalanan  string `gorm:"column:jenis_perjalanan;not null"`
	Keperluan         string `gorm:"column:keperluan;not null"`
	Alat_Angkutan     string `gorm:"column:alat_angkutan;not null"`
	Tempat_Berangkat  string `gorm:"column:tempat_berangkat;not null"`
	Tempat_Tujuan     string `gorm:"column:tempat_tujuan;not null"`
	Tanggal_Berangkat string `gorm:"column:tanggal_berangkat;not null"`
	Tanggal_Kembali   string `gorm:"column:tanggal_kembali;not null"`
	Lama_Perjalanan   string `gorm:"column:lama_perjalanan;not null"`
	Pejabat_Pemberi   string `gorm:"column:pejabat_pemberi;not null"`
	Status            string `gorm:"column:status; not null"`
	File_Surat_Tugas  string `gorm:"column:file_surat_tugas"`
}

type Sppd struct {
	Id            int    `gorm:"column:id;primaryKey;not null;autoIncrement"`
	Template_Sppd string `gorm:"column:template_sppd"`
	Nomor_Sppd    string `gorm:"column:nomor_sppd"`
	Tanggal_Sppd  string `gorm:"column:tanggal_sppd"`
	Tingkat_Biaya string `gorm:"column:tingkat_biaya"`
	Instansi      string `gorm:"column:instansi"`
	Tanda_Tangan  string `gorm:"column:tanda_tangan"`
	Idspt         int    `gorm:"column:idspt"`
}

type JoinSppdSpt struct {
	Id                int    `gorm:"column:id"`
	Template_Sppd     string `gorm:"column:template_sppd"`
	Ditugaskan        string `gorm:"column:ditugaskan"`
	Nomor_Sppd        string `gorm:"column:nomor_sppd"`
	Tanggal_Sppd      string `gorm:"column:tanggal_sppd"`
	Tingkat_Biaya     string `gorm:"column:tingkat_biaya"`
	Instansi          string `gorm:"column:instansi"`
	Tanda_Tangan      string `gorm:"column:tanda_tangan"`
	Nomor_Spt         string `gorm:"column:nomor_spt"`
	Tanggal_Spt       string `gorm:"column:tanggal_spt"`
	Jenis_Perjalanan  string `gorm:"column:jenis_perjalanan"`
	Keperluan         string `gorm:"column:keperluan"`
	Alat_Angkutan     string `gorm:"column:alat_angkutan"`
	Tempat_Berangkat  string `gorm:"column:tempat_berangkat"`
	Tempat_Tujuan     string `gorm:"column:tempat_tujuan"`
	Tanggal_Berangkat string `gorm:"column:tanggal_berangkat"`
	Tanggal_Kembali   string `gorm:"column:tanggal_kembali"`
	Lama_Perjalanan   string `gorm:"column:lama_perjalanan"`
	Pejabat_Pemberi   string `gorm:"column:pejabat_pemberi"`
	Status            string `gorm:"column:status"`
	File_Surat_Tugas  string `gorm:"column:file_surat_tugas"`
}
