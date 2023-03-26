package quotes

import "github.com/gin-gonic/gin"

func (handler *Handler) HandleRandomQuotes(ctx *gin.Context) {

	value, err := handler.quotes.GenerateRandomQuote(ctx)
	if err != nil {
		ctx.JSON(400, err.Error())
	}

	ctx.JSON(200, value)
}
