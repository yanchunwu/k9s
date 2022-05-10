package view

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/derailed/k9s/internal"
	"github.com/derailed/k9s/internal/client"
	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/k9s/internal/render"
	"github.com/derailed/k9s/internal/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rs/zerolog/log"
)

// LogDump presents a directory listing viewer.
type LogDump struct {
	ResourceViewer
}

// NewLogDump returns a new viewer.
func NewLogDump(gvr client.GVR) ResourceViewer {
	s := LogDump{
		ResourceViewer: NewBrowser(gvr),
	}
	s.GetTable().SetBorderFocusColor(tcell.ColorSteelBlue)
	s.GetTable().SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRoyalBlue).Attributes(tcell.AttrNone))
	s.GetTable().SetColorerFn(render.LogDump{}.ColorerFunc())
	s.GetTable().SetSortCol(ageCol, true)
	s.GetTable().SelectRow(1, true)
	s.GetTable().SetEnterFn(s.edit)
	s.SetContextFn(s.dirContext)

	return &s
}

func (s *LogDump) dirContext(ctx context.Context) context.Context {
	dir := filepath.Join(s.App().Config.K9s.GetLogDumpDir(), s.App().Config.K9s.CurrentContext)
	log.Debug().Msgf("SD-DIR %q", dir)
	config.EnsureFullPath(dir, config.DefaultDirMod)

	return context.WithValue(ctx, internal.KeyDir, dir)
}

func (s *LogDump) edit(app *App, model ui.Tabular, gvr, path string) {
	log.Debug().Msgf("LogDump selection is %q", path)

	s.Stop()
	defer s.Start()
	if !edit(app, shellOpts{clear: true, args: []string{path}}) {
		app.Flash().Err(errors.New("Failed to launch editor"))
	}
}
