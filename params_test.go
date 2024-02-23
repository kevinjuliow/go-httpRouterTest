package httpRouter

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id")
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "/product/1", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(bytes)) //Testing , to expect the output = "Product 1"
}

func TestNamedParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id/items/:itemid", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		productId := params.ByName("id")
		itemId := params.ByName("itemid")

		fmt.Fprintf(writer, "Product %s , item %s", productId, itemId)
	})

	req := httptest.NewRequest(http.MethodGet, "/products/1/items/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	resp := recorder.Result()

	bytes, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "Product 1 , item 2", string(bytes))
}

func TestCatchAllParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/picture/*pathfile", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "path : "+params.ByName("pathfile"))
	})

	req := httptest.NewRequest(http.MethodGet, "/picture/document/src/pictures", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	resp := rec.Result()
	bytes, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "path : /document/src/pictures", string(bytes))
}
