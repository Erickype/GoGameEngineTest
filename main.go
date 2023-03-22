package main

import (
	Layers "github.com/Erickype/GoGameEngine/API/Layers/ImGui"
	"github.com/Erickype/GoGameEngine/Core"
	app "github.com/Erickype/GoGameEngineTest/Application"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	exampleLayer := Layers.NewImGui()
	iLayer := Core.ILayer(exampleLayer)

	application := app.Application{}
	application.Construct(&iLayer)
}
