package application

import (
	"github.com/sho-ts/place-api/application/controller"
	ic "github.com/sho-ts/place-api/application/interapter/comment"
	ip "github.com/sho-ts/place-api/application/interapter/post"
	iu "github.com/sho-ts/place-api/application/interapter/user"
	"github.com/sho-ts/place-api/repository"
)

// repository
var userRepository = repository.NewUserRepository()
var postRepository = repository.NewPostRepository()
var storageRepository = repository.NewStorageRepository()
var commentRepository = repository.NewCommentRepository()

// controller
var UserController = controller.NewUserController(
	iu.NewUserCreateInterapter(userRepository),
	iu.NewFindByIdInterapter(userRepository),
	iu.NewFindByDisplayIdInterapter(userRepository),
)
var PostController = controller.NewPostController(
	ip.NewCreatePostInterapter(postRepository, storageRepository),
	ip.NewFindByIdInterapter(postRepository),
	ip.NewFindAllInterapter(postRepository),
)
var CommentController = controller.NewCommentController(
	ic.NewCreateCommentInterapter(commentRepository),
  ic.NewFindAllInterapter(commentRepository),
)
