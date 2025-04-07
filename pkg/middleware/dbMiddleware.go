package middleware

import (
	"ToGoList/pkg/database"
	"github.com/gin-gonic/gin"
)

func DbMiddleware(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
