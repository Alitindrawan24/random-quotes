package quotes

import (
	"encoding/json"
	"os"
	"random-quotes/internal/entity"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetQuotes(ctx *gin.Context) ([]entity.Quotes, error) {

	file, _ := os.Open("data/quotes.json")
	defer file.Close()

	decode := json.NewDecoder(file)
	quotes := []entity.Quotes{}
	err := decode.Decode(&quotes)

	if err != nil {
		return quotes, err
	}

	return quotes, nil
}
