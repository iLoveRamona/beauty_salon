package repositories

import (
	"beauty_salon_bd/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (name, phone, password) VALUES ($1, $2, $3)",
		user.Name, user.Phone, user.Password)
	return err
}

func (r *UserRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		"SELECT id, name, phone, password FROM users WHERE phone = $1", phone).
		Scan(&user.ID, &user.Name, &user.Phone, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UserExists(phone string) (bool, error) {
	var count int
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM users WHERE phone = $1", phone).Scan(&count)
	return count > 0, err
}

func (r *UserRepository) GetUserProfile(phone string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		"SELECT id, name, phone FROM users WHERE phone = $1", phone).
		Scan(&user.ID, &user.Name, &user.Phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
