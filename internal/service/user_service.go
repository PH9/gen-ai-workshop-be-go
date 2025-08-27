package service

import (
	"time"
	"github.com/PH9/gen-ai-workshop-be-go/internal/model"
	"github.com/PH9/gen-ai-workshop-be-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Register(email, password, firstname, lastname, phone, birthday string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	bday, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return err
	}
	user := &model.User{
		Email:     email,
		Password:  string(hash),
		FirstName: firstname,
		LastName:  lastname,
		Phone:     phone,
		Birthday:  bday,
		CreatedAt: time.Now(),
	}
	return s.Repo.Create(user)
}

func (s *UserService) Authenticate(email, password string) (*model.User, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.Repo.FindByID(id)
}
