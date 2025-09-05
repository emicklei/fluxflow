# gostep
gostep is a package to interpret a Go program.
it will parse an existing valid Go program.
The goal is to create an interactive debugger or a REPL for Go, allowing step-by-step execution and inspection of variables and runtime modification of code.

## dependencies
the Go program can use packages from the standard library.
the Go program can use external packages as defined in the go.mod file.
code from packages outside the program will never be interpreted.
packages outside the program are those from the standard library and the external dependencies.
code in the Go program can be organised in local packages.
only code from the Go program will be fully interpreted.
code from external packages are evaluated using the reflect Go package.

# testing
gostep will have near 100% code coverage.

# design
gostep will start by parsing all the files of the Go program includes all its subpackages.
the standard Go parser package is used to create the source AST.
gostep will create a call graph from the source AST.