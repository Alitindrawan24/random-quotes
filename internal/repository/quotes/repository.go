package quotes

import (
	"github.com/spf13/afero"
	"random-quotes/internal/entity"

	"github.com/gin-gonic/gin"
)

type RepositoryInterface interface {
	GetQuotes(ctx *gin.Context) ([]entity.Quotes, error)
}

type Repository struct {
	fs afero.Fs
}

func New() RepositoryInterface {
	return &Repository{
		fs: afero.NewOsFs(),
	}
}
