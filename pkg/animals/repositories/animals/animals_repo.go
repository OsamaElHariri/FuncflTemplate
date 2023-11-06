package animals

import "funcfltemplate/pkg/animals/models"

type AnimalsRepo interface {
	PutAnimal(entity models.Animal) error
	GetAnimal(id string) (models.Animal, bool, error)
}
