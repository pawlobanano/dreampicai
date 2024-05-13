package handler

import (
	"net/http"

	"dreampicai/types"
	"dreampicai/view/home"
)

func HandleHomeIndex(log types.Logger, w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
