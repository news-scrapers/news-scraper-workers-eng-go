package dailyScrapers

import (
	"fmt"
	"news-scrapers-workers-eng-go/models"
	"testing"

	"github.com/joho/godotenv"
)

func TestRecursiveScraperDiario(t *testing.T) {
	t.Skip()
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{UrlBase: "http://localhost:8000", ScraperId: "testScraperElpais", DeviceID: "testDeviceElpais"}
	index := models.ScrapingIndex{ScraperID: "test", PageIndex: 25}
	scraper := TheSunUkFullIndexManager{Config: config}

	//baseUrl := "https://www.amazon.es/gp/bestsellers/?ref_=nav_cs_bestsellers"
	baseUrl := "https://www.eldiario.es/politica/"

	scraper.ScrapNewsInItems(baseUrl, &index)

}
