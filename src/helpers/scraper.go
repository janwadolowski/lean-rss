package helpers

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func downloadPage(url string) goquery.Document {

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", response.StatusCode, response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return *doc
}
