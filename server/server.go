package server

import (
	"encoding/json"
	mutex "github.com/gorilla/mux"
	"net/http"
	"rich/server/middleware"
	"rich/store"
	"time"
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
	router.HandleFunc("/book/{id}", srv.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", srv.DelBook).Methods("DELETE")

	srv.srv.Handler = middleware.Logging(middleware.Validate(router))

	return srv, nil

}
func (srv *BookServer) CreateBook(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srv.s.Create(book)
	response(w, "hello")
}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (srv *BookServer) GetBook(w http.ResponseWriter, r *http.Request) {
	id, ok := mutex.Vars(r)["id"]
	if !ok {
		http.Error(w, "no id found in request.", http.StatusBadRequest)
		return
	}

	book, _ := srv.s.Get(id)

	response(w, book)

}

func (srv *BookServer) DelBook(w http.ResponseWriter, r *http.Request) {
	id, ok := mutex.Vars(r)["id"]
	if !ok {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}

	err := srv.s.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (srv *BookServer) ListenAndServe() (<-chan error, error) {
	c := make(chan error, 1)
	go func() {
		err := srv.srv.ListenAndServe()
		c <- err
	}()

	select {
	case err := <-c:
		return nil, err

	case <-time.After(time.Second):
		return c, nil
	}
}
