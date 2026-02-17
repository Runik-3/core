package convert

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// To avoid testing an external lib, this test focuses on the intermediary
// epub shape the dict gets converted to before being processed by kindlegen.
func TestConvertKindleMobi(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "runik_temp_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp testing dir\n")
	}

	_, err = writeEpub(testDict, tempDir)
	if err != nil {
		t.Fatalf("Failed to convert dict\n")
	}
	convertedDictDir := filepath.Join(tempDir, "temp", "epub")
	fmt.Println("dict dir", convertedDictDir)

	convertedWords := readFileData(filepath.Join(convertedDictDir, "OEBPS", "html", "words.xhtml"))
	expectedWords := getFixtureData(filepath.Join("kindleEpub", "words.xhtml"))
	if !bytes.Equal(convertedWords, expectedWords) {
		t.Fatalf(
			"Kindle epub: converted words.xhtml does not match expected\n\nconverted: \n%s\n\nexpected: \n%s",
			convertedWords,
			expectedWords,
		)
	}

	convertedOpf := readFileData(filepath.Join(convertedDictDir, "OEBPS", "content.opf"))
	expectedOpf := getFixtureData(filepath.Join("kindleEpub", "content.opf"))
	if !bytes.Equal(convertedOpf, expectedOpf) {
		t.Fatalf(
			"Kindle epub: converted opf file does not match expected\n\nconverted: \n%s\n\nexpected: \n%s",
			convertedOpf,
			expectedOpf,
		)
	}

	// Cleanup
	os.RemoveAll(tempDir)
}
