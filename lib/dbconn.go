package lib

import "log"

package lib

import (
"context"
"fmt"
"log"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

var err error

// Role blah
type Role struct {
	GuildID   string `bson:"GuildID,omitempty"`
	ChannelID string `bson:"ChannelID,omitempty"`
	RoleID    string `bson:"RoleID,omitempty"`
	IgnoreID  string `bson:"IgnoreID,omitempty"`
	Phrase    string `bson:"Phrase,omitempty"`
}

// GetClient blah
func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// MonRole mongoDB connection stuff
func MonRole(dbase string, collect string, listens Role) {
	// Connecting to mongoDB
	client := GetClient()
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	collection := client.Database(dbase).Collection(collect)
	insertResult, err := collection.InsertOne(context.TODO(), listens)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted:", insertResult.InsertedID)
}

// MonUpdateRole blah
func MonUpdateRole(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("gemini").Collection("roles")
	update := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Error updating player", err)
	}
	return updatedResult.ModifiedCount
}

// MonReturnOneRole blah
func MonReturnOneRole(client *mongo.Client, filter bson.M) Role {
	var phrase Role
	collection := client.Database("gemini").Collection("roles")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&phrase)
	return phrase
}

// MonReturnAllRole blah
func MonReturnAllRole(client *mongo.Client, filter bson.M) []*Role {
	var roles []*Role
	collection := client.Database("gemini").Collection("roles")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error finding all the things", err)
	}
	for cur.Next(context.TODO()) {
		var role Role
		err = cur.Decode(&role)
		if err != nil {
			log.Fatal("Error Decoding :( ", err)
		}
		roles = append(roles, &role)
	}
	return roles
}

