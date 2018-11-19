package HitBTC

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type httpMock struct {
	mock.Mock
}

func (hm *httpMock) Do(req *http.Request) (*http.Response, error) {
	args := hm.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
