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

func ConvertToJumlahPerjalananPegawai(p model.Pegawai, jumlahPerjalanan int64) responses.JumlahPerjalananPegawaiResponse {
	return responses.JumlahPerjalananPegawaiResponse{
		Nip:              p.Nip,
		Nama:             p.Nama,
		Bidang:           p.Bidang.Singkatan,
		JumlahPerjalanan: jumlahPerjalanan,
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
		Tahun:      p.Tahun,
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
			Tahun:      k.Program.Tahun,
		},
		KodeKegiatan: k.KodeKegiatan,
		NamaKegiatan: k.NamaKegiatan,
		Tahun:        k.Tahun,
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
			Tahun:        s.Kegiatan.Tahun,
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
		BidangId: s.BidangId,
		Bidang: responses.BidangResponse{
			Id:          s.Bidang.Id,
			Nama_Bidang: s.Bidang.Nama_Bidang,
			Singkatan:   s.Bidang.Singkatan,
		},
		Tahun: s.Tahun,
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
				Tahun:        s.SubKegiatan.Kegiatan.Tahun,
			},
			BidangId: s.SubKegiatan.BidangId,
			Bidang: responses.BidangResponse{
				Id:          s.SubKegiatan.Bidang.Id,
				Nama_Bidang: s.SubKegiatan.Bidang.Nama_Bidang,
				Singkatan:   s.SubKegiatan.Bidang.Singkatan,
			},
			KodeSubKegiatan: s.SubKegiatan.KodeSubKegiatan,
			NamaSubKegiatan: s.SubKegiatan.NamaSubKegiatan,
			Tahun:           s.SubKegiatan.Tahun,
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
		Tahun:             s.Tahun,
		PejabatId:         s.PejabatId,
		Pejabat: responses.PejabatResponse{
			Id:       s.Pejabat.Id,
			Nip:      s.Pejabat.Nip,
			Nama:     s.Pejabat.Nama,
			Pangkat:  s.Pejabat.Pangkat,
			Golongan: s.Pejabat.Golongan,
			Eselon:   s.Pejabat.Eselon,
			Jabatan:  s.Pejabat.Jabatan,
		},
		Status:           s.Status,
		StatusSppd:       s.StatusSppd,
		File_Surat_Tugas: s.File_Surat_Tugas,
		UserId:           s.UserId,
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
		Tahun:            s.Tahun,
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
		StatusKwitansi: s.StatusKwitansi,
		SptId:          s.SptId,
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
						Tahun:      s.Spt.SubKegiatan.Kegiatan.Program.Tahun,
					},
					KodeKegiatan: s.Spt.SubKegiatan.Kegiatan.KodeKegiatan,
					NamaKegiatan: s.Spt.SubKegiatan.Kegiatan.NamaKegiatan,
					Tahun:        s.Spt.SubKegiatan.Kegiatan.Tahun,
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
				BidangId: s.Spt.SubKegiatan.BidangId,
				Bidang: responses.BidangResponse{
					Id:          s.Spt.SubKegiatan.Bidang.Id,
					Nama_Bidang: s.Spt.SubKegiatan.Bidang.Nama_Bidang,
					Singkatan:   s.Spt.SubKegiatan.Bidang.Singkatan,
				},
				Tahun: s.Spt.SubKegiatan.Tahun,
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
			Tahun:             s.Spt.Tahun,
			PejabatId:         s.Spt.PejabatId,
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
			Tahun:             d.Spt.Tahun,
			PejabatId:         d.Spt.PejabatId,
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

func ConvertToDataDitugaskanSptResponse(d model.DataDitugaskan) responses.DataDitugaskanSptResponse {
	return responses.DataDitugaskanSptResponse{
		Id:              d.Id,
		SptId:           d.SptId,
		PegawaiId:       d.PegawaiId,
		NamaPegawai:     d.Pegawai.Nama,
		NipPegawai:      d.Pegawai.Nip,
		PangkatPegawai:  d.Pegawai.Pangkat,
		GolonganPegawai: d.Pegawai.Golongan,
		JabatanPegawai:  d.Pegawai.Jabatan,
		StatusSppd:      d.StatusSppd,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
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
			Tahun:             d.Spt.Tahun,
			PejabatId:         d.Spt.PejabatId,
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

func ConvertToKwitansiResponse(k model.Kwitansi) responses.KwitansiResponse {
	return responses.KwitansiResponse{
		Id:     k.Id,
		SppdId: k.SppdId,
		Sppd: responses.SppdResponse{
			Id:            k.Sppd.Id,
			Template_Sppd: k.Sppd.Template_Sppd,
			Jenis:         k.Sppd.Jenis,
			Nomor_Sppd:    k.Sppd.Nomor_Sppd,
			PegawaiId:     k.Sppd.PegawaiId,
			Pegawai: responses.PegawaiResponse{
				Id:            k.Sppd.Pegawai.Id,
				Nip:           k.Sppd.Pegawai.Nip,
				Nama:          k.Sppd.Pegawai.Nama,
				Jenis_Kelamin: k.Sppd.Pegawai.Jenis_Kelamin,
				Status:        k.Sppd.Pegawai.Status,
				Tempat_Lahir:  k.Sppd.Pegawai.Tempat_Lahir,
				Tanggal_Lahir: k.Sppd.Pegawai.Tanggal_Lahir,
				Instansi:      k.Sppd.Pegawai.Instansi,
				BidangId:      k.Sppd.Pegawai.BidangId,
				Bidang: responses.BidangResponse{
					Id:          k.Sppd.Pegawai.Bidang.Id,
					Nama_Bidang: k.Sppd.Pegawai.Bidang.Nama_Bidang,
					Singkatan:   k.Sppd.Pegawai.Bidang.Singkatan,
				},
				Golongan:         k.Sppd.Pegawai.Golongan,
				Eselon:           k.Sppd.Pegawai.Eselon,
				Pangkat:          k.Sppd.Pegawai.Pangkat,
				Jabatan:          k.Sppd.Pegawai.Jabatan,
				StatusPerjalanan: k.Sppd.Pegawai.StatusPerjalanan,
			},
			Tanggal_Sppd:     k.Sppd.Tanggal_Sppd,
			Tempat_Berangkat: k.Sppd.Tempat_Berangkat,
			Tempat_Tujuan:    k.Sppd.Tempat_Tujuan,
			Tahun:            k.Sppd.Tahun,
			Alat_Angkutan:    k.Sppd.Alat_Angkutan,
			Instansi:         k.Sppd.Instansi,
			PejabatId:        k.Sppd.PejabatId,
			Pejabat: responses.PejabatResponse{
				Id:       k.Sppd.Pejabat.Id,
				Nip:      k.Sppd.Pejabat.Nip,
				Nama:     k.Sppd.Pejabat.Nama,
				Pangkat:  k.Sppd.Pejabat.Pangkat,
				Golongan: k.Sppd.Pejabat.Golongan,
				Eselon:   k.Sppd.Pejabat.Eselon,
				Jabatan:  k.Sppd.Pejabat.Jabatan,
			},
			StatusKwitansi: k.Sppd.StatusKwitansi,
			SptId:          k.Sppd.SptId,
			Spt: responses.SptResponse{
				Id:            k.Sppd.Spt.Id,
				Jenis:         k.Sppd.Spt.Jenis,
				Template:      k.Sppd.Spt.Template,
				SubKegiatanId: k.Sppd.Spt.SubKegiatanId,
				SubKegiatan: responses.SubKegiatanResponse{
					Id:         k.Sppd.Spt.SubKegiatan.Id,
					KegiatanId: k.Sppd.Spt.SubKegiatan.KegiatanId,
					Kegiatan: responses.KegiatanResponse{
						Id:        k.Sppd.Spt.SubKegiatan.Kegiatan.Id,
						ProgramId: k.Sppd.Spt.SubKegiatan.Kegiatan.ProgramId,
						Program: responses.ProgramResponse{
							Id:         k.Sppd.Spt.SubKegiatan.Kegiatan.Program.Id,
							Kode:       k.Sppd.Spt.SubKegiatan.Kegiatan.Program.Kode,
							Pembebanan: k.Sppd.Spt.SubKegiatan.Kegiatan.Program.Pembebanan,
							Program:    k.Sppd.Spt.SubKegiatan.Kegiatan.Program.Program,
							Tahun:      k.Sppd.Spt.SubKegiatan.Kegiatan.Program.Tahun,
						},
						KodeKegiatan: k.Sppd.Spt.SubKegiatan.Kegiatan.KodeKegiatan,
						NamaKegiatan: k.Sppd.Spt.SubKegiatan.Kegiatan.NamaKegiatan,
						Tahun:        k.Sppd.Spt.SubKegiatan.Kegiatan.Tahun,
					},
					KodeSubKegiatan: k.Sppd.Spt.SubKegiatan.KodeSubKegiatan,
					NamaSubKegiatan: k.Sppd.Spt.SubKegiatan.NamaSubKegiatan,
					PejabatId:       k.Sppd.Spt.SubKegiatan.PejabatId,
					Pejabat: responses.PejabatResponse{
						Id:       k.Sppd.Spt.SubKegiatan.Pejabat.Id,
						Nip:      k.Sppd.Spt.SubKegiatan.Pejabat.Nip,
						Nama:     k.Sppd.Spt.SubKegiatan.Pejabat.Nama,
						Pangkat:  k.Sppd.Spt.SubKegiatan.Pejabat.Pangkat,
						Golongan: k.Sppd.Spt.SubKegiatan.Pejabat.Golongan,
						Eselon:   k.Sppd.Spt.SubKegiatan.Pejabat.Eselon,
						Jabatan:  k.Sppd.Spt.SubKegiatan.Pejabat.Jabatan,
					},
					BidangId: k.Sppd.Spt.SubKegiatan.BidangId,
					Bidang: responses.BidangResponse{
						Id:          k.Sppd.Spt.SubKegiatan.Bidang.Id,
						Nama_Bidang: k.Sppd.Spt.SubKegiatan.Bidang.Nama_Bidang,
						Singkatan:   k.Sppd.Spt.SubKegiatan.Bidang.Singkatan,
					},
					Tahun: k.Sppd.Spt.SubKegiatan.Tahun,
				},
				Nomor_Spt:   k.Sppd.Spt.Nomor_Spt,
				Tanggal_Spt: k.Sppd.Spt.Tanggal_Spt,
				RekeningId:  k.Sppd.Spt.RekeningId,
				Rekening: responses.RekeningResponse{
					Id:           k.Sppd.Spt.Rekening.Id,
					KodeRekening: k.Sppd.Spt.Rekening.KodeRekening,
					NamaRekening: k.Sppd.Spt.Rekening.NamaRekening,
				},
				Keperluan:         k.Sppd.Spt.Keperluan,
				Tanggal_Berangkat: k.Sppd.Spt.Tanggal_Berangkat,
				Tanggal_Kembali:   k.Sppd.Spt.Tanggal_Kembali,
				Lama_Perjalanan:   k.Sppd.Spt.Lama_Perjalanan,
				Tahun:             k.Sppd.Spt.Tahun,
				PejabatId:         k.Sppd.Spt.PejabatId,
				Status:            k.Sppd.Spt.Status,
				StatusSppd:        k.Sppd.Spt.StatusSppd,
				File_Surat_Tugas:  k.Sppd.Spt.File_Surat_Tugas,
			},
			UserId: k.Sppd.UserId,
		},
		NomorKwitansi: k.NomorKwitansi,
		TanggalBayar:  k.TanggalBayar,
		Keperluan:     k.Keperluan,
		TotalBayar:    k.TotalBayar,
		Tahun:         k.Tahun,
		UserId:        k.UserId,
	}
}

func ConvertToRincianKwitansiResponse(r model.RincianKwitansi) responses.RincianKwitansiResponse {
	return responses.RincianKwitansiResponse{
		Id:         r.Id,
		KwitansiId: r.KwitansiId,
		Kwitansi: responses.KwitansiResponse{
			Id:     r.Kwitansi.Id,
			SppdId: r.Kwitansi.SppdId,
			Sppd: responses.SppdResponse{
				Id:            r.Kwitansi.Sppd.Id,
				Template_Sppd: r.Kwitansi.Sppd.Template_Sppd,
				Jenis:         r.Kwitansi.Sppd.Jenis,
				Nomor_Sppd:    r.Kwitansi.Sppd.Nomor_Sppd,
				PegawaiId:     r.Kwitansi.Sppd.PegawaiId,
				Pegawai: responses.PegawaiResponse{
					Id:            r.Kwitansi.Sppd.Pegawai.Id,
					Nip:           r.Kwitansi.Sppd.Pegawai.Nip,
					Nama:          r.Kwitansi.Sppd.Pegawai.Nama,
					Jenis_Kelamin: r.Kwitansi.Sppd.Pegawai.Jenis_Kelamin,
					Status:        r.Kwitansi.Sppd.Pegawai.Status,
					Tempat_Lahir:  r.Kwitansi.Sppd.Pegawai.Tempat_Lahir,
					Tanggal_Lahir: r.Kwitansi.Sppd.Pegawai.Tanggal_Lahir,
					Instansi:      r.Kwitansi.Sppd.Pegawai.Instansi,
					BidangId:      r.Kwitansi.Sppd.Pegawai.BidangId,
					Bidang: responses.BidangResponse{
						Id:          r.Kwitansi.Sppd.Pegawai.Bidang.Id,
						Nama_Bidang: r.Kwitansi.Sppd.Pegawai.Bidang.Nama_Bidang,
						Singkatan:   r.Kwitansi.Sppd.Pegawai.Bidang.Singkatan,
					},
					Golongan:         r.Kwitansi.Sppd.Pegawai.Golongan,
					Eselon:           r.Kwitansi.Sppd.Pegawai.Eselon,
					Pangkat:          r.Kwitansi.Sppd.Pegawai.Pangkat,
					Jabatan:          r.Kwitansi.Sppd.Pegawai.Jabatan,
					StatusPerjalanan: r.Kwitansi.Sppd.Pegawai.StatusPerjalanan,
				},
				Tanggal_Sppd:     r.Kwitansi.Sppd.Tanggal_Sppd,
				Tempat_Berangkat: r.Kwitansi.Sppd.Tempat_Berangkat,
				Tempat_Tujuan:    r.Kwitansi.Sppd.Tempat_Tujuan,
				Tahun:            r.Kwitansi.Sppd.Tahun,
				Alat_Angkutan:    r.Kwitansi.Sppd.Alat_Angkutan,
				Instansi:         r.Kwitansi.Sppd.Instansi,
				PejabatId:        r.Kwitansi.Sppd.PejabatId,
				Pejabat: responses.PejabatResponse{
					Id:       r.Kwitansi.Sppd.Pejabat.Id,
					Nip:      r.Kwitansi.Sppd.Pejabat.Nip,
					Nama:     r.Kwitansi.Sppd.Pejabat.Nama,
					Pangkat:  r.Kwitansi.Sppd.Pejabat.Pangkat,
					Golongan: r.Kwitansi.Sppd.Pejabat.Golongan,
					Eselon:   r.Kwitansi.Sppd.Pejabat.Eselon,
					Jabatan:  r.Kwitansi.Sppd.Pejabat.Jabatan,
				},
				StatusKwitansi: r.Kwitansi.Sppd.StatusKwitansi,
				SptId:          r.Kwitansi.Sppd.SptId,
				Spt: responses.SptResponse{
					Id:            r.Kwitansi.Sppd.Spt.Id,
					Jenis:         r.Kwitansi.Sppd.Spt.Jenis,
					Template:      r.Kwitansi.Sppd.Spt.Template,
					SubKegiatanId: r.Kwitansi.Sppd.Spt.SubKegiatanId,
					SubKegiatan: responses.SubKegiatanResponse{
						Id:         r.Kwitansi.Sppd.Spt.SubKegiatan.Id,
						KegiatanId: r.Kwitansi.Sppd.Spt.SubKegiatan.KegiatanId,
						Kegiatan: responses.KegiatanResponse{
							Id:        r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.Id,
							ProgramId: r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.ProgramId,
							Program: responses.ProgramResponse{
								Id:         r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.Program.Id,
								Kode:       r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.Program.Kode,
								Pembebanan: r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.Program.Pembebanan,
								Program:    r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.Program.Program,
								Tahun:      r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.Program.Tahun,
							},
							KodeKegiatan: r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.KodeKegiatan,
							NamaKegiatan: r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.NamaKegiatan,
							Tahun:        r.Kwitansi.Sppd.Spt.SubKegiatan.Kegiatan.Tahun,
						},
						KodeSubKegiatan: r.Kwitansi.Sppd.Spt.SubKegiatan.KodeSubKegiatan,
						NamaSubKegiatan: r.Kwitansi.Sppd.Spt.SubKegiatan.NamaSubKegiatan,
						PejabatId:       r.Kwitansi.Sppd.Spt.SubKegiatan.PejabatId,
						Pejabat: responses.PejabatResponse{
							Id:       r.Kwitansi.Sppd.Spt.SubKegiatan.Pejabat.Id,
							Nip:      r.Kwitansi.Sppd.Spt.SubKegiatan.Pejabat.Nip,
							Nama:     r.Kwitansi.Sppd.Spt.SubKegiatan.Pejabat.Nama,
							Pangkat:  r.Kwitansi.Sppd.Spt.SubKegiatan.Pejabat.Pangkat,
							Golongan: r.Kwitansi.Sppd.Spt.SubKegiatan.Pejabat.Golongan,
							Eselon:   r.Kwitansi.Sppd.Spt.SubKegiatan.Pejabat.Eselon,
							Jabatan:  r.Kwitansi.Sppd.Spt.SubKegiatan.Pejabat.Jabatan,
						},
						BidangId: r.Kwitansi.Sppd.Spt.SubKegiatan.BidangId,
						Bidang: responses.BidangResponse{
							Id:          r.Kwitansi.Sppd.Spt.SubKegiatan.Bidang.Id,
							Nama_Bidang: r.Kwitansi.Sppd.Spt.SubKegiatan.Bidang.Nama_Bidang,
							Singkatan:   r.Kwitansi.Sppd.Spt.SubKegiatan.Bidang.Singkatan,
						},
						Tahun: r.Kwitansi.Sppd.Spt.SubKegiatan.Tahun,
					},
					Nomor_Spt:   r.Kwitansi.Sppd.Spt.Nomor_Spt,
					Tanggal_Spt: r.Kwitansi.Sppd.Spt.Tanggal_Spt,
					RekeningId:  r.Kwitansi.Sppd.Spt.RekeningId,
					Rekening: responses.RekeningResponse{
						Id:           r.Kwitansi.Sppd.Spt.Rekening.Id,
						KodeRekening: r.Kwitansi.Sppd.Spt.Rekening.KodeRekening,
						NamaRekening: r.Kwitansi.Sppd.Spt.Rekening.NamaRekening,
					},
					Keperluan:         r.Kwitansi.Sppd.Spt.Keperluan,
					Tanggal_Berangkat: r.Kwitansi.Sppd.Spt.Tanggal_Berangkat,
					Tanggal_Kembali:   r.Kwitansi.Sppd.Spt.Tanggal_Kembali,
					Lama_Perjalanan:   r.Kwitansi.Sppd.Spt.Lama_Perjalanan,
					Tahun:             r.Kwitansi.Sppd.Spt.Tahun,
					PejabatId:         r.Kwitansi.Sppd.Spt.PejabatId,
					Status:            r.Kwitansi.Sppd.Spt.Status,
					StatusSppd:        r.Kwitansi.Sppd.Spt.StatusSppd,
					File_Surat_Tugas:  r.Kwitansi.Sppd.Spt.File_Surat_Tugas,
				},
				UserId: r.Kwitansi.Sppd.UserId,
			},
			NomorKwitansi: r.Kwitansi.NomorKwitansi,
			TanggalBayar:  r.Kwitansi.TanggalBayar,
			Keperluan:     r.Kwitansi.Keperluan,
			TotalBayar:    r.Kwitansi.TotalBayar,
			Tahun:         r.Kwitansi.Tahun,
			UserId:        r.Kwitansi.UserId,
		},
		NamaRincian: r.NamaRincian,
		JumlahBayar: r.JumlahBayar,
		Banyaknya:   r.Banyaknya,
		HasilBayar:  r.HasilBayar,
	}
}
