package device

import (
	"path/filepath"
	"strings"

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

func NewDevice(dir string, appDir string) Device {
	// TODO: Add some validation -- we should throw an error if we're not in a
	// kindle or kobo.
	// Attempt to detect device type from the path name
	maybeDevice := filepath.Base(dir)
	if strings.Contains(strings.ToLower(maybeDevice), "kindle") {
		// TODO - FIXME: pull this from config -- also introduce config
		// /Applications/Kindle\ Previewer\ 3.app/Contents/lib/fc/bin/kindlegen
		kindlegenPath := filepath.FromSlash(`/Applications/Kindle Previewer 3.app/Contents/lib/fc/bin/kindlegen`)
		return Kindle{
			Path:          dir,
			Name:          maybeDevice,
			Type:          "kindle",
			AppDir:        appDir,
			KindleGenPath: kindlegenPath,
		}
	}

	// FIXME: I'd rather this be explicit and throw an error
	// Default to kobo since it's the best supported
	return Kobo{
		Path:   dir,
		Name:   maybeDevice,
		Type:   "kobo",
		AppDir: appDir,
	}
}
