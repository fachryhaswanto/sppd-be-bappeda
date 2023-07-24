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
		BidangId:      p.BidangId,
		Bidang: responses.BidangResponse{
			Id:          p.Bidang.Id,
			Nama_Bidang: p.Bidang.Nama_Bidang,
			Singkatan:   p.Bidang.Singkatan,
		},
		Golongan:         p.Golongan,
		Eselon:           p.Eselon,
		Pangkat:          p.Pangkat,
		Jabatan:          p.Jabatan,
		StatusPerjalanan: p.StatusPerjalanan,
		UserId:           p.UserId,
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

func ConvertToSubKegiatanResponse(s model.SubKegiatan) responses.SubKegiatanResponse {
	return responses.SubKegiatanResponse{
		Id:         s.Id,
		KegiatanId: s.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           s.Kegiatan.Id,
			ProgramId:    s.Kegiatan.ProgramId,
			KodeKegiatan: s.Kegiatan.KodeKegiatan,
			NamaKegiatan: s.Kegiatan.NamaKegiatan,
		},
		KodeSubKegiatan: s.KodeSubKegiatan,
		NamaSubKegiatan: s.NamaSubKegiatan,
		PejabatId:       s.PejabatId,
		Pejabat: responses.PejabatResponse{
			Id:       s.Pejabat.Id,
			Nip:      s.Pejabat.Nip,
			Nama:     s.Pejabat.Nama,
			Pangkat:  s.Pejabat.Pangkat,
			Golongan: s.Pejabat.Golongan,
			Eselon:   s.Pejabat.Eselon,
			Jabatan:  s.Pejabat.Jabatan,
		},
	}
}

func ConvertToRekeningResponse(r model.Rekening) responses.RekeningResponse {
	return responses.RekeningResponse{
		Id:           r.Id,
		KodeRekening: r.KodeRekening,
		NamaRekening: r.NamaRekening,
	}
}

func ConvertToSptResponse(s model.Spt) responses.SptResponse {
	return responses.SptResponse{
		Id:            s.Id,
		Jenis:         s.Jenis,
		Template:      s.Template,
		SubKegiatanId: s.SubKegiatanId,
		SubKegiatan: responses.SubKegiatanResponse{
			Id:         s.SubKegiatan.Id,
			KegiatanId: s.SubKegiatan.KegiatanId,
			Kegiatan: responses.KegiatanResponse{
				Id:           s.SubKegiatan.Kegiatan.Id,
				KodeKegiatan: s.SubKegiatan.Kegiatan.KodeKegiatan,
				NamaKegiatan: s.SubKegiatan.Kegiatan.NamaKegiatan,
			},
			KodeSubKegiatan: s.SubKegiatan.KodeSubKegiatan,
			NamaSubKegiatan: s.SubKegiatan.NamaSubKegiatan,
		},
		Nomor_Spt:   s.Nomor_Spt,
		Tanggal_Spt: s.Tanggal_Spt,
		RekeningId:  s.RekeningId,
		Rekening: responses.RekeningResponse{
			Id:           s.Rekening.Id,
			KodeRekening: s.Rekening.KodeRekening,
			NamaRekening: s.Rekening.NamaRekening,
		},
		Keperluan:         s.Keperluan,
		Tanggal_Berangkat: s.Tanggal_Berangkat,
		Tanggal_Kembali:   s.Tanggal_Kembali,
		Lama_Perjalanan:   s.Lama_Perjalanan,
		Pejabat_Pemberi:   s.Pejabat_Pemberi,
		Status:            s.Status,
		StatusSppd:        s.StatusSppd,
		File_Surat_Tugas:  s.File_Surat_Tugas,
		UserId:            s.UserId,
	}
}

func ConvertToSppdResponse(s model.Sppd) responses.SppdResponse {
	return responses.SppdResponse{
		Id:            s.Id,
		Template_Sppd: s.Template_Sppd,
		Jenis:         s.Jenis,
		Nomor_Sppd:    s.Nomor_Sppd,
		PegawaiId:     s.PegawaiId,
		Pegawai: responses.PegawaiResponse{
			Id:            s.Pegawai.Id,
			Nip:           s.Pegawai.Nip,
			Nama:          s.Pegawai.Nama,
			Jenis_Kelamin: s.Pegawai.Jenis_Kelamin,
			Status:        s.Pegawai.Status,
			Tempat_Lahir:  s.Pegawai.Tempat_Lahir,
			Tanggal_Lahir: s.Pegawai.Tanggal_Lahir,
			Instansi:      s.Pegawai.Instansi,
			BidangId:      s.Pegawai.BidangId,
			Bidang: responses.BidangResponse{
				Id:          s.Pegawai.Bidang.Id,
				Nama_Bidang: s.Pegawai.Bidang.Nama_Bidang,
				Singkatan:   s.Pegawai.Bidang.Singkatan,
			},
			Golongan:         s.Pegawai.Golongan,
			Eselon:           s.Pegawai.Eselon,
			Pangkat:          s.Pegawai.Pangkat,
			Jabatan:          s.Pegawai.Jabatan,
			StatusPerjalanan: s.Pegawai.StatusPerjalanan,
		},
		Tanggal_Sppd:     s.Tanggal_Sppd,
		Tempat_Berangkat: s.Tempat_Berangkat,
		Tempat_Tujuan:    s.Tempat_Tujuan,
		Alat_Angkutan:    s.Alat_Angkutan,
		Instansi:         s.Instansi,
		PejabatId:        s.PejabatId,
		Pejabat: responses.PejabatResponse{
			Id:       s.Pejabat.Id,
			Nip:      s.Pejabat.Nip,
			Nama:     s.Pejabat.Nama,
			Pangkat:  s.Pejabat.Pangkat,
			Golongan: s.Pejabat.Golongan,
			Eselon:   s.Pejabat.Eselon,
			Jabatan:  s.Pejabat.Jabatan,
		},
		SptId: s.SptId,
		Spt: responses.SptResponse{
			Id:            s.Spt.Id,
			Jenis:         s.Spt.Jenis,
			Template:      s.Spt.Template,
			SubKegiatanId: s.Spt.SubKegiatanId,
			SubKegiatan: responses.SubKegiatanResponse{
				Id:         s.Spt.SubKegiatan.Id,
				KegiatanId: s.Spt.SubKegiatan.KegiatanId,
				Kegiatan: responses.KegiatanResponse{
					Id:        s.Spt.SubKegiatan.Kegiatan.Id,
					ProgramId: s.Spt.SubKegiatan.Kegiatan.ProgramId,
					Program: responses.ProgramResponse{
						Id:         s.Spt.SubKegiatan.Kegiatan.Program.Id,
						Kode:       s.Spt.SubKegiatan.Kegiatan.Program.Kode,
						Pembebanan: s.Spt.SubKegiatan.Kegiatan.Program.Pembebanan,
						Program:    s.Spt.SubKegiatan.Kegiatan.Program.Program,
					},
					KodeKegiatan: s.Spt.SubKegiatan.Kegiatan.KodeKegiatan,
					NamaKegiatan: s.Spt.SubKegiatan.Kegiatan.NamaKegiatan,
				},
				KodeSubKegiatan: s.Spt.SubKegiatan.KodeSubKegiatan,
				NamaSubKegiatan: s.Spt.SubKegiatan.NamaSubKegiatan,
				PejabatId:       s.Spt.SubKegiatan.PejabatId,
				Pejabat: responses.PejabatResponse{
					Id:       s.Spt.SubKegiatan.Pejabat.Id,
					Nip:      s.Spt.SubKegiatan.Pejabat.Nip,
					Nama:     s.Spt.SubKegiatan.Pejabat.Nama,
					Pangkat:  s.Spt.SubKegiatan.Pejabat.Pangkat,
					Golongan: s.Spt.SubKegiatan.Pejabat.Golongan,
					Eselon:   s.Spt.SubKegiatan.Pejabat.Eselon,
					Jabatan:  s.Spt.SubKegiatan.Pejabat.Jabatan,
				},
			},
			Nomor_Spt:   s.Spt.Nomor_Spt,
			Tanggal_Spt: s.Spt.Tanggal_Spt,
			RekeningId:  s.Spt.RekeningId,
			Rekening: responses.RekeningResponse{
				Id:           s.Spt.Rekening.Id,
				KodeRekening: s.Spt.Rekening.KodeRekening,
				NamaRekening: s.Spt.Rekening.NamaRekening,
			},
			Keperluan:         s.Spt.Keperluan,
			Tanggal_Berangkat: s.Spt.Tanggal_Berangkat,
			Tanggal_Kembali:   s.Spt.Tanggal_Kembali,
			Lama_Perjalanan:   s.Spt.Lama_Perjalanan,
			Pejabat_Pemberi:   s.Spt.Pejabat_Pemberi,
			Status:            s.Spt.Status,
			StatusSppd:        s.Spt.StatusSppd,
			File_Surat_Tugas:  s.Spt.File_Surat_Tugas,
		},
		UserId: s.UserId,
	}
}

func ConvertToDataDitugaskanResponse(d model.DataDitugaskan) responses.DataDitugaskanResponse {
	return responses.DataDitugaskanResponse{
		Id:    d.Id,
		SptId: d.SptId,
		Spt: responses.SptResponse{
			Id:          d.Spt.Id,
			Jenis:       d.Spt.Jenis,
			Template:    d.Spt.Template,
			Nomor_Spt:   d.Spt.Nomor_Spt,
			Tanggal_Spt: d.Spt.Tanggal_Spt,
			RekeningId:  d.Spt.RekeningId,
			Rekening: responses.RekeningResponse{
				Id:           d.Spt.Rekening.Id,
				KodeRekening: d.Spt.Rekening.KodeRekening,
				NamaRekening: d.Spt.Rekening.NamaRekening,
			},
			Keperluan:         d.Spt.Keperluan,
			Tanggal_Berangkat: d.Spt.Tanggal_Berangkat,
			Tanggal_Kembali:   d.Spt.Tanggal_Kembali,
			Lama_Perjalanan:   d.Spt.Lama_Perjalanan,
			Pejabat_Pemberi:   d.Spt.Pejabat_Pemberi,
			Status:            d.Spt.Status,
			StatusSppd:        d.Spt.StatusSppd,
			File_Surat_Tugas:  d.Spt.File_Surat_Tugas,
		},
		PegawaiId: d.PegawaiId,
		Pegawai: responses.PegawaiResponse{
			Id:            d.Pegawai.Id,
			Nip:           d.Pegawai.Nip,
			Nama:          d.Pegawai.Nama,
			Jenis_Kelamin: d.Pegawai.Jenis_Kelamin,
			Status:        d.Pegawai.Status,
			Tempat_Lahir:  d.Pegawai.Tempat_Lahir,
			Tanggal_Lahir: d.Pegawai.Tanggal_Lahir,
			Instansi:      d.Pegawai.Instansi,
			BidangId:      d.Pegawai.BidangId,
			Bidang: responses.BidangResponse{
				Id:          d.Pegawai.Bidang.Id,
				Nama_Bidang: d.Pegawai.Bidang.Nama_Bidang,
				Singkatan:   d.Pegawai.Bidang.Singkatan,
			},
			Golongan:         d.Pegawai.Golongan,
			Eselon:           d.Pegawai.Eselon,
			Pangkat:          d.Pegawai.Pangkat,
			Jabatan:          d.Pegawai.Jabatan,
			StatusPerjalanan: d.Pegawai.StatusPerjalanan,
		},
		StatusSppd: d.StatusSppd,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}
}

func ConvertToDataPengikutResponse(d model.DataPengikut) responses.DataPengikutResponse {
	return responses.DataPengikutResponse{
		Id:    d.Id,
		SptId: d.SptId,
		Spt: responses.SptResponse{
			Id:          d.Spt.Id,
			Jenis:       d.Spt.Jenis,
			Template:    d.Spt.Template,
			Nomor_Spt:   d.Spt.Nomor_Spt,
			Tanggal_Spt: d.Spt.Tanggal_Spt,
			RekeningId:  d.Spt.RekeningId,
			Rekening: responses.RekeningResponse{
				Id:           d.Spt.Rekening.Id,
				KodeRekening: d.Spt.Rekening.KodeRekening,
				NamaRekening: d.Spt.Rekening.NamaRekening,
			},
			Keperluan:         d.Spt.Keperluan,
			Tanggal_Berangkat: d.Spt.Tanggal_Berangkat,
			Tanggal_Kembali:   d.Spt.Tanggal_Kembali,
			Lama_Perjalanan:   d.Spt.Lama_Perjalanan,
			Pejabat_Pemberi:   d.Spt.Pejabat_Pemberi,
			Status:            d.Spt.Status,
			StatusSppd:        d.Spt.StatusSppd,
			File_Surat_Tugas:  d.Spt.File_Surat_Tugas,
		},
		PegawaiId: d.PegawaiId,
		Pegawai: responses.PegawaiResponse{
			Id:            d.Pegawai.Id,
			Nip:           d.Pegawai.Nip,
			Nama:          d.Pegawai.Nama,
			Jenis_Kelamin: d.Pegawai.Jenis_Kelamin,
			Status:        d.Pegawai.Status,
			Tempat_Lahir:  d.Pegawai.Tempat_Lahir,
			Tanggal_Lahir: d.Pegawai.Tanggal_Lahir,
			Instansi:      d.Pegawai.Instansi,
			BidangId:      d.Pegawai.BidangId,
			Bidang: responses.BidangResponse{
				Id:          d.Pegawai.Bidang.Id,
				Nama_Bidang: d.Pegawai.Bidang.Nama_Bidang,
				Singkatan:   d.Pegawai.Bidang.Singkatan,
			},
			Golongan:         d.Pegawai.Golongan,
			Eselon:           d.Pegawai.Eselon,
			Pangkat:          d.Pegawai.Pangkat,
			Jabatan:          d.Pegawai.Jabatan,
			StatusPerjalanan: d.Pegawai.StatusPerjalanan,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func ConvertToSemuaNamaDitugaskan(s model.Spt) responses.SemuaNamaDitugaskan {
	return responses.SemuaNamaDitugaskan{
		Id:              s.Id,
		Tanggal_Kembali: s.Tanggal_Kembali,
		Status:          s.Status,
	}
}
