package app

import "random-quotes/internal/service/quotes"

type Services struct {
	quotes quotes.ServiceInterface
}

func NewService(repository *Repositories) *Services {
	return &Services{
		quotes: quotes.New(repository.quotes),
	}
}
