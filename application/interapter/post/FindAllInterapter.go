package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindAllInterapter struct {
	PostRepository repository.IPostRepository
}

func NewFindAllInterapter(
	postRepository repository.IPostRepository,
) FindAllInterapter {
	return FindAllInterapter{
		PostRepository: postRepository,
	}
}

func (interapter FindAllInterapter) Handle(i input.FindAllInput) ([]entity.PostsItem, error) {
	posts, err := interapter.PostRepository.FindAll(i.UserId, i.Limit, i.Offset)

	return posts, err
}
