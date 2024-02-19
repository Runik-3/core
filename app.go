package main

import (
	"context"
	"log"
	"path/filepath"

	d "github.com/runik-3/builder/dict"
	wikibot "github.com/runik-3/builder/wikiBot"
	c "github.com/runik-3/core/core"

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
	a.checkAppConfigDirExistsIfNotCreate()
}

func (a *App) checkAppConfigDirExistsIfNotCreate() {
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
func (a *App) BuildDictionary(wikiUrl string, name string, depth int, format string) c.Response[d.Dict] {
	dict, err := d.BuildDictionary(wikiUrl, name, a.dictionaryDir, 10000, depth, "json") // at least for now raw dicts should be json
	if err != nil {
		return c.Response[d.Dict]{Data: d.Dict{}, Error: err.Error()}
	}
	return c.Response[d.Dict]{Data: dict, Error: ""}
}

func (a *App) GetWikiDetails(wikiUrl string) c.Response[wikibot.WikiDetails] {
	details, err := wikibot.GetWikiDetails(wikiUrl)
	if err != nil {
		return c.Response[wikibot.WikiDetails]{Data: wikibot.WikiDetails{}, Error: err.Error()}
	}
	return c.Response[wikibot.WikiDetails]{Data: details, Error: ""}
}

func (a *App) ConvertKoboDictionary(fileName string) c.Response[string] {
	if a.devicePath == "" {
		return c.Response[string]{Data: "", Error: "No device connected"}
	}

	rawDictPath := filepath.Join(a.dictionaryDir, fileName)
	koboDictDir, err := c.FindKoboDictDir(a.devicePath)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}

	dictPath, err := c.ConvertForReader(rawDictPath, koboDictDir)
	if err != nil {
		return c.Response[string]{Data: "", Error: err.Error()}
	}
	return c.Response[string]{Data: dictPath, Error: ""}
}
