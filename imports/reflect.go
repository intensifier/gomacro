// this file was generated by gomacro command: import "reflect"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	pkg "reflect"
)

func Package_reflect() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Append":        ValueOf(pkg.Append),
			"AppendSlice":   ValueOf(pkg.AppendSlice),
			"Array":         ValueOf(pkg.Array),
			"ArrayOf":       ValueOf(pkg.ArrayOf),
			"Bool":          ValueOf(pkg.Bool),
			"BothDir":       ValueOf(pkg.BothDir),
			"Chan":          ValueOf(pkg.Chan),
			"ChanOf":        ValueOf(pkg.ChanOf),
			"Complex128":    ValueOf(pkg.Complex128),
			"Complex64":     ValueOf(pkg.Complex64),
			"Copy":          ValueOf(pkg.Copy),
			"DeepEqual":     ValueOf(pkg.DeepEqual),
			"Float32":       ValueOf(pkg.Float32),
			"Float64":       ValueOf(pkg.Float64),
			"Func":          ValueOf(pkg.Func),
			"FuncOf":        ValueOf(pkg.FuncOf),
			"Indirect":      ValueOf(pkg.Indirect),
			"Int":           ValueOf(pkg.Int),
			"Int16":         ValueOf(pkg.Int16),
			"Int32":         ValueOf(pkg.Int32),
			"Int64":         ValueOf(pkg.Int64),
			"Int8":          ValueOf(pkg.Int8),
			"Interface":     ValueOf(pkg.Interface),
			"Invalid":       ValueOf(pkg.Invalid),
			"MakeChan":      ValueOf(pkg.MakeChan),
			"MakeFunc":      ValueOf(pkg.MakeFunc),
			"MakeMap":       ValueOf(pkg.MakeMap),
			"MakeSlice":     ValueOf(pkg.MakeSlice),
			"Map":           ValueOf(pkg.Map),
			"MapOf":         ValueOf(pkg.MapOf),
			"New":           ValueOf(pkg.New),
			"NewAt":         ValueOf(pkg.NewAt),
			"Ptr":           ValueOf(pkg.Ptr),
			"PtrTo":         ValueOf(pkg.PtrTo),
			"RecvDir":       ValueOf(pkg.RecvDir),
			"Select":        ValueOf(pkg.Select),
			"SelectDefault": ValueOf(pkg.SelectDefault),
			"SelectRecv":    ValueOf(pkg.SelectRecv),
			"SelectSend":    ValueOf(pkg.SelectSend),
			"SendDir":       ValueOf(pkg.SendDir),
			"Slice":         ValueOf(pkg.Slice),
			"SliceOf":       ValueOf(pkg.SliceOf),
			"String":        ValueOf(pkg.String),
			"Struct":        ValueOf(pkg.Struct),
			"StructOf":      ValueOf(pkg.StructOf),
			"Swapper":       ValueOf(pkg.Swapper),
			"TypeOf":        ValueOf(pkg.TypeOf),
			"Uint":          ValueOf(pkg.Uint),
			"Uint16":        ValueOf(pkg.Uint16),
			"Uint32":        ValueOf(pkg.Uint32),
			"Uint64":        ValueOf(pkg.Uint64),
			"Uint8":         ValueOf(pkg.Uint8),
			"Uintptr":       ValueOf(pkg.Uintptr),
			"UnsafePointer": ValueOf(pkg.UnsafePointer),
			"ValueOf":       ValueOf(pkg.ValueOf),
			"Zero":          ValueOf(pkg.Zero),
		}, map[string]Type{
			"ChanDir":      TypeOf((*pkg.ChanDir)(nil)).Elem(),
			"Kind":         TypeOf((*pkg.Kind)(nil)).Elem(),
			"Method":       TypeOf((*pkg.Method)(nil)).Elem(),
			"SelectCase":   TypeOf((*pkg.SelectCase)(nil)).Elem(),
			"SelectDir":    TypeOf((*pkg.SelectDir)(nil)).Elem(),
			"SliceHeader":  TypeOf((*pkg.SliceHeader)(nil)).Elem(),
			"StringHeader": TypeOf((*pkg.StringHeader)(nil)).Elem(),
			"StructField":  TypeOf((*pkg.StructField)(nil)).Elem(),
			"StructTag":    TypeOf((*pkg.StructTag)(nil)).Elem(),
			"Type":         TypeOf((*pkg.Type)(nil)).Elem(),
			"Value":        TypeOf((*pkg.Value)(nil)).Elem(),
			"ValueError":   TypeOf((*pkg.ValueError)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_reflect()
	Binds["reflect"] = binds
	Types["reflect"] = types
}
