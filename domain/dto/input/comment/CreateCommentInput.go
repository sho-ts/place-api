package input

type CreateCommentInput struct {
	CommentId string
	UserId    string
	PostId    string
	Content   string
}

func NewCreateCommentInput(
	commentId string,
	userId string,
	postId string,
	content string,
) CreateCommentInput {
	return CreateCommentInput{
		CommentId: commentId,
		UserId:    userId,
		PostId:    postId,
		Content:   content,
	}
}
