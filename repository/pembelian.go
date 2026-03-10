package repository

import (
	"database/sql"
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
)

type PembelianRepository interface {
	GetAll() ([]model.TrxPembelian, error)
	GetById(id int) (*model.TrxPembelian, error)
	Create(header *model.TrxPembelian, details []model.TrxDetailPembelian) (*model.TrxPembelian, error)
}

type pembelianRepository struct {
	db *sql.DB
}

func NewPembelianRepository(db *sql.DB) PembelianRepository {
	return &pembelianRepository{db: db}
}

// ambil semua
func (r *pembelianRepository) GetAll() ([]model.TrxPembelian, error) {
	query := `SELECT id, trxno, bp_id, tgl, diskon, subtotal, total 
	          FROM trx_pembelian ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.TrxPembelian
	for rows.Next() {
		var p model.TrxPembelian
		if err := rows.Scan(&p.ID, &p.No, &p.PartnerID, &p.Tanggal, &p.Diskon, &p.SubTotal, &p.Total); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// ambil per id
func (r *pembelianRepository) GetById(id int) (*model.TrxPembelian, error) {
	query := `SELECT id, trxno, bp_id, tgl, diskon, subtotal, total 
	          FROM trx_pembelian WHERE id = $1`
	var p model.TrxPembelian
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.No, &p.PartnerID, &p.Tanggal, &p.Diskon, &p.SubTotal, &p.Total)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *pembelianRepository) Create(header *model.TrxPembelian, details []model.TrxDetailPembelian) (*model.TrxPembelian, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // otomatis batalkan kalau gagal

	// STEP 2: INSERT header, ambil id-nya
	queryHeader := `INSERT INTO trx_pembelian (trxno, bp_id, tgl, diskon, subtotal, total)
                    VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = tx.QueryRow(queryHeader, header.No, header.PartnerID, header.Tanggal, header.Diskon, header.SubTotal, header.Total).Scan(&header.ID)
	if err != nil {
		return nil, err
	}

	queryDetail := `INSERT INTO trx_detail_pembelian (purc_id, dno, item_id, qty, harga, diskon, subtotal)
                    VALUES ($1, $2, $3, $4, $5, $6, $7)`
	for i, d := range details {
		d.PurchaseID = header.ID // hubungkan ke header
		d.DetailNum = i + 1      // nomor urut detail
		d.SubTotal = d.Harga * float64(d.Qty)
		_, err = tx.Exec(queryDetail, d.PurchaseID, d.DetailNum, d.ItemID, d.Qty, d.Harga, d.Diskon, d.SubTotal)
		if err != nil {
			return nil, err
		}

		queryStok := `INSERT INTO stok (gudang_id, item_id, qty)
                      VALUES ($1, $2, $3)
                      ON CONFLICT (gudang_id, item_id) DO UPDATE SET qty = stok.qty + $3`
		_, err = tx.Exec(queryStok, d.GudangID, d.ItemID, d.Qty) // dari input user
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return header, nil
}
