package parsers

import (
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aligoren/fiyatine/internal/models"
)

type AmazonParser struct {
	Content io.Reader
}

func (p AmazonParser) parseServiceResponse() []models.ResponseModel {

	doc, err := goquery.NewDocumentFromReader(p.Content)

	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	var items []models.ResponseModel

	doc.Find("#search > div.s-desktop-width-max.s-desktop-content.s-opposite-dir.sg-row > div.s-matching-dir.sg-col-16-of-20.sg-col.sg-col-8-of-12.sg-col-12-of-16 > div > span > div.s-main-slot.s-result-list.s-search-results.sg-row > div > div > div > div > div > div.a-section.a-spacing-small.s-padding-left-small.s-padding-right-small").Each(func(i int, s *goquery.Selection) {

		productTitle := s.Find("h2 a span").Text()
		titleExist := productTitle != ""

		url, _ := s.Find("h2 a").Attr("href")
		urlSlice := strings.Split(url, "/ref")
		url = urlSlice[0]
		priceData := s.Find(".s-price-instructions-style .a-row a .a-offscreen:nth-child(1)").Text()
		prices := strings.Split(priceData, " TL")
		priceData = prices[0]
		priceField, _ := strconv.ParseFloat(strings.Replace(strings.Replace(priceData, ".", "", -1), ",", ".", -1), 64)

		if titleExist && priceField > 0 {
			items = append(items, models.ResponseModel{
				Vendor:     "Amazon",
				Title:      productTitle,
				Url:        fmt.Sprintf("https://www.amazon.com.tr%s", url),
				Price:      fmt.Sprintf("₺%s", priceData),
				PriceField: priceField,
			})
		}
	})

	sort.Slice(items, func(i, j int) bool {
		return items[i].PriceField < items[j].PriceField
	})

	return items
}
