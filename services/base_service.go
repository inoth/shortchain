package services

import (
	"context"
	"shortchain/db"
	"shortchain/model"
)

func Create(entity model.IEntity) error {
	if _, err := db.Instance().Db.Collection(entity.ColName()).InsertOne(context.TODO(), entity); err != nil {
		return err
	}
	return nil
}

func FindOne(filter interface{}, res model.IEntity) error {
	if err := db.Instance().Db.Collection(res.ColName()).FindOne(context.TODO(), filter).Decode(res); err != nil {
		return err
	}
	return nil
}
