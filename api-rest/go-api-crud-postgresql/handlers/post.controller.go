package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/database"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/dto"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/models"
)

func GetAllPost(c *gin.Context) {

	var posts []models.Post
	result := database.DB.Preload("Comments").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Get All Posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
	return
}

func CreatePost(c *gin.Context) {
	var inputPost dto.PostInput
	err := c.BindJSON(&inputPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error Post Not Found"})
		return
	}

	post := models.Post{
		Title:   inputPost.Title,
		Content: inputPost.Content,
	}
	result := database.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error save register Post Created"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post Created"})
	return
}

func UpdatePost(c *gin.Context) {

}
