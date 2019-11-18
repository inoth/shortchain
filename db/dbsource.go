package db

import (
	"context"
	"shortchain/util/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	instance *MongoConnect
)

type MongoConnect struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func Instance() *MongoConnect {
	once.Do(func() {
		instance = &MongoConnect{}
	})
	return instance
}

func (d *MongoConnect) Init(dbname string) (err error) {
	d.Client, err = mongo.Connect(context.TODO(),
		options.Client().ApplyURI(config.Instance().DB.Mongo[0].ConnectStr))
	if err != nil {
		return err
	}
	err = d.Client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	d.Db = d.Client.Database(dbname)
	return nil
}
