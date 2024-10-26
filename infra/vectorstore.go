package infra

import (
	"cmp"
	"log"
	"os"

	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

func GetVectorStore() *weaviate.Store {
	model := GetLanguageModel()

	emb := GetEmbeddingModel(model)

	store, err := weaviate.New(
		weaviate.WithEmbedder(emb),
		weaviate.WithScheme("http"),
		weaviate.WithHost("localhost:"+cmp.Or(os.Getenv("WVPORT"), "9035")),
		weaviate.WithIndexName("Document"),
	)

	if err != nil {
		log.Fatalf("error getting vector store: %s", err)
	}

	return &store
}
