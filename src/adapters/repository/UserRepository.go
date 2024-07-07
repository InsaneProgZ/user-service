package repository

import "github.com/InsaneProgZ/user-service/src/domain/model"

type UserRepository struct {
	Data []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (s *UserRepository) Save(user model.User) error {
	s.Data = append(s.Data, user)
	return nil
}

func (s *UserRepository) Find(username string) *model.User {
	for _, user := range s.Data {
		if(user.Username == username){
			return &user
		}
	}
	return nil
}
