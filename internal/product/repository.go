package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.Query(`SELECT * FROM "main"."products"`)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {

	sqlStatement := `SELECT * FROM "main"."products" WHERE id=$1;`
	row := r.db.QueryRow(sqlStatement, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Exists(ctx context.Context, productCode string) bool {
	sqlStatement := `SELECT product_code FROM "main"."products" WHERE product_code=$1;`
	row := r.db.QueryRow(sqlStatement, productCode)
	err := row.Scan(&productCode)
	if err != nil {
		return false
	}
	return true
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {

	stmt, err := r.db.Prepare(`INSERT INTO "main"."products"("description","expiration_rate","freezing_rate","height","lenght","netweight","product_code","recommended_freezing_temperature","width","id_product_type","id_seller") VALUES (?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	stmt, err := r.db.Prepare(`UPDATE "main"."products" SET "description"=?, "expiration_rate"=?, "freezing_rate"=?, "height"=?, "lenght"=?, "netweight"=?, "product_code"=?, "recommended_freezing_temperature"=?, "width"=?, "id_product_type"=?, "id_seller"=?  WHERE "id"=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID, p.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("product not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(`DELETE FROM "main"."products" WHERE id=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return errors.New("product not found")
	}

	return nil
}
