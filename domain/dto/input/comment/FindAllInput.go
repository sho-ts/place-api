package input

type FindAllInput struct {
	PostId string
	Limit  int
	Offset int
}

func NewFindAllInput(
	postId string,
	limit int,
	offset int,
) FindAllInput {
	return FindAllInput{
		PostId: postId,
		Limit:  limit,
		Offset: offset,
	}
}
