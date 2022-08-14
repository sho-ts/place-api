package input

type CreatePostInput struct {
	PostId  string
	UserId  string
	Caption string
	Urls    []string
}
