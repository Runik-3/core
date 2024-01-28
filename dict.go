package main

import (
	"os"
	"strings"

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

		isRawDict := strings.Contains(entry.Name(), ".json")
		if isRawDict {
			dictName := strings.Split(entry.Name(), ".")[0] // remove file extension
			files = append(files, File{Name: dictName, Size: info.Size()})
		}
	}

	return files
}
