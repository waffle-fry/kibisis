package kibisis

// Database - Wraps multiple databse implementations
type Database interface {
	Conn(host []string, username string, password string) error
	Init(database, collection string) error
	Create(item interface{}) error
	Update(id string, item interface{}) error
	Delete(id string) error
	Find(id string) (interface{}, error)
	FindAll() ([]interface{}, error)
}
