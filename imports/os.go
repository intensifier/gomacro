// this file was generated by gomacro command: import "os"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "os"
	. "reflect"
)

func Package_os() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Args":              ValueOf(&pkg.Args).Elem(),
			"Chdir":             ValueOf(pkg.Chdir),
			"Chmod":             ValueOf(pkg.Chmod),
			"Chown":             ValueOf(pkg.Chown),
			"Chtimes":           ValueOf(pkg.Chtimes),
			"Clearenv":          ValueOf(pkg.Clearenv),
			"Create":            ValueOf(pkg.Create),
			"DevNull":           ValueOf(pkg.DevNull),
			"Environ":           ValueOf(pkg.Environ),
			"ErrClosed":         ValueOf(&pkg.ErrClosed).Elem(),
			"ErrExist":          ValueOf(&pkg.ErrExist).Elem(),
			"ErrInvalid":        ValueOf(&pkg.ErrInvalid).Elem(),
			"ErrNotExist":       ValueOf(&pkg.ErrNotExist).Elem(),
			"ErrPermission":     ValueOf(&pkg.ErrPermission).Elem(),
			"Executable":        ValueOf(pkg.Executable),
			"Exit":              ValueOf(pkg.Exit),
			"Expand":            ValueOf(pkg.Expand),
			"ExpandEnv":         ValueOf(pkg.ExpandEnv),
			"FindProcess":       ValueOf(pkg.FindProcess),
			"Getegid":           ValueOf(pkg.Getegid),
			"Getenv":            ValueOf(pkg.Getenv),
			"Geteuid":           ValueOf(pkg.Geteuid),
			"Getgid":            ValueOf(pkg.Getgid),
			"Getgroups":         ValueOf(pkg.Getgroups),
			"Getpagesize":       ValueOf(pkg.Getpagesize),
			"Getpid":            ValueOf(pkg.Getpid),
			"Getppid":           ValueOf(pkg.Getppid),
			"Getuid":            ValueOf(pkg.Getuid),
			"Getwd":             ValueOf(pkg.Getwd),
			"Hostname":          ValueOf(pkg.Hostname),
			"Interrupt":         ValueOf(&pkg.Interrupt).Elem(),
			"IsExist":           ValueOf(pkg.IsExist),
			"IsNotExist":        ValueOf(pkg.IsNotExist),
			"IsPathSeparator":   ValueOf(pkg.IsPathSeparator),
			"IsPermission":      ValueOf(pkg.IsPermission),
			"Kill":              ValueOf(&pkg.Kill).Elem(),
			"Lchown":            ValueOf(pkg.Lchown),
			"Link":              ValueOf(pkg.Link),
			"LookupEnv":         ValueOf(pkg.LookupEnv),
			"Lstat":             ValueOf(pkg.Lstat),
			"Mkdir":             ValueOf(pkg.Mkdir),
			"MkdirAll":          ValueOf(pkg.MkdirAll),
			"ModeAppend":        ValueOf(pkg.ModeAppend),
			"ModeCharDevice":    ValueOf(pkg.ModeCharDevice),
			"ModeDevice":        ValueOf(pkg.ModeDevice),
			"ModeDir":           ValueOf(pkg.ModeDir),
			"ModeExclusive":     ValueOf(pkg.ModeExclusive),
			"ModeNamedPipe":     ValueOf(pkg.ModeNamedPipe),
			"ModePerm":          ValueOf(pkg.ModePerm),
			"ModeSetgid":        ValueOf(pkg.ModeSetgid),
			"ModeSetuid":        ValueOf(pkg.ModeSetuid),
			"ModeSocket":        ValueOf(pkg.ModeSocket),
			"ModeSticky":        ValueOf(pkg.ModeSticky),
			"ModeSymlink":       ValueOf(pkg.ModeSymlink),
			"ModeTemporary":     ValueOf(pkg.ModeTemporary),
			"ModeType":          ValueOf(pkg.ModeType),
			"NewFile":           ValueOf(pkg.NewFile),
			"NewSyscallError":   ValueOf(pkg.NewSyscallError),
			"O_APPEND":          ValueOf(pkg.O_APPEND),
			"O_CREATE":          ValueOf(pkg.O_CREATE),
			"O_EXCL":            ValueOf(pkg.O_EXCL),
			"O_RDONLY":          ValueOf(pkg.O_RDONLY),
			"O_RDWR":            ValueOf(pkg.O_RDWR),
			"O_SYNC":            ValueOf(pkg.O_SYNC),
			"O_TRUNC":           ValueOf(pkg.O_TRUNC),
			"O_WRONLY":          ValueOf(pkg.O_WRONLY),
			"Open":              ValueOf(pkg.Open),
			"OpenFile":          ValueOf(pkg.OpenFile),
			"PathListSeparator": ValueOf(pkg.PathListSeparator),
			"PathSeparator":     ValueOf(pkg.PathSeparator),
			"Pipe":              ValueOf(pkg.Pipe),
			"Readlink":          ValueOf(pkg.Readlink),
			"Remove":            ValueOf(pkg.Remove),
			"RemoveAll":         ValueOf(pkg.RemoveAll),
			"Rename":            ValueOf(pkg.Rename),
			"SEEK_CUR":          ValueOf(pkg.SEEK_CUR),
			"SEEK_END":          ValueOf(pkg.SEEK_END),
			"SEEK_SET":          ValueOf(pkg.SEEK_SET),
			"SameFile":          ValueOf(pkg.SameFile),
			"Setenv":            ValueOf(pkg.Setenv),
			"StartProcess":      ValueOf(pkg.StartProcess),
			"Stat":              ValueOf(pkg.Stat),
			"Stderr":            ValueOf(&pkg.Stderr).Elem(),
			"Stdin":             ValueOf(&pkg.Stdin).Elem(),
			"Stdout":            ValueOf(&pkg.Stdout).Elem(),
			"Symlink":           ValueOf(pkg.Symlink),
			"TempDir":           ValueOf(pkg.TempDir),
			"Truncate":          ValueOf(pkg.Truncate),
			"Unsetenv":          ValueOf(pkg.Unsetenv),
		}, map[string]Type{
			"File":         TypeOf((*pkg.File)(nil)).Elem(),
			"FileInfo":     TypeOf((*pkg.FileInfo)(nil)).Elem(),
			"FileMode":     TypeOf((*pkg.FileMode)(nil)).Elem(),
			"LinkError":    TypeOf((*pkg.LinkError)(nil)).Elem(),
			"PathError":    TypeOf((*pkg.PathError)(nil)).Elem(),
			"ProcAttr":     TypeOf((*pkg.ProcAttr)(nil)).Elem(),
			"Process":      TypeOf((*pkg.Process)(nil)).Elem(),
			"ProcessState": TypeOf((*pkg.ProcessState)(nil)).Elem(),
			"Signal":       TypeOf((*pkg.Signal)(nil)).Elem(),
			"SyscallError": TypeOf((*pkg.SyscallError)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_os()
	Binds["os"] = binds
	Types["os"] = types
}
