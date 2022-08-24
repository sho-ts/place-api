package application

import (
	"github.com/sho-ts/place-api/application/controller"
	ic "github.com/sho-ts/place-api/application/interapter/comment"
	ifo "github.com/sho-ts/place-api/application/interapter/follow"
	il "github.com/sho-ts/place-api/application/interapter/like"
	ip "github.com/sho-ts/place-api/application/interapter/post"
	iu "github.com/sho-ts/place-api/application/interapter/user"
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
	iu.NewUserCreateInterapter(userRepository),
	iu.NewFindByIdInterapter(userRepository),
	iu.NewFindByDisplayIdInterapter(userRepository),
	iu.NewChangeProfileInterapter(userRepository, storageRepository),
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
var LikeController = controller.NewLikeController(
	il.NewToggleLikeInterapter(likeRepository),
)
var FollowController = controller.NewFollowController(
	ifo.NewToggleFollowInterapter(followRepository),
)
