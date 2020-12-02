package kibisis

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDb - The MongoDB instance
type MongoDb struct {
	Client     mongo.Client
	Database   mongo.Database
	Collection mongo.Collection
}

// Conn - Creates a database client
func (mongodb *MongoDb) Conn(host []string, username string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		AuthMechanism: "PLAIN",
		Username:      username,
		Password:      password,
	}
	clientOpts := options.Client().ApplyURI(host[0]).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	mongodb.Client = *client

	return nil
}

// Init - Connects to the target database using the client
func (mongodb *MongoDb) Init(database, collection string) error {
	mongodb.Collection = *mongodb.Client.Database(database).Collection(collection)

	return nil
}

// Create - Inserts an item into the database
func (mongodb *MongoDb) Create(item interface{}) error {
	ctx := context.Background()
	_, err := mongodb.Collection.InsertOne(ctx, item)
	if err != nil {
		return fmt.Errorf("Error inserting item: %v", err)
	}

	return nil
}

// Find - Get a single item from the database
func (mongodb *MongoDb) Find(id string) (interface{}, error) {

	return nil, fmt.Errorf("Error finding item")
}

// FindAll - Get a collection of items from the database
func (mongodb *MongoDb) FindAll(where []string, sort []string, limit int) ([]interface{}, error) {

	return nil, fmt.Errorf("Error finding items")
}

// Update - Update an item in the database
func (mongodb *MongoDb) Update(id string, item interface{}) error {

	return nil
}

// Delete - Delete an item from the database
func (mongodb *MongoDb) Delete(id string) error {

	return nil
}