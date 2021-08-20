package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/VladMak/ServiceAPI"
	"log"
)

type ChatServer struct {
	messages []*Message
	users []*ServiceAPI.User
}

// Структура ответа чата
type ChatResponse struct {
	ChatId int `json:"chatId" binding:"required"`
	UserId int `json:"userId" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// Получаем данные с клиента в формате JSON
func (ch *ChatServer) Test(c *gin.Context) {
	var chatResponse ChatResponse
	if err := c.BindJSON(&chatResponse); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(chatResponse.ChatId, chatResponse.UserId, chatResponse.Message)
	c.JSON(http.StatusOK, map[string]interface{}{
		"test": "Yes",
	})
}

//func (ch *ChatServer) GetMessageById() {}
//func (ch *ChatServer) GetMessagesAll() {}

func (ch *ChatServer) SendMessage(c *gin.Context) {
	var chatResponse ChatResponse
	// Получаем ID клиента, который отправил сообщение, ID чата и само сообщение (еще потом наверное надо будет добавить время)
	if err := c.BindJSON(&chatResponse); err != nil {
		log.Println(err)
		return
	}



	c.JSON(http.StatusOK, map[string]interface{}{
		"idUs": "sd",
	})
}