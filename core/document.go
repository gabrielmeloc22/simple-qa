package core

import (
	"context"
	"simple-qa/util"

	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

func ProcessDocuments(documents []string, store *weaviate.Store) ([]string, error) {
	ctx := context.Background()

	docs := util.Map(documents, func(document string) schema.Document {
		return schema.Document{PageContent: document}
	})

	return store.AddDocuments(ctx, docs)
}
