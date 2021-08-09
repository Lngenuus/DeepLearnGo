package server

import (
	http "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

func RegisterSwagger(srv *http.Server) {
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", h)
}
