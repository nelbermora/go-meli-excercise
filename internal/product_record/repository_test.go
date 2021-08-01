package productrecord

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	record := domain.ProductRecord{
		LastUpdate:    time.Now(),
		PurchasePrice: 123,
		SalePrice:     123,
		ProductId:     2,
	}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO product_records").
		WithArgs(record.LastUpdate, record.PurchasePrice, record.SalePrice, record.ProductId).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	idInserted, errResult := mockedRepo.Save(ctx, record)

	assert.Equal(t, 3, idInserted)
	assert.Nil(t, errResult)
}

func TestSaveConflict(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	record := domain.ProductRecord{
		LastUpdate:    time.Now(),
		PurchasePrice: 123,
		SalePrice:     123,
		ProductId:     2,
	}
	testError := errors.New("not inserted")
	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO product_records").
		WithArgs(record.LastUpdate, record.PurchasePrice, record.SalePrice, record.ProductId).
		WillReturnError(testError)
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	_, errResult := mockedRepo.Save(ctx, record)

	assert.Equal(t, testError, errResult)

}
