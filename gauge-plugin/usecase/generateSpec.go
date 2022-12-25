package usecase

import (
	"play-cdc/domain"
	"play-cdc/repository"
)

func GenerateSpec() {
	for _, e := range repository.GetProviderEnvs() {
		requests := repository.LoadRecordedRequests()

		contracts := domain.ToContracts(requests)

		consumerName := repository.GetConsumerName()
		if consumerName == "" {
			panic("Error: consumer name is empty!")
		}

		spec := contracts.ToSpec(consumerName, e.ProviderName)
		spec.SortScenarios()

		outputBasePath := repository.GetOutputBasePath()
		if outputBasePath == "" {
			panic("Error: output base path is empty!")
		}
		outputPath := domain.OutputPath(outputBasePath, e.ProviderName)

		err := repository.SaveSpec(spec, domain.SpecFilePath(outputPath, consumerName))
		if err != nil {
			panic(err)
		}

		repository.SaveRequestBodies(contracts, domain.RequestBodiesPath(outputPath, consumerName))
	}
}