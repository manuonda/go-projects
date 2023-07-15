package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/database"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/dto"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/models"
)

func CreateComent(c *gin.Context) {

	var comment dto.CommentRequest
	err := c.BindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := database.DB.Create(&models.Comment{
		Content: comment.Comment,
		PostID:  comment.PostID,
	})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created Message"})
	return

}

func GetAll(c *gin.Context) {
	var comments []models.Comment
	err := database.DB.Preload("Post").Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
	return

}
