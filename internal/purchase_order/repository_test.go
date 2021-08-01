package purchaseorder

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
	model := domain.PurchaseOrder{
		ID:              1,
		OrderNumber:     "order#1",
		OrderDate:       time.Now(),
		TrackingCode:    "abs23",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   2,
	}
	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO purchase_orders").
		WithArgs(model.OrderNumber, model.OrderDate, model.TrackingCode, model.BuyerId, model.ProductRecordId, model.OrderStatusId).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	idInserted, errResult := mockedRepo.Save(ctx, model)

	assert.Nil(t, errResult)
	assert.Equal(t, 3, idInserted)
}

func TestSaveConflict(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	model := domain.PurchaseOrder{
		ID:              1,
		OrderNumber:     "",
		OrderDate:       time.Now(),
		TrackingCode:    "abs23",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   2,
	}
	testError := errors.New("not inserted")
	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO purchase_orders").
		WithArgs(model.OrderNumber, model.OrderDate, model.TrackingCode, model.BuyerId, model.ProductRecordId, model.OrderStatusId).
		WillReturnError(testError)
	mock.ExpectCommit()

	mockedRepo := NewRepository(db)
	ctx := context.Background()

	_, errResult := mockedRepo.Save(ctx, model)

	assert.Equal(t, testError, errResult)
}
