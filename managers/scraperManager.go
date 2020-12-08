package managers

import "news-scrapers-workers-eng-go/models"

type ScraperManager interface {
	StartScraping(config models.ScrapingConfig)
}
