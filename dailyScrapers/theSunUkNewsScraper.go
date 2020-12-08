package dailyScrapers

import (
	uuid "github.com/nu7hatch/gouuid"
	"news-scrapers-workers-eng-go/models"
	"strings"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type TheSunUkNewsScraper struct {
	Config models.ScrapingConfig
}


func (scraper *TheSunUkNewsScraper) ScrapNewUrl(urlNew models.UrlNew) models.NewScraped {
	result := models.NewScraped{}

	ajaxUrl := urlNew.Url
	// Instantiate default collector
	c := colly.NewCollector(

	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})
	date := ""
	tags := []string{}


	c.OnHTML("meta", func(e *colly.HTMLElement) {
		if (e.Attr("property") == "article:published_time") {
			date = e.Attr("content")
		}
		if (e.Attr("property") == "article:tag") {
			if strings.Contains(e.Attr("content"), ","){
				tags = strings.Split(e.Attr("content"), ",")

			} else {
				tags = append(tags,e.Attr("content"))
			}
		}
	})

	c.OnHTML(".article-switcheroo", func(e *colly.HTMLElement) {
			headline := ""
			content := ""
			tags := []string{}


		e.ForEach(".article__content article__content--intro", func(_ int, elem *colly.HTMLElement) {
				headline = strings.TrimSpace(elem.Text)
			})

			e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
				content = content + " " + elem.Text
			})


			layout := "2006-01-02"
			date = strings.Split(date, "T")[0]
			t, _ := time.Parse(layout, date)
			result.Url=urlNew.Url
			result.Headline=headline
			result.ScraperID = scraper.Config.ScraperId

			result.NewsPaper = "thesun.uk"
			result.Content = strings.TrimSpace(content)
			result.Date = t
			result.DateString = date
			result.Tags = tags
			u, _ := uuid.NewV4()
			result.ID = u.String()

			log.Println("obtained new with headline " + headline)


	})


	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong:", err)
	})

	c.Visit(ajaxUrl)
	c.Wait()

	return result

}
