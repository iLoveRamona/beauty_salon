package services

import (
	"beauty_salon_bd/models"
	"beauty_salon_bd/repositories"
	"beauty_salon_bd/utils"
	"fmt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(user *models.User) error {
	exists, err := s.userRepo.UserExists(user.Phone)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserExists
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(phone, password string) (*models.User, error) {
	user, err := s.userRepo.FindByPhone(phone)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (s *AuthService) GetUserProfile(phone string) (*models.User, error) {
	user, err := s.userRepo.GetUserProfile(phone)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %v", err)
	}
	return user, nil
}

var (
	ErrUserExists         = fmt.Errorf("user already exists")
	ErrInvalidCredentials = fmt.Errorf("invalid credentials")
)
