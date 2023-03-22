package Application

import (
	"github.com/Erickype/GoGameEngine/API/Platform/Windows"
	core "github.com/Erickype/GoGameEngine/Core"
)

type _ interface {
	Construct(layer *core.ILayer)
}

type Application struct {
	*core.Application
}

func (a *Application) Construct(layer *core.ILayer) {
	core.CreateApplication(layer, Windows.CreateAbstractWindow("Game Test", 800, 600))
}
