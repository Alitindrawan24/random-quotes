package app

import "random-quotes/internal/handler/quotes"

type Handlers struct {
	Quotes *quotes.Handler
}

func NewHandler(service *Services) *Handlers {
	return &Handlers{
		Quotes: quotes.New(service.quotes),
	}
}
