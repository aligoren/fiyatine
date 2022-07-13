package parsers

import "github.com/aligoren/fiyatine/internal/models"

type BaseParser struct {
	ParserService
}

func (p BaseParser) Parse() []models.ResponseModel {
	return p.ParserService.parseServiceResponse()
}
