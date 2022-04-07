package service

import (
	"ilanver/internal/repository"
	"ilanver/request"
)

type IAddressService interface {
	Update(request request.UpdateAddress) error
}

type AddressService struct {
	RepoAddress repository.IAddressRepo
	Repository  repository.IRepository
}

func NewAddressService(repo repository.IAddressRepo, repository repository.IRepository) IAddressService {
	return &AddressService{
		RepoAddress: repo,
		Repository:  repository,
	}
}

func (s *AddressService) Update(request request.UpdateAddress) error {

	address, err := s.RepoAddress.GetByID(request.ID)
	if err != nil {
		return err
	}

	address.Districtfk = uint(request.District)
	address.Detail = request.Detail

	err = s.RepoAddress.Update(&address)

	return err
}
