package input

import "github.com/InsaneProgZ/user-service/src/application/domain"

type UserPort interface {
	CreateUser(user domain.CreateUser)
}
