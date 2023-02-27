package store

type Book struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Authors string `json:"authors"`
	Press   string `json:"press"`
}

type Store interface {
	Create(name string, book *Book) error
	Update(name string, book *Book) error
	Get(name string) (*Book, error)
	GetAll() ([]*Book, error)
	Delete(name string) error
}
