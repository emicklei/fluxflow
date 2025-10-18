package internal

import (
	"fmt"
	"os"
	"reflect"
)

var trace = os.Getenv("TRACE") != ""

type Env interface {
	valueLookUp(name string) reflect.Value
	typeLookUp(name string) reflect.Type
	valueOwnerOf(name string) Env
	set(name string, value reflect.Value)
	newChild() Env
	depth() int
	getParent() Env
	addConstOrVar(cv ConstOrVar)
	rootPackageEnv() *PkgEnvironment
}

type PkgEnvironment struct {
	Env
	declTable    map[string]CanDeclare
	inits        []FuncDecl
	packageTable map[string]Package // path -> Package
}

func newBuiltinsEnvironment(parent Env) Env {
	return &Environment{
		parent:     parent,
		valueTable: builtinsMap,
	}
}

func newPkgEnvironment(parent Env) Env {
	return &PkgEnvironment{
		Env:          newEnvironment(parent),
		declTable:    map[string]CanDeclare{},
		packageTable: map[string]Package{},
	}
}
func (p *PkgEnvironment) addInit(f FuncDecl) {
	p.inits = append(p.inits, f)
}

func (p *PkgEnvironment) addConstOrVar(cv ConstOrVar) {
	p.declTable[cv.Name.Name] = CanDeclare(cv)
}

// registerPackage adds the package to the package table.
// It returns true if the package was newly added, false if it already existed.
func (p *PkgEnvironment) registerPackage(pkg Package) bool {
	if _, exists := p.packageTable[pkg.Path]; exists {
		return false
	}
	p.packageTable[pkg.Path] = pkg
	return true
}

func (p *PkgEnvironment) rootPackageEnv() *PkgEnvironment {
	if p.getParent() == nil {
		return p
	}
	if _, ok := p.getParent().(*Environment); ok {
		return p
	}
	return p.getParent().rootPackageEnv()
}

func (p *PkgEnvironment) String() string {
	return fmt.Sprintf("PkgEnvironment(pkgs=%d)", len(p.packageTable))
}

func (p *PkgEnvironment) newChild() Env {
	return newEnvironment(p)
}

type Environment struct {
	parent     Env
	valueTable map[string]reflect.Value
}

func newEnvironment(parentOrNil Env) Env {
	return &Environment{
		parent:     parentOrNil,
		valueTable: map[string]reflect.Value{},
	}
}

func (e *Environment) getParent() Env {
	return e.parent
}

func (e *Environment) depth() int {
	if e.parent == nil {
		return 0
	}
	return e.parent.depth() + 1
}

func (e *Environment) String() string {
	return fmt.Sprintf("-- env[depth=%d,len=%d]", e.depth(), len(e.valueTable))
}

func (e *Environment) newChild() Env {
	return newEnvironment(e)
}
func (e *Environment) valueLookUp(name string) reflect.Value {
	v, ok := e.valueTable[name]
	if !ok {
		if e.parent == nil {
			return reflect.Value{}
		}
		return e.parent.valueLookUp(name)
	}
	return v
}

func (e *Environment) typeLookUp(name string) reflect.Type {
	v, ok := builtinTypesMap[name]
	if !ok {
		return nil
	}
	return v
}

func (e *Environment) valueOwnerOf(name string) Env {
	_, ok := e.valueTable[name]
	if !ok {
		if e.parent == nil {
			return nil
		}
		return e.parent.valueOwnerOf(name)
	}
	return e
}

func (e *Environment) set(name string, value reflect.Value) {
	if trace {
		fmt.Println(e, name, "=", value.Interface())
	}
	e.valueTable[name] = value
}

func (e *Environment) addConstOrVar(cv ConstOrVar) {}

func (e *Environment) rootPackageEnv() *PkgEnvironment {
	if e.parent == nil {
		return nil
	}
	return e.parent.rootPackageEnv()
}
