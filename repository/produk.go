package repository

import (
	"database/sql"
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
)

type ProdukRepository interface {
	FindAll() ([]model.Produk, error)
	FindByID(id int) (*model.Produk, error)
	Create(p *model.Produk) (*model.Produk, error)
	Update(p *model.Produk) (*model.Produk, error)
	Delete(id int) error
}

type produkRepository struct {
	db *sql.DB
}

func NewProdukRepository(db *sql.DB) ProdukRepository {
	return &produkRepository{db: db}
}

// ambil semua
func (r *produkRepository) FindAll() ([]model.Produk, error) {
	query := `SELECT id, kode, nama_produk FROM master_produk ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Produk
	for rows.Next() {
		var p model.Produk
		if err := rows.Scan(&p.ID, &p.Kode, &p.NamaProduk); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// ambil per id
func (r *produkRepository) FindByID(id int) (*model.Produk, error) {
	query := `SELECT id, kode, nama_produk FROM master_produk WHERE id = $1`
	var p model.Produk
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Kode, &p.NamaProduk)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// simpan
func (r *produkRepository) Create(p *model.Produk) (*model.Produk, error) {
	query := `INSERT INTO master_produk (kode, nama_produk)
	          VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, p.Kode, p.NamaProduk).Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// ubah
func (r *produkRepository) Update(p *model.Produk) (*model.Produk, error) {
	query := `UPDATE master_produk SET kode=$1, nama_produk=$2
	          WHERE id=$3 RETURNING id, kode, nama_produk`
	err := r.db.QueryRow(query, p.Kode, p.NamaProduk, p.ID).
		Scan(&p.ID, &p.Kode, &p.NamaProduk)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

// buang
func (r *produkRepository) Delete(id int) error {
	query := `DELETE FROM master_produk WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("produk tidak ditemukan")
	}
	return nil
}
