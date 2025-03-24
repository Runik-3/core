package device

import (
	"path/filepath"

	c "github.com/runik-3/core/core"
)

type Device interface {
	GetPath() string
	ReaderType() string
	DictionaryDir() (string, error)
	GetDictionaries() ([]c.File, error)
	DeleteDictionary(string) error
	ConvertDictionary(string) (string, error)
}

func NewDevice(dir string, appDir string) Device {
	// TODO: Expand this function to detect reader type from the path name.
	// Maybe also accept an input for when the consumer knows their device type.
	//
	// Note: be sure to also include a DisplayName prop to Device. Can we infer
	// this from the file path on all operating systems including windows?

	// FIXME: pull this from config -- also introduce config
	// /Applications/Kindle\ Previewer\ 3.app/Contents/lib/fc/bin/kindlegen
	kindlegenPath := filepath.FromSlash(`/Applications/Kindle Previewer 3.app/Contents/lib/fc/bin/kindlegen`)
	return Kindle{Path: dir, Type: "kobo", AppDir: appDir, KindleGenPath: kindlegenPath}
}
