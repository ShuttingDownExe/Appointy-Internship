package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AccessDB(userid , passwd string) *mongo.Client {
	var err error
	var client *mongo.Client

	AtlasURI := "mongodb+srv://"+userid+":"+passwd+"@storagecluster.ow0m3.mongodb.net/test"
	opts := options.Client()
	opts.ApplyURI(AtlasURI)
	opts.SetMaxPoolSize(1)

	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		panic(err)
	}
	return client
}
