package repository

import (
	"database/sql"
	"errors"

	"github.com/12Ndraaa/restapi-jualbeli/model"
)

type GudangRepository interface {
	FindAll() ([]model.Gudang, error)
	FindByID(id int) (*model.Gudang, error)
	Create(p *model.Gudang) (*model.Gudang, error)
	Update(p *model.Gudang) (*model.Gudang, error)
	Delete(id int) error
}

type gudangRepository struct {
	db *sql.DB
}

func NewGudangRepository(db *sql.DB) GudangRepository {
	return &gudangRepository{db: db}
}

// ambil semua
func (r *gudangRepository) FindAll() ([]model.Gudang, error) {
	query := `SELECT id, kode, nama FROM master_gudang ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Gudang
	for rows.Next() {
		var p model.Gudang
		if err := rows.Scan(&p.ID, &p.Kode, &p.Nama); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// ambil per id
func (r *gudangRepository) FindByID(id int) (*model.Gudang, error) {
	query := `SELECT id, kode, nama FROM master_gudang WHERE id = $1`
	var p model.Gudang
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Kode, &p.Nama)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// simpan
func (r *gudangRepository) Create(p *model.Gudang) (*model.Gudang, error) {
	query := `INSERT INTO master_gudang (kode, nama)
	          VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, p.Kode, p.Nama).Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// ubah
func (r *gudangRepository) Update(p *model.Gudang) (*model.Gudang, error) {
	query := `UPDATE master_gudang SET kode=$1, nama=$2
	          WHERE id=$3 RETURNING id, kode, nama`
	err := r.db.QueryRow(query, p.Kode, p.Nama, p.ID).
		Scan(&p.ID, &p.Kode, &p.Nama)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

// buang
func (r *gudangRepository) Delete(id int) error {
	query := `DELETE FROM master_gudang WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gudang tidak ditemukan")
	}
	return nil
}
