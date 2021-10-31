package forms

type PostForm struct {
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
	UserId    int    `json:"userId" validate:"required"`
}
