package core

import (
	"log"
	"os"
	"strings"
	"time"

	j "encoding/json"
	d "github.com/runik-3/builder/dict"
)

type File struct {
	Name      string
	Display   string
	Extension string
	Size      int64
	Modified  time.Time
}

func GetUserConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return configDir
}

func MkdirIfNotExists(path string) {
	appDirErr := os.MkdirAll(path, os.ModePerm)
	if appDirErr != nil {
		log.Fatal(appDirErr)
	}
}

// Read raw dictionary and unmarshal as Dict
func DictFromFile(path string) (d.Dict, error) {
	rawDict, err := os.ReadFile(path)
	if err != nil {
		return d.Dict{}, err
	}

	dict := d.Dict{}
	err = j.Unmarshal(rawDict, &dict)
	if err != nil {
		return d.Dict{}, err
	}

	return dict, nil
}

func GetFilesFromPath(path string) ([]File, error) {
	files := []File{}

	fileEntries, err := os.ReadDir(path)
	if err != nil {
		return []File{}, err
	}
	for _, entry := range fileEntries {
		info, err := entry.Info()
		if err != nil {
			return []File{}, err
		}

		fileParts := strings.Split(entry.Name(), ".")
		ext := ""
		if len(fileParts) > 1 {
			ext = fileParts[1]
		}
		files = append(files, File{
			Name:      entry.Name(),
			Display:   fileParts[0],
			Extension: ext,
			Size:      info.Size(),
			Modified:  info.ModTime(),
		})
	}

	return files, nil
}
