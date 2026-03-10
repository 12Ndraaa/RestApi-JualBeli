package service

import (
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
)

type MitraBisnisService interface {
	GetAll() ([]model.MitraBisnis, error)
	GetById(id int) (*model.MitraBisnis, error)
	Create(p *model.MitraBisnis) (*model.MitraBisnis, error)
	Update(p *model.MitraBisnis) (*model.MitraBisnis, error)
	Delete(id int) error
}

type mitrabisnisService struct {
	repo repository.MitraBisnisRepository
}

func NewMitraBisnisService(repo repository.MitraBisnisRepository) MitraBisnisService {
	return &mitrabisnisService{repo: repo}
}

// ambil semua
func (s *mitrabisnisService) GetAll() ([]model.MitraBisnis, error) {
	return s.repo.FindAll()
}

// ambil per id
func (s *mitrabisnisService) GetById(id int) (*model.MitraBisnis, error) {
	mitra_bisnis, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if mitra_bisnis == nil {
		return nil, errors.New("produk ga ditemukan")
	}
	return mitra_bisnis, nil
}

// cek & ubah
func (s *mitrabisnisService) Update(p *model.MitraBisnis) (*model.MitraBisnis, error) {
	if p.Kode == "" || p.Nama == "" || p.TipeMitra == "" {
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
func (s *mitrabisnisService) Create(p *model.MitraBisnis) (*model.MitraBisnis, error) {
	// Validasi: field tidak boleh kosong
	if p.Kode == "" || p.Nama == "" || p.TipeMitra == "" {
		return nil, errors.New("kode, nama, dan tipe mitra gboleh ksong")
	}
	return s.repo.Create(p)
}

// hapus
func (s *mitrabisnisService) Delete(id int) error {
	return s.repo.Delete(id)
}
