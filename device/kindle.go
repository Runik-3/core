package device

import (
	c "github.com/runik-3/core/core"
)

type Kindle struct {
	Path string
	Type string
}

func (k Kindle) GetPath() string {
	return k.Path
}
func (k Kindle) ReaderType() string {
	return k.Type
}

func (k Kindle) DictionaryDir() (string, error) {
	return "", nil
}

func (k Kindle) GetDictionaries() ([]c.File, error) {
	return []c.File{}, nil
}

func (k Kindle) DeleteDictionary(name string) error {
	return nil
}

func (k Kindle) ConvertDictionary(rawDictPath string) (string, error) {
	return "", nil
}
