package application

import (
	"github.com/sho-ts/place-api/application/controller"
	ic "github.com/sho-ts/place-api/application/interactor/comment"
	ifo "github.com/sho-ts/place-api/application/interactor/follow"
	il "github.com/sho-ts/place-api/application/interactor/like"
	ip "github.com/sho-ts/place-api/application/interactor/post"
	iu "github.com/sho-ts/place-api/application/interactor/user"
	"github.com/sho-ts/place-api/repository"
)

// repository
var userRepository = repository.NewUserRepository()
var postRepository = repository.NewPostRepository()
var storageRepository = repository.NewStorageRepository()
var commentRepository = repository.NewCommentRepository()
var likeRepository = repository.NewLikeRepository()
var followRepository = repository.NewFollowRepository()

// controller
var UserController = controller.NewUserController(
	iu.NewUserCreateInteractor(userRepository),
	iu.NewFindByIdInteractor(userRepository),
	iu.NewFindByDisplayIdInteractor(userRepository),
	iu.NewChangeProfileInteractor(userRepository, storageRepository),
)
var PostController = controller.NewPostController(
	ip.NewCreatePostInteractor(postRepository, storageRepository),
	ip.NewFindByIdInteractor(postRepository),
	ip.NewFindAllInteractor(postRepository),
)
var CommentController = controller.NewCommentController(
	ic.NewCreateCommentInteractor(commentRepository),
	ic.NewFindAllInteractor(commentRepository),
)
var LikeController = controller.NewLikeController(
	il.NewToggleLikeInteractor(likeRepository),
)
var FollowController = controller.NewFollowController(
	ifo.NewToggleFollowInteractor(followRepository),
  ifo.NewGetFollowsByDisplayIdInteractor(followRepository),
  ifo.NewGetFollowersByDisplayIdInteractor(followRepository),
)
