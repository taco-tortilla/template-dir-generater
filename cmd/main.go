package main

import (
	"log"

	"github.com/taco-tortilla/template-dir-generater/infrastructure"
	handler "github.com/taco-tortilla/template-dir-generater/interface"
	"github.com/taco-tortilla/template-dir-generater/usecase"
)

func main() {
	fileSystem := infrastructure.NewFileSystem()
	fileUsecase := usecase.NewFileUsecase(fileSystem)
	cli_handler := handler.NewCLIHandler(fileUsecase)

	if err := cli_handler.Run(); err != nil {
		log.Fatal(err)
	}

}
