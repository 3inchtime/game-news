package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func GetFami() {
	url := "https://www.famitsu.com/search/?category=news"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: status code %d", resp.StatusCode)
		return
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	var newsList []string
	newsList = getFamiNewsList(html, newsList)

	var wg sync.WaitGroup
	for i := 0; i < len(newsList); i++ {
		wg.Add(1)
		go getFamiNews(newsList[i], &wg)
	}
	wg.Wait()
}

func getFamiNewsList(html *goquery.Document, newsList []string) []string {
	// '//h2[@class="card__title"]/a/@href'
	html.Find("h2[class=card__title]>a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		newsList = append(newsList, url)
	})
	return newsList
}

func getFamiNews(url string, wg *sync.WaitGroup) {
	fmt.Println(url)
	wg.Done()
}
