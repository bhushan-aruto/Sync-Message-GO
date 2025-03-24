// package router

// import (
// 	"github.com/labstack/echo/v4"
// 	"github.com/vithsutra/ca-chat-sync-message-service/internals/app"
// 	"github.com/vithsutra/ca-chat-sync-message-service/internals/delivery/http/handler"
// 	"github.com/vithsutra/ca-chat-sync-message-service/internals/domain"
// )

// func InitRoutes(e *echo.Echo, messagerepo domain.Messagerepo) {

// 	messageService, ok := messagerepo.(*app.MessageService)
// 	if !ok {
// 		e.Logger.Error("Failed to cast messagerepo to MessageService")
// 		return
// 	}

// 	messageOrderHandler := handler.NewMessageHandler(messageService)

// 	e.GET("/messages/:userId/:messageId", messageOrderHandler.FetchMessages)
// }

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/vithsutra/ca-chat-sync-message-service/internals/app"
	"github.com/vithsutra/ca-chat-sync-message-service/internals/delivery/http/handler"
)

func InitRoutes(e *echo.Echo, messageService *app.MessageService) {

	messageOrderHandler := handler.NewMessageHandler(messageService)

	e.GET("/messages/:user_id", messageOrderHandler.FetchMessages)

}
