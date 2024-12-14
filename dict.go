package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	d "github.com/runik-3/builder/dict"
	c "github.com/runik-3/core/core"
)

type File struct {
	Name      string
	Display   string
	Extension string
	Size      int64
	Modified  time.Time
}

func (a *App) GetLocalDictionaries() c.Response[[]File] {
	files, err := getDictFilesFromPath(a.dictionaryDir)
	if err != nil {
		return c.Response[[]File]{Data: []File{}, Error: err.Error()}
	}

	localDicts := []File{}
	for _, file := range files {
		if file.Extension == "json" {
			localDicts = append(localDicts, file)
		}
	}

	return c.Response[[]File]{Data: localDicts, Error: ""}
}

func (a *App) GetDeviceDictionaries() c.Response[[]File] {
	// if no device connected, do no work and return empty with no error
	if a.devicePath == "" {
		return c.Response[[]File]{Data: []File{}, Error: ""}
	}

	// TODO: This is a kobo specific solution. Generalize when we have
	// support for other readers.
	koboDictDir, _ := c.FindKoboDictDir(a.devicePath)
	files, err := getDictFilesFromPath(koboDictDir)
	if err != nil {
		return c.Response[[]File]{Data: []File{}, Error: err.Error()}
	}

	deviceDicts := []File{}
	for _, file := range files {
		if file.Extension == "zip" {
			deviceDicts = append(deviceDicts, file)
		}
	}

	return c.Response[[]File]{Data: deviceDicts, Error: ""}
}

func (a *App) ReadLocalDictionary(name string) c.Response[d.Dict] {
	path := filepath.Join(a.dictionaryDir, name)

	data, err := os.ReadFile(path)
	if err != nil {
		return c.Response[d.Dict]{Data: d.Dict{}, Error: err.Error()}
	}

	var dict d.Dict
	err = json.Unmarshal(data, &dict)

	return c.Response[d.Dict]{Data: dict, Error: ""}
}

func (a *App) DeleteLocalDictFile(name string) c.Response[string] {
	dictFilePath := filepath.Join(a.dictionaryDir, name)
	err := os.Remove(dictFilePath)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}
	return c.Response[string]{Data: "", Error: ""}
}

func (a *App) DeleteDeviceDictFile(name string) c.Response[string] {
	if a.devicePath == "" {
		return c.Response[string]{Data: "", Error: "No device connected"}
	}

	// TODO: This is a kobo specific solution. Generalize when we have
	// support for other readers.
	koboDictDir, _ := c.FindKoboDictDir(a.devicePath)
	dictFilePath := filepath.Join(koboDictDir, name)
	err := os.Remove(dictFilePath)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}
	return c.Response[string]{Data: "", Error: ""}
}

func getDictFilesFromPath(path string) ([]File, error) {

	files := []File{}

	fileEntries, err := os.ReadDir(path)
	if err != nil {
		return []File{}, err
	}
	for _, entry := range fileEntries {
		info, err := entry.Info()
		if err != nil {
			return []File{}, err
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

	return files, nil
}
