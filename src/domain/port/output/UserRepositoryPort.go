package input

import "github.com/InsaneProgZ/user-service/src/domain"

type UserRepositoryPort interface {
	save(user domain.CreateUser)
}
