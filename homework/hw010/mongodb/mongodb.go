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
	operationsTimeOut = 3 * time.Second
)

// Represents the contactsBook model
type contactsBook struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
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

// Close ...
func (c contactsBook) Close() {
	log.Println("mongo: book disconnecting ...")
	defer log.Println("mongo: book disconnecting - done")
	err := c.client.Disconnect(ctxDefault)
	if err != nil {
		return
	}
}

// Drop ...
func (c contactsBook) Drop() (err error) {
	err = c.collection.Drop(ctxDefault)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("mongo: database dropped")
	return
}

// Truncate ...
func (c contactsBook) Truncate() (err error) {
	_, err = c.collection.DeleteMany(ctxDefault, bson.D{})
	if err != nil {
		return
	}
	return
}

// Create ...
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

// AssignContactToGroup ...
func (c contactsBook) AssignContactToGroup(contact *types.Contact, gr types.Group) (n *types.Contact) {
	n = new(types.Contact)
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	var (
		stmt   = bson.M{"$set": bson.M{"group": gr}}
		filter = bson.M{"uuid": &contact.UUID}
	)
	rtn := c.collection.FindOneAndUpdate(ctx, filter, stmt).Decode(&n)
	if rtn != nil {
		if rtn == mongo.ErrNoDocuments {
			return
		}
		log.Println(rtn)
	}
	return
}

// FindByGroup ...
func (c contactsBook) FindByGroup(gr types.Group) (contacts []*types.Contact, err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
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
	defer cur.Close(ctxDefault)

	if err = cur.All(ctx, &contacts); err != nil {
		log.Println(err)
		return
	}
	if err = cur.Err(); err != nil {
		return
	}
	return
}

// NewContactsBook ...
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
