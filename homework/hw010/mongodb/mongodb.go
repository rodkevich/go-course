package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	cb "github.com/rodkevich/go-course/homework/hw009/book"
	"github.com/rodkevich/go-course/homework/hw009/book/types"
)

// Represents the contactsBook model
type contactsBook struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

// Up ...
func (c contactsBook) Up() error {
	err := c.client.Ping(c.ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongo: book UP operation done")
	return nil
}

// Close ...
func (c contactsBook) Close() {
	log.Println("mongo: book disconnecting ...")
	defer log.Println("mongo: book disconnecting - done")
	err := c.client.Disconnect(c.ctx)
	if err != nil {
		return
	}
}

// Drop ...
func (c contactsBook) Drop() error {
	err := c.collection.Drop(c.ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongo: database dropped")
	return nil
}

// Truncate ...
func (c contactsBook) Truncate() error {
	_, err := c.collection.DeleteMany(c.ctx, bson.D{})
	if err != nil {
		return err
	}
	return nil
}

// Create ...
func (c contactsBook) Create(contact *cb.Contact) (recordID string, err error) {
	rtn, err := c.collection.InsertOne(c.ctx, contact)
	if err != nil {
		return
	}
	recordID = rtn.InsertedID.(primitive.ObjectID).Hex()
	return
}

// AssignContactToGroup ...
func (c contactsBook) AssignContactToGroup(contact *cb.Contact, gr types.Group) (new *cb.Contact) {
	var (
		stmt   = bson.M{"$set": bson.M{"group": gr}}
		filter = bson.M{"uuid": &contact.UUID}
	)
	rtn := c.collection.FindOneAndUpdate(c.ctx, filter, stmt).Decode(&new)
	if rtn != nil {
		if rtn == mongo.ErrNoDocuments {
			return nil
		}
		log.Println(rtn)
	}
	return
}

// FindByGroup ...
func (c contactsBook) FindByGroup(gr types.Group) ([]*cb.Contact, error) {
	sort := options.Find() // just curiosity ^^ if it doesn't break a stmt
	sort.SetSort(bson.D{{"uuid", -1}})
	cur, err := c.collection.Find(
		c.ctx,
		bson.D{{"group", gr}},
		sort,
	)
	if err != nil {
		return nil, err
	}
	defer cur.Close(c.ctx)

	var rtn []*cb.Contact
	if err = cur.All(context.Background(), &rtn); err != nil {
		log.Println(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	for _, record := range rtn {
		fmt.Println(record)
	}
	return rtn, nil
}

// NewContactsBook ...
func NewContactsBook() (cb.ContactsBookDataSource, error) {
	var mongoURL = os.Getenv("MONGO_URL")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("contacts-book").Collection("contacts")
	return &contactsBook{
		client:     client,
		collection: database,
		ctx:        ctx,
	}, nil
}

//
// func (c *CountriesCollection) Find(filter interface{}, opts ...*options.FindOptions) ([]Country, error) {
// 	cur, err := c.collection.Find(c.ctx, filter, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cur.Close(c.ctx)
//
// 	var results []Country
// 	cur.All(c.ctx, &results)
// 	if err := cur.Err(); err != nil {
// 		return nil, err
// 	}
// }

// func (c *CountriesCollection) InsertOne(country Country) error {
// 	_, err := c.collection.InsertOne(c.ctx, country)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//

//
// 	return results, nil
// }
//
// func (c *CountriesCollection) Remove(filter interface{}) (int, error) {
// 	deleteResult, err := c.collection.DeleteMany(c.ctx, filter)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return int(deleteResult.DeletedCount), nil
// }
