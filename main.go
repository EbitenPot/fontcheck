package main

//go:generate gofmt -s -w .

import (
	"github.com/EbitenPot/fontcheck/form"
	"github.com/ying32/govcl/vcl"
)


func main() {
    vcl.Application.SetScaled(true)
	vcl.Application.SetName("fontcheck")
	// vcl.Application.SetIcon(getIcon())
    vcl.Application.SetTitle("fontcheck")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&form.Init)
	vcl.Application.Run()
}
