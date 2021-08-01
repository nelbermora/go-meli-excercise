package purchaseorder

import (
	"context"
	"database/sql"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

const (
	insertPurchaseOrders = `INSERT INTO purchase_orders
									(order_number,order_date,order_date,buyer_id,product_record_id,order_status_id) 
								  VALUES (?,?,?,?,?,?)`
)

type Repository interface {
	Save(ctx context.Context, l domain.PurchaseOrder) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, p domain.PurchaseOrder) (int, error) {
	stmt, err := r.db.Prepare(insertPurchaseOrders)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&p.OrderNumber, &p.OrderNumber, &p.TrackingCode, &p.BuyerId, &p.ProductRecordId, &p.OrderStatusId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
