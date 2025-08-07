package main

import (
	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/sdlbackend"
	"github.com/AllenDang/cimgui-go/examples/common"
)

func main() {
	var currentBackend backend.Backend[sdlbackend.SDLWindowFlags]
	common.Initialize()
	currentBackend, _ = backend.CreateBackend(sdlbackend.NewSDLBackend())
	currentBackend.SetAfterCreateContextHook(common.AfterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(common.BeforeDestroyContext)

	currentBackend.Run(common.Loop)
}
