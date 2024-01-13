package main

import (
	"context"
	"log"
	"path/filepath"

	c "github.com/runik-3/core/core"

	"github.com/runik-3/builder/dict"
	wikibot "github.com/runik-3/builder/wikiBot"
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
func (a *App) BuildDictionary(wikiUrl string, name string, depth int, format string) dict.Dict {
	// TODO: size of 5 for testing
	return dict.BuildDictionary(wikiUrl, name, a.dictionaryDir, 5, depth, "json") // at least for now raw dicts should be json
}

func (a *App) GetWikiDetails(wikiUrl string) wikibot.WikiDetails {
	return wikibot.GetWikiDetails(wikiUrl)
}
