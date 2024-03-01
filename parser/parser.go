package parser

import (
	"fmt"
	"net/http"
	"nosuparser/logger"
	"strings"

	"github.com/gocolly/colly"
)

type News struct {
	Title string
	Data  string
	Link  string
}

func ParseNosuNews(startPage, endPage uint) ([]News, error) {
	newsList := make([]News, 0, (endPage-startPage+1)*10)

	c := colly.NewCollector(
		colly.AllowedDomains("www.nosu.ru"),
	)

	c.OnHTML("div.news-item", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		title := e.ChildText(".title")
		date := e.ChildText(".date")
		newsList = append(newsList, News{Title: title, Link: link, Data: date})

	})

	c.OnRequest(func(r *colly.Request) {
		logger.WriteLogs("Visiting", r.URL.String())
	})

	header := http.Header{}
	header.Set("authority", "www.nosu.ru")
	header.Set("accept", "*/*")
	header.Set("accept-language", "ru,en;q=0.9")
	header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	header.Set("cookie", "PHPSESSID=e41450a0ea388cac10c478c9acd9057e; _ym_uid=170922193429642890; _ym_d=1709221934; _ym_isad=1")
	header.Set("dnt", "1")
	header.Set("origin", "https://www.nosu.ru")
	header.Set("referer", "https://www.nosu.ru/category/news/")
	header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "YaBrowser";v="24.1", "Yowser";v="2.5"`)
	header.Set("sec-ch-ua-mobile", "?0")
	header.Set("sec-ch-ua-platform", `"Linux"`)
	header.Set("sec-fetch-dest", "empty")
	header.Set("sec-fetch-mode", "cors")
	header.Set("sec-fetch-site", "same-origin")
	header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36")
	header.Set("x-requested-with", "XMLHttpRequest")

	for i := startPage; i <= endPage; i++ {
		err := c.Request(
			http.MethodPost,
			"https://www.nosu.ru/category/news/",
			strings.NewReader(fmt.Sprintf("paged=%v", i)),
			nil,
			header,
		)
		if err != nil {
			return newsList, err
		}
	}

	return newsList, nil
}
