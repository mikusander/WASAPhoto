package api

import(
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	
	//creo l'oggetto user
	var user User

	//decodifico il body e metto i campi dentro user
	err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Impossibile leggere il corpo della richiesta", http.StatusBadRequest)
        return
    }

    // Chiudere il corpo della richiesta dopo la lettura
    defer r.Body.Close()

	id, error := rt.db.CheckUserExists(user.Username)
	if(error != nil){
		http.Error(w, "Errore nel db", http.StatusBadRequest)
	}

	if (id == ""){
		_, error = rt.db.CreateUser(user.Username)
		if(error != nil){
			http.Error(w,  "Errore nella creazioone dell'utente", http.StatusBadRequest)
		}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}