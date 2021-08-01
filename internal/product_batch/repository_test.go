package productbatch

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
	batch := domain.ProdcutBatch{
		BatchNumber:       "123",
		CurrentQuantity:   1,
		InitialQuantity:   1,
		ProductId:         1,
		SectionId:         1,
		CurrentTemp:       123,
		MinTemperature:    123,
		DueDate:           time.Now(),
		ManufacturingDate: time.Now(),
	}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO product_batches").
		WithArgs(batch.BatchNumber, batch.CurrentQuantity, batch.CurrentTemp, batch.DueDate, batch.InitialQuantity, batch.ManufacturingDate, batch.MinTemperature, batch.ProductId, batch.SectionId).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	idInserted, errResult := mockedRepo.Save(ctx, batch)

	assert.Equal(t, 3, idInserted)
	assert.Nil(t, errResult)
}

func TestSaveConflict(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	batch := domain.ProdcutBatch{
		BatchNumber:       "123",
		CurrentQuantity:   1,
		InitialQuantity:   1,
		ProductId:         1,
		SectionId:         1,
		CurrentTemp:       123,
		MinTemperature:    123,
		DueDate:           time.Now(),
		ManufacturingDate: time.Now(),
	}
	testError := errors.New("not inserted")
	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO product_batches").
		WithArgs(batch.BatchNumber, batch.CurrentQuantity, batch.CurrentTemp, batch.DueDate, batch.InitialQuantity, batch.ManufacturingDate, batch.MinTemperature, batch.ProductId, batch.SectionId).
		WillReturnError(testError)
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	_, errResult := mockedRepo.Save(ctx, batch)

	assert.Equal(t, testError, errResult)

}
