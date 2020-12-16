package dao

import (
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	user, err := GetUser("u1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}
