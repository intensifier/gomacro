// this file was generated by gomacro command: import "net/http/httptrace"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "net/http/httptrace"
	. "reflect"
)

func Package_net_http_httptrace() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"ContextClientTrace": ValueOf(pkg.ContextClientTrace),
			"WithClientTrace":    ValueOf(pkg.WithClientTrace),
		}, map[string]Type{
			"ClientTrace":      TypeOf((*pkg.ClientTrace)(nil)).Elem(),
			"DNSDoneInfo":      TypeOf((*pkg.DNSDoneInfo)(nil)).Elem(),
			"DNSStartInfo":     TypeOf((*pkg.DNSStartInfo)(nil)).Elem(),
			"GotConnInfo":      TypeOf((*pkg.GotConnInfo)(nil)).Elem(),
			"WroteRequestInfo": TypeOf((*pkg.WroteRequestInfo)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_net_http_httptrace()
	Binds["net/http/httptrace"] = binds
	Types["net/http/httptrace"] = types
}
