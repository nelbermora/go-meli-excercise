package inboundorder

import (
	"context"
	"database/sql"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

const (
	insertInboundOrders = `INSERT INTO inbound_orders
									(order_date,order_number,employee_id,product_batch_id,warehouse_id) 
								  VALUES (?,?,?,?,?)`
)

type Repository interface {
	Save(ctx context.Context, l domain.InboundOrder) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, p domain.InboundOrder) (int, error) {
	stmt, err := r.db.Prepare(insertInboundOrders)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&p.OrderDate, &p.OrderNumber, &p.EmployeeId, &p.ProductBatchId, &p.WarehouseId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
