package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/aligoren/fiyatine/internal/models"
	"github.com/aligoren/fiyatine/internal/parsers"
)

type HepsiBuradaService struct {
	SearchParams models.ProductSearchModel
}

func (service HepsiBuradaService) buildUrl() string {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   "hepsiburada.com",
		Path:   "ara",
	}

	query := requestUrl.Query()

	query.Add("q", service.SearchParams.ProductName)

	requestUrl.RawQuery = query.Encode()

	return requestUrl.String()
}

func (service HepsiBuradaService) searchProduct() {

	baseUrl := service.buildUrl()

	rqeuest, err := http.NewRequest(http.MethodGet, baseUrl, nil)

	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	rqeuest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	rqeuest.Header.Add("referer", "https://www.hepsiburada.com/")
	rqeuest.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	response, err := http.DefaultClient.Do(rqeuest)
	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	parser := parsers.BaseParser{
		ParserService: parsers.HepsiBuradaParser{
			Content: string(body),
		},
	}

	parser.Parse()

}
