package main

import (
	"log"

	"github.com/taco-tortilla/template-dir-generater/config"
	"github.com/taco-tortilla/template-dir-generater/usecase"
)

func main() {
	loader := config.NewConfigLoader("config.yml")
	config, err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	dirPath := config.Directory.Path + config.Directory.Name

	fileUsecase := usecase.NewFileUsecase()
	if err := fileUsecase.CreateTemplateDirectory(dirPath, config.Files); err != nil {
		log.Fatal(err)
	}

}
