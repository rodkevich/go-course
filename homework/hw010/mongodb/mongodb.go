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

var (
	ctxDefault        = context.Background()
	operationsTimeOut = 10 * time.Second
)

// Represents the contactsBook data-source
type contactsBook struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func (c contactsBook) String() string {
	return "Mongo"
}

// Up ...
func (c contactsBook) Up() (err error) {
	err = c.client.Ping(ctxDefault, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("mongo: book UP operation done")
	return
}

// Close current data-source
func (c contactsBook) Close() {
	log.Println("mongo: book disconnecting ...")
	defer log.Println("mongo: book disconnecting - done")
	err := c.client.Disconnect(ctxDefault)
	if err != nil {
		return
	}
}

// Drop current data-source
func (c contactsBook) Drop() (err error) {
	err = c.collection.Drop(ctxDefault)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("mongo: database dropped")
	return
}

// Truncate current data-source
func (c contactsBook) Truncate() (err error) {
	filter := bson.D{}
	_, err = c.collection.DeleteMany(ctxDefault, filter)
	if err != nil {
		return
	}
	return
}

// Create new record in current data-source
func (c contactsBook) Create(contact *types.Contact) (recordID string, err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	rtn, err := c.collection.InsertOne(ctx, contact)
	if err != nil {
		return
	}
	recordID = rtn.InsertedID.(primitive.ObjectID).Hex()
	return
}

// AssignContactToGroup change contact's group to required
func (c contactsBook) AssignContactToGroup(contact *types.Contact, gr types.Group) (rtn *types.Contact) {
	_, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	var (
		stmt   = bson.M{"$set": bson.M{"group": gr}}
		filter = bson.M{"uuid": &contact.UUID}
		// add opts... other way will get old not updated document in return (default: before)
		after  = options.After
		opt    = options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
		}
	)
	res := c.collection.FindOneAndUpdate(
		ctxDefault,
		filter,
		stmt,
		&opt,
	).Decode(&rtn)
	if res != nil {
		if res == mongo.ErrNoDocuments {
			return
		}
		log.Println("mongo_Decode() unmarshal err: ", res)
	}
	return
}

// FindByGroup look up with group filter
func (c contactsBook) FindByGroup(gr types.Group) (rtn []*types.Contact, err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	sort := options.Find() // add sort
	sort.SetSort(bson.D{{"uuid", -1}})
	cur, err := c.collection.Find(
		c.ctx,
		bson.D{{"group", gr}},
		sort,
	)
	if err != nil {
		return
	}
	if err = cur.All(ctx, &rtn); err != nil {
		log.Println(err)
		return
	}
	if err = cur.Err(); err != nil {
		return
	}
	return
}

// NewContactsBook create new data-source instance
func NewContactsBook() (ds cb.ContactBookDataSource, err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	var mongoURL = os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Println(err)
		return
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	database := client.Database("contacts-book").Collection("contacts")
	ds = &contactsBook{
		client:     client,
		collection: database,
		ctx:        ctxDefault,
	}
	return
}
