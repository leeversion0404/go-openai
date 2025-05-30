package main

import (
	"context"
	"fmt"
	"os"

	"github.com/leeversion0404/go-openai"
)

func main() {
	client := openai.NewClient(os.Getenv("apikey"))

	respUrl, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         "buatkan saya gambar wanita sedang bermain gitar di taman",
			Size:           openai.CreateImageSize256x256,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	fmt.Println(respUrl.Data[0].URL)
}
