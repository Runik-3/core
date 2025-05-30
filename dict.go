package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	d "github.com/runik-3/builder/dict"
	c "github.com/runik-3/core/core"
)

func (a *App) BuildDictionary(wikiUrl string, name string, depth int, format string) c.Response[d.Dict] {
	dict, err := d.BuildDictionary(wikiUrl, d.GeneratorOptions{
		Name:         name,
		Output:       a.dictionaryDir,
		EntryLimit:   25000,
		Depth:        depth,
		Format:       "json",
		ProgressHook: a.emitProgress,
	})
	if err != nil {
		return c.Response[d.Dict]{Data: d.Dict{}, Error: err.Error()}
	}
	return c.Response[d.Dict]{Data: dict, Error: ""}
}

func (a *App) ConvertDictionary(fileName string) c.Response[string] {
	if a.device == nil || a.device.GetPath() == "" {
		return c.Response[string]{Data: "", Error: "No device connected"}
	}

	rawDictPath := filepath.Join(a.dictionaryDir, fileName)
	dictPath, err := a.device.ConvertDictionary(rawDictPath)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}
	return c.Response[string]{Data: dictPath, Error: ""}
}

func (a *App) GetLocalDictionaries() c.Response[[]c.File] {
	files, err := c.GetFilesFromPath(a.dictionaryDir)
	if err != nil {
		return c.Response[[]c.File]{Data: []c.File{}, Error: err.Error()}
	}

	localDicts := []c.File{}
	for _, file := range files {
		if file.Extension == "json" {
			localDicts = append(localDicts, file)
		}
	}

	return c.Response[[]c.File]{Data: localDicts, Error: ""}
}

func (a *App) GetDeviceDictionaries() c.Response[[]c.File] {
	// if no device connected, do no work and return empty with no error
	if a.device == nil || a.device.GetPath() == "" {
		return c.Response[[]c.File]{Data: []c.File{}, Error: ""}
	}

	deviceDicts, err := a.device.GetDictionaries()
	if err != nil {
		return c.Response[[]c.File]{Data: []c.File{}, Error: err.Error()}
	}

	return c.Response[[]c.File]{Data: deviceDicts, Error: ""}
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

func (a *App) WriteLocalDictionary(dict d.Dict) c.Response[string] {
	path, err := dict.Write(a.dictionaryDir, "json")
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}
	return c.Response[string]{Data: path, Error: ""}
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
	if a.device == nil || a.device.GetPath() == "" {
		return c.Response[string]{Data: "", Error: "No device connected"}
	}

	err := a.device.DeleteDictionary(name)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}
	return c.Response[string]{Data: "", Error: ""}
}
