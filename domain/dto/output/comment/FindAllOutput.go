package output

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type FindAllOutput struct {
	Items []entity.Comment `json:"items"`
	Total int64            `json:"total"`
}

func NewFindAllOutput(
	items []entity.Comment,
	total int64,
) FindAllOutput {
	return FindAllOutput{
		Items: items,
		Total: total,
	}
}
