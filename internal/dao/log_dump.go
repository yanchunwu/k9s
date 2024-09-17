package dao

import (
	"context"
	"errors"
	"os"

	"github.com/derailed/k9s/internal"
	"github.com/derailed/k9s/internal/render"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	_ Accessor = (*LogDump)(nil)
	_ Nuker    = (*LogDump)(nil)

	// InvalidCharsRX contains invalid filename characters.
)

// LogDump represents a scraped resources.
type LogDump struct {
	NonResource
}

// Delete a LogDump.
func (d *LogDump) Delete(path string, cascade, force bool) error {
	return os.Remove(path)
}

// List returns a collection of log dumps.
func (d *LogDump) List(ctx context.Context, _ string) ([]runtime.Object, error) {
	dir, ok := ctx.Value(internal.KeyDir).(string)
	if !ok {
		return nil, errors.New("no logdump dir found in context")
	}

	ff, err := os.ReadDir(SanitizeFilename(dir))
	if err != nil {
		return nil, err
	}

	oo := make([]runtime.Object, len(ff))
	for i, f := range ff {
		if fi, err := f.Info(); err == nil {
			oo[i] = render.FileRes{File: fi, Dir: dir}
		}
	}

	return oo, nil
}

// Helpers...

// // SanitizeFilename sanitizes the dump filename.
// func SanitizeFilename(name string) string {
// 	return invalidPathCharsRX.ReplaceAllString(name, "-")
// }
