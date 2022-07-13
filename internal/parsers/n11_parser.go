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

type N11Parser struct {
	Content io.Reader
}

func (p N11Parser) parseServiceResponse() []models.ResponseModel {

	doc, err := goquery.NewDocumentFromReader(p.Content)

	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	var items []models.ResponseModel

	doc.Find(".list-ul .column .columnContent .pro").Each(func(i int, s *goquery.Selection) {
		aTag := s.Find("a")

		productTitle, titleExist := aTag.Attr("title")
		url, _ := aTag.Attr("href")
		urlSlice := strings.Split(url, "?")
		url = urlSlice[0]

		priceContainer := s.Find(".proDetail .priceContainer div span ins")

		priceData := strings.Replace(priceContainer.Text(), " TL", "", 1)
		priceField, _ := strconv.ParseFloat(strings.Replace(priceData, ",", ".", 1), 64)

		if titleExist {
			items = append(items, models.ResponseModel{
				Title:      productTitle,
				Url:        url,
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
