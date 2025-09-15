package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/srbhgalinde/url-shortner/internal/models"
)

const Baseurl = "http://localhost:8080/"

func shortenURLHandler(c *gin.Context) {
	var req models.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id := uuid.New().String()[:8]

	backhalf := req.Backhalf
	if backhalf == "" {
		backhalf = id
	}

	shortURL := fmt.Sprintf("%s/%s", Baseurl, backhalf)

	c.JSON(http.StatusCreated, gin.H{
		"id":       id,
		"shortUrl": shortURL,
	})
}
