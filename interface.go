package kibisis

// Database - Wraps multiple databse implementations
type Database interface {
	Conn(host []string, username string, password string) error
	Init(database, collection string) error
	Create(item interface{}) (string, error)
	Update(id string, item interface{}) error
	Delete(id string) error
	Find(id string) (interface{}, error)
	FindAll(where []string, sort []string, limit int) ([]interface{}, error)
}
