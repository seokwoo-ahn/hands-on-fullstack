package rest

import (
	"encoding/json"
	"errors"
	"hands-on/RestfulServer/backend/src/dblayer"
	"hands-on/RestfulServer/backend/src/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockdbLayer := dblayer.NewMockDBLayerWithData()
	h := NewHandlerWithDB(mockdbLayer)
	const productsURL string = "/products"
	type errMSG struct {
		Error string `json:"error"`
	}
	tests := []struct {
		name             string
		inErr            error
		outStatusCode    int
		expectedRespBody interface{}
	}{
		{
			"getproductsnoerrors",
			nil,
			http.StatusOK,
			mockdbLayer.GetMockProductData(),
		},
		{
			"getproductswitherror",
			errors.New("get products error"),
			http.StatusInternalServerError,
			errMSG{Error: "get products error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockdbLayer.SetError(tt.inErr)
			req := httptest.NewRequest(http.MethodGet, productsURL, nil)
			w := httptest.NewRecorder()
			_, engine := gin.CreateTestContext(w)
			engine.GET(productsURL, h.GetProducts)
			engine.ServeHTTP(w, req)
			response := w.Result()
			if response.StatusCode != tt.outStatusCode {
				t.Errorf("received status code %d does not match expected status code %d", response.StatusCode, tt.outStatusCode)
			}
			var respBody interface{}
			if tt.inErr != nil {
				var errmsg errMSG
				json.NewDecoder(response.Body).Decode(&errmsg)
				respBody = errmsg
			} else {
				products := []models.Product{}
				json.NewDecoder(response.Body).Decode(&products)
				respBody = products
			}

			if !reflect.DeepEqual(respBody, tt.expectedRespBody) {
				t.Logf("%+v , %+v", respBody, tt.expectedRespBody)
				t.Errorf("Received HTTP response body %+v does not match expected HTTP response Body %+v", respBody, tt.expectedRespBody)
			}
		})
	}
}
