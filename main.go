package main

import (
	"github.com/Erickype/GoGameEngine/Core"
	app "github.com/Erickype/GoGameEngineTest/Application"
	"github.com/Erickype/GoGameEngineTest/Layers"
)

func main() {

	exampleLayer := Layers.ExampleLayer{}
	exampleLayer.Construct("Example")
	iLayer := Core.ILayer(&exampleLayer)

	application := app.Application{}
	application.Construct(&iLayer)
}
