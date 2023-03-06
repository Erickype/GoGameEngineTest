package Layers

import (
	"github.com/Erickype/GoGameEngine/Common"
	core "github.com/Erickype/GoGameEngine/Core"
	"github.com/Erickype/GoGameEngine/Events"
)

type ExampleLayer struct {
	core.Layer
}

func (e *ExampleLayer) OnUpdate() {
	//Common.ClientLogger.Info(e.GetName(), "Layer: Update")
}

func (e *ExampleLayer) OnEvent(event *Events.IEvent) {
	Common.ClientLogger.Trace(event)
}
