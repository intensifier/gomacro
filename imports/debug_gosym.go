// this file was generated by gomacro command: import "debug/gosym"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "debug/gosym"
	. "reflect"
)

func Package_debug_gosym() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"NewLineTable": ValueOf(pkg.NewLineTable),
			"NewTable":     ValueOf(pkg.NewTable),
		}, map[string]Type{
			"DecodingError":    TypeOf((*pkg.DecodingError)(nil)).Elem(),
			"Func":             TypeOf((*pkg.Func)(nil)).Elem(),
			"LineTable":        TypeOf((*pkg.LineTable)(nil)).Elem(),
			"Obj":              TypeOf((*pkg.Obj)(nil)).Elem(),
			"Sym":              TypeOf((*pkg.Sym)(nil)).Elem(),
			"Table":            TypeOf((*pkg.Table)(nil)).Elem(),
			"UnknownFileError": TypeOf((*pkg.UnknownFileError)(nil)).Elem(),
			"UnknownLineError": TypeOf((*pkg.UnknownLineError)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_debug_gosym()
	Binds["debug/gosym"] = binds
	Types["debug/gosym"] = types
}
