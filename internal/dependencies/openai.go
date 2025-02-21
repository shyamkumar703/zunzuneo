package dependencies

import (
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var (
	openaiClient *openai.Client
)

func GetOpenAIClient() (*openai.Client, error) {
	var err error
	once.Do(func() {
		key := os.Getenv("OPEN_AI_KEY")
		openaiClient = openai.NewClient(option.WithAPIKey(key))
	})
	return openaiClient, err
}
