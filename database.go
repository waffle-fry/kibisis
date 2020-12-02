package kibisis

import "errors"

// GetDriver - Returns the database driver
func GetDriver(name string) (Database, error) {
	switch name {
	case "arangoDB":
		var driver ArangoDb
		return &driver, nil
	case "mongoDB":
		var driver MongoDb
		return &driver, nil
	default:
		return nil, errors.New("Database driver not found")
	}
}
