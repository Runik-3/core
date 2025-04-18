package device

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	path "path/filepath"

	d "github.com/runik-3/builder/dict"
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
	if k.KindleGenPath == "" {
		return "", errors.New("Kindlegen not configured. Please include a path to kindlegen in settings.")
	}

	dict, err := c.DictFromFile(rawDictPath)
	if err != nil {
		return "", err
	}

	epubOpfPath, err := writeEpub(dict, k.AppDir)
	if err != nil {
		return "", err
	}

	// Generate dictionary with kindlegen
	cmd := exec.Command(k.KindleGenPath, epubOpfPath, "-o", fmt.Sprintf("%s.mobi", dict.Name), "-gen_ff_mobi7")
	// This will likely return an error code because we're not specifying a
	// cover image (EmbeddedCover tag) and kindlegen doesn't like that.
	// So let's confirm that the file actually generated. If not, we throw.
	cmd.Run()
	convertedDictionary := path.Join(k.AppDir, "temp", "epub", "OEBPS", fmt.Sprintf("%s.mobi", dict.Name))
	_, err = os.Stat(convertedDictionary)
	if err != nil {
		return "", errors.New("Failed to convert dictionary for Kindle. Make sure you've configured the kindlegen path correctly in settings.")
	}

	// send to device
	convertedDictContents, err := os.ReadFile(convertedDictionary)
	if err != nil {
		return "", err
	}

	deviceDictDir, err := k.DictionaryDir()
	if err != nil {
		return "", err
	}

	convertedDictPath := path.Join(deviceDictDir, fmt.Sprintf("%s.mobi", dict.Name))
	err = os.WriteFile(convertedDictPath, convertedDictContents, os.ModePerm)
	if err != nil {
		return "", err
	}

	// cleanup temp dir after sending to device
	err = os.RemoveAll(path.Join(k.AppDir, "temp", "epub"))
	if err != nil {
		return "", err
	}

	return convertedDictPath, nil
}

const OPF_TEMPLATE string = `<?xml version="1.0" encoding="UTF-8"?>
<package xmlns="http://www.idpf.org/2007/opf" xmlns:opf="http://www.idpf.org/2007/opf" version="3.0" unique-identifier="BookID"> 
  <metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
   <dc:title>{{.Name}}</dc:title>
   <dc:language>{{.Lang}}</dc:language>
   <dc:publisher>runik</dc:publisher>
  </metadata>
  <x-metadata>
   <DictionaryInLanguage>{{.Lang}}</DictionaryInLanguage>
   <DictionaryOutLanguage>{{.Lang}}</DictionaryOutLanguage>
  </x-metadata>
  <manifest>
   <item id="cover" href="html/cover.xhtml" media-type="application/xhtml+xml"/>
   <item id="words" href="html/words.xhtml" media-type="application/xhtml+xml"/>
  </manifest>
  <spine>
   <itemref idref="cover"/>
   <itemref idref="words"/>
  </spine>
</package>`

const WORDS_TEMPLATE string = `<html xmlns:math="http://exslt.org/math" xmlns:svg="http://www.w3.org/2000/svg" xmlns:tl="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf"
xmlns:saxon="http://saxon.sf.net/" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns:cx="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf" xmlns:dc="http://purl.org/dc/elements/1.1/"
xmlns:mbp="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf" xmlns:mmc="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf" xmlns:idx="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf">
<head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"></head>
<body>
<mbp:frameset>
  {{range .Lexicon}}
    <idx:entry scriptable="yes" spell="yes">
      <idx:orth>{{.Word}}</idx:orth>
      <p>{{.Definition}}</p>
    </idx:entry>
  {{end}}
</mbp:frameset>
</body>
</html>`

const COVER_TEMPLATE string = `<html xmlns:math="http://exslt.org/math" xmlns:svg="http://www.w3.org/2000/svg" xmlns:tl="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf"
xmlns:saxon="http://saxon.sf.net/" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns:cx="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf" xmlns:dc="http://purl.org/dc/elements/1.1/"
xmlns:mbp="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf" xmlns:mmc="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf" xmlns:idx="https://kindlegen.s3.amazonaws.com/AmazonKindlePublishingGuidelines.pdf">
<head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"></head>
<body>
  <h1 style="text-align: center;">{{.Name}}</h1>
  <h4 style="text-align: center;">Generated by <a href="https://runik.app">runik</a></h4>
</body>
</html>`

func writeEpub(dict d.Dict, appDir string) (string, error) {
	// Scaffold epub directory
	tempEpub := path.Join(appDir, "temp", "epub")
	c.MkdirIfNotExists(path.Join(tempEpub, "META-INF"))
	c.MkdirIfNotExists(path.Join(tempEpub, "OEBPS", "html"))

	// Write mimetype
	err := os.WriteFile(path.Join(tempEpub, "mimetype"), []byte("application/epub+zip"), os.ModePerm)
	if err != nil {
		return "", err
	}

	// Write container.xml
	const CONTAINER_XML string = `<?xml version="1.0" encoding="UTF-8"?>
<container xmlns="urn:oasis:names:tc:opendocument:xmlns:container" version="1.0">
	<rootfiles>
		<rootfile full-path="OEBPS/content.opf" media-type="application/oebps-package+xml"/>
	</rootfiles>
</container>`
	err = os.WriteFile(path.Join(tempEpub, "META-INF", "container.xml"), []byte(CONTAINER_XML), os.ModePerm)
	if err != nil {
		return "", err
	}

	// Write content.opf
	tmpl, err := template.New("epub_opf").Parse(OPF_TEMPLATE)
	if err != nil {
		return "", err
	}

	opfPath := path.Join(tempEpub, "OEBPS", "content.opf")
	opfFile, err := os.Create(opfPath)
	if err != nil {
		return "", err
	}
	defer opfFile.Close()

	err = tmpl.Execute(opfFile, dict)
	if err != nil {
		return "", err
	}

	// Write words.xhtml
	tmpl, err = template.New("epub_words").Parse(WORDS_TEMPLATE)
	if err != nil {
		return "", err
	}

	wordsFile, err := os.Create(path.Join(tempEpub, "OEBPS", "html", "words.xhtml"))
	if err != nil {
		return "", err
	}
	defer wordsFile.Close()

	err = tmpl.Execute(wordsFile, dict)
	if err != nil {
		return "", err
	}

	// Write cover.xhtml
	tmpl, err = template.New("epub_cover").Parse(COVER_TEMPLATE)
	if err != nil {
		return "", err
	}

	coverFile, err := os.Create(path.Join(tempEpub, "OEBPS", "html", "cover.xhtml"))
	if err != nil {
		return "", err
	}
	defer coverFile.Close()

	err = tmpl.Execute(coverFile, dict)
	if err != nil {
		return "", err
	}

	return opfPath, nil
}
