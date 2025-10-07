package device

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/runik-3/core/convert"
	c "github.com/runik-3/core/core"
)

type Kobo struct {
	Path   string
	Name   string
	Type   string
	AppDir string
}

func (k Kobo) GetPath() string {
	return k.Path
}

func (k Kobo) ReaderName() string {
	return k.Name
}

func (k Kobo) ReaderType() string {
	return k.Type
}

func (k Kobo) DictionaryDir() (string, error) {
	if k.Path == "" {
		return "", errors.New("No device selected.")
	}
	pathCustomDict := filepath.Join(k.Path, ".kobo", "custom-dict")
	pathDict := filepath.Join(k.Path, ".kobo", "dict")

	_, err := os.Stat(pathCustomDict)
	if err == nil {
		return pathCustomDict, nil
	}
	_, err = os.Stat(pathDict)
	if err == nil {
		return pathDict, nil
	}

	return "", errors.New("No dictionary location detected on device.")
}

func (k Kobo) GetDictionaries() ([]c.File, error) {
	koboDictDir, _ := k.DictionaryDir()
	files, err := c.GetFilesFromPath(koboDictDir)
	if err != nil {
		return []c.File{}, err
	}

	deviceDicts := []c.File{}
	for _, file := range files {
		if file.Extension == "zip" {
			// strip dicthtml-[r]
			file.Display = strings.Split(file.Display, "dicthtml-[r]")[1]
			deviceDicts = append(deviceDicts, file)
		}
	}

	return deviceDicts, nil
}

func (k Kobo) DeleteDictionary(name string) error {
	koboDictDir, _ := k.DictionaryDir()
	dictFilePath := filepath.Join(koboDictDir, name)
	err := os.Remove(dictFilePath)
	if err != nil {
		return err
	}
	return nil
}

func (k Kobo) ConvertDictionary(rawDictPath string) (string, error) {
	koboDictDir, err := k.DictionaryDir()
	if err != nil {
		return "", err
	}

	convertedDictPath, err := convert.KoboDicthtml(rawDictPath, koboDictDir, k.AppDir)
	if err != nil {
		return "", err
	}
	return convertedDictPath, nil
}
