package service

import (
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
)

type PenjualanService interface {
	GetAll() ([]model.TrxPenjualan, error)
	GetById(id int) (*model.TrxPenjualan, error)
	Create(header *model.TrxPenjualan, details []model.TrxDetailPenjualan) (*model.TrxPenjualan, error)
}

type penjualanService struct {
	repo repository.PenjualanRepository
}

func NewPenjualanService(repo repository.PenjualanRepository) PenjualanService {
	return &penjualanService{repo: repo}
}

// ambil semua
func (s *penjualanService) GetAll() ([]model.TrxPenjualan, error) {
	return s.repo.GetAll()
}

// ambil per id
func (s *penjualanService) GetById(id int) (*model.TrxPenjualan, error) {
	penjualan, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if penjualan == nil {
		return nil, errors.New("produk ga ditemukan")
	}
	return penjualan, nil
}

// Create
func (s *penjualanService) Create(header *model.TrxPenjualan, details []model.TrxDetailPenjualan) (*model.TrxPenjualan, error) {
	if header.No == "" || header.PartnerID == 0 || header.Tanggal.IsZero() {
		return nil, errors.New("no, id, dan tanggal")
	}
	// Hitung subtotal per detail dan total header
	var total float64
	for i, d := range details {
		stok, err := s.repo.GetStok(d.GudangID, d.ItemID)
		if err != nil {
			return nil, err
		}
		if stok < d.Qty {
			return nil, errors.New("stok tidak mencukupi")
		}
		details[i].SubTotal = d.Harga * float64(d.Qty)
		total += details[i].SubTotal
	}
	header.SubTotal = total
	header.Total = total - header.Diskon
	return s.repo.Create(header, details)
}
