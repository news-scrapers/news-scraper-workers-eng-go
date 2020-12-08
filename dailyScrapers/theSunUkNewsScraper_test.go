package dailyScrapers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"news-scrapers-workers-eng-go/models"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestScraperTheSunUk(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	config := models.ScrapingConfig{ScraperId: "testScraperdiario", DeviceID: "testDevicediario"}
	scraper := TheSunUkNewsScraper{Config: config}
	baseUrl := "https://www.thesun.co.uk/tvandshowbiz/13409249/mark-wright-found-car-stolen-essex"
	newUrl := models.UrlNew{baseUrl, time.Now()}

	result := scraper.ScrapNewUrl(newUrl)
	assert.NotEqual(t, result.Content, "")
	assert.NotEmpty(t, result.Tags)
	assert.NotEqual(t, result.Date.IsZero(),true)
	assert.NotEqual(t, result.Headline, "")
	fmt.Println(result)

}
