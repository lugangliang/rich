package server

import (
	mutex "github.com/gorilla/mux"
	"net/http"
	"rich/server/middleware"
	"rich/store"
)

type BookServer struct {
	s   store.Store
	srv http.Server
}

func NewBookStoreServer(addr string, s store.Store) (*BookServer, error) {

	srv := &BookServer{
		s: s,
		srv: http.Server{
			Addr: addr,
		},
	}

	router := mutex.NewRouter()
	router.HandleFunc("/book", srv.CreateBook).Methods("POST")

	srv.srv.Handler = middleware.Logging(middleware.Validate(router))

	return srv, nil

}
func (srv *BookServer) CreateBook(http.ResponseWriter, *http.Request) {

}
