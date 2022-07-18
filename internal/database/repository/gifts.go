package repository

import (
	"context"
	
	"wedding_gifts/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GiftsRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewGiftsRepository(client *mongo.Client) GiftsRepository {
	db := client.Database("wedding_gifts")
	coll := db.Collection("gifts")

	return GiftsRepository{
		db: db,
		coll: coll,
	}
}

func (rp GiftsRepository) Create(data model.Gift) error {
	if _, err := rp.coll.InsertOne(context.TODO(), data); err != nil {
		return err
	}

	return nil
}

func (rp GiftsRepository) List(filter any) ([]model.Gift, error) {
	var gifts []model.Gift

	cursor, err := rp.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &gifts); err != nil {
		return nil, err
	}

	return gifts, nil
}

func (rp GiftsRepository) Get(filter any) (model.Gift, error) {
	var gift model.Gift
	
	res := rp.coll.FindOne(context.TODO(), filter)
	
	if err := res.Decode(&gift); err != nil {
		return gift, err
	}
	
	return gift, nil
}

func (rp GiftsRepository) Update(filter any, payload bson.D) error {
	if _, err := rp.coll.UpdateOne(context.TODO(), filter, payload); err != nil {
		return err
	}

	return nil
}

func (rp GiftsRepository) Delete(filter any) error {
	if _, err := rp.coll.DeleteOne(context.TODO(), filter); err != nil {
		return err
	}
	return nil
}
