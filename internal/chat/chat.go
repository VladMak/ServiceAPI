package chat

import (
	"fmt"
)

type Chat struct {
	chatId string
}

func (c *Chat) Test(c *gin.Context) {
	fmt.Println("Hello")
	c.JSON(http.StatusOK, map[string]interface{}{
		"test": "Yes",
	})
}