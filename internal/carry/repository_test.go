package carry

import (
	"context"
	"errors"
	"testing"

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
	c := domain.Carry{
		ID:        3,
		Cid:       "cid#44",
		Company:   "Company Test",
		Address:   "Baker Avenue",
		Batch:     1,
		Telephone: "4112-123123",
		Locality:  1,
	}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO carries").
		WithArgs(c.ID, c.Cid, c.Batch, c.Company, c.Address, c.Telephone, c.Locality).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	idInserted, errResult := mockedRepo.Save(ctx, c)

	assert.Equal(t, c.ID, idInserted)
	assert.Nil(t, errResult)
}

func TestSaveConflict(t *testing.T) {
	testError := errors.New("not inserted")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	c := domain.Carry{
		ID:        3,
		Cid:       "cid#44",
		Company:   "Company Test",
		Address:   "Baker Avenue",
		Batch:     1,
		Telephone: "4112-123123",
		Locality:  1,
	}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO carries").
		WithArgs(c.ID, c.Cid, c.Batch, c.Company, c.Address, c.Telephone, c.Locality).
		WillReturnError(testError)
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	_, errResult := mockedRepo.Save(ctx, c)

	assert.Equal(t, testError, errResult)

}
