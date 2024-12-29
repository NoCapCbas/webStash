// service domain logic, contains only business specific logic
package users

import (
	"github.com/NoCapCbas/webStash/internal/users/model"
	"github.com/NoCapCbas/webStash/internal/users/repository"
	"github.com/NoCapCbas/webStash/internal/common/publisher"
)

UserService interface {
	SignUp(*User) (*User, error)
	Update(user *User) error
	Deactivate(id int) error
	Reactivate(id int) error
}

type UserServiceImpl struct {
	repo UserPostgresRepository
	pub  publisher.Publisher
}

func NewUserService(repo repository.UserPostgresRepository, pub publisher.Publisher) UserService {
	return &UserServiceImpl{
		repo: repo,
		pub:  pub,
	}
}

func (s *UserServiceImpl) SignUp(user *User) (*User, error) {
	_, err := s.repo.Create(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = s.pub.Publish(context.Background(), "user.signedup", user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) Update(user *User) error {
	_, err := s.repo.Update(user)
	if err != nil {
		log.Error(err)
		return err
	}
	err = s.pub.Publish(context.Background(), "user.updated", user)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *UserServiceImpl) Verify(user *User) error {
	_, err := s.repo.Verify(user)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.pub.Publish(context.Background(), "user.verified", user)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

