# fluxflow

[![Go](https://github.com/emicklei/fluxflow/actions/workflows/go.yml/badge.svg)](https://github.com/emicklei/fluxflow/actions/workflows/go.yml)

a Go interpreter that eventually should serve a live coding and debug experience.

## status

This is work in progress.
See [examples](./examples) for runnable examples using the `fluxflow` cli.
See [status](STATUS.md) for the supported Go language features.

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
    fluxflow.Run("path/to/program")        
}
```

&copy; 2025. https://ernestmicklei.com . MIT License