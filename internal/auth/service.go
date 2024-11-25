package auth

import (
	"github.com/Nyamerka/coffee_shop/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(name, email, password string) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
		RegisterDate: time.Now(),
	}

	if err := postgres.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func Login(email, password string) (*models.User, error) {
	var user models.User
	if err := postgres.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}
