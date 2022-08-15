package input

type CreateCommentInput struct {
	Id      string
	UserId  string
	PostId  string
	Content string
}