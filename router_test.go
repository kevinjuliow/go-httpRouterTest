package httpRouter

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
