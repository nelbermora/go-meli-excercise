package productrecord

import (
	"context"
	"database/sql"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

const (
	insertProductRecord = `INSERT INTO product_records
									(last_update_date,purchase_price,sale_price,product_id) 
								  VALUES (?,?,?,?)`
)

type Repository interface {
	Save(ctx context.Context, l domain.ProductRecord) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, l domain.ProductRecord) (int, error) {
	stmt, err := r.db.Prepare(insertProductRecord)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&l.LastUpdate, &l.PurchasePrice, &l.SalePrice, &l.ProductId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
