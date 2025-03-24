package device

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	d "github.com/runik-3/builder/dict"
	c "github.com/runik-3/core/core"

	"github.com/pgaskin/dictutil/dictgen"
	"github.com/pgaskin/dictutil/kobodict"
	_ "github.com/pgaskin/dictutil/kobodict/marisa"
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

	convertedDictPath, err := convertForKobo(rawDictPath, koboDictDir)
	if err != nil {
		return "", err
	}
	return convertedDictPath, nil
}

func convertForKobo(pathToRawDict string, outputDir string) (string, error) {
	dict, err := c.DictFromFile(pathToRawDict)
	if err != nil {
		return "", err
	}

	df, err := d.Format("df", dict)
	if err != nil {
		return "", err
	}

	dictFile, err := dictgen.ParseDictFile(strings.NewReader(df))
	if err != nil {
		return "", err
	}

	err = dictFile.Validate()
	if err != nil {
		return "", err
	}

	outFileName := path.Join(outputDir, getDeviceReadableName(dict.Name))
	file, err := os.Create(outFileName)
	defer file.Close()
	if err != nil {
		return "", err
	}

	kw := kobodict.NewWriter(file)

	err = dictFile.WriteDictzip(kw, new(dictgen.ImageHandlerRemove), dictgen.ImageFuncFilesystem)
	if err != nil {
		return "", err
	}

	defer kw.Close()

	return outFileName, nil
}

// TODO - Dictionary naming needs more investigation.
//
// Due to differences in firmware, this seems like the best way to handle
// custom dictionary names. Newer Kobos don't appear to have a dictionary
// table in the sqlite database to inject custom dict names based on iso
// language codes, so we're falling back on using the dicthtml prefixed
// dictionary name. Because we don't want the first couple of letters of
// dict title to be misread as a language code this also needs to be
// prefixed with a character, hence the [r].
//
// This isn't ideal because users will see `dicthtml` prefix is not very
// friendly.
func getDeviceReadableName(dictName string) string {
	return "dicthtml-[r]" + dictName + ".zip"
}
