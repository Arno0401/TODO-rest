package repository

import (
	"arno/internal/models"
	"arno/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *Repository) SignUpUser(name, login, pass string) error {
	var existingUser models.Users
	if err := r.conn.Where("login = ?", login).First(&existingUser).Error; err == nil {
		return fmt.Errorf("user with this login already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	hashedPassword, err := utils.HashPassword(pass)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	if err := r.conn.Create(&models.Users{UserName: name, Login: login, Password: hashedPassword, Role: "user"}).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *Repository) IsExistingUser(login, pass string) (bool, error) {
	var existingUser models.Users

	if err := r.conn.Where("login = ?", login).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, utils.InternalError
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(pass))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *Repository) GetUser(login string) (*models.Users, error) {
	var user models.Users
	if err := r.conn.Where("login = ?", login).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.UserNotFoundError
		}
		return nil, utils.InternalError
	}
	return &user, nil
}

func (r *Repository) GetUserByID(userID int) (*models.Users, error) {
	var user models.Users
	err := r.conn.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) ChangePassword(userID int, oldPassword, newPassword string) error {
	var user models.Users

	if err := r.conn.First(&user, userID).Error; err != nil {
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(user.Password))
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	if err := r.conn.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
 