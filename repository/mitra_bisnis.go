package repository

import (
	"database/sql"
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
)

type MitraBisnisRepository interface {
	FindAll() ([]model.MitraBisnis, error)
	FindByID(id int) (*model.MitraBisnis, error)
	Create(p *model.MitraBisnis) (*model.MitraBisnis, error)
	Update(p *model.MitraBisnis) (*model.MitraBisnis, error)
	Delete(id int) error
}

type mitrabisnisRepository struct {
	db *sql.DB
}

func NewMitraBisnisRepository(db *sql.DB) MitraBisnisRepository {
	return &mitrabisnisRepository{db: db}
}

// ambil semua
func (r *mitrabisnisRepository) FindAll() ([]model.MitraBisnis, error) {
	query := `SELECT id, kode, nama, tipe_mitra FROM master_mitra_bisnis ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.MitraBisnis
	for rows.Next() {
		var p model.MitraBisnis
		if err := rows.Scan(&p.ID, &p.Kode, &p.Nama, &p.TipeMitra); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// ambil per id
func (r *mitrabisnisRepository) FindByID(id int) (*model.MitraBisnis, error) {
	query := `SELECT id, kode, nama, tipe_mitra FROM master_mitra_bisnis WHERE id = $1`
	var p model.MitraBisnis
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Kode, &p.Nama, &p.TipeMitra)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// simpan
func (r *mitrabisnisRepository) Create(p *model.MitraBisnis) (*model.MitraBisnis, error) {
	query := `INSERT INTO master_mitra_bisnis (kode, nama, tipe_mitra)
	          VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, p.Kode, p.Nama, p.TipeMitra).Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// ubah
func (r *mitrabisnisRepository) Update(p *model.MitraBisnis) (*model.MitraBisnis, error) {
	query := `UPDATE master_mitra_bisnis SET kode=$1, nama=$2, tipe_mitra=$3
	          WHERE id=$4 RETURNING id, kode, nama, tipe_mitra`
	err := r.db.QueryRow(query, p.Kode, p.Nama, p.TipeMitra, p.ID).
		Scan(&p.ID, &p.Kode, &p.Nama, &p.TipeMitra)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

// buang
func (r *mitrabisnisRepository) Delete(id int) error {
	query := `DELETE FROM master_mitra_bisnis WHERE id = $1`
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
