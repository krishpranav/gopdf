package services

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

func SendToChatGPT(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Extract meaningful insights from the following text"},
			{Role: "user", Content: text},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func ProcessTextWithChatGPT(text string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, // Use GPT-4 or GPT-3.5
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a helpful AI that summarizes and extracts key points."},
				{Role: "user", Content: text},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}