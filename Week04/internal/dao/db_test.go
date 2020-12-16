package dao

import (
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	dao, err := NewUserDao(db)
	if err != nil {
		log.Fatal(err)
	}

	u, _ := dao.GetUser("u1")
	log.Printf("u=%v", u)
}
