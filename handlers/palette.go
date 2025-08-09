package handlers

import (
	"pixel-battle-backend/api"
	"pixel-battle-backend/constants"
)

func GetPaletteData() []api.Color {
	return constants.ColorsPalette
}
