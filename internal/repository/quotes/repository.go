package quotes

import (
	"random-quotes/internal/entity"

	"github.com/gin-gonic/gin"
)

type RepositoryInterface interface {
	GetQuotes(ctx *gin.Context) ([]entity.Quotes, error)
}

type Repository struct {
}

func New() RepositoryInterface {
	return &Repository{}
}
