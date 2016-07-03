package main

import (
  //"fmt"
  "os"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  file,err := os.Create("items.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  // Change the link below to your link
  doc,err := goquery.NewDocument("http://www.listchallenges.com/top-100-modern-horror")
  if err != nil {
    panic(err)
  }

  doc.Find(".col-list-items .clearfix .list-item .item-click-area").Each(func(i int, s *goquery.Selection) {
    itemName := s.Find("div .item-name").Text()
    file.WriteString(itemName + "\n")
  })
}
