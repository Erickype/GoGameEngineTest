package main

import (
	overlay "github.com/Erickype/GoGameEngine/API/Layers/ImGui"
	"github.com/Erickype/GoGameEngine/API/Platform/Windows"
	"github.com/Erickype/GoGameEngine/Core"
	"github.com/Erickype/GoGameEngineTest/Application"
	"github.com/Erickype/GoGameEngineTest/Layers"
)

func main() {
	//Create and initializes the application
	application := Application.Application{Application: &Core.Application{}}
	application.Construct(Windows.CreateAbstractWindow("Game Test", 800, 600))

	//Push the example layer
	exampleLayer := Layers.ExampleLayer{}
	exampleLayer.Construct("Example")
	iLayer := Core.ILayer(&exampleLayer)
	application.PushLayer(&iLayer)

	//Push the imGui overlay
	imGuiLayer := Core.ILayer(overlay.NewImGui())
	application.PushOverlay(&imGuiLayer)

	//Run the application
	application.Run()

	//Destroy the application
	application.Destroy()
}
