package handler

import (
	"errors"
	"log"
	"os"

	"github.com/taco-tortilla/template-dir-generater/config"
	"github.com/taco-tortilla/template-dir-generater/usecase"
)

type CLIHandler struct {
	fileUsecase *usecase.FileUsecase
}

func NewCLIHandler(fileUsecase *usecase.FileUsecase) *CLIHandler {
	return &CLIHandler{fileUsecase: fileUsecase}
}

func (cli *CLIHandler) Run() error {

	loader := config.NewConfigLoader("config.yml")
	config, err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	var dirPath string

	args := os.Args

	if len(args) > 3 {
		return errors.New(`
		Error: Too many arguments
		Expeceted args: <dir-path> <number-of-creation>

		Example:
			./tdg ../sample 2
		`)
	}

	if len(args) > 1 {
		dirPath = args[1]
	} else {
		dirPath = config.Directory.Path + config.Directory.Name
	}

	return cli.fileUsecase.CreateTemplateDirectory(dirPath, config.Files)

}
