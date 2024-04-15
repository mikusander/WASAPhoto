package api


import(
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	
	//prendo il nuovo username
	var user User

	//decodifico il body e metto i campi dentro user
	err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Impossibile leggere il corpo della richiesta", http.StatusBadRequest)
        return
    }

    // Chiudere il corpo della richiesta dopo la lettura
    defer r.Body.Close()

	//prendo il vecchio username
	oldUsername := ps.ByName("user_name")
	
	userID, err := rt.db.CheckUserExists(oldUsername)
	if(err != nil){
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if(userID == 0){
		http.Error(w, "L'utente non esiste", http.StatusBadRequest)
		return
	}

	user.ID = userID

	id, err := rt.db.CheckUserExists(user.Username)
	if(err != nil){
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if(id != 0){
		http.Error(w, "Username già in uso", http.StatusBadRequest)
		return
	}

	//cambio dell'username
	err = rt.db.SetNewUsername(userID, user.Username)
	if err != nil{
		http.Error(w, "errore nel cambio della password", http.StatusBadRequest)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}