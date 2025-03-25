package device

import (
	"errors"
	"path/filepath"

	c "github.com/runik-3/core/core"
)

type Device interface {
	GetPath() string
	ReaderName() string
	ReaderType() string
	DictionaryDir() (string, error)
	GetDictionaries() ([]c.File, error)
	DeleteDictionary(string) error
	ConvertDictionary(string) (string, error)
}

// TODO: This func needs to take config
func NewDevice(dir string, appDir string) (Device, error) {
	name := filepath.Base(dir)

	if isKobo(dir) {
		return Kobo{
			Path:   dir,
			Name:   name,
			Type:   "kobo",
			AppDir: appDir,
		}, nil
	}

	if isKindle(dir) {
		kindlegenPath := filepath.FromSlash(`/Applications/Kindle Previewer 3.app/Contents/lib/fc/bin/kindlegen`)
		return Kindle{
			Path:          dir,
			Name:          name,
			Type:          "kindle",
			AppDir:        appDir,
			KindleGenPath: kindlegenPath,
		}, nil
	}

	return Kobo{}, errors.New("Could not detect a supported e-reader")
}

func isKindle(dir string) bool {
	kindle := Kindle{Path: dir}
	_, err := kindle.DictionaryDir()
	if err != nil {
		return false
	}
	return true
}

func isKobo(dir string) bool {
	kobo := Kobo{Path: dir}
	_, err := kobo.DictionaryDir()
	if err != nil {
		return false
	}
	return true
}
