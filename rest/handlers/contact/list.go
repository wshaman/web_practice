package contact

import (
	"net/http"

	"github.com/wshaman/web_practice/rest/lib/ds"
	"github.com/wshaman/web_practice/rest/lib/response"
)

func listAll(w http.ResponseWriter) {
	d := ds.GetYP()
	c, err := d.List()
	if err != nil {
		response.ThrowError(w, 500, err.Error())
	}
	response.WriteJSON(w, c)
}

func listFiltered(w http.ResponseWriter, phonePart string) {
	d := ds.GetYP()
	c, err := d.FindByPhone(phonePart)
	if err != nil {
		response.ThrowError(w, 500, err.Error())
	}
	response.WriteJSON(w, c)
}

func HandleList(w http.ResponseWriter, r *http.Request) {
	phoneFilter, ok := r.URL.Query()["phone"]
	if !ok {
		listAll(w)
		return
	}
	listFiltered(w, phoneFilter[0])
}
