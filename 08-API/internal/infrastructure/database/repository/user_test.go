package repository

import (
	"testing"

	"gorm.io/gorm"

	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"github.com/KelpGF/Go-Expert/08-APIs/test/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := setupUser()

	user := defaultUser()
	userRepository := NewUserRepository(db)
	err := userRepository.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.True(t, userFound.ComparePassword("password"))
}

func TestFindByEmail(t *testing.T) {
	db := setupUser()

	user := defaultUser()
	userRepository := NewUserRepository(db)
	db.Create(user)

	userFound, err := userRepository.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

func setupUser() *gorm.DB {
	db := database.Setup()

	db.AutoMigrate(&entity.User{})

	return db
}

func defaultUser() *entity.User {
	return entity.NewUser("user", "user@mail.com", "password")
}
