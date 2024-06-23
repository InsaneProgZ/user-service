package input

import "github.com/InsaneProgZ/user-service/src/domain"

type UserPort interface {
	CreateUser(user domain.CreateUser)
}
