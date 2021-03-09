package store

import "fl_ru/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(user *model.User) error
}

type SessionRepository interface {
	Create(session *model.Session) error
	Find(session *model.Session) error
}
