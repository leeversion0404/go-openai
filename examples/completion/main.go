package main

import (
	"context"
	"fmt"
	"os"

	"github.com/leeversion0404/go-openai"
)

func main() {
	client := openai.NewClient(os.Getenv("apikey"))
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:     openai.GPT3Babbage002,
			MaxTokens: 5,
			Prompt:    "Bulan Sutena",
		},
	)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
