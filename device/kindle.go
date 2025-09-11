package device

import (
	"errors"
	"os"
	path "path/filepath"

	"github.com/runik-3/core/convert"
	c "github.com/runik-3/core/core"
)

type Kindle struct {
	Path          string
	Name          string
	Type          string
	AppDir        string
	KindleGenPath string
}

func (k Kindle) GetPath() string {
	return k.Path
}

func (k Kindle) ReaderName() string {
	return k.Name
}

func (k Kindle) ReaderType() string {
	return k.Type
}

func (k Kindle) DictionaryDir() (string, error) {
	if k.Path == "" {
		return "", errors.New("No device selected.")
	}
	pathDict := path.Join(k.Path, "documents", "dictionaries")

	_, err := os.Stat(pathDict)
	if err != nil {
		return "", errors.New("No dictionary location detected on device.")
	}

	return pathDict, nil
}

func (k Kindle) GetDictionaries() ([]c.File, error) {
	kindleDictDir, _ := k.DictionaryDir()
	files, err := c.GetFilesFromPath(kindleDictDir)
	if err != nil {
		return []c.File{}, err
	}

	deviceDicts := []c.File{}
	for _, file := range files {
		if file.Extension == "mobi" {
			deviceDicts = append(deviceDicts, file)
		}
	}

	return deviceDicts, nil
}

func (k Kindle) DeleteDictionary(name string) error {
	kindleDictDir, _ := k.DictionaryDir()
	dictFilePath := path.Join(kindleDictDir, name)
	err := os.Remove(dictFilePath)
	if err != nil {
		return err
	}
	return nil
}

func (k Kindle) ConvertDictionary(rawDictPath string) (string, error) {
	deviceDir, err := k.DictionaryDir()
	if err != nil {
		return "", err
	}
	return convert.KindleMobi(rawDictPath, deviceDir, k.AppDir, k.KindleGenPath)
}
