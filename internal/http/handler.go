package http

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/srbhgalinde/url-shortner/internal/models"
)

const (
	Baseurl = "http://localhost:8080/"
)

var (
	urlStore           = make(map[string]string) // backhalf -> originalURL
	reverseStore       = make(map[string]string) // originalURL -> backhalf
	metricsStore       = make(map[string]int)    // Domain -> Count
	metricsMu          sync.Mutex                // for concurrent RW metricsStore
	validBackhalfRegex = regexp.MustCompile(`[^a-zA-Z0-9_-]`)
)

func shortenURLHandler(c *gin.Context) {
	var req models.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// If URL was already shortened, return respective backhalf
	if existing, ok := reverseStore[req.URL]; ok {
		shortURL := fmt.Sprintf("%s%s", Baseurl, existing)
		c.JSON(http.StatusOK, gin.H{
			"shortUrl": shortURL,
			"backhalf": existing,
		})
		return
	}

	parsedURL := req.URL
	if !regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+\-.]*://`).MatchString(parsedURL) {
		parsedURL = "http://" + parsedURL
	}
	parsed, err := url.Parse(parsedURL)
	if err != nil || parsed.Host == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL - Host parsing issue"})
		return
	}
	domain := parsed.Host

	// Eles, generate new backhalf
	backhalf := req.Backhalf
	if backhalf == "" {
		backhalf = uuid.New().String()[:8]
	} else {
		// Sanitize backhalf
		backhalf = validBackhalfRegex.ReplaceAllString(backhalf, "")
		if backhalf == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid backhalf"})
			return
		}
	}

	// Ensure backhalf is unique
	if _, exists := urlStore[backhalf]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Backhalf already taken"})
		return
	}

	// Save mappings
	urlStore[backhalf] = req.URL
	reverseStore[req.URL] = backhalf

	shortURL := fmt.Sprintf("%s%s", Baseurl, backhalf)
	c.JSON(http.StatusCreated, gin.H{
		"shortUrl": shortURL,
		"backhalf": backhalf,
	})
	go updateDomainMetrics(domain)
}

func redirectHandler(c *gin.Context) {
	backhalf := c.Param("backhalf")

	originalURL, exists := urlStore[backhalf]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	// Redirect to orignal url
	c.Redirect(http.StatusFound, originalURL)
}

func metricsHandler(c *gin.Context) {
	if len(metricsStore) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No URLs shortened",
		})
	}
	topDomains := getTopDomains(3)
	c.JSON(http.StatusOK, gin.H{
		"message": "Top 3 domains",
		"data":    topDomains,
	})

}
