package dependencies

import (
	"os"
	"sync"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var (
	openaiClient *openai.Client
	openaiOnce   sync.Once
)

func GetOpenAIClient() *openai.Client {
	openaiOnce.Do(func() {
		key := os.Getenv("OPEN_AI_KEY")
		openaiClient = openai.NewClient(option.WithAPIKey(key))
	})
	return openaiClient
}
