package parsers

import "github.com/aligoren/fiyatine/internal/models"

type ParserService interface {
	parseServiceResponse() []models.ResponseModel
}
