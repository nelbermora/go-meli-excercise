package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/seller"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type sellerServiceMock struct{
	mock.Mock
}

func (s *sellerServiceMock) Get(ctx context.Context, id int) (domain.Seller, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *sellerServiceMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
	args := s.Called(ctx)
	return args.Get(0).([]domain.Seller), args.Error(1)
}

func (s *sellerServiceMock) Store(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (domain.Seller, error) {
	args := s.Called(ctx, cid, companyName, address, telephone, localityID)
	if args.Get(0) != nil {
		return args.Get(0).(domain.Seller), args.Error(1)
	}
	return domain.Seller{}, args.Error(1)
}

func (s *sellerServiceMock) Update(ctx context.Context, id, cid int, companyName, address, telephone string, localityID int) (domain.Seller, error) {
	args := s.Called(ctx, id, cid, companyName, address, telephone, localityID)
	return args.Get(0).(domain.Seller), args.Error(1)
}

func (s *sellerServiceMock) Delete(ctx context.Context, id int) error {
	return s.Called(ctx, id).Error(0)
}

type mockRequest struct {
	SellerID    int    `json:"seller_id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

var sellerAPIv = "/api/v1"
func TestSeller_Create_OK(t *testing.T) {
	r := mockRequest{0,2,"test company", "test address", "1234", 1}
	body, _ := json.Marshal(r)
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	s := domain.Seller{ID: 1, CID: r.CID, CompanyName: r.CompanyName, Address: r.Address, Telephone: r.Telephone, LocalityID: r.LocalityID}
	svcMock := &sellerServiceMock{}
	svcMock.On("Store", req.Context(), r.CID, r.CompanyName, r.Address, r.Telephone, r.LocalityID).Return(s, nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	sellerHandler := NewSeller(svcMock)

	sellerHandler.Store()(c)
	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, `{"data":{"id":1,"cid":2,"company_name":"test company","address":"test address","telephone":"1234","locality_id":1}}`, rr.Body.String())
}

func TestSeller_Create_ErrValueRequired(t *testing.T) {
	r := mockRequest{0,0,"test company", "test address", "1234", 1}
	body, _ := json.Marshal(r)
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	s := domain.Seller{ID: 1, CID: r.CID, CompanyName: r.CompanyName, Address: r.Address, Telephone: r.Telephone, LocalityID: r.LocalityID}
	svcMock := &sellerServiceMock{}
	svcMock.On("Store", req.Context(), r.CID, r.CompanyName, r.Address, r.Telephone, r.LocalityID).Return(s, nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	sellerHandler := NewSeller(svcMock)

	sellerHandler.Store()(c)
	assert.Equal(t, 422, rr.Code)
	assert.Equal(t, `{"code":"unprocessable_entity","message":"cid can not be empty"}`, rr.Body.String())
}

func TestSeller_Create_ErrCIDExist(t *testing.T) {
	r := mockRequest{0,1,"test company", "test address", "1234", 1}
	body, _ := json.Marshal(r)
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	svcMock := &sellerServiceMock{}
	svcMock.On("Store", req.Context(), r.CID, r.CompanyName, r.Address, r.Telephone, r.LocalityID).Return(nil, seller.UNIQUE)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	sellerHandler := NewSeller(svcMock)

	sellerHandler.Store()(c)
	assert.Equal(t, 409, rr.Code)
	assert.Equal(t, `{"code":"conflict","message":"The cid field has already been taken."}`, rr.Body.String())
}