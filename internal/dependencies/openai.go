package dependencies

import (
	"context"
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

func RequestLLM(prompt string, ctx *context.Context) (*string, error) {
	client := GetOpenAIClient()
	if client == nil {
		panic("open ai client is nil after injection")
	}
	var chatContext context.Context
	if ctx == nil {
		chatContext = context.TODO()
	} else {
		chatContext = *ctx
	}
	chatCompletion, err := client.Chat.Completions.New(
		chatContext,
		openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			}),
			Model: openai.F(openai.ChatModelGPT4o),
		},
	)
	return &chatCompletion.Choices[0].Message.Content, err
}
