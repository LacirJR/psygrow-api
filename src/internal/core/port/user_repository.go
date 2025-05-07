package port

import "github.com/LacirJR/psygrow-api/src/internal/core/model"

type UserRepository interface {
	Save(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
