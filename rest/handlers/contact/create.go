package contact

import (
	"encoding/json"
	"net/http"

	"github.com/wshaman/stub_contacts"
	"github.com/wshaman/web_practice/rest/lib/ds"
	"github.com/wshaman/web_practice/rest/lib/response"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	d := ds.GetYP()
	c := stub_contacts.Contact{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		response.ThrowError(w, 404, "wrong body sent")
		return
	}
	c.ID = 0
	uid, err := d.Save(c)
	if err != nil {
		response.ThrowError(w, 500, err.Error())
		return
	}
	response.WriteJSON(w, uid)
}
