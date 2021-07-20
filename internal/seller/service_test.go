package seller

import (
	"context"
	"errors"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type sellerRepoMock struct{
	mock.Mock
}

func (r *sellerRepoMock) Get(ctx context.Context, id int) (domain.Seller, error) {
	args := r.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(domain.Seller), args.Error(1)
	}
	return domain.Seller{}, args.Error(1)
}

func (r *sellerRepoMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
	args := r.Called(ctx)
	return args.Get(0).([]domain.Seller), args.Error(1)
}

func (r *sellerRepoMock) Exists(ctx context.Context,  cid int) bool {
	args := r.Called(ctx, cid)
	return args.Bool(0)
}

func (r *sellerRepoMock) Save(ctx context.Context, s domain.Seller) (int, error) {
	args := r.Called(ctx, s)
	return args.Int(0), args.Error(1)
}

func (r *sellerRepoMock) Update(ctx context.Context, s domain.Seller) error {
	args := r.Called(ctx, s)
	return args.Error(0)
}

func (r *sellerRepoMock) Delete(ctx context.Context, id int) error {
	return r.Called(ctx, id).Error(0)
}

func TestService_Create_OK(t *testing.T) {
	s := domain.Seller{CID: 20, CompanyName: "company example", Address: "ad 12", Telephone: "123", LocalityID: 2}
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Save", ctx, s).Return(20, nil)
	repoMock.On("Exists", ctx, s.CID).Return(false)
	sellerService := NewService(repoMock)

	seller, err := sellerService.Store(ctx, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)

	assert.Nil(t, err)
	assert.Equal(t, 20, seller.ID)
}

func TestService_Create_ErrCIDExist(t *testing.T) {
	s := domain.Seller{CID: 20, CompanyName: "company example", Address: "ad 12", Telephone: "123", LocalityID: 2}
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Save", ctx, s).Return(20, nil)
	repoMock.On("Exists", ctx, s.CID).Return(true)
	sellerService := NewService(repoMock)

	seller, err := sellerService.Store(ctx, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "The cid field has already been taken.")
	assert.Equal(t, 0, seller.ID)
}


func TestSeller_GetAll_OK(t *testing.T) {
	s := []domain.Seller{{ID: 1, CID: 23, CompanyName: "test 1", Address: "a1", Telephone: "22", LocalityID: 2}, {ID: 2, CID: 33, CompanyName: "test 2", Address: "a2", Telephone: "44", LocalityID: 5}}
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("GetAll", ctx).Return(s, nil)
	sellerService := NewService(repoMock)

	seller, err := sellerService.GetAll(ctx)

	assert.Nil(t, err)
	assert.Equal(t, s, seller)
}


func TestSeller_Get_OK(t *testing.T) {
	s := domain.Seller{ID: 1, CID: 23, CompanyName: "test 1", Address: "a1", Telephone: "22", LocalityID: 2}
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Get", ctx, 1).Return(s, nil)
	sellerService := NewService(repoMock)

	seller, err := sellerService.Get(ctx, 1)

	assert.Nil(t, err)
	assert.Equal(t, s, seller)
}

func TestSeller_Get_Err(t *testing.T) {
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Get", ctx, 1).Return(nil, errors.New("Has been an error in TestSeller_Get_Err"))
	sellerService := NewService(repoMock)

	seller, err := sellerService.Get(ctx, 1)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Has been an error in TestSeller_Get_Err")
	assert.Equal(t, 0, seller.ID)
}

func TestSeller_Update_OK(t *testing.T) {
	s := domain.Seller{ID: 1, CID: 23, CompanyName: "test 1", Address: "a1", Telephone: "22", LocalityID: 2}
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Update",ctx, s).Return( nil)
	repoMock.On("Exists", ctx, s.CID).Return(false)
	sellerService := NewService(repoMock)

	seller, err := sellerService.Update(ctx, s.ID,s.CID,s.CompanyName, s.Address, s.Telephone, s.LocalityID)

	assert.Nil(t, err)
	assert.Equal(t, s, seller)
}

func TestSeller_Update_Err(t *testing.T) {
	s := domain.Seller{ID: 1, CID: 23, CompanyName: "test 1", Address: "a1", Telephone: "22", LocalityID: 2}
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Update", ctx, s).Return(NOT_FOUND)
	repoMock.On("Exists", ctx, s.CID).Return(false)

	sellerService := NewService(repoMock)

	_, err := sellerService.Update(ctx, s.ID,s.CID,s.CompanyName, s.Address, s.Telephone, s.LocalityID)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Seller not found.")
}

func TestSeller_Delete_OK(t *testing.T) {
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Delete", ctx, 1).Return(nil)
	sellerService := NewService(repoMock)

	err := sellerService.Delete(ctx, 1)

	assert.Nil(t, err)
}

func TestSeller_Delete_Err(t *testing.T) {
	repoMock := &sellerRepoMock{}
	ctx := context.Background()
	repoMock.On("Delete", ctx, 1).Return(NOT_FOUND)
	sellerService := NewService(repoMock)

	err := sellerService.Delete(ctx, 1)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Seller not found.")
}