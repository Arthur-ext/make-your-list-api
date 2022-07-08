package controller

// any's must be overwritten to struct's when models are created
type GiftsRepository interface{
	Create(data any) error
	List() ([]any, error)
	Get(id int) (any, error)
	Update(data any) error
	Delete(id int)
}

type Gifts struct {
	repository *GiftsRepository
}

func NewGifts() *Gifts {
	return &Gifts{
		
	}
}