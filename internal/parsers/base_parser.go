package parsers

type BaseParser struct {
	ParserService
}

func (p BaseParser) Parse() {
	p.ParserService.parseServiceResponse()
}
