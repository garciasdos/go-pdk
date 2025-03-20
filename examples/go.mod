module github.com/Kong/gp-pdk/examples

go 1.21

require github.com/garciasdos/go-pdk v0.12.1

require (
	github.com/ugorji/go/codec v1.2.12 // indirect
	google.golang.org/protobuf v1.36.2 // indirect
)

replace github.com/garciasdos/go-pdk => ../
