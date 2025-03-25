package main

import (
	"context"
	"log"
	"path/filepath"

	_ "embed"

	wikibot "github.com/runik-3/builder/wikiBot"
	c "github.com/runik-3/core/core"
	dev "github.com/runik-3/core/device"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed version
var version string

// App struct
type App struct {
	ctx context.Context
	// store app config files
	runikDir string
	// store raw dictionary files
	dictionaryDir string
	// the selected reader device
	device dev.Device
	config c.Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Based on GNOME design standard https://developer.gnome.org/hig/guidelines/adaptive.html
const MIN_WIDTH = 1024
const MIN_HEIGHT = 600

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowSetMinSize(ctx, MIN_WIDTH, MIN_HEIGHT)

	a.checkAppConfigDirExistsIfNotCreate()
}

// pull out into config.go
func (a *App) checkAppConfigDirExistsIfNotCreate() {
	configDir := c.GetUserConfigDir()

	a.runikDir = filepath.Join(configDir, "runik")
	a.dictionaryDir = filepath.Join(a.runikDir, "dictionaries")

	c.MkdirIfNotExists(a.runikDir)
	c.MkdirIfNotExists(a.dictionaryDir)

	config, err := c.GetOrCreateConfig(a.runikDir)
	if err != nil {
		log.Fatal(err)
	}
	a.config = config
}

func (a *App) GetConfig() c.Config {
	return a.config
}

func (a *App) CheckForUpdate() bool {
	return c.UpdateAvailable(version)
}

func (a *App) SelectDevice() c.Response[dev.Device] {
	deviceDir := a.selectDirectory(runtime.OpenDialogOptions{})

	device, err := dev.NewDevice(deviceDir, a.runikDir)
	if err != nil {
		return c.Response[dev.Device]{Data: dev.Kobo{}, Error: err.Error()}
	}
	a.device = device

	return c.Response[dev.Device]{Data: a.device, Error: ""}
}

func (a *App) selectDirectory(options runtime.OpenDialogOptions) string {
	dirPath, err := runtime.OpenDirectoryDialog(a.ctx, options)
	if err != nil {
		log.Fatal(err)
	}

	return dirPath
}

type GeneratorProgress struct {
	Processed int
	Total     int
}

// Sends generator progress to frontend
func (a *App) emitProgress(processed int, total int) {
	runtime.EventsEmit(a.ctx, "progressUpdate", GeneratorProgress{Processed: processed, Total: total})
}

func (a *App) GetWikiDetails(wikiUrl string) c.Response[wikibot.WikiDetails] {
	details, err := wikibot.GetWikiDetails(wikiUrl)
	if err != nil {
		return c.Response[wikibot.WikiDetails]{Data: wikibot.WikiDetails{}, Error: err.Error()}
	}
	return c.Response[wikibot.WikiDetails]{Data: details, Error: ""}
}
