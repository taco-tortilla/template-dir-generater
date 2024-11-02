package infrastructure

import (
	"io"
	"os"

	"github.com/taco-tortilla/template-dir-generater/config"
)

const SRC_TEMPLATE_FILE_PATH = "templates"

type FileSystem struct{}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (fs *FileSystem) CreateDirectory(dirPath string) error {
	if err := os.Mkdir(dirPath, 0755); err != nil {
		return err
	}
	return nil
}

func (fs *FileSystem) CopyFile(files []config.File, dirPath string) error {
	for _, f := range files {
		file := f.File
		name := f.Name

		if name == "" {
			name = file
		}

		src, err := os.Open(SRC_TEMPLATE_FILE_PATH + "/" + file)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(dirPath + "/" + name)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return nil
}
