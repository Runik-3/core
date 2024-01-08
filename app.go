package main

import (
	"context"
	c "github.com/runik-3/core/core"
	"log"
	"path/filepath"

	"github.com/runik-3/builder/pkg/builder"
	"github.com/runik-3/builder/pkg/dict"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	// store app config files
	runikDir string
	// store raw dictionary files
	dictionaryDir string
	// path to the user selected reader device
	devicePath string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.checkAppDirExistsIfNotCreate()
}

func (a *App) checkAppDirExistsIfNotCreate() {
	configDir := c.GetUserConfigDir()

	a.runikDir = filepath.Join(configDir, "runik")
	a.dictionaryDir = filepath.Join(a.runikDir, "dictionaries")

	// populate app config directories
	c.MkdirIfNotExists(a.runikDir)
	c.MkdirIfNotExists(a.dictionaryDir)
}

func (a *App) SelectDevice() string {
	a.devicePath = a.selectDirectory(runtime.OpenDialogOptions{})
	return a.devicePath
}

func (a *App) selectDirectory(options runtime.OpenDialogOptions) string {
	dirPath, err := runtime.OpenDirectoryDialog(a.ctx, options)
	if err != nil {
		log.Fatal(err)
	}

	return dirPath
}

// Builds runik dictionary
func (a *App) BuildDictionary(wikiUrl string, name string, output string, entryLimit int, depth int, format string) dict.Dict {
	return builder.BuildDictionary(wikiUrl, name, output, entryLimit, depth, format)
}
