package implementation

import (
	"FL_2/model"
	"FL_2/store"
)

type ResponseUseCase struct {
	store      store.Store
	mediaStore store.MediaStore
}

func (r *ResponseUseCase) Create(response model.Response) (*model.Response, error) {
	user, err := r.store.User().FindByID(response.UserID)
	if err != nil {
		return nil, err
	}
	response.UserLogin = user.Login
	response.UserImg = user.Img
	id, err := r.store.Response().Create(response)
	response.ID = id
	if err != nil {
		return nil, err
	}
	img, err := r.mediaStore.Image().GetImage(response.UserImg)
	if err != nil {
		return nil, err
	}
	response.UserImg = string(img)
	return &response, nil
}

func (r *ResponseUseCase) FindByOrderID(id uint64) ([]model.Response, error) {
	responses, err := r.store.Response().FindById(id)
	if err != nil {
		return nil, err
	}
	for _, response := range responses {
		img, err := r.mediaStore.Image().GetImage(response.UserImg)
		if err != nil {
			return nil, err
		}
		response.UserImg = string(img)
	}
	if responses == nil {
		return []model.Response{}, nil
	}
	return responses, nil
}

func (r *ResponseUseCase) Change(response model.Response) (*model.Response, error) {
	changedResponse, err := r.store.Response().Change(response)
	if err != nil {
		return nil, err
	}
	return changedResponse, nil
}

func (r *ResponseUseCase) Delete(response model.Response) error {
	err := r.store.Response().Delete(response)
	if err != nil {
		return err
	}
	return nil
}
