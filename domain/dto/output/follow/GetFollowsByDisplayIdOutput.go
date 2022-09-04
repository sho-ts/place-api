package output

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type GetFollowsByDisplayIdOutput struct {
	Items []entity.Follow `json:"items"`
	Total int64           `json:"total"`
}

func NewGetFollowsByDisplayIdOutput(
	items []entity.Follow,
	total int64,
) GetFollowsByDisplayIdOutput {
	return GetFollowsByDisplayIdOutput{
		Items: items,
		Total: total,
	}
}
