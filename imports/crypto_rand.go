// this file was generated by gomacro command: import "crypto/rand"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "crypto/rand"
	. "reflect"
)

func Package_crypto_rand() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Int":    ValueOf(pkg.Int),
			"Prime":  ValueOf(pkg.Prime),
			"Read":   ValueOf(pkg.Read),
			"Reader": ValueOf(&pkg.Reader).Elem(),
		}, map[string]Type{}
}

func init() {
	binds, types := Package_crypto_rand()
	Binds["crypto/rand"] = binds
	Types["crypto/rand"] = types
}
