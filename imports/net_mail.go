// this file was generated by gomacro command: import "net/mail"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "net/mail"
	. "reflect"
)

func Package_net_mail() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"ErrHeaderNotPresent": ValueOf(&pkg.ErrHeaderNotPresent).Elem(),
			"ParseAddress":        ValueOf(pkg.ParseAddress),
			"ParseAddressList":    ValueOf(pkg.ParseAddressList),
			"ParseDate":           ValueOf(pkg.ParseDate),
			"ReadMessage":         ValueOf(pkg.ReadMessage),
		}, map[string]Type{
			"Address":       TypeOf((*pkg.Address)(nil)).Elem(),
			"AddressParser": TypeOf((*pkg.AddressParser)(nil)).Elem(),
			"Header":        TypeOf((*pkg.Header)(nil)).Elem(),
			"Message":       TypeOf((*pkg.Message)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_net_mail()
	Binds["net/mail"] = binds
	Types["net/mail"] = types
}
