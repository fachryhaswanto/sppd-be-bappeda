package helper

import (
	"sppd/model"
	"sppd/responses"
)

func ConvertToUserResponse(u model.User) responses.UserResponse {

	return responses.UserResponse{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}

}

func ConvertToKabupatenResponse(k model.Kabupaten) responses.KabupatenResponse {

	return responses.KabupatenResponse{
		Id:   k.Id,
		Kode: k.Kode,
		Nama: k.Nama,
	}

}

func ConvertToInstansiResponse(d model.Instansi) responses.InstansiResponse {

	return responses.InstansiResponse{
		Id:            d.Id,
		Instansi:      d.Instansi,
		Kode_Instansi: d.Kode_Instansi,
		Singkatan:     d.Singkatan,
	}
}

func ConvertToBidangResponse(b model.Bidang) responses.BidangResponse {

	return responses.BidangResponse{
		Id:          b.Id,
		Nama_Bidang: b.Nama_Bidang,
		Singkatan:   b.Singkatan,
	}

}

func ConvertToPejabatResponse(p model.Pejabat) responses.PejabatResponse {
	return responses.PejabatResponse{
		Id:       p.Id,
		Nip:      p.Nip,
		Nama:     p.Nama,
		Pangkat:  p.Pangkat,
		Golongan: p.Golongan,
		Eselon:   p.Eselon,
		Jabatan:  p.Jabatan,
	}
}

func ConvertToPegawaiResponse(p model.Pegawai) responses.PegawaiResponse {

	return responses.PegawaiResponse{
		Id:            p.Id,
		Nip:           p.Nip,
		Nama:          p.Nama,
		Jenis_Kelamin: p.Jenis_Kelamin,
		Status:        p.Status,
		Tempat_Lahir:  p.Tempat_Lahir,
		Tanggal_Lahir: p.Tanggal_Lahir,
		Instansi:      p.Instansi,
		Bidang:        p.Bidang,
		Golongan:      p.Golongan,
		Eselon:        p.Eselon,
		Pangkat:       p.Pangkat,
		Jabatan:       p.Jabatan,
	}

}

func ConvertToPegawaiByNameResponse(p model.Pegawai) responses.PegawaiByNameResponse {

	return responses.PegawaiByNameResponse{
		Nip:      p.Nip,
		Nama:     p.Nama,
		Golongan: p.Golongan,
		Eselon:   p.Eselon,
		Pangkat:  p.Pangkat,
		Jabatan:  p.Jabatan,
	}

}

func ConvertToProgramResponse(p model.Program) responses.ProgramResponse {
	return responses.ProgramResponse{
		Id:         p.Id,
		Kode:       p.Kode,
		Pembebanan: p.Pembebanan,
		Program:    p.Program,
	}
}

func ConvertToKegiatanResponse(k model.Kegiatan) responses.KegiatanResponse {
	return responses.KegiatanResponse{
		Id:        k.Id,
		ProgramId: k.ProgramId,
		Program: responses.ProgramResponse{
			Id:         k.Program.Id,
			Kode:       k.Program.Kode,
			Pembebanan: k.Program.Pembebanan,
			Program:    k.Program.Program,
		},
		KodeKegiatan: k.KodeKegiatan,
		NamaKegiatan: k.NamaKegiatan,
	}
}

func ConvertToSubKegiatanResponse(s model.Sub_Kegiatan) responses.SubKegiatanResponse {
	return responses.SubKegiatanResponse{
		Id:           s.Id,
		Tahun:        s.Tahun,
		Nama_Program: s.Nama_Program,
		Kode:         s.Kode,
		Kegiatan:     s.Kegiatan,
	}
}

func ConvertToSptResponse(s model.Spt) responses.SptResponse {
	return responses.SptResponse{
		Id:                s.Id,
		Template:          s.Template,
		Nomor_Spt:         s.Nomor_Spt,
		Tanggal_Spt:       s.Tanggal_Spt,
		Ditugaskan:        s.Ditugaskan,
		Jenis_Perjalanan:  s.Jenis_Perjalanan,
		Keperluan:         s.Keperluan,
		Alat_Angkutan:     s.Alat_Angkutan,
		Tempat_Berangkat:  s.Tempat_Berangkat,
		Tempat_Tujuan:     s.Tempat_Tujuan,
		Tanggal_Berangkat: s.Tanggal_Berangkat,
		Tanggal_Kembali:   s.Tanggal_Kembali,
		Lama_Perjalanan:   s.Lama_Perjalanan,
		Pejabat_Pemberi:   s.Pejabat_Pemberi,
		Status:            s.Status,
		File_Surat_Tugas:  s.File_Surat_Tugas,
	}
}

func ConvertToSppdResponse(s model.Sppd) responses.SppdResponse {
	return responses.SppdResponse{
		Id:            s.Id,
		Template_Sppd: s.Template_Sppd,
		Nomor_Sppd:    s.Nomor_Sppd,
		Tanggal_Sppd:  s.Tanggal_Sppd,
		Tingkat_Biaya: s.Tingkat_Biaya,
		Instansi:      s.Instansi,
		Tanda_Tangan:  s.Tanda_Tangan,
		Idspt:         s.Idspt,
	}
}

func ConvertToJoinResponse(s model.JoinSppdSpt) responses.JoinResponse {
	return responses.JoinResponse{
		Id:                s.Id,
		Template_Sppd:     s.Template_Sppd,
		Ditugaskan:        s.Ditugaskan,
		Nomor_Sppd:        s.Nomor_Sppd,
		Tanggal_Sppd:      s.Tanggal_Sppd,
		Tingkat_Biaya:     s.Tingkat_Biaya,
		Instansi:          s.Instansi,
		Tanda_Tangan:      s.Tanda_Tangan,
		Nomor_Spt:         s.Nomor_Spt,
		Tanggal_Spt:       s.Tanggal_Spt,
		Jenis_Perjalanan:  s.Jenis_Perjalanan,
		Keperluan:         s.Keperluan,
		Alat_Angkutan:     s.Alat_Angkutan,
		Tempat_Berangkat:  s.Tempat_Berangkat,
		Tempat_Tujuan:     s.Tempat_Tujuan,
		Tanggal_Berangkat: s.Tanggal_Berangkat,
		Tanggal_Kembali:   s.Tanggal_Kembali,
		Lama_Perjalanan:   s.Lama_Perjalanan,
		Pejabat_Pemberi:   s.Pejabat_Pemberi,
		Status:            s.Status,
	}
}

func ConvertToSemuaNamaDitugaskan(s model.Spt) responses.SemuaNamaDitugaskan {
	return responses.SemuaNamaDitugaskan{
		Id:              s.Id,
		Tanggal_Kembali: s.Tanggal_Kembali,
		Ditugaskan:      s.Ditugaskan,
		Status:          s.Status,
	}
}
