package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	d "github.com/runik-3/builder/dict"
	"github.com/runik-3/core/convert"
	c "github.com/runik-3/core/core"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) BuildDictionary(wikiUrl string, name string, depth int, format string) c.Response[d.Dict] {
	// handle name collisions
	dicts := a.GetLocalDictionaries()
	if dicts.Error != "" {
		return c.Response[d.Dict]{Data: d.Dict{}, Error: dicts.Error}
	}
	for _, dict := range dicts.Data {
		if dict.Display == name {
			return c.Response[d.Dict]{Data: d.Dict{}, Error: fmt.Sprintf("A dictionary called '%s' already exists, try again with a unique name.", name)}
		}
	}

	dict, err := d.BuildDictionary(wikiUrl, d.GeneratorOptions{
		Name:         name,
		Output:       a.dictionaryDir,
		EntryLimit:   300000, // This should cover some of the largest fan wikis
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

func (a *App) ExportDictionary(dictType string, dicts []string) c.Response[string] {
	outDir := a.selectDirectory(runtime.OpenDialogOptions{})
	if outDir == "" {
		return c.Response[string]{Data: "", Error: "No export destination selected."}
	}

	var err error
	for _, dictName := range dicts {
		dictPath := filepath.Join(a.dictionaryDir, dictName)

		switch dictType {
		case "stardict":
			_, err = convert.StarDict(dictPath, outDir, a.runikDir)

		case "dicthtml":
			_, err = convert.KoboDicthtml(dictPath, outDir, a.runikDir)

		case "mobi":
			_, err = convert.KindleMobi(dictPath, outDir, a.runikDir, a.config.KindlegenPath)

		default:
			return c.Response[string]{Data: "", Error: fmt.Sprintf("%s is not a valid dictionary format", dictType)}
		}
	}

	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}

	return c.Response[string]{Data: "Export successful", Error: ""}
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

/*
dictName: the current dict name including extension
newName: raw name not including extension (eg. .json)
*/
func (a *App) RenameLocalDictionary(dictName string, newName string) c.Response[string] {
	originalPath := filepath.Join(a.dictionaryDir, dictName)

	// Load existing dict into memory
	dict, err := c.DictFromFile(originalPath)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}

	// Change name
	dict.Name = newName
	// Write dict to new path
	newPath, err := dict.Write(a.dictionaryDir, "json")
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}

	// Delete old file
	err = os.Remove(originalPath)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}

	return c.Response[string]{Data: newPath, Error: ""}
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
