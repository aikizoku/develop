package main

import (
	"context"
	"fmt"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func main() {
	ocr("aov.jpg")
}

func ocr(filename string) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		panic(err)
	}

	texts, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		panic(err)
	}

	for _, text := range texts {
		fmt.Println(text.Description)
	}
}
