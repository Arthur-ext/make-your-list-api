package app

import (
	"wedding_gifts/internal/app/controller"
	"wedding_gifts/internal/database"
)

type APP struct {
	giftsController controller.Gifts
}

func NewAPP() *APP {
	db, err := database.NewMongoClient()
	if err != nil {
		panic(err)
	}

	err = db.CheckConn()
	if err != nil {
		panic(err)
	}

	return &APP{}
}