package getanimal

import (
	"errors"
	"funcfltemplate/pkg/animals/models"
	"funcfltemplate/pkg/animals/repositories/animals"

	"github.com/samber/do"
)

func Invoke(injector *do.Injector, id string) (models.Animal, error) {
	animalsRepo := do.MustInvoke[animals.AnimalsRepo](injector)

	entity, found, err := animalsRepo.GetAnimal(id)
	if err != nil {
		return models.Animal{}, err
	}
	if !found {
		return models.Animal{}, errors.New("invalid id")
	}

	return entity, nil
}
