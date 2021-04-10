package store

import "FL_2/model"

type ResponseRepository interface {
	Create(response model.Response) (uint64, error)
	FindById(id uint64) ([]model.Response, error)
}
