// this file was generated by gomacro command: import "unsafe"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	pkg "unsafe"
)

func Package_unsafe() (map[string]Value, map[string]Type) {
	return map[string]Value{}, map[string]Type{
			"Pointer": TypeOf((*pkg.Pointer)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_unsafe()
	Binds["unsafe"] = binds
	Types["unsafe"] = types
}
