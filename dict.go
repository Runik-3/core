package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/runik-3/core/core"
)

type File struct {
	Name      string
	Display   string
	Extension string
	Size      int64
	Modified  time.Time
}

func (a *App) GetDictFiles() []File {
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

		fileParts := strings.Split(entry.Name(), ".")
		files = append(files, File{
			Name:      entry.Name(),
			Display:   fileParts[0],
			Extension: fileParts[1],
			Size:      info.Size(),
			Modified:  info.ModTime(),
		})
	}

	return files
}

func (a *App) DeleteDictFile(name string) core.Response[string] {
	dictFilePath := filepath.Join(a.dictionaryDir, name)
	err := os.Remove(dictFilePath)
	if err != nil {
		return core.Response[string]{Data: "", Error: err.Error()}
	}
	return core.Response[string]{Data: "", Error: ""}
}
