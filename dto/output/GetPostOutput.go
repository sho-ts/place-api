package output

import (
	"github.com/sho-ts/place-api/entity"
	"time"
)

type GetPostOutput struct {
	PostId    string           `json:"postId"`
	Caption   string           `json:"caption"`
	CreatedAt time.Time        `json:"createdAt"`
	Liked     int              `json:"liked"`
	Files     []entity.Storage `json:"files"`
	User      entity.User      `json:"user"`
}

func NewGetPostOutput(
	postId string,
	caption string,
	createdAt time.Time,
	liked int,
	files []entity.Storage,
	user entity.User,
) GetPostOutput {
	return GetPostOutput{
		PostId:    postId,
		Caption:   caption,
		CreatedAt: createdAt,
		Liked:     liked,
		Files:     files,
		User:      user,
	}
}
