package managers

import (
	"news-scrapers-workers-eng-go/dailyScrapers"
	"news-scrapers-workers-eng-go/models"
	"news-scrapers-workers-eng-go/utils"
	"sync"

	log "github.com/sirupsen/logrus"
)

type ScraperManagerAllDailyNewspapers struct {
}

func (mainScraper ScraperManagerAllDailyNewspapers) StartScraping(config models.ScrapingConfig) {

	scraperDiarioEs := dailyScrapers.TheSunUkFullIndexManager{Config: config}

	log.Info("using daily Scrapers:")
	for {
		var wg sync.WaitGroup

		scrapAll := utils.StringInSlice("all", config.NewsPaper)

		if utils.StringInSlice("eldiario.es", config.NewsPaper) || scrapAll {
			go mainScraper.ScrapOneIteration(scraperDiarioEs, "eldiario.es", config, &wg)
			wg.Add(1)
		}
		//add the rest

		wg.Wait()
		log.Info("-------------------------------------------------------------------------------------------------")
		log.Info("-------------------Finished one iteration, all news from page scraped----------------------------")
		log.Info("-------------------------------------------------------------------------------------------------")
	}

}

func (mainScraper *ScraperManagerAllDailyNewspapers) ScrapOneIteration(scraper dailyScrapers.DailyScraper, source string, config models.ScrapingConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Info("starting scraping using " + source)
	scrapingIndex, err := models.GetCurrentIndex(config.ScraperId, source, "daily")

	if scrapingIndex==nil || err != nil {
		scrapingIndex = models.CreateScrapingIndex(config, source)
	}

	scrapingIndex.UpdateUrls(config, source)

	index := scrapingIndex.UrlIndex
	if index >= len(scrapingIndex.StartingUrls) {
		index = 0
	}

	log.Printf("starting with url number %d", index)

	nextUrl := scrapingIndex.StartingUrls[index]
	scraper.ScrapNewsInItems(nextUrl, scrapingIndex)

	//models.SaveMany(results, config)

	index = index + 1

	scrapingIndex.UrlIndex = index

	scrapingIndex.Save()
}
