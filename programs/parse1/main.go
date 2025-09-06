package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fset := token.NewFileSet()
	dirPath := "../test1"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	pkg := &ast.Package{
		Name:  "main",
		Files: make(map[string]*ast.File),
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			filePath := filepath.Join(dirPath, file.Name())
			f, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
			if err != nil {
				log.Printf("could not parse file %s: %v", filePath, err)
				continue
			}
			pkg.Files[filePath] = f
			if pkg.Name == "" {
				pkg.Name = f.Name.Name
			}
		}
	}

	fmt.Println("Package:", pkg.Name)
	err = ast.Print(fset, pkg)
	if err != nil {
		log.Fatal(err)
	}
}
