package convert

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	d "github.com/runik-3/builder/dict"
)

var testDict d.Dict = d.Dict{
	Name:   "Test",
	ApiUrl: "example.com",
	Lang:   "en",
	Lexicon: d.Lexicon{
		d.Entry{Word: "Simple", Definition: "This is a simple entry"},
		d.Entry{Word: "Simple Two", Definition: "Another simple entry"},
		d.Entry{
			Word:       "Synonyms",
			Definition: "This is an entry with synonyms",
			Synonyms:   []string{"First", "Second", "Third"},
		},
	},
}

func TestConvertStarDict(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "runik_temp_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp testing dir\n")
	}
	dict, err := testDict.Write(tempDir, "json")
	if err != nil {
		t.Fatalf("Failed to write test dict\n")
	}

	convertedDictDir, err := StarDict(dict, tempDir, "") // app config dir not needed for stardict
	if err != nil {
		t.Fatalf("Failed to convert dict\n")
	}
	fmt.Println("dict dir", convertedDictDir)

	convertedDict := readFileData(filepath.Join(convertedDictDir, "Test.dict"))
	expectedDict := getFixtureData(filepath.Join("stardict", "stardict.dict"))
	if !bytes.Equal(convertedDict, expectedDict) {
		t.Fatalf(
			"Stardict: converted dict does not match expected\n\nconverted: \n%b\n\nexpected: \n%b",
			convertedDict,
			expectedDict,
		)
	}

	convertedIdx := readFileData(filepath.Join(convertedDictDir, "Test.idx"))
	expectedIdx := getFixtureData(filepath.Join("stardict", "stardict.idx"))
	if !bytes.Equal(convertedIdx, expectedIdx) {
		t.Fatalf(
			"Stardict: converted idx does not match expected\n\nconverted: \n%b\n\nexpected: \n%b",
			convertedIdx,
			expectedIdx,
		)
	}

	convertedIfo := readFileData(filepath.Join(convertedDictDir, "Test.ifo"))
	expectedIfo := getFixtureData(filepath.Join("stardict", "stardict.ifo"))
	if !bytes.Equal(convertedIfo, expectedIfo) {
		t.Fatalf(
			"Stardict: converted ifo does not match expected\n\nconverted: \n%b\n\nexpected: \n%b",
			convertedIfo,
			expectedIfo,
		)
	}

	convertedSyn := readFileData(filepath.Join(convertedDictDir, "Test.syn"))
	expectedSyn := getFixtureData(filepath.Join("stardict", "stardict.syn"))
	if !bytes.Equal(convertedSyn, expectedSyn) {
		t.Fatalf(
			"Stardict: converted syn does not match expected\n\nconverted: \n%b\n\nexpected: \n%b",
			convertedSyn,
			expectedSyn,
		)
	}

	// Cleanup
	os.RemoveAll(tempDir)
}

// TODO split out into common test file
func readFileData(path string) []byte {
	text, _ := os.ReadFile(path)
	return text
}

func getFixtureData(fixture string) []byte {
	pathToFixture, _ := filepath.Abs(filepath.Join("fixtures", fixture))
	return readFileData(pathToFixture)
}
