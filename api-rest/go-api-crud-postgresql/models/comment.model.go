package models

import "fmt"

type Comment struct {
	ID      uint   `json:"id" gorm:"id"`
	Content string `json:"content" gorm:"content"`
	PostID  uint   `json:"post_id" gorm:"column:post_id"`
	Post    Post
}

func (comment Comment) ToString() string {
	return fmt.Sprintf("ID %d, content %s, Post ID %d, Post Post", comment.ID, comment.Content, comment.PostID)

}
