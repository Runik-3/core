package core

import (
	"os"
	"path"
	"strings"

	j "encoding/json"

	"github.com/pgaskin/dictutil/dictgen"
	"github.com/pgaskin/dictutil/kobodict"
	_ "github.com/pgaskin/dictutil/kobodict/marisa"
	d "github.com/runik-3/builder/dict"
)

// currently only supports Kobo readers
func ConvertForReader(pathToRawDict string, outputDir string) (string, error) {
	// read raw dictionary and unmarshal as Dict
	rawDict, err := os.ReadFile(pathToRawDict)
	if err != nil {
		return "", err
	}

	dict := d.Dict{}
	err = j.Unmarshal(rawDict, &dict)
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
// table in the sqlite database to inject custom dict names so we're
// falling back on a prefix.
//
// This isn't ideal because the `dicthtml` prefix is not very user friendly.
// Long term,
func getDeviceReadableName(dictName string) string {
	return "dicthtml-[r]" + dictName + ".zip"
}
