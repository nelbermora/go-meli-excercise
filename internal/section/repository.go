package section

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Repository encapsulates the storage of a section.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, cid int) bool
	ExistsById(ctx context.Context, sectionId int) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section) error
	Delete(ctx context.Context, id int) error
	GetProductBatchBySection(ctx context.Context, id int) ([]domain.Section, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Section, error) {
	rows, err := r.db.Query(`SELECT * FROM "main"."sections"`)
	if err != nil {
		return nil, err
	}

	var sections []domain.Section

	for rows.Next() {
		s := domain.Section{}
		_ = rows.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
		sections = append(sections, s)
	}

	return sections, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Section, error) {

	sqlStatement := `SELECT * FROM "main"."sections" WHERE id=$1;`
	row := r.db.QueryRow(sqlStatement, id)
	s := domain.Section{}
	err := row.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		return domain.Section{}, err
	}

	return s, nil
}

func (r *repository) Exists(ctx context.Context, sectionNumber int) bool {
	sqlStatement := `SELECT section_number FROM "main"."sections" WHERE section_number=$1;`
	row := r.db.QueryRow(sqlStatement, sectionNumber)
	err := row.Scan(&sectionNumber)
	if err != nil {
		return false
	}
	return true
}

func (r *repository) ExistsById(ctx context.Context, sectionId int) bool {
	sqlStatement := `SELECT id FROM "main"."sections" WHERE id=$1;`
	row := r.db.QueryRow(sqlStatement, sectionId)
	err := row.Scan(&sectionId)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Section) (int, error) {

	stmt, err := r.db.Prepare(`INSERT INTO "main"."sections"("section_number","current_temperature","minimum_temperature","current_capacity","minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id") VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Section) error {
	stmt, err := r.db.Prepare(`UPDATE "main"."sections" SET "section_number"=?, "current_temperature"=?, "minimum_temperature"=?, "current_capacity"=?, "minimum_capacity"=?, "maximum_capacity"=?, "warehouse_id"=?, "product_type_id"=?  WHERE "id"=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID, &s.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("section not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(`DELETE FROM "main"."sections" WHERE id=?`)
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
		return errors.New("section not found")
	}

	return nil
}

func (r *repository) GetProductBatchBySection(ctx context.Context, id int) ([]domain.Section, error) {
	queryProdBatchesBySection := `SELECT s.*,(
										SELECT count(id)
										FROM product_batches
										WHERE product_batches.section_id = s.id) as cant 
							  FROM sections s
							 WHERE s.id = IFNULL(?,s.ID)`
	var sections []domain.Section
	var rows *sql.Rows
	var err error
	if id > 0 {
		rows, err = r.db.Query(queryProdBatchesBySection, id)
	} else {
		rows, err = r.db.Query(queryProdBatchesBySection, nil)
	}

	for rows.Next() {
		l := domain.Section{}
		err = rows.Scan(&l.ID, &l.SectionNumber, &l.CurrentTemperature, &l.MinimumTemperature, &l.CurrentCapacity, &l.MinimumCapacity, &l.MaximumCapacity, &l.WarehouseID, &l.ProductTypeID, &l.ProductBatchesCount)
		sections = append(sections, l)
	}

	if err != nil {
		return sections, err
	}
	return sections, nil
}
