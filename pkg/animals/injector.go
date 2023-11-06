package animals

import (
	"funcfltemplate/pkg/animals/repositories/animals"
	"funcfltemplate/pkg/infrastructure_provider"

	"github.com/samber/do"
)

func getInjector() *do.Injector {

	injector := do.New()
	defer injector.Shutdown()

	do.Provide(injector, func(i *do.Injector) (animals.AnimalsRepo, error) {
		return &animals.AnimalsRepoDynamo{
			Client:    infrastructure_provider.GetDynamo(),
			TableName: infrastructure_provider.GetTableName(),
		}, nil
	})

	return injector
}
