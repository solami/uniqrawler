package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	// ブラウザはChromeを指定して起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	if err := page.Navigate("https://www.uniqlo.com/jp/ja/search"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	search := page.FindByID("Search")
	search.Fill("test")
	if err := page.FindByClass("uq-ec-search__button").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	time.Sleep(1 * time.Second)

	names := page.AllByClass("uq-ec-product-tile__end-product-name")
	count, _ := names.Count()
	for i := 0; i < count; i++ {
		name, _ := names.At(i).Text()
		fmt.Println(name)
	}
	time.Sleep(10 * time.Second)
}
