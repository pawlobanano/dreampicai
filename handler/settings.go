package handler

import (
	"dreampicai/types"
	"dreampicai/view/settings"
	"net/http"
)

func HandleSettingsIndex(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	return render(w, r, settings.Index(user))
}
