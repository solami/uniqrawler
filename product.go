package main

import (
	"fmt"

	"github.com/sclevine/agouti"
)

// 製品名を取得するためのクラス
const nameClass = "uq-ec-product-tile__end-product-name"

// Product 製品
type Product struct {
	name string
}

func (p *Product) String() string {
	return fmt.Sprintf("name: %s", p.name)
}

// ToProduct 要素を製品に変換する
func ToProduct(selection *agouti.Selection) (*Product, error) {
	// 製品の名称を取得
	nameElement := selection.FindByClass(nameClass)
	name, err := nameElement.Text()
	if err != nil {
		return nil, err
	}

	// 製品を生成して返却
	return &Product{
		name: name,
	}, nil
}
