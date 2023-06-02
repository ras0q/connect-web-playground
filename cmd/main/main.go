package main

import (
	"net/http"

	"github.com/ras0q/connect-web-playground/bufgen/go/api/proto/protoconnect"
	"github.com/ras0q/connect-web-playground/internal/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	rh := handler.NewReadyHandler()
	mux := http.NewServeMux()
	path, handler := protoconnect.NewReadyServiceHandler(rh)
	mux.Handle(path, handler)
	if err := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		panic(err)
	}
}
