// this file was generated by gomacro command: import "compress/flate"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "compress/flate"
	. "reflect"
)

func Package_compress_flate() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"BestCompression":    ValueOf(pkg.BestCompression),
			"BestSpeed":          ValueOf(pkg.BestSpeed),
			"DefaultCompression": ValueOf(pkg.DefaultCompression),
			"HuffmanOnly":        ValueOf(pkg.HuffmanOnly),
			"NewReader":          ValueOf(pkg.NewReader),
			"NewReaderDict":      ValueOf(pkg.NewReaderDict),
			"NewWriter":          ValueOf(pkg.NewWriter),
			"NewWriterDict":      ValueOf(pkg.NewWriterDict),
			"NoCompression":      ValueOf(pkg.NoCompression),
		}, map[string]Type{
			"CorruptInputError": TypeOf((*pkg.CorruptInputError)(nil)).Elem(),
			"InternalError":     TypeOf((*pkg.InternalError)(nil)).Elem(),
			"ReadError":         TypeOf((*pkg.ReadError)(nil)).Elem(),
			"Reader":            TypeOf((*pkg.Reader)(nil)).Elem(),
			"Resetter":          TypeOf((*pkg.Resetter)(nil)).Elem(),
			"WriteError":        TypeOf((*pkg.WriteError)(nil)).Elem(),
			"Writer":            TypeOf((*pkg.Writer)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_compress_flate()
	Binds["compress/flate"] = binds
	Types["compress/flate"] = types
}
