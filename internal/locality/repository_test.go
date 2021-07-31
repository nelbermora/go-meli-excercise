package locality

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
	myLoc := domain.Locality{
		ID:       3,
		Name:     "test Locality",
		Province: "province",
		Country:  "country",
	}

	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO localities").
		WithArgs(myLoc.ID, myLoc.Name, myLoc.Province, myLoc.Country).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	idInserted, errResult := mockedRepo.Save(ctx, myLoc)

	assert.Equal(t, myLoc.ID, idInserted)
	assert.Nil(t, errResult)
}

func TestSaveConflict(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	myLoc := domain.Locality{
		ID:       3,
		Name:     "test Locality",
		Province: "province",
		Country:  "country",
	}
	testError := errors.New("not inserted")
	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO localities").
		WithArgs(myLoc.ID, myLoc.Name, myLoc.Province, myLoc.Country).
		WillReturnError(testError)
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)
	ctx := context.Background()

	_, errResult := mockedRepo.Save(ctx, myLoc)

	assert.Equal(t, testError, errResult)

}
