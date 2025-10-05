# fluxflow

[![Go](https://github.com/emicklei/fluxflow/actions/workflows/go.yml/badge.svg)](https://github.com/emicklei/fluxflow/actions/workflows/go.yml)

a Go interpreter that will serve a live code and debug experience

## install

    go install github.com/emicklei/fluxflow/cmd/fluxflow@latest

## Use CLI

    fluxflow run .

## Use as package

    ```go
    package main

    import (
        "github.com/emicklei/fluxflow"
    )

    func main() {
        pkg, _ := fluxflow.LoadPackage(".")
        time, _ := fluxflow.Call(pkg,"addDays",1)
        fmt.Println(time)
    }    
    ```

## Features

| Feature | Implemented |
|---|---|
| Imports | ✅ |
| Variable declaration `var` | ✅ |
| Constant declaration `const` | ✅ |
| Assignment `=`, `:=` | ✅ |
| Functions `func` | ✅ |
| Function calls | ✅ |
| `return` statement | ✅ |
| `if/else` statements | ✅ |
| `for` loops | ✅ |
| `for..range` loops | ✅ |
| Basic literal types (int,string,rune,...) | ✅ |
| Composite type `array` | ✅ |
| Composite type `slice` | ✅ |
| Package `init` | ✅ |
| Binary and Unary Operators | ✅ |
| Composite type `struct` | ✅ |
| Composite type `map` | ✅ |
| Unsigned integer arithmetic | ✅ |
| `switch` statement | ✅ |
| Pointers | ⬜ |
| Interfaces | ⬜ |
| Methods | ⬜ |
| Goroutines `go` | ⬜ |
| Channels `chan` | ⬜ |
| `select` statement | ⬜ |
| Anonymous `func` | ⬜ |
| Function literal | ✅ |
| Type alias | ✅ |
| defer | ✅ |



&copy; 2025. https://ernestmicklei.com . MIT License