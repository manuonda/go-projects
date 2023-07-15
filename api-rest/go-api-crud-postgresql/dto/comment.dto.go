package dto

// https://learningprogramming.net/golang/gorm/fetch-data-using-one-to-many-relationship-in-gorm/
type CommentRequest struct {
	Title   string `json:"title"`
	Comment string `json:"comment"`
	PostID  uint   `json:"post_id"`
}
