package core

import (
	"context"
	"fmt"
	"simple-qa/util"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

func QueryDocument(query string, store *weaviate.Store, model *googleai.GoogleAI) (*string, error) {
	ctx := context.Background()
	docs, err := store.SimilaritySearch(ctx, query, 5)

	if err != nil {
		return nil, fmt.Errorf("error searching for similar documents: %w", err)
	}

	docsContents := util.Map(docs, func(doc schema.Document) string {
		return doc.PageContent
	})

	ragQuery := fmt.Sprintf(ragTemplateStr, query, strings.Join(docsContents, "\n"))
	response, err := llms.GenerateFromSinglePrompt(ctx, model, ragQuery)

	if err != nil {
		return nil, fmt.Errorf("error generating response: %w", err)
	}

	return &response, nil
}

const ragTemplateStr = `
You are an specialist in QA, you should answer the given query within these rules:
If the question relates to the context, answer it using the context.
If the answer does not exist in the context, you should answer it with your knowledge and say that your answer it's not in the context.
If the question is unrelated to the context, you should not answer it at all.

Question:
%s

Context:
%s
`
