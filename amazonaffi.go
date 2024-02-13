// ./amazonaffi B018WNIBJS

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/buger/jsonparser"

	paapi5 "github.com/goark/pa-api"
	"github.com/goark/pa-api/entity"
	"github.com/goark/pa-api/query"
)

func main() {

	length := len(os.Args)
	if length < 2 {
		fmt.Printf("Error: %s <ASIN>\n", os.Args[0])
		return
	}

	for i := 1; i < length; i++ {
		asinID := os.Args[i]
		makeHTML(asinID, length > 2)
	}
}

func makeHTML(asinID string, maketd bool) {

	amazonAffiliateID := os.Getenv("PA_ASSOCIATE_TAG")

	//Create client
	client := paapi5.New(
		paapi5.WithMarketplace(paapi5.LocaleJapan),
	).CreateClient(
		amazonAffiliateID,
		os.Getenv("PA_ACCESS_KEY"),
		os.Getenv("PA_SECRET_KEY"),
	)

	//Make query for images and info

	q := query.NewGetItems(
		client.Marketplace(),
		client.PartnerTag(),
		client.PartnerType(),
	).ASINs([]string{asinID}).EnableImages().EnableItemInfo()

	//Requet and response
	body, err := client.RequestContext(context.Background(), q)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	//Decode JSON
	res, err := entity.DecodeResponse(body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// fmt.Printf("json: %s\n", res.String())
	data := []byte(res.String())

	// <a href="https://www.amazon.co.jp/dp/B08N6LR285/?tag=gikohadiary-22" target="_blank"><img src="https://m.media-amazon.com/images/I/31b2W4qu5tL._SL200_.jpg" alt="B08N6LR285" border="0" /><br>最新 Apple MacBook Air Apple M1 Chip (13インチPro, 8GB RAM, 256GB SSD) - スペースグレイ</a>

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if maketd {
			fmt.Print("<td>")
		}
		fmt.Printf("<a href=\"https://www.amazon.co.jp/dp/%s/?tag=%s\" target=\"_blank\">",
			asinID, amazonAffiliateID)
		singlevalue, _, _, _ := jsonparser.Get(value, "Images")
		smalljson, _, _, _ := jsonparser.Get(singlevalue, "Primary", "Small")
		url, _, _, _ := jsonparser.Get(smalljson, "URL")
		fmt.Printf("<img src=\"%s\" alt=\"%s\" border=\"0\" />\n", string(url), asinID)
		fmt.Printf("<br>")
		iteminfo, _, _, _ := jsonparser.Get(value, "ItemInfo")
		proinfo, _, _, _ := jsonparser.Get(iteminfo, "Title", "DisplayValue")
		fmt.Printf("%s</a>\n", string(proinfo))
		if maketd {
			fmt.Println("</td>")
		}

	}, "ItemsResult", "Items")

}
