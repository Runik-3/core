package main

import (
	"os"

	"github.com/labstack/gommon/log"
)

type File struct {
	Name string
	Size int64
}

func (a *App) GetRawDicts() []File {
	files := []File{}

	dirEntries, err := os.ReadDir(a.dictionaryDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range dirEntries {
		info, infoErr := entry.Info()
		if infoErr != nil {
			log.Fatal(infoErr)
		}

		files = append(files, File{Name: entry.Name(), Size: info.Size()})
	}

	return files
}
