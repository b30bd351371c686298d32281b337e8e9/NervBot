package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Client represents a connection to Google's Cloud Datastore instance
type Client struct {
	Client  *datastore.Client
	Context context.Context
}

type Entity struct {
	Value string
}

// New creates a new Cloud Datastore client and returns a reference to a new DataStoreClient
func New(ctx context.Context, project string) (*Client, error) {
	client, err := datastore.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client:  client,
		Context: ctx,
	}, nil
}

// Get will get a value from a Cloud Datastore instance
func (d *Client) Get(kind string, key string) (string, error) {
	k := datastore.NameKey(kind, key, nil)
	e := new(Entity)

	if err := d.Client.Get(d.Context, k, e); err != nil {
		return "", err
	}
	return e.Value, nil
}
