package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto user
	var user User

	// decodifico il body e metto i campi dentro user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiudere il corpo della richiesta dopo la lettura
	defer r.Body.Close()

	id, error := rt.db.CheckUserExists(user.Username)
	if error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// se l'utente non esiste lo creo
	if id == 0 {
		id, error = rt.db.CreateUser(user.Username)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}
		user.ID = id
	} else {
		user.ID = id
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
