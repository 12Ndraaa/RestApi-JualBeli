package service

import (
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
)

type ProdukService interface {
	GetAll() ([]model.Produk, error)
	GetById(id int) (*model.Produk, error)
	Create(p *model.Produk) (*model.Produk, error)
	Update(p *model.Produk) (*model.Produk, error)
	Delete(id int) error
}

type produkService struct {
	repo repository.ProdukRepository
}

func NewProdukService(repo repository.ProdukRepository) ProdukService {
	return &produkService{repo: repo}
}

// ambil semua
func (s *produkService) GetAll() ([]model.Produk, error) {
	return s.repo.FindAll()
}

// ambil per id
func (s *produkService) GetById(id int) (*model.Produk, error) {
	produk, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if produk == nil {
		return nil, errors.New("produk ga ditemukan")
	}
	return produk, nil
}

// cek & ubah
func (s *produkService) Update(p *model.Produk) (*model.Produk, error) {
	if p.Kode == "" || p.NamaProduk == "" {
		return nil, errors.New("kode dan nama produk gboleh ksong")
	}
	result, err := s.repo.Update(p)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("produk ga ada")
	}
	return result, nil
}

// cek & simpan
func (s *produkService) Create(p *model.Produk) (*model.Produk, error) {
	// Validasi: field tidak boleh kosong
	if p.Kode == "" || p.NamaProduk == "" {
		return nil, errors.New("kode dan nama produk tidak boleh kosong")
	}
	return s.repo.Create(p)
}

// hapus
func (s *produkService) Delete(id int) error {
	return s.repo.Delete(id)
}
