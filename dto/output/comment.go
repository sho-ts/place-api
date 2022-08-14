package output

type GetCommentsOutput struct {
	Id       string `json:"commentId"`
	Content  string `json:"content"`
	UserId   string `json:"userId"`
	PostId   string `json:"postId"`
	Avatar   string `json:"avatar"`
	UserName string `json:"name"`
}