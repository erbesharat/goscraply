package main

import (
  //"fmt"
  "os"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  file,err := os.Create("books.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  // Change the link below to your link
  doc,err := goquery.NewDocument("https://www.goodreads.com/shelf/show/horror?page=1")
  if err != nil {
    panic(err)
  }

  doc.Find(".elementList .left").Each(func(i int, s *goquery.Selection) {
    bookName := s.Find(".bookTitle").Text()
    author := s.Find(".authorName").Text()
    file.WriteString(bookName + "-" + author + "\n")
  })
}
