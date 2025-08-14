package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/genai"
)

func GeminiCall(query string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(query),
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return result.Text(), nil
}
