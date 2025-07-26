package middleware

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GzipRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.GetHeader("Content-Encoding"), "gzip") {
			gr, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid gzip body"})
				return
			}
			defer func() {
				if err := gr.Close(); err != nil {
					_ = c.Error(fmt.Errorf("gzip reader close error: %w", err))
				}
			}()
			body, err := io.ReadAll(gr)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to read gzip body"})
				return
			}
			c.Request.Body = io.NopCloser(strings.NewReader(string(body)))
		}
		c.Next()
	}
}
