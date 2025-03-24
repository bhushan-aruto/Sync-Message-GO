package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vithsutra/ca-chat-sync-message-service/internals/app"
)

type messageHandler struct {
	Service *app.MessageService
}

func NewMessageHandler(service *app.MessageService) *messageHandler {
	return &messageHandler{
		Service: service,
	}
}

func (h *messageHandler) FetchMessages(c echo.Context) error {

	userId := c.Param("user_id")
	// currentMessageId := c.QueryParam("current_message_id")
	// currentMessageId := c.Param("current_message_id")
	currentMessageId := c.QueryParam("current_message_id")

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing user_id"})
	}
	if currentMessageId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing current_message_id"})
	}
	message, err := h.Service.GetOrderedMessages(userId, currentMessageId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error fetching messages"})
	}
	return c.JSON(http.StatusOK, message)

}
