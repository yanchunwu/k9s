package render

import (
	"fmt"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
)

// LogDump renders a logdumps to log.
type LogDump struct {
	Base
}

// ColorerFunc colors a resource row.
func (LogDump) ColorerFunc() ColorerFunc {
	return func(ns string, _ Header, re RowEvent) tcell.Color {
		return tcell.ColorNavajoWhite
	}
}

// Header returns a header row.
func (LogDump) Header(ns string) Header {
	return Header{
		HeaderColumn{Name: "NAME"},
		HeaderColumn{Name: "DIR"},
		HeaderColumn{Name: "VALID", Wide: true},
		HeaderColumn{Name: "AGE", Time: true, Decorator: AgeDecorator},
	}
}

// Render renders a K8s resource to log.
func (b LogDump) Render(o interface{}, ns string, r *Row) error {
	f, ok := o.(FileRes)
	if !ok {
		return fmt.Errorf("expecting logdumper, but got %T", o)
	}

	r.ID = filepath.Join(f.Dir, f.File.Name())
	r.Fields = Fields{
		f.File.Name(),
		f.Dir,
		"",
		timeToAge(f.File.ModTime()),
	}

	return nil
}

// ----------------------------------------------------------------------------
// Helpers...

// func timeToAge(timestamp time.Time) string {
// 	return time.Since(timestamp).String()
// }
//
// // FileRes represents a file resource.
// type FileRes struct {
// 	File os.FileInfo
// 	Dir  string
// }
//
// // GetObjectKind returns a schema object.
// func (c FileRes) GetObjectKind() schema.ObjectKind {
// 	return nil
// }
//
// // DeepCopyObject returns a container copy.
// func (c FileRes) DeepCopyObject() runtime.Object {
// 	return c
// }
