package kibisis

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// credential := options.Credential{
	// 	AuthMechanism: "PLAIN",
	// 	Username:      username,
	// 	Password:      password,
	// }
	clientOpts := options.Client().ApplyURI(host[0]) //.SetAuth(credential)
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
func (mongodb *MongoDb) Create(item interface{}) (string, error) {
	ctx := context.Background()
	res, err := mongodb.Collection.InsertOne(ctx, item)
	if err != nil {
		return "", fmt.Errorf("Error inserting item: %v", err)
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

// Find - Get a single item from the database
func (mongodb *MongoDb) Find(id string) (interface{}, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.M{"_id": objectID}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result bson.M

	err = mongodb.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("Error finding item: %v", err)
	}

	return result, nil
}

// FindAll - Get a collection of items from the database
func (mongodb *MongoDb) FindAll(where []string, sort []string, limit int) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := mongodb.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("Error finding items: %v", err)
	}
	defer cur.Close(ctx)
	var results []interface{}

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("Error finding items: %v", err)
		}

		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("Error finding items: %v", err)
	}

	return results, nil
}

// Update - TODO: Update an item in the database
func (mongodb *MongoDb) Update(id string, item interface{}) error {

	return nil
}

// Delete - TODO: Delete an item from the database
func (mongodb *MongoDb) Delete(id string) error {

	return nil
}
