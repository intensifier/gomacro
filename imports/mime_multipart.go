// this file was generated by gomacro command: import "mime/multipart"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "mime/multipart"
	. "reflect"
)

func Package_mime_multipart() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"NewReader": ValueOf(pkg.NewReader),
			"NewWriter": ValueOf(pkg.NewWriter),
		}, map[string]Type{
			"File":       TypeOf((*pkg.File)(nil)).Elem(),
			"FileHeader": TypeOf((*pkg.FileHeader)(nil)).Elem(),
			"Form":       TypeOf((*pkg.Form)(nil)).Elem(),
			"Part":       TypeOf((*pkg.Part)(nil)).Elem(),
			"Reader":     TypeOf((*pkg.Reader)(nil)).Elem(),
			"Writer":     TypeOf((*pkg.Writer)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_mime_multipart()
	Binds["mime/multipart"] = binds
	Types["mime/multipart"] = types
}
