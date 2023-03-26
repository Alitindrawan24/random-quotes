package app

import "random-quotes/internal/repository/quotes"

type Repositories struct {
	quotes quotes.RepositoryInterface
}

func NewRepository() *Repositories {
	return &Repositories{
		quotes: quotes.New(),
	}
}
