// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * call_ret1.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	. "github.com/cosmos72/gomacro/base"
)

func call0ret1(c *Call) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	funvar := c.Funvar
	var funupn, funindex int
	if funvar != nil {
		funupn = funvar.Upn
		funindex = funvar.Desc.Index()
	}
	kret := expr.Type.Out(0).Kind()
	var cachedfunv r.Value
	var call I
	switch kret {
	case r.Bool:

		if funvar == nil {
			call = func(env *Env) bool {
				fun := exprfun(env).Interface().(func() bool)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) bool {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() bool)
				return fun()
			}
		} else {

			var cachedfun func() bool

			call = func(env *Env) bool {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() bool)
				}
				return cachedfun()
			}
		}
	case r.Int:
		if funvar == nil {
			call = func(env *Env) int {
				fun := exprfun(env).Interface().(func() int)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) int {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() int)
				return fun()
			}
		} else {

			var cachedfun func() int

			call = func(env *Env) int {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() int)
				}
				return cachedfun()
			}
		}
	case r.Int8:
		if funvar == nil {
			call = func(env *Env) int8 {
				fun := exprfun(env).Interface().(func() int8)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) int8 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() int8)
				return fun()
			}
		} else {

			var cachedfun func() int8

			call = func(env *Env) int8 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() int8)
				}
				return cachedfun()
			}
		}
	case r.Int16:
		if funvar == nil {
			call = func(env *Env) int16 {
				fun := exprfun(env).Interface().(func() int16)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) int16 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() int16)
				return fun()
			}
		} else {

			var cachedfun func() int16

			call = func(env *Env) int16 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() int16)
				}
				return cachedfun()
			}
		}
	case r.Int32:
		if funvar == nil {
			call = func(env *Env) int32 {
				fun := exprfun(env).Interface().(func() int32)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) int32 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() int32)
				return fun()
			}
		} else {

			var cachedfun func() int32

			call = func(env *Env) int32 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() int32)
				}
				return cachedfun()
			}
		}
	case r.Int64:
		if funvar == nil {
			call = func(env *Env) int64 {
				fun := exprfun(env).Interface().(func() int64)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) int64 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() int64)
				return fun()
			}
		} else {

			var cachedfun func() int64

			call = func(env *Env) int64 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() int64)
				}
				return cachedfun()
			}
		}
	case r.Uint:
		if funvar == nil {
			call = func(env *Env) uint {
				fun := exprfun(env).Interface().(func() uint)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) uint {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() uint)
				return fun()
			}
		} else {

			var cachedfun func() uint

			call = func(env *Env) uint {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() uint)
				}
				return cachedfun()
			}
		}
	case r.Uint8:
		if funvar == nil {
			call = func(env *Env) uint8 {
				fun := exprfun(env).Interface().(func() uint8)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) uint8 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() uint8)
				return fun()
			}
		} else {

			var cachedfun func() uint8

			call = func(env *Env) uint8 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() uint8)
				}
				return cachedfun()
			}
		}
	case r.Uint16:
		if funvar == nil {
			call = func(env *Env) uint16 {
				fun := exprfun(env).Interface().(func() uint16)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) uint16 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() uint16)
				return fun()
			}
		} else {

			var cachedfun func() uint16

			call = func(env *Env) uint16 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() uint16)
				}
				return cachedfun()
			}
		}
	case r.Uint32:
		if funvar == nil {
			call = func(env *Env) uint32 {
				fun := exprfun(env).Interface().(func() uint32)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) uint32 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() uint32)
				return fun()
			}
		} else {

			var cachedfun func() uint32

			call = func(env *Env) uint32 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() uint32)
				}
				return cachedfun()
			}
		}
	case r.Uint64:
		if funvar == nil {
			call = func(env *Env) uint64 {
				fun := exprfun(env).Interface().(func() uint64)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) uint64 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() uint64)
				return fun()
			}
		} else {

			var cachedfun func() uint64

			call = func(env *Env) uint64 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() uint64)
				}
				return cachedfun()
			}
		}
	case r.Uintptr:
		if funvar == nil {
			call = func(env *Env) uintptr {
				fun := exprfun(env).Interface().(func() uintptr)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) uintptr {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() uintptr)
				return fun()
			}
		} else {

			var cachedfun func() uintptr

			call = func(env *Env) uintptr {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() uintptr)
				}
				return cachedfun()
			}
		}
	case r.Float32:
		if funvar == nil {
			call = func(env *Env) float32 {
				fun := exprfun(env).Interface().(func() float32)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) float32 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() float32)
				return fun()
			}
		} else {

			var cachedfun func() float32

			call = func(env *Env) float32 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() float32)
				}
				return cachedfun()
			}
		}
	case r.Float64:
		if funvar == nil {
			call = func(env *Env) float64 {
				fun := exprfun(env).Interface().(func() float64)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) float64 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() float64)
				return fun()
			}
		} else {

			var cachedfun func() float64

			call = func(env *Env) float64 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() float64)
				}
				return cachedfun()
			}
		}
	case r.Complex64:
		if funvar == nil {
			call = func(env *Env) complex64 {
				fun := exprfun(env).Interface().(func() complex64)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) complex64 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() complex64)
				return fun()
			}
		} else {

			var cachedfun func() complex64

			call = func(env *Env) complex64 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() complex64)
				}
				return cachedfun()
			}
		}
	case r.Complex128:
		if funvar == nil {
			call = func(env *Env) complex128 {
				fun := exprfun(env).Interface().(func() complex128)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) complex128 {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() complex128)
				return fun()
			}
		} else {

			var cachedfun func() complex128

			call = func(env *Env) complex128 {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() complex128)
				}
				return cachedfun()
			}
		}
	case r.String:
		if funvar == nil {
			call = func(env *Env) string {
				fun := exprfun(env).Interface().(func() string)
				return fun()
			}
		} else if funupn != 1 {
			call = func(env *Env) string {
				o := env
				for i := 0; i < funupn; i++ {
					o = o.Outer
				}

				fun := o.Binds[funindex].Interface().(func() string)
				return fun()
			}
		} else {

			var cachedfun func() string

			call = func(env *Env) string {
				funv := env.Outer.Binds[funindex]
				if cachedfunv != funv {
					cachedfunv = funv
					cachedfun = funv.Interface().(func() string)
				}
				return cachedfun()
			}
		}
	default:
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			return funv.Call(ZeroValues)[0]
		}

	}
	return call
}
func call1ret1(c *Call) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	funvar := c.Funvar
	var cachedfunv r.Value
	var funupn, funindex int
	if funvar != nil {
		funupn = funvar.Upn
		funindex = funvar.Desc.Index()
	}
	karg0 := expr.Type.In(0).Kind()
	kret := expr.Type.Out(0).Kind()
	arg0 := c.Args[0]
	argfun := c.Argfuns[0]
	var call I
	switch kret {
	case r.Bool:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) bool

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(bool) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(bool) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) bool

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) bool

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int8) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int8) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) bool

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int16) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int16) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) bool

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int32) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int32) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) bool

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int64) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(int64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) bool

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) bool

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint8) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint8) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) bool

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint16) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint16) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) bool

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint32) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint32) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) bool

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint64) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uint64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) bool

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uintptr) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(uintptr) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) bool

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(float32) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(float32) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) bool

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(float64) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(float64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) bool

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(complex64) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(complex64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) bool

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(complex128) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(complex128) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) bool

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(string) bool)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) bool)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) bool)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) bool {
							fun := exprfun(env).Interface().(func(string) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) bool {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) bool)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) bool {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) bool)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) bool {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Bool()
			}
		}
	case r.Int:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) int

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(bool) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(bool) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) int

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) int

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int8) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int8) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) int

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int16) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int16) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) int

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int32) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int32) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) int

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int64) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(int64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) int

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) int

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint8) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint8) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) int

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint16) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint16) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) int

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint32) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint32) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) int

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint64) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uint64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) int

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uintptr) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(uintptr) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) int

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(float32) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(float32) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) int

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(float64) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(float64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) int

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(complex64) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(complex64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) int

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(complex128) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(complex128) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) int

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(string) int)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) int {
							fun := exprfun(env).Interface().(func(string) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) int {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return int(ret0.Int())
			}
		}
	case r.Int8:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) int8

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(bool) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(bool) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) int8

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) int8

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int8) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int8) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) int8

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int16) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int16) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) int8

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int32) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int32) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) int8

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int64) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) int8

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) int8

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint8) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint8) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) int8

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint16) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint16) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) int8

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint32) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint32) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) int8

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint64) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uint64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) int8

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uintptr) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(uintptr) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) int8

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(float32) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(float32) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) int8

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(float64) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(float64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) int8

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(complex64) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(complex64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) int8

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(complex128) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(complex128) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) int8

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(string) int8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(string) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) int8 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return int8(ret0.Int())
			}
		}
	case r.Int16:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) int16

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(bool) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(bool) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) int16

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) int16

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int8) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int8) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) int16

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int16) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int16) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) int16

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int32) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int32) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) int16

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int64) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) int16

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) int16

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint8) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint8) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) int16

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint16) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint16) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) int16

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint32) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint32) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) int16

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint64) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uint64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) int16

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uintptr) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(uintptr) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) int16

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(float32) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(float32) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) int16

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(float64) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(float64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) int16

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(complex64) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(complex64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) int16

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(complex128) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(complex128) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) int16

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(string) int16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(string) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) int16 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return int16(ret0.Int())
			}
		}
	case r.Int32:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) int32

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(bool) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(bool) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) int32

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) int32

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int8) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int8) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) int32

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int16) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int16) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) int32

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int32) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int32) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) int32

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int64) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(int64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) int32

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) int32

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint8) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint8) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) int32

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint16) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint16) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) int32

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint32) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint32) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) int32

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint64) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uint64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) int32

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uintptr) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(uintptr) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) int32

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(float32) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(float32) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) int32

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(float64) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(float64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) int32

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(complex64) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(complex64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) int32

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(complex128) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(complex128) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) int32

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(string) int32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) int32 {
							fun := exprfun(env).Interface().(func(string) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) int32 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return int32(ret0.Int())
			}
		}
	case r.Int64:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) int64

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(bool) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(bool) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) int64

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) int64

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int8) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int8) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) int64

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int16) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int16) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) int64

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int32) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int32) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) int64

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int64) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(int64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) int64

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) int64

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint8) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint8) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) int64

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint16) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint16) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) int64

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint32) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint32) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) int64

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint64) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uint64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) int64

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uintptr) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(uintptr) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) int64

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(float32) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(float32) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) int64

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(float64) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(float64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) int64

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(complex64) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(complex64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) int64

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(complex128) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(complex128) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) int64

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(string) int64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) int64 {
							fun := exprfun(env).Interface().(func(string) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) int64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) int64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) int64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Int()
			}
		}
	case r.Uint:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) uint

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(bool) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(bool) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) uint

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) uint

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int8) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int8) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) uint

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int16) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int16) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) uint

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int32) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int32) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) uint

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int64) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(int64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) uint

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) uint

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint8) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint8) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) uint

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint16) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint16) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) uint

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint32) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint32) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) uint

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint64) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uint64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) uint

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uintptr) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(uintptr) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) uint

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(float32) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(float32) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) uint

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(float64) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(float64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) uint

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(complex64) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(complex64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) uint

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(complex128) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(complex128) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) uint

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(string) uint)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) uint {
							fun := exprfun(env).Interface().(func(string) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) uint {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return uint(ret0.Uint())
			}
		}

	case r.Uint8:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) uint8

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(bool) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(bool) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) uint8

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) uint8

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int8) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int8) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) uint8

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int16) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int16) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) uint8

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int32) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int32) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) uint8

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int64) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(int64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) uint8

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) uint8

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint8) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint8) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) uint8

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint16) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint16) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) uint8

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint32) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint32) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) uint8

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint64) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) uint8

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uintptr) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uintptr) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) uint8

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(float32) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(float32) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) uint8

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(float64) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(float64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) uint8

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(complex64) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(complex64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) uint8

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(complex128) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(complex128) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) uint8

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(string) uint8)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint8)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint8)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(string) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint8 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint8)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) uint8 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return uint8(ret0.Uint())
			}
		}

	case r.Uint16:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) uint16

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(bool) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(bool) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) uint16

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) uint16

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int8) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int8) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) uint16

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int16) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int16) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) uint16

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int32) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int32) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) uint16

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int64) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(int64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) uint16

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) uint16

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint8) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint8) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) uint16

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint16) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint16) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) uint16

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint32) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint32) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) uint16

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint64) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) uint16

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uintptr) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uintptr) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) uint16

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(float32) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(float32) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) uint16

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(float64) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(float64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) uint16

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(complex64) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(complex64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) uint16

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(complex128) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(complex128) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) uint16

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(string) uint16)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint16)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint16)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(string) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint16 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint16)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) uint16 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return uint16(ret0.Uint())
			}
		}

	case r.Uint32:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) uint32

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(bool) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(bool) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) uint32

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) uint32

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int8) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int8) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) uint32

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int16) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int16) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) uint32

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int32) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) uint32

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int64) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(int64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) uint32

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) uint32

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint8) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint8) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) uint32

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint16) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint16) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) uint32

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint32) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) uint32

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint64) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) uint32

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uintptr) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uintptr) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) uint32

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(float32) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(float32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) uint32

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(float64) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(float64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) uint32

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(complex64) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(complex64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) uint32

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(complex128) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(complex128) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) uint32

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(string) uint32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(string) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) uint32 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return uint32(ret0.Uint())
			}
		}

	case r.Uint64:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) uint64

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(bool) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(bool) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) uint64

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) uint64

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int8) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int8) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) uint64

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int16) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int16) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) uint64

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int32) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int32) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) uint64

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int64) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(int64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) uint64

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) uint64

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint8) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint8) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) uint64

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint16) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint16) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) uint64

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint32) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint32) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) uint64

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint64) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uint64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) uint64

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uintptr) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(uintptr) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) uint64

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(float32) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(float32) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) uint64

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(float64) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(float64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) uint64

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(complex64) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(complex64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) uint64

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(complex128) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(complex128) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) uint64

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(string) uint64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) uint64 {
							fun := exprfun(env).Interface().(func(string) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uint64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uint64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) uint64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Uint()
			}
		}

	case r.Uintptr:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(bool) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(bool) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int8) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int8) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int16) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int16) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int32) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int32) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int64) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(int64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint8) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint8) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint16) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint16) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint32) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint32) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint64) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uint64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uintptr) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uintptr) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(float32) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(float32) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(float64) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(float64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(complex64) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(complex64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(complex128) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(complex128) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) uintptr

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(string) uintptr)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uintptr)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uintptr)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(string) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) uintptr {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uintptr)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) uintptr {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return uintptr(ret0.Uint())
			}
		}

	case r.Float32:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) float32

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(bool) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(bool) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) float32

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) float32

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int8) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int8) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) float32

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int16) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int16) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) float32

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int32) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) float32

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int64) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(int64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) float32

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) float32

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint8) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint8) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) float32

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint16) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint16) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) float32

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint32) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) float32

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint64) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uint64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) float32

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uintptr) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(uintptr) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) float32

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(float32) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(float32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) float32

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(float64) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(float64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) float32

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(complex64) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(complex64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) float32

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(complex128) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(complex128) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) float32

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(string) float32)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) float32)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) float32)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(string) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float32 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) float32)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) float32 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return float32(ret0.Float())
			}
		}

	case r.Float64:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) float64

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(bool) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(bool) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) float64

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) float64

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int8) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int8) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) float64

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int16) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int16) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) float64

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int32) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int32) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) float64

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int64) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(int64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) float64

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) float64

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint8) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint8) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) float64

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint16) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint16) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) float64

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint32) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint32) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) float64

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint64) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uint64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) float64

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uintptr) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(uintptr) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) float64

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(float32) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(float32) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) float64

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(float64) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(float64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) float64

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(complex64) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(complex64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) float64

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(complex128) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(complex128) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) float64

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(string) float64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) float64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) float64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(string) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) float64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) float64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) float64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Float()
			}
		}

	case r.Complex64:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) complex64

				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(bool) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(bool) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) complex64

				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) complex64

				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int8) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int8) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) complex64

				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int16) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int16) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) complex64

				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int32) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int32) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) complex64

				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int64) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(int64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) complex64

				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) complex64

				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint8) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint8) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) complex64

				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint16) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint16) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) complex64

				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint32) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint32) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) complex64

				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint64) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uint64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) complex64

				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uintptr) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(uintptr) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) complex64

				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(float32) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(float32) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) complex64

				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(float64) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(float64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) complex64

				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(complex64) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(complex64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) complex64

				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(complex128) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(complex128) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) complex64

				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(string) complex64)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) complex64)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) complex64)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(string) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex64 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) complex64)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) complex64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return complex64(ret0.Complex())
			}
		}

	case r.Complex128:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) complex128
				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(bool) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(bool) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) complex128
				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) complex128
				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int8) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int8) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) complex128
				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int16) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int16) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) complex128
				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int32) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int32) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) complex128
				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int64) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(int64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) complex128
				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) complex128
				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint8) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint8) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) complex128
				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint16) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint16) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) complex128
				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint32) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint32) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) complex128
				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint64) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uint64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) complex128
				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uintptr) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(uintptr) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) complex128
				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(float32) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(float32) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) complex128
				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(float64) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(float64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) complex128
				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(complex64) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(complex64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) complex128
				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(complex128) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(complex128) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) complex128
				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(string) complex128)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) complex128)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) complex128)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(string) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) complex128 {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) complex128)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) complex128 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Complex()
			}
		}

	case r.String:
		switch karg0 {
		case r.Bool:

			{
				var cachedfun func(bool) string
				if arg0.Const() {
					argconst := arg0.Value.(bool)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(bool) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) bool)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(bool) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(bool) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int:

			{
				var cachedfun func(int) string
				if arg0.Const() {
					argconst := arg0.Value.(int)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int8:

			{
				var cachedfun func(int8) string
				if arg0.Const() {
					argconst := arg0.Value.(int8)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int8) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int8)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int8) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int8) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int16:

			{
				var cachedfun func(int16) string
				if arg0.Const() {
					argconst := arg0.Value.(int16)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int16) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int16)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int16) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int16) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int32:

			{
				var cachedfun func(int32) string
				if arg0.Const() {
					argconst := arg0.Value.(int32)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int32) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int32)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int32) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int32) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Int64:

			{
				var cachedfun func(int64) string
				if arg0.Const() {
					argconst := arg0.Value.(int64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int64) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) int64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(int64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(int64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint:

			{
				var cachedfun func(uint) string
				if arg0.Const() {
					argconst := arg0.Value.(uint)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint8:

			{
				var cachedfun func(uint8) string
				if arg0.Const() {
					argconst := arg0.Value.(uint8)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint8) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint8)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint8) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint8) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint16:

			{
				var cachedfun func(uint16) string
				if arg0.Const() {
					argconst := arg0.Value.(uint16)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint16) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint16)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint16) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint16) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint32:

			{
				var cachedfun func(uint32) string
				if arg0.Const() {
					argconst := arg0.Value.(uint32)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint32) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint32)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint32) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint32) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uint64:

			{
				var cachedfun func(uint64) string
				if arg0.Const() {
					argconst := arg0.Value.(uint64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint64) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uint64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uint64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uint64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Uintptr:

			{
				var cachedfun func(uintptr) string
				if arg0.Const() {
					argconst := arg0.Value.(uintptr)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uintptr) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) uintptr)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(uintptr) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(uintptr) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float32:

			{
				var cachedfun func(float32) string
				if arg0.Const() {
					argconst := arg0.Value.(float32)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(float32) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float32)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(float32) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float32) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Float64:

			{
				var cachedfun func(float64) string
				if arg0.Const() {
					argconst := arg0.Value.(float64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(float64) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) float64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(float64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(float64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex64:

			{
				var cachedfun func(complex64) string
				if arg0.Const() {
					argconst := arg0.Value.(complex64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(complex64) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex64)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(complex64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex64) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.Complex128:

			{
				var cachedfun func(complex128) string
				if arg0.Const() {
					argconst := arg0.Value.(complex128)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(complex128) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) complex128)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(complex128) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(complex128) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		case r.String:

			{
				var cachedfun func(string) string
				if arg0.Const() {
					argconst := arg0.Value.(string)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(string) string)
							return fun(argconst)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) string)
							return fun(argconst)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) string)
							}
							return cachedfun(argconst)
						}
					}

				} else {
					argfun := arg0.Fun.(func(env *Env) string)
					if funvar == nil {
						call = func(env *Env) string {
							fun := exprfun(env).Interface().(func(string) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funupn != 1 {
						call = func(env *Env) string {
							o := env
							for i := 0; i < funupn; i++ {
								o = o.Outer
							}

							fun := o.Binds[funindex].Interface().(func(string) string)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) string {
							funv := env.Outer.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) string)
							}
							arg := argfun(env)
							return cachedfun(arg)
						}
					}

				}
			}
		default:
			call = func(env *Env) string {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.String()
			}
		}

	default:
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			return funv.Call(argv)[0]
		}

	}
	return call
}
func call2ret1(c *Call) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	karg0 := expr.Type.In(0).Kind()
	karg1 := expr.Type.In(1).Kind()
	kret := expr.Type.Out(0).Kind()
	args := c.Args
	argfuns := c.Argfuns
	var call I
	switch kret {
	case r.Bool:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(bool, bool) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(bool, bool) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(bool, bool) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(bool, bool) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int, int) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int, int) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int, int) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int, int) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int8, int8) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int8, int8) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int8, int8) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int8, int8) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int16, int16) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int16, int16) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int16, int16) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int16, int16) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int32, int32) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int32, int32) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int32, int32) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int32, int32) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int64, int64) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int64, int64) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int64, int64) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(int64, int64) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint, uint) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint, uint) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint, uint) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint, uint) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint8, uint8) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint8, uint8) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint8, uint8) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint8, uint8) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint16, uint16) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint16, uint16) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint16, uint16) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint16, uint16) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint32, uint32) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint32, uint32) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint32, uint32) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint32, uint32) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint64, uint64) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint64, uint64) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint64, uint64) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uint64, uint64) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float32, float32) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float32, float32) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float32, float32) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float32, float32) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float64, float64) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float64, float64) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float64, float64) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(float64, float64) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex64, complex64) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex64, complex64) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex64, complex64) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex64, complex64) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex128, complex128) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex128, complex128) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex128, complex128) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(complex128, complex128) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(string, string) bool)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(string, string) bool)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(string, string) bool)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) bool {
									fun := exprfun(env).Interface().(func(string, string) bool)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) bool {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Bool()
				}
			}

		}
	case r.Int:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(bool, bool) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(bool, bool) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(bool, bool) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(bool, bool) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int, int) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int, int) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int, int) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int, int) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int8, int8) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int8, int8) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int8, int8) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int8, int8) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int16, int16) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int16, int16) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int16, int16) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int16, int16) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int32, int32) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int32, int32) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int32, int32) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int32, int32) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int64, int64) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int64, int64) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int64, int64) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(int64, int64) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint, uint) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint, uint) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint, uint) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint, uint) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint8, uint8) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint8, uint8) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint8, uint8) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint8, uint8) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint16, uint16) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint16, uint16) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint16, uint16) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint16, uint16) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint32, uint32) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint32, uint32) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint32, uint32) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint32, uint32) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint64, uint64) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint64, uint64) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint64, uint64) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uint64, uint64) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float32, float32) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float32, float32) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float32, float32) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float32, float32) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float64, float64) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float64, float64) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float64, float64) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(float64, float64) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex64, complex64) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex64, complex64) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex64, complex64) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex64, complex64) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex128, complex128) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex128, complex128) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex128, complex128) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(complex128, complex128) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(string, string) int)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(string, string) int)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(string, string) int)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int {
									fun := exprfun(env).Interface().(func(string, string) int)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int(ret0.Int())
				}
			}

		}
	case r.Int8:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(bool, bool) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(bool, bool) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(bool, bool) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(bool, bool) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int, int) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int, int) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int, int) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int, int) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int8, int8) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int8, int8) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int8, int8) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int8, int8) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int16, int16) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int16, int16) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int16, int16) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int16, int16) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int32, int32) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int32, int32) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int32, int32) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int32, int32) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int64, int64) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int64, int64) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int64, int64) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(int64, int64) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint, uint) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint, uint) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint, uint) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint, uint) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float32, float32) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float32, float32) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float32, float32) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float32, float32) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float64, float64) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float64, float64) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float64, float64) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(float64, float64) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(string, string) int8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(string, string) int8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(string, string) int8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int8 {
									fun := exprfun(env).Interface().(func(string, string) int8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int8 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int8(ret0.Int())
				}
			}

		}
	case r.Int16:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(bool, bool) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(bool, bool) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(bool, bool) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(bool, bool) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int, int) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int, int) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int, int) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int, int) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int8, int8) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int8, int8) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int8, int8) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int8, int8) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int16, int16) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int16, int16) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int16, int16) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int16, int16) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int32, int32) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int32, int32) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int32, int32) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int32, int32) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int64, int64) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int64, int64) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int64, int64) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(int64, int64) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint, uint) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint, uint) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint, uint) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint, uint) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float32, float32) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float32, float32) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float32, float32) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float32, float32) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float64, float64) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float64, float64) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float64, float64) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(float64, float64) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(string, string) int16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(string, string) int16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(string, string) int16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int16 {
									fun := exprfun(env).Interface().(func(string, string) int16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int16 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int16(ret0.Int())
				}
			}

		}
	case r.Int32:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(bool, bool) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(bool, bool) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(bool, bool) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(bool, bool) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int, int) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int, int) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int, int) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int, int) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int8, int8) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int8, int8) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int8, int8) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int8, int8) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int16, int16) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int16, int16) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int16, int16) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int16, int16) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int32, int32) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int32, int32) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int32, int32) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int32, int32) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int64, int64) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int64, int64) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int64, int64) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(int64, int64) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint, uint) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint, uint) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint, uint) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint, uint) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float32, float32) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float32, float32) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float32, float32) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float32, float32) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float64, float64) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float64, float64) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float64, float64) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(float64, float64) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(string, string) int32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(string, string) int32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(string, string) int32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int32 {
									fun := exprfun(env).Interface().(func(string, string) int32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int32(ret0.Int())
				}
			}

		}
	case r.Int64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(bool, bool) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(bool, bool) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(bool, bool) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(bool, bool) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int, int) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int, int) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int, int) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int, int) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int8, int8) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int8, int8) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int8, int8) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int8, int8) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int16, int16) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int16, int16) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int16, int16) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int16, int16) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int32, int32) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int32, int32) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int32, int32) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int32, int32) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int64, int64) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int64, int64) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int64, int64) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(int64, int64) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint, uint) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint, uint) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint, uint) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint, uint) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float32, float32) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float32, float32) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float32, float32) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float32, float32) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float64, float64) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float64, float64) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float64, float64) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(float64, float64) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(string, string) int64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(string, string) int64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(string, string) int64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int64 {
									fun := exprfun(env).Interface().(func(string, string) int64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Int()
				}
			}

		}
	case r.Uint:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(bool, bool) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(bool, bool) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(bool, bool) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(bool, bool) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int, int) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int, int) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int, int) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int, int) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int8, int8) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int8, int8) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int8, int8) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int8, int8) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int16, int16) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int16, int16) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int16, int16) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int16, int16) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int32, int32) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int32, int32) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int32, int32) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int32, int32) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int64, int64) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int64, int64) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int64, int64) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(int64, int64) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint, uint) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint, uint) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint, uint) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint, uint) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float32, float32) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float32, float32) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float32, float32) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float32, float32) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float64, float64) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float64, float64) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float64, float64) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(float64, float64) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(string, string) uint)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(string, string) uint)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(string, string) uint)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint {
									fun := exprfun(env).Interface().(func(string, string) uint)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint(ret0.Uint())
				}
			}

		}
	case r.Uint8:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(bool, bool) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(bool, bool) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(bool, bool) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(bool, bool) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int, int) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int, int) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int, int) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int, int) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int8, int8) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int8, int8) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int8, int8) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int8, int8) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int16, int16) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int16, int16) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int16, int16) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int16, int16) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int32, int32) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int32, int32) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int32, int32) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int32, int32) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int64, int64) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int64, int64) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int64, int64) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(int64, int64) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint, uint) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint, uint) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint, uint) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint, uint) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float32, float32) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float32, float32) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float32, float32) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float32, float32) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float64, float64) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float64, float64) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float64, float64) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(float64, float64) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(string, string) uint8)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(string, string) uint8)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(string, string) uint8)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint8 {
									fun := exprfun(env).Interface().(func(string, string) uint8)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint8 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint8(ret0.Uint())
				}
			}

		}
	case r.Uint16:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(bool, bool) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(bool, bool) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(bool, bool) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(bool, bool) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int, int) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int, int) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int, int) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int, int) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int8, int8) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int8, int8) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int8, int8) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int8, int8) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int16, int16) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int16, int16) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int16, int16) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int16, int16) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int32, int32) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int32, int32) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int32, int32) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int32, int32) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int64, int64) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int64, int64) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int64, int64) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(int64, int64) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint, uint) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint, uint) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint, uint) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint, uint) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float32, float32) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float32, float32) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float32, float32) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float32, float32) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float64, float64) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float64, float64) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float64, float64) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(float64, float64) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(string, string) uint16)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(string, string) uint16)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(string, string) uint16)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint16 {
									fun := exprfun(env).Interface().(func(string, string) uint16)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint16 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint16(ret0.Uint())
				}
			}

		}
	case r.Uint32:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(bool, bool) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(bool, bool) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(bool, bool) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(bool, bool) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int, int) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int, int) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int, int) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int, int) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int8, int8) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int8, int8) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int8, int8) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int8, int8) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int16, int16) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int16, int16) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int16, int16) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int16, int16) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int32, int32) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int32, int32) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int32, int32) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int32, int32) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int64, int64) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int64, int64) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int64, int64) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(int64, int64) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint, uint) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint, uint) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint, uint) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint, uint) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float32, float32) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float32, float32) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float32, float32) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float32, float32) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float64, float64) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float64, float64) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float64, float64) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(float64, float64) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(string, string) uint32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(string, string) uint32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(string, string) uint32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint32 {
									fun := exprfun(env).Interface().(func(string, string) uint32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint32(ret0.Uint())
				}
			}

		}
	case r.Uint64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(bool, bool) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(bool, bool) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(bool, bool) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(bool, bool) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int, int) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int, int) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int, int) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int, int) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int8, int8) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int8, int8) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int8, int8) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int8, int8) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int16, int16) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int16, int16) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int16, int16) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int16, int16) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int32, int32) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int32, int32) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int32, int32) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int32, int32) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int64, int64) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int64, int64) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int64, int64) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(int64, int64) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint, uint) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint, uint) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint, uint) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint, uint) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float32, float32) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float32, float32) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float32, float32) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float32, float32) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float64, float64) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float64, float64) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float64, float64) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(float64, float64) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(string, string) uint64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(string, string) uint64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(string, string) uint64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint64 {
									fun := exprfun(env).Interface().(func(string, string) uint64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Uint()
				}
			}

		}
	case r.Uintptr:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(bool, bool) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(bool, bool) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(bool, bool) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(bool, bool) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int, int) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int, int) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int, int) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int, int) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int8, int8) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int8, int8) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int8, int8) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int8, int8) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int16, int16) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int16, int16) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int16, int16) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int16, int16) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int32, int32) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int32, int32) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int32, int32) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int32, int32) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int64, int64) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int64, int64) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int64, int64) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(int64, int64) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint, uint) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint, uint) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint, uint) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint, uint) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint8, uint8) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint8, uint8) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint8, uint8) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint8, uint8) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint16, uint16) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint16, uint16) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint16, uint16) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint16, uint16) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint32, uint32) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint32, uint32) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint32, uint32) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint32, uint32) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint64, uint64) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint64, uint64) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint64, uint64) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uint64, uint64) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float32, float32) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float32, float32) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float32, float32) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float32, float32) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float64, float64) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float64, float64) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float64, float64) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(float64, float64) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex64, complex64) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex64, complex64) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex64, complex64) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex64, complex64) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex128, complex128) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex128, complex128) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex128, complex128) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(complex128, complex128) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(string, string) uintptr)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(string, string) uintptr)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(string, string) uintptr)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uintptr {
									fun := exprfun(env).Interface().(func(string, string) uintptr)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uintptr {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uintptr(ret0.Uint())
				}
			}

		}
	case r.Float32:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(bool, bool) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(bool, bool) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(bool, bool) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(bool, bool) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int, int) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int, int) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int, int) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int, int) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int8, int8) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int8, int8) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int8, int8) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int8, int8) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int16, int16) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int16, int16) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int16, int16) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int16, int16) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int32, int32) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int32, int32) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int32, int32) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int32, int32) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int64, int64) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int64, int64) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int64, int64) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(int64, int64) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint, uint) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint, uint) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint, uint) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint, uint) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float32, float32) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float32, float32) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float32, float32) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float32, float32) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float64, float64) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float64, float64) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float64, float64) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(float64, float64) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(string, string) float32)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(string, string) float32)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(string, string) float32)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float32 {
									fun := exprfun(env).Interface().(func(string, string) float32)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) float32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return float32(ret0.Float())
				}
			}

		}
	case r.Float64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(bool, bool) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(bool, bool) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(bool, bool) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(bool, bool) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int, int) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int, int) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int, int) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int, int) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int8, int8) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int8, int8) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int8, int8) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int8, int8) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int16, int16) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int16, int16) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int16, int16) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int16, int16) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int32, int32) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int32, int32) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int32, int32) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int32, int32) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int64, int64) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int64, int64) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int64, int64) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(int64, int64) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint, uint) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint, uint) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint, uint) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint, uint) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float32, float32) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float32, float32) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float32, float32) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float32, float32) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float64, float64) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float64, float64) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float64, float64) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(float64, float64) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(string, string) float64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(string, string) float64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(string, string) float64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float64 {
									fun := exprfun(env).Interface().(func(string, string) float64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) float64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Float()
				}
			}

		}
	case r.Complex64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(bool, bool) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(bool, bool) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(bool, bool) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(bool, bool) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int, int) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int, int) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int, int) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int, int) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int8, int8) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int8, int8) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int8, int8) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int8, int8) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int16, int16) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int16, int16) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int16, int16) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int16, int16) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int32, int32) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int32, int32) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int32, int32) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int32, int32) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int64, int64) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int64, int64) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int64, int64) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(int64, int64) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint, uint) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint, uint) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint, uint) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint, uint) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float32, float32) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float32, float32) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float32, float32) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float32, float32) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float64, float64) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float64, float64) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float64, float64) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(float64, float64) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(string, string) complex64)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(string, string) complex64)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(string, string) complex64)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex64 {
									fun := exprfun(env).Interface().(func(string, string) complex64)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) complex64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return complex64(ret0.Complex())
				}
			}

		}
	case r.Complex128:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(bool, bool) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(bool, bool) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(bool, bool) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(bool, bool) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int, int) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int, int) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int, int) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int, int) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int8, int8) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int8, int8) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int8, int8) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int8, int8) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int16, int16) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int16, int16) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int16, int16) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int16, int16) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int32, int32) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int32, int32) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int32, int32) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int32, int32) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int64, int64) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int64, int64) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int64, int64) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(int64, int64) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint, uint) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint, uint) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint, uint) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint, uint) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint8, uint8) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint16, uint16) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint32, uint32) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uint64, uint64) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float32, float32) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float32, float32) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float32, float32) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float32, float32) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float64, float64) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float64, float64) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float64, float64) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(float64, float64) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex64, complex64) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(complex128, complex128) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(string, string) complex128)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(string, string) complex128)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(string, string) complex128)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex128 {
									fun := exprfun(env).Interface().(func(string, string) complex128)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) complex128 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Complex()
				}
			}

		}
	case r.String:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(bool, bool) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(bool, bool) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(bool, bool) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(bool, bool) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int, int) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int, int) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int, int) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int, int) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int8, int8) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int8, int8) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int8, int8) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int8, int8) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int16, int16) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int16, int16) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int16, int16) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int16, int16) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int32, int32) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int32, int32) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int32, int32) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int32, int32) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int64, int64) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int64, int64) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int64, int64) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(int64, int64) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint, uint) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint, uint) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint, uint) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint, uint) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint8, uint8) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint8, uint8) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint8, uint8) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint8, uint8) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint16, uint16) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint16, uint16) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint16, uint16) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint16, uint16) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint32, uint32) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint32, uint32) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint32, uint32) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint32, uint32) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint64, uint64) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint64, uint64) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint64, uint64) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uint64, uint64) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(uintptr, uintptr) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float32, float32) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float32, float32) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float32, float32) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float32, float32) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float64, float64) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float64, float64) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float64, float64) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(float64, float64) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex64, complex64) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex64, complex64) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex64, complex64) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex64, complex64) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex128, complex128) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex128, complex128) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex128, complex128) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(complex128, complex128) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(string, string) string)
									return fun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(string, string) string)
									arg1 := arg1fun(env)
									return fun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(string, string) string)
									arg0 := arg0fun(env)
									return fun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) string {
									fun := exprfun(env).Interface().(func(string, string) string)
									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return fun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) string {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.String()
				}
			}

		}
	default:
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			return funv.Call(argv)[0]
		}

	}
	return call
}
func callnret1(c *Call) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	kret := expr.Type.Out(0).Kind()
	argfuns := c.Argfuns
	var call I
	switch kret {
	case r.Bool:
		call = func(env *Env) bool {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Bool()
		}
	case r.Int:
		call = func(env *Env) int {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int(ret0.Int())
		}
	case r.Int8:
		call = func(env *Env) int8 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int8(ret0.Int())
		}
	case r.Int16:
		call = func(env *Env) int16 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int16(ret0.Int())
		}
	case r.Int32:
		call = func(env *Env) int32 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int32(ret0.Int())
		}
	case r.Int64:
		call = func(env *Env) int64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Int()
		}
	case r.Uint:
		call = func(env *Env) uint {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint(ret0.Uint())
		}
	case r.Uint8:
		call = func(env *Env) uint8 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint8(ret0.Uint())
		}
	case r.Uint16:
		call = func(env *Env) uint16 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint16(ret0.Uint())
		}

	case r.Uint32:
		call = func(env *Env) uint32 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint32(ret0.Uint())
		}

	case r.Uint64:
		call = func(env *Env) uint64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Uint()
		}

	case r.Uintptr:
		call = func(env *Env) uintptr {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uintptr(ret0.Uint())
		}

	case r.Float32:
		call = func(env *Env) float32 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return float32(ret0.Float())
		}

	case r.Float64:
		call = func(env *Env) float64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Float()
		}

	case r.Complex64:
		call = func(env *Env) complex64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return complex64(ret0.Complex())
		}

	case r.Complex128:
		call = func(env *Env) complex128 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Complex()
		}

	case r.String:
		call = func(env *Env) string {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.String()
		}

	default:
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0

		}

	}
	return call
}
func call_ret1(c *Call) I {
	expr := c.Fun
	var call I

	switch expr.Type.NumIn() {
	case 0:
		call = call0ret1(c)
	case 1:
		call = call1ret1(c)
	case 2:
		call = call2ret1(c)
	default:
		call = callnret1(c)
	}
	return call
}