package output

import "github.com/InsaneProgZ/user-service/src/domain/model"

type UserRepositoryPort interface {
	Save(user model.User) error
	Find(username string) *model.User
}
