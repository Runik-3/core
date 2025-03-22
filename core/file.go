package core

import (
	"log"
	"os"
	"strings"
	"time"
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
