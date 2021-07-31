package carry

import (
	"context"
	"database/sql"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

const (
	insertCarry = `INSERT INTO carries(id,cid,batch_number,company_name,address,telephone,locality_id) VALUES (?,?,?,?,?,?,?)`
	queryExists = `SELECT id from carries where id = ?`
)

type Repository interface {
	Save(ctx context.Context, c domain.Carry) (int, error)
	Exists(ctx context.Context, id int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, c domain.Carry) (int, error) {
	stmt, err := r.db.Prepare(insertCarry)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&c.ID, &c.Cid, &c.Batch, &c.Company, &c.Address, &c.Telephone, &c.Locality)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(queryExists, id)
	err := row.Scan(&id)
	return err == nil
}
