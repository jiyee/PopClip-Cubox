package main

import (
	"log"
	"os"
	"strings"

	"github.com/jiyee/cubox-cli/client"
)

func main() {
	api := os.Getenv("POPCLIP_OPTION_API")
	tag := os.Getenv("POPCLIP_OPTION_TAG")
	folder := os.Getenv("POPCLIP_OPTION_FOLDER")
	rawContent := os.Getenv("POPCLIP_TEXT")
	browserTitle := os.Getenv("POPCLIP_BROWSER_TITLE")
	browserURL := os.Getenv("POPCLIP_BROWSER_URL")
	selectedText := strings.TrimSpace(rawContent)

	if api == "" || browserURL == "" {
		log.Fatalf("lack of necessary arguments")
	}

	var tags []string
	if tag != "" {
		tags = strings.Split(tag, ",")
	}

	link := client.Link{
		Type:        "url",
		API:         api,
		Content:     browserURL,
		Title:       browserTitle,
		Description: selectedText,
		Tags:        tags,
		Folder:      folder,
	}

	responseMessage, err := link.Submit(true)

	if err != nil {
		switch err.(type) {
		case *client.ResponseError:
			re, _ := err.(*client.ResponseError)

			log.Println(err)

			if re.StatusCode >= 400 && re.StatusCode < 500 {
				os.Exit(2)
			} else {
				os.Exit(1)
			}
		default:
			log.Println(err)
			os.Exit(1)
		}
	}

	log.Println(responseMessage)
	os.Exit(0)
}
