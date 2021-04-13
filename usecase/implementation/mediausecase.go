package implementation

import (
	"FL_2/model"
	"FL_2/store"
	"github.com/pkg/errors"
	"strconv"
)
const(
	mediaUseCaseError = "Media use case error."
)
type MediaUseCase struct {
	store      store.Store
	mediaStore store.MediaStore
}

func (s *MediaUseCase) GetImage(imageInfo interface{}) (*model.User, error) {
	return nil, nil
}

func (s *MediaUseCase) SetImage(imageInfo interface{}, image []byte) (*model.User, error) {
	u := imageInfo.(*model.User)
	imageId, err := s.mediaStore.Image().SetImage(strconv.FormatUint(u.ID, 10), image)
	if err != nil {
		return nil, errors.Wrap(err, mediaUseCaseError)
	}
	u.Img = imageId
	u, err = s.store.User().ChangeUser(*u)
	if err != nil {
		return nil, errors.Wrap(err,mediaUseCaseError)
	}
	u.Img = string(image)
	return u, err
}
