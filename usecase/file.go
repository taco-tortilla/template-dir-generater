package usecase

import (
	"github.com/taco-tortilla/template-dir-generater/config"
	"github.com/taco-tortilla/template-dir-generater/infrastructure"
)

type FileUsecase struct {
	fs infrastructure.FileSystem
}

func NewFileUsecase(fileSystem *infrastructure.FileSystem) *FileUsecase {
	return &FileUsecase{fs: *fileSystem}
}

func (fu *FileUsecase) CreateTemplateDirectory(dirPath string, files []config.File) error {
	if err := fu.fs.CreateDirectory(dirPath); err != nil {
		return err
	}

	if err := fu.fs.CopyFile(files, dirPath); err != nil {
		return err
	}

	return nil
}
