package services

import (
	"fmt"
	"os"

	services "github.com/Phk13/go-user-service"
)

var UserService services.Service

func Setup() {
	UserService = services.CreateService(fmt.Sprintf("%s/api/", os.Getenv("USERS_MS")))
}
