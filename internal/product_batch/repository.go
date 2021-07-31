package productbatch

import (
	"context"
	"database/sql"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

const (
	insertProductBatch = `INSERT INTO product_batches
									(batch_number,current_quantity,current_temperature,due_date,initial_quantity,manufacturing_date,minimum_temperature,product_id,section_id) 
								  VALUES (?,?,?,?,?,?,?,?,?)`
	queryExists  = `SELECT id from product_batches where id = ?`
	queryGetById = `SELECT * from product_batches where id = ?`
)

type Repository interface {
	Save(ctx context.Context, l domain.ProdcutBatch) (int, error)
	Exists(ctx context.Context, id int) bool
	Get(ctx context.Context, id int) (domain.ProdcutBatch, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, l domain.ProdcutBatch) (int, error) {
	stmt, err := r.db.Prepare(insertProductBatch)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&l.BatchNumber, &l.CurrentQuantity, &l.CurrentTemp, &l.DueDate, &l.InitialQuantity, &l.ManufacturingDate, &l.MinTemperature, &l.ProductId, &l.SectionId)
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

func (r *repository) Get(ctx context.Context, id int) (domain.ProdcutBatch, error) {
	row := r.db.QueryRow(queryGetById, id)
	l := domain.ProdcutBatch{}
	err := row.Scan(&l.ID, &l.BatchNumber, &l.CurrentQuantity, &l.CurrentTemp, &l.DueDate, &l.InitialQuantity, &l.ManufacturingDate, &l.MinTemperature, &l.ProductId, &l.SectionId)
	if err != nil {
		return domain.ProdcutBatch{}, err
	}
	return l, nil
}
