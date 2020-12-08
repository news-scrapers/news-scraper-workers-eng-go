package dailyScrapers

import (
	"news-scrapers-workers-eng-go/models"
)

type DailyScraper interface {
	ScrapNewsInItems(baseUrl string, scrapingIndex *models.ScrapingIndex)
}
