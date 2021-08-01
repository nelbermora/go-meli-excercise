package inboundorder

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
	model := domain.InboundOrder{
		OrderDate:      time.Now(),
		OrderNumber:    "123ab",
		EmployeeId:     1,
		ProductBatchId: 2,
		WarehouseId:    3,
	}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO inbound_orders").
		WithArgs(model.OrderDate, model.OrderNumber, model.EmployeeId, model.ProductBatchId, model.WarehouseId).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	idInserted, errResult := mockedRepo.Save(ctx, model)

	assert.Nil(t, errResult)
	assert.Equal(t, 3, idInserted)
}

func TestSaveConflict(t *testing.T) {
	testError := errors.New("not inserted")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	model := domain.InboundOrder{
		OrderDate:      time.Now(),
		OrderNumber:    "",
		EmployeeId:     1,
		ProductBatchId: 2,
		WarehouseId:    3,
	}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO inbound_orders").
		WithArgs(model.OrderDate, model.OrderNumber, model.EmployeeId, model.ProductBatchId, model.WarehouseId).
		WillReturnError(testError)
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	_, errResult := mockedRepo.Save(ctx, model)

	assert.Equal(t, testError, errResult)
}
