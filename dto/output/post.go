package output

import (
	"github.com/sho-ts/place-api/entity"
)

type GetPostResponseOutput struct {
	PostId string `json:"postId"`
	UserId string `json:"userId"`
  Caption string `json:"caption"`
  Files []entity.Storage `json:"files"`
}
