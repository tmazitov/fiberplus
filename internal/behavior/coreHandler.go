package behavior

import mods "github.com/tmazitov/fiberplus/internal/mods"

type CoreHandler[Services any, Input any, Output any] struct {
	DefaultHandler[Services]
}

func (h *CoreHandler[Services, Input, Output]) Init() {
	h.introMods = []mods.Mod{
		&mods.JsonInputMod[Input]{},
	}
	h.outroMods = []mods.Mod{
		&mods.JsonOutputMod[Output]{},
	}
}
