package application

import (
	"github.com/sho-ts/place-api/application/controller"
	iu "github.com/sho-ts/place-api/application/interapter/user"
	"github.com/sho-ts/place-api/repository"
)

// repository
var userRepository = repository.NewUserRepository()

// controller
var UserController = controller.NewUserController(
	iu.NewUserCreateInterapter(userRepository),
	iu.NewFindByIdInterapter(userRepository),
	iu.NewFindByDisplayIdInterapter(userRepository),
)
