package service

import (
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	user := GetUser("u1")
	log.Println(user)
}
