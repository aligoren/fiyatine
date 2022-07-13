package parsers

import "fmt"

type HepsiBuradaParser struct {
	Content string
}

func (p HepsiBuradaParser) parseServiceResponse() {
	fmt.Printf("%v", p.Content)
}
