/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * compile.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/gls"
)

func NewComp(outer *Comp, code *Code) *Comp {
	if outer == nil {
		return &Comp{UpCost: 1}
	}
	c := Comp{
		UpCost:      1,
		Depth:       outer.Depth + 1,
		Outer:       outer,
		CompGlobals: outer.CompGlobals,
	}
	// Debugf("NewComp(%p->%p) %s", outer, &c, debug.Stack())
	if code != nil {
		c.Code = *code
	}
	return &c
}

func (c *Comp) TopComp() *Comp {
	for ; c != nil; c = c.Outer {
		if c.Outer == nil {
			break
		}
	}
	return c
}

func (c *Comp) FileComp() *Comp {
	for ; c != nil; c = c.Outer {
		outer := c.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return c
}

func NewIrGlobals() *IrGlobals {
	return &IrGlobals{
		Globals: *NewGlobals(),
		gls:     make(map[uintptr]*Run),
	}
}

func (g *IrGlobals) glsGet(goid uintptr) *Run {
	g.lock.Lock()
	ret := g.gls[goid]
	g.lock.Unlock()
	return ret
}

func (run *Run) getRun4Goid(goid uintptr) *Run {
	g := run.IrGlobals
	ret := g.glsGet(goid)
	if ret == nil {
		ret = run.new(goid)
		ret.glsStore()
	}
	return ret
}

func (tg *Run) glsStore() {
	g := tg.IrGlobals
	goid := tg.goid
	g.lock.Lock()
	g.gls[goid] = tg
	g.lock.Unlock()
}

func (tg *Run) glsDel() {
	g := tg.IrGlobals
	goid := tg.goid
	g.lock.Lock()
	delete(g.gls, goid)
	g.lock.Unlock()
}

func (run *Run) new(goid uintptr) *Run {
	return &Run{
		IrGlobals: run.IrGlobals,
		goid:      goid,
		// Interrupt, Signal, PoolSize and Pool are zero-initialized, fine with that
	}
}

// if a function Env only declares ignored binds, it gets this scratch buffers
var ignoredBinds = []r.Value{Nil}
var ignoredIntBinds = []uint64{0}

// common part between NewEnv() and newEnv4Func()
func newEnv(run *Run, outer *Env, nbind int, nintbind int) *Env {
	pool := &run.Pool // pool is an array, do NOT copy it!
	index := run.PoolSize - 1
	var env *Env
	if index >= 0 {
		run.PoolSize = index
		env = pool[index]
		pool[index] = nil
	} else {
		env = &Env{}
	}
	if nbind <= 1 {
		env.Vals = ignoredBinds
	} else if cap(env.Vals) < nbind {
		env.Vals = make([]r.Value, nbind)
	} else {
		env.Vals = env.Vals[0:nbind]
	}
	if nintbind <= 1 {
		env.Ints = ignoredIntBinds
	} else if cap(env.Ints) < nintbind {
		env.Ints = make([]uint64, nintbind)
	} else {
		env.Ints = env.Ints[0:nintbind]
	}
	env.Outer = outer
	env.Run = run
	env.FileEnv = outer.FileEnv
	return env
}

// return a new, nested Env with given number of binds and intbinds
func NewEnv(outer *Env, nbind int, nintbind int) *Env {
	run := outer.Run

	// manually inline
	// env := newEnv(run, outer, nbind, nintbind)
	var env *Env
	{
		pool := &run.Pool // pool is an array, do NOT copy it!
		index := run.PoolSize - 1
		if index >= 0 {
			run.PoolSize = index
			env = pool[index]
			pool[index] = nil
		} else {
			env = &Env{}
		}
		if nbind <= 1 {
			env.Vals = ignoredBinds
		} else if cap(env.Vals) < nbind {
			env.Vals = make([]r.Value, nbind)
		} else {
			env.Vals = env.Vals[0:nbind]
		}
		if nintbind <= 1 {
			env.Ints = ignoredIntBinds
		} else if cap(env.Ints) < nintbind {
			env.Ints = make([]uint64, nintbind)
		} else {
			env.Ints = env.Ints[0:nintbind]
		}
		env.Outer = outer
		env.Run = run
		env.FileEnv = outer.FileEnv
	}
	env.IP = outer.IP
	env.Code = outer.Code
	env.DebugPos = outer.DebugPos
	env.CallDepth = outer.CallDepth
	// this is a nested *Env, not a function body: to obtain the caller function,
	// follow env.Outer.Outer... chain until you find an *Env with non-nil Caller
	// env.Caller = nil
	// DebugCallStack Debugf("NewEnv(%p->%p) nbind=%d nintbind=%d calldepth: %d->%d", outer, env, nbind, nintbind, outer.CallDepth, env.CallDepth)
	run.CurrEnv = env
	return env
}

func newEnv4Func(outer *Env, nbind int, nintbind int, debugComp *Comp) *Env {
	goid := gls.GoID()
	run := outer.Run
	if run.goid != goid {
		// no luck... get the correct ThreadGlobals for goid
		run = run.getRun4Goid(goid)
	}
	// manually inline
	// env := newEnv(run, outer, nbind, nintbind)
	var env *Env
	{
		pool := &run.Pool // pool is an array, do NOT copy it!
		index := run.PoolSize - 1
		if index >= 0 {
			run.PoolSize = index
			env = pool[index]
			pool[index] = nil
		} else {
			env = &Env{}
		}
		if nbind <= 1 {
			env.Vals = ignoredBinds
		} else if cap(env.Vals) < nbind {
			env.Vals = make([]r.Value, nbind)
		} else {
			env.Vals = env.Vals[0:nbind]
		}
		if nintbind <= 1 {
			env.Ints = ignoredIntBinds
		} else if cap(env.Ints) < nintbind {
			env.Ints = make([]uint64, nintbind)
		} else {
			env.Ints = env.Ints[0:nintbind]
		}
		env.Outer = outer
		env.Run = run
		env.FileEnv = outer.FileEnv
	}
	env.DebugComp = debugComp
	caller := run.CurrEnv
	env.Caller = caller
	if caller == nil {
		env.CallDepth = 1
	} else {
		env.CallDepth = caller.CallDepth + 1
	}
	// DebugCallStack Debugf("newEnv4Func(%p->%p) nbind=%d nintbind=%d calldepth: %d->%d", caller, env, nbind, nintbind, env.CallDepth-1, env.CallDepth)
	run.CurrEnv = env
	return env
}

func (env *Env) MarkUsedByClosure() {
	for ; env != nil && !env.UsedByClosure; env = env.Outer {
		env.UsedByClosure = true
	}
}

// FreeEnv tells the interpreter that given nested *Env is no longer needed.
func (env *Env) FreeEnv() {
	run := env.Run
	run.CurrEnv = env.Outer
	env.freeEnv(run)
}

// freeEnv4Func tells the interpreter that given function body *Env is no longer needed.
func (env *Env) freeEnv4Func() {
	run := env.Run
	run.CurrEnv = env.Caller
	env.freeEnv(run)
}

func (env *Env) freeEnv(run *Run) {
	// DebugCallStack Debugf("FreeEnv(%p->%p), calldepth: %d->%d", env, caller, env.CallDepth, caller.CallDepth)
	if env.UsedByClosure {
		// in use, cannot recycle
		return
	}
	n := run.PoolSize
	if n >= poolCapacity {
		return
	}
	if env.IntAddressTaken {
		env.Ints = nil
		env.IntAddressTaken = false
	}
	env.Outer = nil
	env.Code = nil
	env.DebugPos = nil
	env.DebugComp = nil
	env.Caller = nil
	env.Run = nil
	env.FileEnv = nil
	run.Pool[n] = env // pool is an array, be careful NOT to copy it!
	run.PoolSize = n + 1
}

func (env *Env) Top() *Env {
	for ; env != nil; env = env.Outer {
		if env.Outer == nil {
			break
		}
	}
	return env
}

func (env *Env) Up(n int) *Env {
	for ; n >= 3; n -= 3 {
		env = env.Outer.Outer.Outer
	}
	switch n {
	case 2:
		env = env.Outer
		fallthrough
	case 1:
		env = env.Outer
	}
	return env
}

// combined Parse + MacroExpandCodeWalk
func (c *Comp) Parse(src string) Ast {
	c.Line = 0
	nodes := c.ParseBytes([]byte(src))
	forms := anyToAst(nodes, "Parse")

	forms, _ = c.MacroExpandCodewalk(forms)
	if c.Options&OptShowMacroExpand != 0 {
		c.Debugf("after macroexpansion: %v", forms.Interface())
	}
	return forms
}

func (c *Comp) Compile(in Ast) *Expr {
	switch form := in.(type) {
	case nil:
		return nil
	case AstWithNode:
		return c.CompileNode(form.Node())
	case AstWithSlice:
		n := form.Size()
		var list []*Expr
		for i := 0; i < n; i++ {
			e := c.Compile(form.Get(i))
			if e != nil {
				list = append(list, e)
			}
		}
		return exprList(list, c.CompileOptions())
	}
	c.Errorf("unsupported Ast node, expecting <AstWithNode> or <AstWithSlice>, found %v <%v>", in, r.TypeOf(in))
	return nil
}

// compileExpr is a wrapper for Compile
// that guarantees Code does not get clobbered/cleared.
// Used by Comp.Quasiquote
func (c *Comp) compileExpr(in Ast) *Expr {
	cf := NewComp(c, nil)
	cf.UpCost = 0
	cf.Depth--
	return cf.Compile(in)
}

func (c *Comp) CompileNode(node ast.Node) *Expr {
	if n := c.Code.Len(); n != 0 {
		c.Warnf("Compile: discarding %d previously compiled statements from code buffer", n)
	}
	c.Code.Clear()
	if node == nil {
		return nil
	}
	c.Pos = node.Pos()
	switch node := node.(type) {
	case ast.Decl:
		c.Decl(node)
	case ast.Expr:
		return c.Expr(node, nil)
	case *ast.ExprStmt:
		// special case of statement
		return c.Expr(node.X, nil)
	case ast.Stmt:
		c.Stmt(node)
	case *ast.File:
		c.File(node)
	default:
		c.Errorf("unsupported node type, expecting <ast.Decl>, <ast.Expr>, <ast.Stmt> or <*ast.File>, found %v <%v>", node, r.TypeOf(node))
		return nil
	}
	return c.Code.AsExpr()
}

func (c *Comp) File(node *ast.File) {
	c.Name = node.Name.Name
	for _, decl := range node.Decls {
		c.Decl(decl)
	}
}

func (c *Comp) Append(stmt Stmt, pos token.Pos) {
	c.Code.Append(stmt, pos)
}

func (c *Comp) append(stmt Stmt) {
	c.Code.Append(stmt, c.Pos)
}
