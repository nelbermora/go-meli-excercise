package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
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
	return args.Get(0).(domain.Seller), args.Error(1)
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
func TestSeller_Create(t *testing.T) {
	r := mockRequest{0,2,"test company", "test address", "1234", 1}

	body, _ := json.Marshal(r)
	fmt.Println(string(body))
	req := httptest.NewRequest(http.MethodPost, sellerAPIv + "/sellers/", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	s := domain.Seller{ID: 1, CID: r.CID, CompanyName: r.CompanyName, Address: r.Address, Telephone: r.Telephone, LocalityID: r.LocalityID}
	svcMock := &sellerServiceMock{}
	svcMock.On("Store", req.Context(), r.CID, r.CompanyName, r.Address, r.Telephone, r.LocalityID).Return(s, nil)

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	sellerHandler := NewSeller(svcMock)

	sellerHandler.Store()(c)
	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, `{aa}`, rr.Body.String())

}