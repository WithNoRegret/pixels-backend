package handlers

import (
	"encoding/json"
	"net/http"

	"pixel-battle-backend/constants"
)

// PaletteHandler godoc
// @Summary Get color palette
// @Description Returns list of available colors for the game
// @Tags palette
// @Produce json
// @Success 200 {array} models.Color
// @Router /palette [get]
func PaletteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(constants.ColorsPalette); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
