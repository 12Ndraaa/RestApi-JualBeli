package service 

import (
	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
)

type LaporanService interface {
    RekapPerFaktur() ([]model.RekapPerFaktur, error)
    RekapPerItem() ([]model.RekapPerItem, error)
    LihatStok() ([]model.StokInfo, error)
}

type laporanService struct {
    repo repository.LaporanRepository
}

func NewLaporanService(repo repository.LaporanRepository) LaporanService {
    return &laporanService{repo: repo}
}

func (s *laporanService) RekapPerFaktur() ([]model.RekapPerFaktur, error) {
    return s.repo.RekapPerFaktur()
}
func (s *laporanService) RekapPerItem() ([]model.RekapPerItem, error) {
    return s.repo.RekapPerItem()
}
func (s *laporanService) LihatStok() ([]model.StokInfo, error) {
    return s.repo.LihatStok()
}