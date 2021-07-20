package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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
	if args.Get(0) != nil {
		return args.Get(0).(domain.Seller), args.Error(1)
	}
	return domain.Seller{}, args.Error(1)
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
	if args.Get(0) != nil {
		return args.Get(0).(domain.Seller), args.Error(1)
	}
	return domain.Seller{}, args.Error(1)
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

func TestSeller_GetAll_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, sellerAPIv + "/sellers", nil)
	rr := httptest.NewRecorder()
	s := []domain.Seller{{ID: 1, CID: 23, CompanyName: "test 1", Address: "a1", Telephone: "22", LocalityID: 2}, {ID: 2, CID: 33, CompanyName: "test 2", Address: "a2", Telephone: "44", LocalityID: 5}}
	svcMock := &sellerServiceMock{}
	svcMock.On("GetAll", req.Context()).Return(s, nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	sellerHandler := NewSeller(svcMock)

	sellerHandler.GetAll()(c)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, `{"data":[{"id":1,"cid":23,"company_name":"test 1","address":"a1","telephone":"22","locality_id":2},{"id":2,"cid":33,"company_name":"test 2","address":"a2","telephone":"44","locality_id":5}]}`, rr.Body.String())
}

func TestSeller_Get_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, sellerAPIv + "/sellers/1", nil)
	rr := httptest.NewRecorder()
	s := domain.Seller{ID: 1, CID: 23, CompanyName: "test 1", Address: "a1", Telephone: "22", LocalityID: 2}
	svcMock := &sellerServiceMock{}
	svcMock.On("Get", req.Context(),1).Return(s, nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	sellerHandler := NewSeller(svcMock)

	sellerHandler.Get()(c)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, `{"data":{"id":1,"cid":23,"company_name":"test 1","address":"a1","telephone":"22","locality_id":2}}`, rr.Body.String())
}

func TestSeller_Get_Err(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, sellerAPIv + "/sellers/1", nil)
	rr := httptest.NewRecorder()
	svcMock := &sellerServiceMock{}
	svcMock.On("Get", req.Context(),1).Return(nil, errors.New(""))
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	sellerHandler := NewSeller(svcMock)

	sellerHandler.Get()(c)
	assert.Equal(t, 404, rr.Code)
	assert.Equal(t, `{"code":"not_found","message":"Seller not found"}`, rr.Body.String())
}

func TestSeller_Update_OK(t *testing.T) {
	r := mockRequest{1,2,"test company", "test address", "1234", 1}
	body, _ := json.Marshal(r)
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	s := domain.Seller{ID: 1, CID: r.CID, CompanyName: r.CompanyName, Address: r.Address, Telephone: r.Telephone, LocalityID: r.LocalityID}
	svcMock := &sellerServiceMock{}
	svcMock.On("Update", req.Context(), r.SellerID, r.CID, r.CompanyName, r.Address, r.Telephone, r.LocalityID).Return(s, nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	sellerHandler := NewSeller(svcMock)

	sellerHandler.Update()(c)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, `{"id":1,"cid":2,"company_name":"test company","address":"test address","telephone":"1234","locality_id":1}`, rr.Body.String())
}

func TestSeller_Update_Err(t *testing.T) {
	r := mockRequest{1,2,"test company", "test address", "1234", 1}
	body, _ := json.Marshal(r)
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	svcMock := &sellerServiceMock{}
	svcMock.On("Update", req.Context(), r.SellerID, r.CID, r.CompanyName, r.Address, r.Telephone, r.LocalityID).Return(nil, seller.NOT_FOUND)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	sellerHandler := NewSeller(svcMock)

	sellerHandler.Update()(c)
	assert.Equal(t, 404, rr.Code)
	assert.Equal(t, `{"code":"not_found","message":"Seller not found."}`, rr.Body.String())
}

func TestSeller_Delete_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers/1", nil)
	rr := httptest.NewRecorder()
	svcMock := &sellerServiceMock{}
	svcMock.On("Delete", req.Context(), 1).Return(nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	sellerHandler := NewSeller(svcMock)

	sellerHandler.Delete()(c)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, `{"code":"ok","message":"The seller has been deleted"}`, rr.Body.String())
}

func TestSeller_Delete_Err(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers/1", nil)
	rr := httptest.NewRecorder()
	svcMock := &sellerServiceMock{}
	svcMock.On("Delete", req.Context(), 1).Return(seller.NOT_FOUND)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	sellerHandler := NewSeller(svcMock)

	sellerHandler.Delete()(c)
	assert.Equal(t, 404, rr.Code)
	assert.Equal(t, `{"code":"not_found","message":"Seller not found."}`, rr.Body.String())
}