package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuonda/go-projects/api-rest/go-api-mongodb/models"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// type PostResult struct {
// 	id        string `json:"id"`
// 	title     string `json:"title"`
// 	content   string `json:"content"`
// 	createdAt string `json:"createdAt"`
// }

func FindBooks(c *gin.Context) {

}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error : ": err.Error()})
	}

	post := models.POST{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func GetProductoById(c *gin.Context) {
	var input models.POST
	id := c.Param("id")

	err := models.DB.Where("id = ?", id).Find(&input).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})

}

func GetAll(c *gin.Context) {
	var posts []models.POST
	models.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func Update(c *gin.Context) {
	var post models.POST
	id := c.Param("id")
	err := models.DB.Where("id  = ?", id).Find(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post Not Found"})
	}

	//Validate Input
	var postInput CreatePostInput
	error := c.ShouldBindJSON(&postInput)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al obtener los datos"})
		return
	}

	models.DB.Model(&post).Updates(models.POST{
		Title:   postInput.Title,
		Content: postInput.Content,
	})

	c.JSON(http.StatusOK, gin.H{"message ": "Update Post"})

}
