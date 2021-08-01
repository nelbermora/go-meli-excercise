package buyer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Repository encapsulates the storage of a buyer.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, cardNumberID string) (domain.Buyer, error)
	Exists(ctx context.Context, cardNumberID string) bool
	ExistsById(ctx context.Context, id int) bool
	Save(ctx context.Context, b domain.Buyer) (int, error)
	Update(ctx context.Context, b domain.Buyer) error
	Delete(ctx context.Context, cardNumberID string) error
	GetPurchaseByBuyer(ctx context.Context, id int) ([]domain.Buyer, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	rows, err := r.db.Query(`SELECT * FROM "main"."buyers"`)
	if err != nil {
		return nil, err
	}

	var buyers []domain.Buyer

	for rows.Next() {
		b := domain.Buyer{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
		buyers = append(buyers, b)
	}

	return buyers, nil
}

func (r *repository) Get(ctx context.Context, cardNumberID string) (domain.Buyer, error) {

	sqlStatement := `SELECT * FROM "main"."buyers" WHERE card_number_id=$1;`
	row := r.db.QueryRow(sqlStatement, cardNumberID)
	b := domain.Buyer{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return domain.Buyer{}, err
	}

	return b, nil
}

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	sqlStatement := `SELECT card_number_id FROM "main"."buyers" WHERE card_number_id=$1;`
	row := r.db.QueryRow(sqlStatement, cardNumberID)
	err := row.Scan(&cardNumberID)
	if err != nil {
		return false
	}
	return true
}

func (r *repository) ExistsById(ctx context.Context, id int) bool {
	sqlStatement := `SELECT id FROM "main"."buyers" WHERE id=$1;`
	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}

func (r *repository) Save(ctx context.Context, b domain.Buyer) (int, error) {

	stmt, err := r.db.Prepare(`INSERT INTO "main"."buyers"("card_number_id","first_name","last_name") VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, b domain.Buyer) error {
	stmt, err := r.db.Prepare(`UPDATE "main"."buyers" SET "first_name"=?, "last_name"=?  WHERE "card_number_id"=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&b.FirstName, &b.LastName, &b.CardNumberID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("buyer not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, cardNumberID string) error {
	stmt, err := r.db.Prepare(`DELETE FROM "main"."buyers" WHERE card_number_id=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(cardNumberID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return errors.New("buyer not found")
	}

	return nil
}

func (r *repository) GetPurchaseByBuyer(ctx context.Context, id int) ([]domain.Buyer, error) {
	ProdsByPurchase := `SELECT b.*,(
								SELECT count(id)
								FROM purchase_orders
								WHERE purchase_orders.buyer_id = b.id) as cant 
						FROM buyers b
					   WHERE b.id = IFNULL(?,b.id)`
	var rows *sql.Rows
	var err error
	if id > 0 {
		rows, err = r.db.Query(ProdsByPurchase, id)
	} else {
		rows, err = r.db.Query(ProdsByPurchase, nil)
	}

	if err != nil {
		return nil, err
	}

	var buyers []domain.Buyer

	for rows.Next() {
		b := domain.Buyer{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName, &b.PurchaseOrdersCount)
		buyers = append(buyers, b)
	}

	return buyers, nil
}

