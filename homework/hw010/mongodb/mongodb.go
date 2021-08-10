package mongodb

import (
	"context"
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
func (c contactsBook) Up() (err error) {
	err = c.client.Ping(c.ctx, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("mongo: book UP operation done")
	return
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
func (c contactsBook) Drop() (err error) {
	err = c.collection.Drop(c.ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("mongo: database dropped")
	return
}

// Truncate ...
func (c contactsBook) Truncate() (err error) {
	_, err = c.collection.DeleteMany(c.ctx, bson.D{})
	if err != nil {
		return
	}
	return
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
			return
		}
		log.Println(rtn)
	}
	return
}

// FindByGroup ...
func (c contactsBook) FindByGroup(gr types.Group) (contacts []*cb.Contact, err error) {
	sort := options.Find() // just curiosity ^^ if it doesn't break a stmt
	sort.SetSort(bson.D{{"uuid", -1}})
	cur, err := c.collection.Find(
		c.ctx,
		bson.D{{"group", gr}},
		sort,
	)
	if err != nil {
		return
	}
	defer cur.Close(c.ctx)

	if err = cur.All(context.Background(), &contacts); err != nil {
		log.Println(err)
		return
	}
	if err = cur.Err(); err != nil {
		return
	}
	return
}

// NewContactsBook ...
func NewContactsBook() (cb.ContactBookDataSource, error) {
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
