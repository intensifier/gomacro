// this file was generated by gomacro command: import "html/template"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "html/template"
	. "reflect"
)

func Package_html_template() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"ErrAmbigContext":     ValueOf(pkg.ErrAmbigContext),
			"ErrBadHTML":          ValueOf(pkg.ErrBadHTML),
			"ErrBranchEnd":        ValueOf(pkg.ErrBranchEnd),
			"ErrEndContext":       ValueOf(pkg.ErrEndContext),
			"ErrNoSuchTemplate":   ValueOf(pkg.ErrNoSuchTemplate),
			"ErrOutputContext":    ValueOf(pkg.ErrOutputContext),
			"ErrPartialCharset":   ValueOf(pkg.ErrPartialCharset),
			"ErrPartialEscape":    ValueOf(pkg.ErrPartialEscape),
			"ErrRangeLoopReentry": ValueOf(pkg.ErrRangeLoopReentry),
			"ErrSlashAmbig":       ValueOf(pkg.ErrSlashAmbig),
			"HTMLEscape":          ValueOf(pkg.HTMLEscape),
			"HTMLEscapeString":    ValueOf(pkg.HTMLEscapeString),
			"HTMLEscaper":         ValueOf(pkg.HTMLEscaper),
			"IsTrue":              ValueOf(pkg.IsTrue),
			"JSEscape":            ValueOf(pkg.JSEscape),
			"JSEscapeString":      ValueOf(pkg.JSEscapeString),
			"JSEscaper":           ValueOf(pkg.JSEscaper),
			"Must":                ValueOf(pkg.Must),
			"New":                 ValueOf(pkg.New),
			"OK":                  ValueOf(pkg.OK),
			"ParseFiles":          ValueOf(pkg.ParseFiles),
			"ParseGlob":           ValueOf(pkg.ParseGlob),
			"URLQueryEscaper":     ValueOf(pkg.URLQueryEscaper),
		}, map[string]Type{
			"CSS":       TypeOf((*pkg.CSS)(nil)).Elem(),
			"Error":     TypeOf((*pkg.Error)(nil)).Elem(),
			"ErrorCode": TypeOf((*pkg.ErrorCode)(nil)).Elem(),
			"FuncMap":   TypeOf((*pkg.FuncMap)(nil)).Elem(),
			"HTML":      TypeOf((*pkg.HTML)(nil)).Elem(),
			"HTMLAttr":  TypeOf((*pkg.HTMLAttr)(nil)).Elem(),
			"JS":        TypeOf((*pkg.JS)(nil)).Elem(),
			"JSStr":     TypeOf((*pkg.JSStr)(nil)).Elem(),
			"Template":  TypeOf((*pkg.Template)(nil)).Elem(),
			"URL":       TypeOf((*pkg.URL)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_html_template()
	Binds["html/template"] = binds
	Types["html/template"] = types
}
