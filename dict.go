package main

import (
	c "github.com/runik-3/core/core"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	Name      string
	Display   string
	Extension string
	Size      int64
	Modified  time.Time
}

func (a *App) GetDictFiles() c.Response[[]File] {
	files := []File{}

	fileEntries, err := os.ReadDir(a.dictionaryDir)
	if err != nil {
		return c.Response[[]File]{Data: []File{}, Error: err.Error()}
	}
	for _, entry := range fileEntries {
		info, err := entry.Info()
		if err != nil {
			return c.Response[[]File]{Data: []File{}, Error: err.Error()}
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

	return c.Response[[]File]{Data: files, Error: ""}
}

func (a *App) DeleteDictFile(name string) c.Response[string] {
	dictFilePath := filepath.Join(a.dictionaryDir, name)
	err := os.Remove(dictFilePath)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}
	return c.Response[string]{Data: "", Error: ""}
}
