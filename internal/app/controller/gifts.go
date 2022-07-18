package controller

import (
	"errors"
	"time"
	"wedding_gifts/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GiftStore interface{
	Create(data model.Gift) error
	List(filter any) ([]model.Gift, error)
	Get(filter any) (model.Gift, error)
	Update(filter any, payload bson.D) error
	Delete(filter any) error
}

type GiftsController struct {
	repo GiftStore
}

func NewGifts(store GiftStore) GiftsController {
	return GiftsController{
		repo: store,
	}
}

func (c *GiftsController) Create(data model.Gift) error {
	data.ID = primitive.NewObjectID()
	data.CreatedAt = time.Now()

	return c.repo.Create(data)
}

func (c *GiftsController) List() ([]model.Gift, error) {
	res, err := c.repo.List(bson.D{})
	
	return res, err
}

func (c *GiftsController) Get(id string) (model.Gift, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Gift{}, err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	
	return c.repo.Get(filter)
}

func (c *GiftsController) UpdateAssigner(data map[string]string) error {
	id, ok := data["id"]
	if !ok {
		return errors.New("not have a key id")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}

	assigned, ok := data["assigned"]
	if !ok {
		return errors.New("not have a key assigned")
	}
	payload := bson.D{{
		Key: "$set",
		Value: bson.M{
			"assigned": assigned,
			"updatedAt": time.Now(),
		},
	}}

	if err := c.repo.Update(filter, payload); err != nil {
		return err
	}

	return nil
}

func (c *GiftsController) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	
	return c.repo.Delete(filter)
}