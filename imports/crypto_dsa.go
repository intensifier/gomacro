// this file was generated by gomacro command: import "crypto/dsa"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "crypto/dsa"
	. "reflect"
)

func Package_crypto_dsa() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"ErrInvalidPublicKey": ValueOf(&pkg.ErrInvalidPublicKey).Elem(),
			"GenerateKey":         ValueOf(pkg.GenerateKey),
			"GenerateParameters":  ValueOf(pkg.GenerateParameters),
			"L1024N160":           ValueOf(pkg.L1024N160),
			"L2048N224":           ValueOf(pkg.L2048N224),
			"L2048N256":           ValueOf(pkg.L2048N256),
			"L3072N256":           ValueOf(pkg.L3072N256),
			"Sign":                ValueOf(pkg.Sign),
			"Verify":              ValueOf(pkg.Verify),
		}, map[string]Type{
			"ParameterSizes": TypeOf((*pkg.ParameterSizes)(nil)).Elem(),
			"Parameters":     TypeOf((*pkg.Parameters)(nil)).Elem(),
			"PrivateKey":     TypeOf((*pkg.PrivateKey)(nil)).Elem(),
			"PublicKey":      TypeOf((*pkg.PublicKey)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_crypto_dsa()
	Binds["crypto/dsa"] = binds
	Types["crypto/dsa"] = types
}
