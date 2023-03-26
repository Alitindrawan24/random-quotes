package quotes

import "random-quotes/internal/service/quotes"

type Handler struct {
	quotes quotes.ServiceInterface
}

func New(quotes quotes.ServiceInterface) *Handler {
	return &Handler{quotes}
}
