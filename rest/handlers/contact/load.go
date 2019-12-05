package contact

import (
	"net/http"

	"github.com/wshaman/web_practice/rest/lib/ds"
	"github.com/wshaman/web_practice/rest/lib/response"
)

func HandleLoad(w http.ResponseWriter, r *http.Request) {
	id, err := idFromVars(r)
	if err != nil {
		response.ThrowError(w, 400, "no correct id given")
		return
	}
	d := ds.GetYP()
	c, err := d.Load(id)
	if err != nil {
		response.ThrowError(w, 404, "not found")
		return
	}
	response.WriteJSON(w, c)
}
