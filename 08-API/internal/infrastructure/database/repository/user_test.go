package repository

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/KelpGF/Go-Expert/08-APIs/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := setup()

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
	db := setup()

	user := defaultUser()
	userRepository := NewUserRepository(db)
	userRepository.Create(user)

	userFound, err := userRepository.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

func setup() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	return db
}

func defaultUser() *entity.User {
	return entity.NewUser("user", "user@mail.com", "password")
}
