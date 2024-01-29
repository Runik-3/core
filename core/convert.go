package core

import (
	"fmt"
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
		fmt.Println(err.Error())
	}

	dict := d.Dict{}
	err = j.Unmarshal(rawDict, &dict)
	if err != nil {
		fmt.Println(err.Error())
	}

	df, err := d.Format("df", dict)
	if err != nil {
		fmt.Println(err.Error())
	}

	dictFile, err := dictgen.ParseDictFile(strings.NewReader(df))
	if err != nil {
		fmt.Println(err.Error())
	}

	err = dictFile.Validate()
	if err != nil {
		fmt.Println(err.Error())
	}

	outFileName := path.Join(outputDir, "dicthtml-"+dict.Name+".zip")
	file, err := os.Create(outFileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	kw := kobodict.NewWriter(file)

	err = dictFile.WriteDictzip(kw, new(dictgen.ImageHandlerRemove), dictgen.ImageFuncFilesystem)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer kw.Close()

	return "", nil
}
