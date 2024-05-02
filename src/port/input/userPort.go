package input

import (
	"planzin/user/src/application/domain"
)

type UserPort interface {
	CreateUser(c domain.CreateUser)
}
