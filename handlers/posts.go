package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rootkill/genie/models"
)

func (h *PostHandler) GetPosts(c *gin.Context) {
	var posts []models.Post
	h.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&post)
	c.JSON(http.StatusCreated, post)
}
