package server

import (
	"bookstore/pkg/db"
	"bookstore/pkg/handler"
	"net/http"
)

type Server struct {
	http.Server
	handler handler.Handler
}

func New(conn *db.Conn) *Server {
	h := handler.NewHandler(conn)
	return &Server{
		Server: http.Server{
			Addr:    ":8080",
			Handler: h,
		},
		handler: h,
	}
}

func (s *Server) ListenAndServe() error {
	return s.Server.ListenAndServe()
}