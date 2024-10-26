package infra

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/googleai"
)

const modelName = "gemini-1.5-flash"

func GetLanguageModel() *googleai.GoogleAI {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	llm, err := googleai.New(ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultModel(modelName),
	)

	if err != nil {
		log.Fatalf("error getting language model: %s", err)
	}

	return llm
}

func GetEmbeddingModel(model embeddings.EmbedderClient) *embeddings.EmbedderImpl {
	emb, err := embeddings.NewEmbedder(model)

	if err != nil {
		log.Fatalf("error getting embedding model: %s", err)
	}

	return emb
}
