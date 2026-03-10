package service

import (
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
)

type PembelianService interface {
	GetAll() ([]model.TrxPembelian, error)
	GetById(id int) (*model.TrxPembelian, error)
	Create(header *model.TrxPembelian, details []model.TrxDetailPembelian) (*model.TrxPembelian, error)
}

type pembelianService struct {
	repo repository.PembelianRepository
}

func NewPembelianService(repo repository.PembelianRepository) PembelianService {
	return &pembelianService{repo: repo}
}

// ambil semua
func (s *pembelianService) GetAll() ([]model.TrxPembelian, error) {
	return s.repo.GetAll()
}

// ambil per id
func (s *pembelianService) GetById(id int) (*model.TrxPembelian, error) {
	pembelian, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if pembelian == nil {
		return nil, errors.New("produk ga ditemukan")
	}
	return pembelian, nil
}

// Create
func (s *pembelianService) Create(header *model.TrxPembelian, details []model.TrxDetailPembelian) (*model.TrxPembelian, error) {
	if header.No == "" || header.PartnerID == 0 || header.Tanggal.IsZero() {
		return nil, errors.New("no, id, dan tanggal")
	}
	// Hitung subtotal per detail dan total header
	var total float64
	for i := range details {
		details[i].SubTotal = details[i].Harga * float64(details[i].Qty)
		total += details[i].SubTotal
	}
	header.SubTotal = total
	header.Total = total - header.Diskon
	return s.repo.Create(header, details)
}
