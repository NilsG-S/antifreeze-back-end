package db

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

type Conn struct {
	context context.Context
	client  *datastore.Client
}

var instance *Conn
var once sync.Once

// TODO(NilsG-S): Don't use singletons
func GetInstance() (*Conn, error) {
	var err error

	once.Do(func() {
		var (
			ctx context.Context = context.Background()
			cli *datastore.Client
		)

		// $DATASTORE_PROJECT_ID is used when second arg is empty
		cli, err = datastore.NewClient(ctx, "")
		if err != nil {
			err = fmt.Errorf("Couldn't create client: %v", err)
		}

		instance = &Conn{
			context: ctx,
			client:  cli,
		}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

type Test struct {
	Value string
}

func (c *Conn) Testing() error {
	k := datastore.IncompleteKey("Test", nil)
	e := new(Test)
	e.Value = "Testing"

	if _, err := c.client.Put(c.context, k, e); err != nil {
		return fmt.Errorf("Couldn't insert test entity: %v", err)
	}

	return nil
}

func (c *Conn) TestingGet() error {
	q := datastore.NewQuery("Test")
	for i := c.client.Run(c.context, q); ; {
		var t Test
		key, err := i.Next(&t)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Error when iterating test query: %v", err)
		}
		fmt.Println(key, t)
	}

	return nil
}