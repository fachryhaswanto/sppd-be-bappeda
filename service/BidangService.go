package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type BidangService interface {
	FindAll() ([]model.Bidang, error)
	FindById(id int) (model.Bidang, error)
	Create(bidang request.CreateBidangRequest) (model.Bidang, error)
	Update(id int, bidang request.UpdateBidangRequest) (model.Bidang, error)
	Delete(id int) (model.Bidang, error)
}

type bidangService struct {
	bidangRepository repository.BidangRepository
}

func NewBidangService(repository repository.BidangRepository) *bidangService {
	return &bidangService{repository}
}

func (s *bidangService) FindAll() ([]model.Bidang, error) {
	var dinass, err = s.bidangRepository.FindAll()

	return dinass, err
}

func (s *bidangService) FindById(id int) (model.Bidang, error) {
	var dinas, err = s.bidangRepository.FindById(id)

	return dinas, err
}

func (s *bidangService) Create(bidangRequest request.CreateBidangRequest) (model.Bidang, error) {
	var bidang = model.Bidang{
		Nama_Bidang: bidangRequest.Nama_Bidang,
		Singkatan:   bidangRequest.Singkatan,
	}

	var newBidang, err = s.bidangRepository.Create(bidang)

	return newBidang, err
}

func (s *bidangService) Update(id int, bidangRequest request.UpdateBidangRequest) (model.Bidang, error) {
	var bidang, err = s.bidangRepository.FindById(id)

	bidang.Nama_Bidang = bidangRequest.Nama_Bidang
	bidang.Singkatan = bidangRequest.Singkatan

	updatedBidang, err := s.bidangRepository.Update(bidang)

	return updatedBidang, err
}

func (s *bidangService) Delete(id int) (model.Bidang, error) {
	var bidang, err = s.bidangRepository.FindById(id)

	deletedBidang, err := s.bidangRepository.Delete(bidang)

	return deletedBidang, err
}
