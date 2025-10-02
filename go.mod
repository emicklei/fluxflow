module github.com/emicklei/fluxflow

go 1.25

replace github.com/emicklei/structexplorer => ../structexplorer

require golang.org/x/tools v0.37.0

require github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect

require (
	github.com/emicklei/structexplorer v0.0.0
	golang.org/x/mod v0.28.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
)
