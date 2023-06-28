package starters

import (
	"mongodb.com/users/application/handlers"
)

func StartHandlers() {
	users.NewUserHandler()
}
