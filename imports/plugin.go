// this file was generated by gomacro command: import "plugin"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "plugin"
	. "reflect"
)

func Package_plugin() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Open": ValueOf(pkg.Open),
		}, map[string]Type{
			"Plugin": TypeOf((*pkg.Plugin)(nil)).Elem(),
			"Symbol": TypeOf((*pkg.Symbol)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_plugin()
	Binds["plugin"] = binds
	Types["plugin"] = types
}
