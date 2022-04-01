package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	keywordFile := "keywords.list"
	// キーワードリストを取得
	keywords, err := LoadKeywords(keywordFile)
	if err != nil {
		log.Fatalf("Failed to load keywords. %v", err)
	}

	// webdriverの起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	// キーワードごとに検索実行
	for _, keyword := range keywords {
		// ヘッダー出力
		fmt.Printf("-----------------\n")
		fmt.Printf("keyword: %s\n", keyword)
		fmt.Printf("-----------------\n")
		// 検索実行
		products, err := Search(page, keyword)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
		// 検索結果を出力
		for _, product := range products {
			fmt.Println(product.String())
		}
	}
}

// LoadKeywords ファイルパスからキーワードリストを取得
func LoadKeywords(path string) ([]string, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	r := bufio.NewReader(fd)

	var lines []string
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

// 検索ページ
const baseURL = "https://www.uniqlo.com/jp/ja/search"

// 検索バーを特定するID
const searchID = "Search"

// 検索ボタンを特定するクラス
const searchButtonClass = "uq-ec-search__button"

// 製品エリアを特定するクラス
const productAreaClass = "uq-ec-product-tile-resize-wrapper"

// Search 指定したkeywordの製品一覧を返す
func Search(page *agouti.Page, keyword string) ([]*Product, error) {
	// 検索ページに移動
	if err := page.Navigate(baseURL); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	// 検索バーを取得
	search := page.FindByID(searchID)
	// 検索バーにキーワードをセット
	search.Fill(keyword)
	// 検索実行
	if err := page.FindByClass(searchButtonClass).Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	// 製品を表す要素群を取得
	elements := page.AllByClass(productAreaClass)
	// 要素群を製品の配列に変換
	count, _ := elements.Count()
	var products []*Product
	for i := 0; i < count; i++ {
		product, err := ToProduct(elements.At(i))
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
