package kibisis

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ArangoDb struct {
	Client     driver.Client
	Database   driver.Database
	Collection driver.Collection
}

// Conn - Creates a database client
func (arangodb *ArangoDb) Conn() error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	arangodb.Client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "root"),
	})
	if err != nil {
		return fmt.Errorf("Failed to create database client: %v", err)
	}

	return nil
}

// Init - Connects to the target database using the client
func (arangodb *ArangoDb) Init(database, collection string) error {
	var err error

	arangodb.Database, err = arangodb.Client.Database(nil, database)

	if err != nil {
		return fmt.Errorf("Failed to initialise database: %v", err)
	}

	arangodb.Collection, err = arangodb.Database.Collection(nil, collection)

	return nil
}

// Create - Inserts an item into the database
func (arangodb *ArangoDb) Create(item interface{}) error {
	ctx := context.Background()
	_, err := arangodb.Collection.CreateDocument(ctx, item)
	if err != nil {
		return fmt.Errorf("Error inserting item: %v", err)
	}

	return nil
}

// Find - Get a single item from the database
func (arangodb *ArangoDb) Find(id string) (interface{}, error) {
	// ctx := context.Background()
	// var item interface{}

	// _, err := arangodb.Collection.ReadDocument(ctx, "16717", &item)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error finding item: %v", err)
	// }

	// return item, nil

	ctx := context.Background()
	query := fmt.Sprintf("FOR d in %v LIMIT 1 RETURN d", arangodb.Collection.Name())
	cursor, err := arangodb.Database.Query(ctx, query, nil)
	if err != nil {
		return nil, fmt.Errorf("Error finding items: %v", err)
	}

	defer cursor.Close()
	for {
		var item interface{}
		_, err := cursor.ReadDocument(ctx, &item)
		if err != nil {
			return nil, fmt.Errorf("Error fetching item: %v", err)
		}

		return item, nil
	}
}

// FindAll - Get a collection of items from the database
func (arangodb *ArangoDb) FindAll() ([]interface{}, error) {
	ctx := context.Background()
	query := fmt.Sprintf("FOR d in %v RETURN d", arangodb.Collection.Name())
	cursor, err := arangodb.Database.Query(ctx, query, nil)
	if err != nil {
		return nil, fmt.Errorf("Error finding items: %v", err)
	}

	defer cursor.Close()
	var items []interface{}
	for {
		var item interface{}
		_, err := cursor.ReadDocument(ctx, &item)
		items = append(items, item)

		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("Error fetching items: %v", err)
		}
	}

	return items, nil
}

// Update - Update an item in the database
func (arangodb *ArangoDb) Update(id string, item interface{}) error {
	ctx := context.Background()
	_, err := arangodb.Collection.UpdateDocument(ctx, id, item)
	if err != nil {
		return fmt.Errorf("Error updating item: %v", err)
	}

	return nil
}

// Delete - Delete an item from the database
func (arangodb *ArangoDb) Delete(id string) error {
	ctx := context.Background()
	_, err := arangodb.Collection.RemoveDocument(ctx, id)
	if err != nil {
		return fmt.Errorf("Error deleting item: %v", err)
	}

	return nil
}
