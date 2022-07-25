package actions

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"news-bot/db"

	"github.com/PuerkitoBio/goquery"
)

func HabrGo(tag string) {
	urlMain := "https://habr.com"
	URL := fmt.Sprintf("%s/ru/hub/%s/", urlMain, tag)

	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	linkAll := doc.Find(".tm-articles-list__item")
	link, _ := linkAll.Find(".tm-article-snippet__readmore").Attr("href")
	linkText, _ := linkAll.Find(".tm-article-snippet__title-link").Find("span").Html()
	fmt.Println(linkText, link)
	linkMD5Sum := md5.Sum([]byte(link))

	text := fmt.Sprintf(`[%s]<a href\=\"%s%s\">%s</a>`, tag, urlMain, link, linkText)
	fmt.Println(text)
	db.CheckSiteNewsBot(URL, link, text, fmt.Sprintf("%x", linkMD5Sum))
}
