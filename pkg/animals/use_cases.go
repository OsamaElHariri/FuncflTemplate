package animals

import (
	"github.com/samber/do"
)

type UseCases struct {
	Injector *do.Injector
}

func NewUseCases() UseCases {
	return UseCases{
		Injector: getInjector(),
	}
}

func (usecases *UseCases) MakeSound() string {
	return "*Animal noises"
}
