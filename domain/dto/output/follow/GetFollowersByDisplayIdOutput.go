package output

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type GetFollowersByDisplayIdOutput struct {
	Items []entity.Follow `json:"items"`
	Total int64           `json:"total"`
}

func NewGetFollowersByDisplayIdOutput(
	items []entity.Follow,
	total int64,
) GetFollowersByDisplayIdOutput {
	return GetFollowersByDisplayIdOutput{
		Items: items,
		Total: total,
	}
}
