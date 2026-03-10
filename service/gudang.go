package service

import (
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
)

type GudangService interface {
	GetAll() ([]model.Gudang, error)
	GetById(id int) (*model.Gudang, error)
	Create(p *model.Gudang) (*model.Gudang, error)
	Update(p *model.Gudang) (*model.Gudang, error)
	Delete(id int) error
}

type gudangService struct {
	repo repository.GudangRepository
}

func NewGudangService(repo repository.GudangRepository) GudangService {
	return &gudangService{repo: repo}
}

// ambil semua
func (s *gudangService) GetAll() ([]model.Gudang, error) {
	return s.repo.FindAll()
}

// ambil per id
func (s *gudangService) GetById(id int) (*model.Gudang, error) {
	gudang, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if gudang == nil {
		return nil, errors.New("produk ga ditemukan")
	}
	return gudang, nil
}

// cek & ubah
func (s *gudangService) Update(p *model.Gudang) (*model.Gudang, error) {
	if p.Kode == "" || p.Nama == "" {
		return nil, errors.New("kode, nama, dan tipe mitra gboleh ksong")
	}
	result, err := s.repo.Update(p)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("mitra bisnis ga ada")
	}
	return result, nil
}

// cek & simpan
func (s *gudangService) Create(p *model.Gudang) (*model.Gudang, error) {
	// Validasi: field tidak boleh kosong
	if p.Kode == "" || p.Nama == "" {
		return nil, errors.New("kode, nama, dan tipe mitra gboleh ksong")
	}
	return s.repo.Create(p)
}

// hapus
func (s *gudangService) Delete(id int) error {
	return s.repo.Delete(id)
}
