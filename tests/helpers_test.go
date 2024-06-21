package tests

import (
	"WeekendPOS/app/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetFirstUser(t *testing.T) *entity.User {
	user := new(entity.User)
	err := db.First(user).Error
	assert.Nil(t, err)
	return user
}
