package output

type CountOutput struct {
	Count int64 `json:"count"`
}

func NewCountOutput(count int64) CountOutput {
	return CountOutput{
		Count: count,
	}
}
