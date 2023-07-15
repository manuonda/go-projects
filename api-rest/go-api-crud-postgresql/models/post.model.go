package models

import "fmt"

type Post struct {
	ID      uint   `json:"id" gorm:"id"`
	Title   string `json:"title" gorm:"title"`
	Content string `json:"content" gorm:"content"`

	Comments []Comment `gorm:"ForeignKey:post_id"`
}

func (post Post) ToString() string {
	return fmt.Sprintf("id: %d\n title: %s\n content: %s", post.ID, post.Title, post.Content)
}
