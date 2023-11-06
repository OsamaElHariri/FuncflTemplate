package models

import (
	"time"

	"github.com/segmentio/ksuid"
)

type Animal struct {
	Id        string
	Name      string
	Sound     string
	CreatedAt string
}

func NewAnimal(name, sound string) Animal {
	return Animal{
		Id:        "animal#" + ksuid.New().String(),
		Name:      name,
		Sound:     sound,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
}
