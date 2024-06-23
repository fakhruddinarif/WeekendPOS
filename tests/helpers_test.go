package tests

import (
	"WeekendPOS/app/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ClearUsers() {
	err := db.Where("id IS NOT NULL").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func GetFirstUser(t *testing.T) *entity.User {
	user := new(entity.User)
	err := db.First(user).Error
	assert.Nil(t, err)
	return user
}
