package helpers

import (
	"lean-rss/main/src/models"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type articleParser interface {
	image(html goquery.Document) string
	title(html goquery.Document) string
	body(html goquery.Document) string
	author(html goquery.Document) string
	language(html goquery.Document) string
	tags(html goquery.Document) []string
	url(url string) string
	publishedTime(html goquery.Document) time.Time
	scrappedTime(time time.Time) time.Time
	parseArticle(html goquery.Document, url string, scraping_time time.Time, parser articleParser) models.Article
}

// The most general type of parser used by default for URLS,
// which don't have their specific parser assigned.
// For that reason it is more error-prone and has less filtering capabilities than others.
type GenericArticleParser struct {
}

func (parser GenericArticleParser) title(html goquery.Document) string {
	return ""
}

// Returns URL to a cover image of the article.
func (parser GenericArticleParser) image(html goquery.Document) string {
	return ""
}

func (parser GenericArticleParser) body(html goquery.Document) string {
	return ""
}

func (parser GenericArticleParser) author(html goquery.Document) string {
	return ""
}

func (parser GenericArticleParser) language(html goquery.Document) string {
	return ""
}

func (parser GenericArticleParser) tags(html goquery.Document) []string {
	return []string{}
}

// Returns cleaned up, validated URL to the article.
func (parser GenericArticleParser) url(url string) string {
	return url
}

func (parser GenericArticleParser) publishedTime(html goquery.Document) time.Time {
	return time.Now()
}

// Returns formatted time of scrapping of the article
func (parser GenericArticleParser) scrappedTime(time_ time.Time) time.Time {
	return time.Now()
}

func (parser GenericArticleParser) parseArticle(html goquery.Document, url string, scraping_time time.Time) models.Article {
	return models.Article{
		Image:         parser.image(html),
		Title:         parser.title(html),
		Body:          parser.body(html),
		Author:        parser.author(html),
		Language:      parser.language(html),
		Tags:          parser.tags(html),
		URL:           parser.url(url),
		PublishedTime: parser.publishedTime(html),
		ScrappedTime:  parser.scrappedTime(scraping_time),
	}
}

func ParseArticle(html goquery.Document, url string, scraping_time time.Time, parser articleParser) models.Article {
	return parser.parseArticle(html, url, scraping_time)

}

func mapProviderToParser(url string) articleParser {
	providersDomain := justDomain(url)
	providerToParserMapping := map[string]interface{}{
		"spidersweb.pl":     "",
		"thehackernews.com": "",
		"antyweb.pl":        "",
		"default":           GenericArticleParser{},
	}
	return providerToParserMapping[providersDomain]
}

// Converts a full URL to just host for purpose of matching with a parser.
func justDomain(url_ string) string {
	urlStruct, err := url.Parse(url_)
	if err != nil {
		log.Fatal(err)
	}
	return urlStruct.Host
}
