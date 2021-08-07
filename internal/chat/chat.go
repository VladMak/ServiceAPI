package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Chat struct {
	chatId string
}

func (ch *Chat) Test(c *gin.Context) {
	fmt.Println("Hello")
	c.JSON(http.StatusOK, map[string]interface{}{
		"test": "Yes",
	})
}