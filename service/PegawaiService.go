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
	FindBySearch(whereClause map[string]interface{}) ([]model.Pegawai, error)
	Create(pegawai request.CreatePegawaiRequest) (model.Pegawai, error)
	Update(id int, pegawai request.UpdatePegawaiRequest) (model.Pegawai, error)
	UpdateBatchStatusPerjalanan(updateBatchRequest request.UpdatePegawaiBatchStatusPerjalanan) error
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

func (s *pegawaiService) FindBySearch(whereClause map[string]interface{}) ([]model.Pegawai, error) {
	var pegawais, err = s.pegawaiRepository.FindBySearch(whereClause)

	return pegawais, err
}

func (s *pegawaiService) Create(pegawaiRequest request.CreatePegawaiRequest) (model.Pegawai, error) {
	var pegawai = model.Pegawai{
		Nip:              pegawaiRequest.Nip,
		Nama:             pegawaiRequest.Nama,
		Jenis_Kelamin:    pegawaiRequest.Jenis_Kelamin,
		Status:           pegawaiRequest.Status,
		Tempat_Lahir:     pegawaiRequest.Tempat_Lahir,
		Tanggal_Lahir:    pegawaiRequest.Tanggal_Lahir,
		Instansi:         pegawaiRequest.Instansi,
		BidangId:         pegawaiRequest.BidangId,
		Golongan:         pegawaiRequest.Golongan,
		Eselon:           pegawaiRequest.Eselon,
		Pangkat:          pegawaiRequest.Pangkat,
		Jabatan:          pegawaiRequest.Jabatan,
		StatusPerjalanan: "Tidak Dalam Perjalanan",
		UserId:           pegawaiRequest.UserId,
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
	pegawai.BidangId = pegawaiRequest.BidangId
	pegawai.Golongan = pegawaiRequest.Golongan
	pegawai.Eselon = pegawaiRequest.Eselon
	pegawai.Pangkat = pegawaiRequest.Pangkat
	pegawai.Jabatan = pegawaiRequest.Jabatan
	pegawai.StatusPerjalanan = pegawaiRequest.StatusPerjalanan
	pegawai.UserId = pegawaiRequest.UserId

	updatedPegawai, err := s.pegawaiRepository.Update(pegawai)

	return updatedPegawai, err
}

func (s *pegawaiService) UpdateBatchStatusPerjalanan(updateBatchRequest request.UpdatePegawaiBatchStatusPerjalanan) error {
	var err = s.pegawaiRepository.UpdateBatchStatusPerjalanan(updateBatchRequest.Id, updateBatchRequest.StatusPerjalanan)

	return err
}

func (s *pegawaiService) Delete(id int) (model.Pegawai, error) {
	var pegawai, err = s.pegawaiRepository.FindById(id)

	deletedBidang, err := s.pegawaiRepository.Delete(pegawai)

	return deletedBidang, err
}
