package Layers

import (
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Input"
	"github.com/Erickype/GoGameEngine/API/Log"
	core "github.com/Erickype/GoGameEngine/Core"
)

type ExampleLayer struct {
	core.Layer
}

func (e *ExampleLayer) OnUpdate() {
	Log.GetClientInstance().Trace(e.GetName(), " layer: update")
	if (*Input.GetInputInstance()).IsKeyPressed((int)(Input.KeySpace), (*core.ApplicationInstance.GetPlatform()).GetWindowPtr()) {
		Log.GetClientInstance().Warn("Key pressed")
	}
}

func (e *ExampleLayer) OnEvent(event *Events.IEvent) {
	Log.GetClientInstance().Info(e.GetName(), " layer event (event): ", (*event).GetName())
}
