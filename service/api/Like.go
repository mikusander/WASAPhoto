package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto like
	var like Like

	// prendo l'username dal path
	username := ps.ByName("username")

	// verifico se esiste l'utente
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

	// assegno id e username all'oggetto like
	like.Ownerphoto = userID
	like.Usernameowner = username

	// prendo l'id della photo dal path
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifico se la photo esiste
	id, err := rt.db.CheckPhotoExists(photoid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if id == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// assegno l'id della photo all'oggetto like
	like.PhotoID = photoid

	// prendo l'username dal path
	username = ps.ByName("likeid")

	// verifico se l'utente esiste
	userID, err = rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// vedo se l'utente è autorizzato
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// assegno l'id e l'username all'oggetto like
	like.UserID = userID
	like.Username = username

	// verifico se il lije che si vuole aggiungere già esiste
	exist, err := rt.db.CheckLikeExists(like.UserID, like.PhotoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se già esiste ritorno errore
	if exist == true {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// aggiungo il like al database
	err = rt.db.AddLike(like.UserID, like.PhotoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(like)
}

func (rt *_router) isLikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto like
	var like Like

	// prendo l'username dal path
	username := ps.ByName("username")

	// verifico se esiste l'utente
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

	// assegno id e username all'oggetto like
	like.Ownerphoto = userID
	like.Usernameowner = username

	// prendo l'id della photo dal path
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifico se la photo esiste
	id, err := rt.db.CheckPhotoExists(photoid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if id == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// assegno l'id della photo all'oggetto like
	like.PhotoID = photoid

	// prendo l'username dal path
	username = ps.ByName("likeid")

	// verifico se l'utente esiste
	userID, err = rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// vedo se l'utente è autorizzato
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// assegno l'id e l'username all'oggetto like
	like.UserID = userID
	like.Username = username

	// verifico se il lije che si vuole aggiungere già esiste
	exist, err := rt.db.CheckLikeExists(like.UserID, like.PhotoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se già esiste ritorno errore
	if exist == true {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(like)
	} else {
		like.UserID = 0
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(like)
	}

}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto like
	var like Like

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
	// controllo se l'utente ha i permessi per effettuare quell'azione
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// assegno all'oggetto like l'user id dell'utente
	like.UserID = userID

	// prendo l'id della photo dal path
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifico se la photo esiste
	id, err := rt.db.CheckPhotoExists(photoid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non eiste ritorno errore
	if id == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// assegno l'id della photo all'oggetto like
	like.PhotoID = photoid

	// prendo l'username dal path
	username = ps.ByName("likeid")

	// verifico se l'utente esiste
	userID, err = rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// assegno l'id dell'utente all'oggetto like
	like.Ownerphoto = userID

	// verifico se il like che si vuole levare esiste
	exist, err := rt.db.CheckLikeExists(like.UserID, like.PhotoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if exist == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// rimuovo il like
	err = rt.db.DeleteLike(like.UserID, like.PhotoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
