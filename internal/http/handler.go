package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/srbhgalinde/url-shortner/internal/models"
)

const (
	Baseurl = "http://localhost:8080/"
)

var (
	urlStore     = make(map[string]string) // backhalf -> originalURL
	reverseStore = make(map[string]string) // originalURL -> backhalf
)

func shortenURLHandler(c *gin.Context) {
	var req models.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// If URL was already shortened, return respective backhalf
	if existing, ok := reverseStore[req.URL]; ok {
		shortURL := fmt.Sprintf("%s/%s", Baseurl, existing)
		c.JSON(http.StatusOK, gin.H{
			"shortUrl": shortURL,
			"backhalf": existing,
		})
		return
	}

	// eles, generate new backhalf
	backhalf := req.Backhalf
	if backhalf == "" {
		backhalf = uuid.New().String()[:8]
	}

	// Ensure backhalf is unique
	if _, exists := urlStore[backhalf]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "backhalf already taken"})
		return
	}

	// Save mappings
	urlStore[backhalf] = req.URL
	reverseStore[req.URL] = backhalf

	shortURL := fmt.Sprintf("%s/%s", Baseurl, backhalf)
	c.JSON(http.StatusCreated, gin.H{
		"shortUrl": shortURL,
		"backhalf": backhalf,
	})
}
