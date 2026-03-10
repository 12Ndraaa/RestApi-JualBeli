package repository

import (
	"database/sql"

	"github.com/12Ndraaa/restapi-jualbeli/model"
)

type LaporanRepository interface {
	RekapPerFaktur() ([]model.RekapPerFaktur, error)
	RekapPerItem() ([]model.RekapPerItem, error)
	LihatStok() ([]model.StokInfo, error)
}

type laporanRepository struct {
	db *sql.DB
}

func NewLaporanRepository(db *sql.DB) LaporanRepository {
	return &laporanRepository{db: db}
}

// ambil semua
func (r *laporanRepository) RekapPerFaktur() ([]model.RekapPerFaktur, error) {
	query := `SELECT trxno, tgl, bp_id, total FROM trx_penjualan ORDER BY tgl DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.RekapPerFaktur
	for rows.Next() {
		var p model.RekapPerFaktur
		if err := rows.Scan(&p.No, &p.Tanggal, &p.PartnerID, &p.Total); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func (r *laporanRepository) RekapPerItem() ([]model.RekapPerItem, error) {
	query := `SELECT mp.nama_produk, SUM(td.qty), SUM(td.subtotal)
			  FROM trx_detail_penjualan td
			  JOIN master_produk mp ON td.item_id = mp.id
			  GROUP BY mp.nama_produk`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.RekapPerItem
	for rows.Next() {
		var p model.RekapPerItem
		if err := rows.Scan(&p.NamaItem, &p.Qty, &p.Total); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func (r *laporanRepository) LihatStok() ([]model.StokInfo, error) {
	query := `SELECT mg.nama, mp.nama_produk, s.qty
			  FROM stok s
			  JOIN master_gudang mg ON s.gudang_id = mg.id
			  JOIN master_produk mp ON s.item_id = mp.id
			  ORDER BY mg.nama`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.StokInfo
	for rows.Next() {
		var p model.StokInfo
		if err := rows.Scan(&p.Nama, &p.NamaItem, &p.Qty); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}
