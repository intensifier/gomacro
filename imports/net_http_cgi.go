// this file was generated by gomacro command: import "net/http/cgi"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "net/http/cgi"
	. "reflect"
)

func Package_net_http_cgi() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Request":        ValueOf(pkg.Request),
			"RequestFromMap": ValueOf(pkg.RequestFromMap),
			"Serve":          ValueOf(pkg.Serve),
		}, map[string]Type{
			"Handler": TypeOf((*pkg.Handler)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_net_http_cgi()
	Binds["net/http/cgi"] = binds
	Types["net/http/cgi"] = types
}
