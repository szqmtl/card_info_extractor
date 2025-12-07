package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	cardInfo, err := ExtractIdentityInfo(ctx, client, "XXXXXX.png")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("card: %+v\n", cardInfo)
	}
}

type IdentityInfo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Issuer     string `json:"issuer"`
	Expiration string `json:"expiration"`
}

func ExtractIdentityInfo(ctx context.Context, client *genai.Client, file string) (*IdentityInfo, error) {
	prompt := "Detect if the image is an Identity card. If yes, extract the issuer, name, the ID, and the expiration of the image in pure json format. If no, return NONE"

	contentType, err := getContentType(file)
	if err != nil {
		return nil, err
	}
	imgBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	parts := []*genai.Part{
		genai.NewPartFromBytes(imgBytes, contentType),
		genai.NewPartFromText(prompt),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", contents, nil)

	if err != nil {
		return nil, err
	}
	// fmt.Println(result.Text())

	if strings.EqualFold(result.Text(), "NONE") {
		return nil, fmt.Errorf("%s is not a valid Identity card", file)
	}

	jsonString, err := extractJsonFromMarkdown(result.Text())
	if err != nil {
		return nil, err
	}

	var info IdentityInfo
	err = json.Unmarshal([]byte(jsonString), &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func getContentType(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = io.ReadFull(file, buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func extractJsonFromMarkdown(markdownText string) (string, error) {
	re := regexp.MustCompile("(?s)```json\\s*\\n(.+?)\\n\\s*```")
	match := re.FindStringSubmatch(markdownText)

	if len(match) > 1 {
		return match[1], nil
	}

	return "", fmt.Errorf("no json block found in markdown text")
}
