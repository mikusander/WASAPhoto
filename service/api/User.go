package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo il nuovo username
	var user User

	// decodifico il body e metto i campi dentro user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiudere il corpo della richiesta dopo la lettura
	defer r.Body.Close()

	// prendo il vecchio username
	oldUsername := ps.ByName("username")

	// verifico se il vecchio username esiste
	userID, err := rt.db.CheckUserExists(oldUsername)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se già esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// controllo se l'utente ha i permessi
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// assegno all'oggetto user il suo id
	user.ID = userID

	// verifico se il nuovo username è già usato
	id, err := rt.db.CheckUserExists(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se è già in uso ritorno errore
	if id != 0 {
		http.Error(w, "Username già in uso", http.StatusBadRequest)
		return
	}

	// cambio dell'username
	err = rt.db.SetNewUsername(userID, user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto myProfile
	var myProfile MyProfile

	//creo l'oggetto User
	var user User

	// prendo l'username dal path
	username := ps.ByName("username")

	// verifico se l'user esiste
	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// controllo se l'utente ha i permessi
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// assegno l'id e l'username all'oggetto utente
	user.ID = userID
	user.Username = username

	// assegno all'oggetto my profile l'oggetto user
	myProfile.UserOwner = user

	// conto il nuomero di persone che seguono l'utente
	myProfile.NumFollow, err = rt.db.CountFollow(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// conto il numero di persone che segue l'utente
	myProfile.NumFollowing, err = rt.db.CountFollowing(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// conto il numero di photo che ha l'utente
	myProfile.NumPhoto, err = rt.db.CountPhoto(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// estraggo dal database tutte le photo dell'utente
	myProfile.ListPhoto, err = rt.db.GetListPhoto(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(myProfile)

}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto my stream
	var myStream MyStream

	//creo l'oggetto User
	var user User

	// prendo l'username dal path
	username := ps.ByName("username")

	// verifico se l'utente esiste
	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// controllo se l'utente ha i permessi
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// assegno all'oggetto user id e username
	user.ID = userID
	user.Username = username

	// assegno all'oggetto my stream l'user
	myStream.UserOwner = user

	// estraggo dal db tutte le persone che l'utente segue
	listFollow, err := rt.db.GetFollowPerson(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// estraggo dal database la lista delle photo
	myStream.ListPhoto, err = rt.db.GetAllPhoto(listFollow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(myStream)

}
