package prod

import (
	"net/http"

	"github.com/yeddaTech/TaskManager/templates"
)

// MyWebApp è la funzione che renderizza la tua dashboard
func MyWebApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	// Qui puoi passare i dati della tua dashboard (i tuoi server)
	templates.Index().Render(r.Context(), w)
}
