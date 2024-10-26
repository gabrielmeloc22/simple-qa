package context

import (
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

type Context struct {
	VectorStore   *weaviate.Store
	LanguageModel *googleai.GoogleAI
}
