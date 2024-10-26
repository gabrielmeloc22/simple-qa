package routes

import (
	"simple-qa/server/context"
	"simple-qa/server/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, context context.Context) {
	handlers := handlers.Handler{Context: &context}

	r.POST("/document", handlers.DocumentPost)
	r.POST("/query", handlers.QueryPost)
}
