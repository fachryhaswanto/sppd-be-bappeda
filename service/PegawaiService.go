package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type PegawaiService interface {
	FindAll() ([]model.Pegawai, error)
	FindById(id int) (model.Pegawai, error)
	FindByName(name string) (model.Pegawai, error)
	Create(pegawai request.CreatePegawaiRequest) (model.Pegawai, error)
	Update(id int, pegawai request.UpdatePegawaiRequest) (model.Pegawai, error)
	Delete(id int) (model.Pegawai, error)
}

type pegawaiService struct {
	pegawaiRepository repository.PegawaiRepository
}

func NewPegawaiService(pegawaiRepository repository.PegawaiRepository) *pegawaiService {
	return &pegawaiService{pegawaiRepository}
}

func (s *pegawaiService) FindAll() ([]model.Pegawai, error) {
	var pegawais, err = s.pegawaiRepository.FindAll()

	return pegawais, err
}

func (s *pegawaiService) FindById(id int) (model.Pegawai, error) {
	var pegawai, err = s.pegawaiRepository.FindById(id)

	return pegawai, err
}

func (s *pegawaiService) FindByName(name string) (model.Pegawai, error) {
	var pegawai, err = s.pegawaiRepository.FindByName(name)

	return pegawai, err
}

func (s *pegawaiService) Create(pegawaiRequest request.CreatePegawaiRequest) (model.Pegawai, error) {
	var pegawai = model.Pegawai{
		Nip:           pegawaiRequest.Nip,
		Nama:          pegawaiRequest.Nama,
		Jenis_Kelamin: pegawaiRequest.Jenis_Kelamin,
		Status:        pegawaiRequest.Status,
		Tempat_Lahir:  pegawaiRequest.Tempat_Lahir,
		Tanggal_Lahir: pegawaiRequest.Tanggal_Lahir,
		Instansi:      pegawaiRequest.Instansi,
		Bidang:        pegawaiRequest.Bidang,
		Golongan:      pegawaiRequest.Golongan,
		Eselon:        pegawaiRequest.Eselon,
		Pangkat:       pegawaiRequest.Pangkat,
		Jabatan:       pegawaiRequest.Jabatan,
	}

	var newPegawai, err = s.pegawaiRepository.Create(pegawai)

	return newPegawai, err
}

func (s *pegawaiService) Update(id int, pegawaiRequest request.UpdatePegawaiRequest) (model.Pegawai, error) {
	var pegawai, err = s.pegawaiRepository.FindById(id)

	pegawai.Nip = pegawaiRequest.Nip
	pegawai.Nama = pegawaiRequest.Nama
	pegawai.Jenis_Kelamin = pegawaiRequest.Jenis_Kelamin
	pegawai.Status = pegawaiRequest.Status
	pegawai.Tempat_Lahir = pegawaiRequest.Tempat_Lahir
	pegawai.Tanggal_Lahir = pegawaiRequest.Tanggal_Lahir
	pegawai.Instansi = pegawaiRequest.Instansi
	pegawai.Bidang = pegawaiRequest.Bidang
	pegawai.Golongan = pegawaiRequest.Golongan
	pegawai.Eselon = pegawaiRequest.Eselon
	pegawai.Pangkat = pegawaiRequest.Pangkat
	pegawai.Jabatan = pegawaiRequest.Jabatan

	updatedPegawai, err := s.pegawaiRepository.Update(pegawai)

	return updatedPegawai, err
}

func (s *pegawaiService) Delete(id int) (model.Pegawai, error) {
	var pegawai, err = s.pegawaiRepository.FindById(id)

	deletedBidang, err := s.pegawaiRepository.Delete(pegawai)

	return deletedBidang, err
}
