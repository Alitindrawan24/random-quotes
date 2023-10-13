package quotes

import (
	"encoding/json"
	"fmt"

	"random-quotes/internal/entity"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetQuotes(ctx *gin.Context) ([]entity.Quotes, error) {

	file, err := repository.fs.Open("data/quotes.json")
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	decode := json.NewDecoder(file)
	quotes := []entity.Quotes{}
	err = decode.Decode(&quotes)

	if err != nil {
		return quotes, err
	}

	return quotes, nil
}
