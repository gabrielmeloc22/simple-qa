package main

import (
	"log"
	"simple-qa/infra"
	"simple-qa/server/context"
	"simple-qa/server/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	vectorStore := infra.GetVectorStore()
	languageModel := infra.GetLanguageModel()

	context := context.Context{VectorStore: vectorStore, LanguageModel: languageModel}

	routes.Routes(r, context)

	r.Run()
}
