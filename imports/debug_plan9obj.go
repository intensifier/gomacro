// this file was generated by gomacro command: import "debug/plan9obj"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "debug/plan9obj"
	. "reflect"
)

func Package_debug_plan9obj() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Magic386":   ValueOf(pkg.Magic386),
			"Magic64":    ValueOf(pkg.Magic64),
			"MagicAMD64": ValueOf(pkg.MagicAMD64),
			"MagicARM":   ValueOf(pkg.MagicARM),
			"NewFile":    ValueOf(pkg.NewFile),
			"Open":       ValueOf(pkg.Open),
		}, map[string]Type{
			"File":          TypeOf((*pkg.File)(nil)).Elem(),
			"FileHeader":    TypeOf((*pkg.FileHeader)(nil)).Elem(),
			"Section":       TypeOf((*pkg.Section)(nil)).Elem(),
			"SectionHeader": TypeOf((*pkg.SectionHeader)(nil)).Elem(),
			"Sym":           TypeOf((*pkg.Sym)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_debug_plan9obj()
	Binds["debug/plan9obj"] = binds
	Types["debug/plan9obj"] = types
}
