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

func TestPanicHandler(t *testing.T) {

	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprint(writer, "panic", i)
	}

	router.GET("/products", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups")
	})

	req := httptest.NewRequest(http.MethodGet, "/products", nil)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	resp := rec.Result()
	bytes, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "panicUps", string(bytes))
}

func TestPanicHandler2(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprint(writer, i)
	}

	router.GET("/products", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Panic Handler")
	})

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
