package application

import (
	"github.com/sho-ts/place-api/application/controller"
	interapter "github.com/sho-ts/place-api/application/interapter/user"
	"github.com/sho-ts/place-api/repository"
)

// repository
var userRepository = repository.NewUserRepository()

// controller
var UserController = controller.NewUserController(
	interapter.NewUserCreateInterapter(userRepository),
	interapter.NewFindByIdInterapter(userRepository),
	interapter.NewFindByDisplayIdInterapter(userRepository),
)