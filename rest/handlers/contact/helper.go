package contact

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func idFromVars(r *http.Request) (uint, error) {
	idString := mux.Vars(r)["id"]
	i, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
