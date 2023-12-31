package main

import (
	"context"

	"github.com/runik-3/builder/pkg/builder"
	"github.com/runik-3/builder/pkg/dict"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Builds runik dictionary
func (a *App) BuildDictionary(wikiUrl string, name string, output string, entryLimit int, depth int, format string) dict.Dict {
	return builder.BuildDictionary(wikiUrl, name, output, entryLimit, depth, format)
}
