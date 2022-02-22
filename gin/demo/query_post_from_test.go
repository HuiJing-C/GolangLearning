package demo

import (
	"github.com/gin-gonic/gin"
	"testing"
)

// POST /post?id=1234&page=1 HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// name=manu&message=this_is_great
func TestQueryAndPostFrom(t *testing.T) {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		age := c.DefaultPostForm("age", "18")
		println(id, page, name, message, age)
	})
	r.Run(":8080")
}
