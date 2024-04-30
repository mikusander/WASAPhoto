package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto di tipo follow
	var follow Follow

	// prendo l'username personale
	follow.PersonalUserId = ps.ByName("username")

	// estraggo l'id dal database
	userID, err := rt.db.CheckUserExists(follow.PersonalUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// verifico se l'utente esiste
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// eseguo il controllo di sicurezza
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// aggiungo all'oggetto follow l'username personale dell'utente
	follow.FollowUserId = ps.ByName("followid")

	// verifico se l'utente che vuole seguire esiste ed estraggo l'utente
	userID, err = rt.db.CheckUserExists(follow.FollowUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// se non esiste l'utente da seguire ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// verifico se segue già l'utente che vuole seguire
	isFollow, err := rt.db.CheckFollow(follow.PersonalUserId, follow.FollowUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if isFollow == true {
		http.Error(w, "Utente già seguito", http.StatusBadRequest)
		return
	}

	// aggiungo al database il nuovo follow
	err = rt.db.NewFollow(follow.PersonalUserId, follow.FollowUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(follow)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto follow
	var follow Follow

	// prendo l'username personale
	follow.PersonalUserId = ps.ByName("username")

	// estraggo l'id dal database
	userID, err := rt.db.CheckUserExists(follow.PersonalUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// verifico se l'utente esiste
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// eseguo il controllo di sicurezza
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// aggiungo all'oggetto follow l'username personale dell'utente
	follow.FollowUserId = ps.ByName("followid")

	// verifico se l'utente che vuole seguire esiste ed estraggo l'utente
	userID, err = rt.db.CheckUserExists(follow.FollowUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// se non esiste l'utente da seguire ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// verifico se segue già l'utente che vuole non seguire più
	isFollow, err := rt.db.CheckFollow(follow.PersonalUserId, follow.FollowUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if isFollow == false {
		http.Error(w, "Utente non seguito", http.StatusBadRequest)
		return
	}

	// rimuovo dal database il follow
	err = rt.db.RemoveFollow(follow.PersonalUserId, follow.FollowUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
