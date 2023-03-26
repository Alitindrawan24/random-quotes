package quotes

import (
	"math/rand"
	"random-quotes/internal/entity"

	"github.com/gin-gonic/gin"
)

func (service *Service) GenerateRandomQuote(ctx *gin.Context) (entity.Quotes, error) {
	quotes, err := service.quotes.GetQuotes(ctx)
	if err != nil {
		return entity.Quotes{}, err
	}

	rand := rand.Intn(len(quotes))
	quote := quotes[len(quotes)-rand]

	return quote, nil
}
