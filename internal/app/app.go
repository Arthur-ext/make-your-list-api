package app

import (
	"wedding_gifts/internal/app/controller"
	"wedding_gifts/internal/database"
	"wedding_gifts/internal/database/repository"
)

type APP struct {
	Gifts controller.GiftsController
}

func NewAPP() APP {
	db, err := database.NewMongoClient()
	if err != nil {
		panic(err)
	}
	if err = db.CheckConn(); err != nil {
		panic(err)
	}

	giftRepo := repository.NewGiftsRepository(db.Client)
	giftController := controller.NewGifts(giftRepo)
	
	return APP{
		Gifts: giftController,
	}
}