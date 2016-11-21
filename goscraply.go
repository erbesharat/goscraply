package goscraply

import (
	"github.com/PuerkitoBio/goquery"
)

// Goodreads scraper
// Gets a category and scraps top 50 books on that category
func goodreads(category string) []string {
	var books []string
	url := "https://www.goodreads.com/shelf/show/" + category + "?page=1"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	doc.Find(".elementList .left").Each(func(i int, s *goquery.Selection) {
		bookName := s.Find(".bookTitle").Text()
		author := s.Find(".authorName").Text()
		result := bookName + "-" + author
		books = append(books, result)
	})
	return books
}

// Listchallenges scraper
// Gets a category and scraps top 50 items on that category
func lastchallenges(link string) []string {
	var items []string
	url := link
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	doc.Find(".col-list-items .clearfix .list-item .item-click-area").Each(func(i int, s *goquery.Selection) {
		itemName := s.Find("div .item-name").Text()
		items = append(items, itemName)
	})
	return items
}
