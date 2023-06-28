package users

var UserHandler *UserHandlerStruct

type UserHandlerStruct struct{}

func NewUserHandler() *UserHandlerStruct {
	UserHandler = &UserHandlerStruct{}

	return UserHandler
}
