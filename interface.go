package kibisis

// Database - Wraps multiple databse implementations
type Database interface {
	Conn() error
	Init(database, collection string) error
	Create(item interface{}) error
	Update(id string, item interface{}) error
	Delete(id string) error
	Find(id string) (interface{}, error)
	FindAll() ([]interface{}, error)
}
