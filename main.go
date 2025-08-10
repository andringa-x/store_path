package main

import (
    "context"
    "fmt"
    "log"
    "google.golang.org/genai"
)

func main() {
	var query string = "What is 5 plus 5"

	geminiCall(query)
}

func geminiCall(query string) {
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
    }

    fmt.Println(result.Text())
}
