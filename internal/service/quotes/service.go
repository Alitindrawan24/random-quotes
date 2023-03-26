package quotes

import (
	"random-quotes/internal/entity"
	"random-quotes/internal/repository/quotes"

	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	GenerateRandomQuote(ctx *gin.Context) (entity.Quotes, error)
}

type Service struct {
	quotes quotes.RepositoryInterface
}

func New(quotes quotes.RepositoryInterface) ServiceInterface {
	return &Service{quotes}
}
