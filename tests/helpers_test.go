package tests

import (
	"WeekendPOS/app/entity"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func ClearAll() {
	ClearCategories()
	ClearUsers()
}

func ClearUsers() {
	err := db.Where("id IS NOT NULL").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func ClearCategories() {
	err := db.Where("id IS NOT NULL").Delete(&entity.Category{}).Error
	if err != nil {
		log.Fatalf("Failed clear category data : %+v", err)
	}
}

func CreateCategories(user *entity.User, total int) {
	for i := 0; i < total; i++ {
		category := &entity.Category{
			Name:   "Category " + strconv.Itoa(i),
			UserId: user.ID,
		}
		err := db.Create(category).Error
		if err != nil {
			log.Fatalf("Failed create category data : %+v", err)
		}
	}
}

func GetFirstUser(t *testing.T) *entity.User {
	user := new(entity.User)
	err := db.First(user).Error
	assert.Nil(t, err)
	return user
}
