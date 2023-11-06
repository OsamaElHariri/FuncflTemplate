package animals

import "funcfltemplate/pkg/animals/models"

type AnimalDb struct {
	PK        string
	SK        string
	Name      string
	Sound     string
	CreatedAt string
}

func fromAppEntity(entity models.Animal) AnimalDb {
	return AnimalDb{
		PK:        entity.Id,
		SK:        entity.Id,
		Name:      entity.Name,
		Sound:     entity.Sound,
		CreatedAt: entity.CreatedAt,
	}
}

func toAppEntity(entity AnimalDb) models.Animal {
	return models.Animal{
		Id:        entity.PK,
		Name:      entity.Name,
		Sound:     entity.Sound,
		CreatedAt: entity.CreatedAt,
	}
}
