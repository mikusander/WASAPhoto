package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto ban
	var ban Ban

	// prendo l'username personale
	ban.PersonalUserID = ps.ByName("username")

	// verifico se l'username esiste
	userID, err := rt.db.CheckUserExists(ban.PersonalUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// effettuo l'autentificazione
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// prendo l'username dell'utente da bannare
	ban.BanUserID = ps.ByName("banid")

	// verifico se l'utente esiste
	userID, err = rt.db.CheckUserExists(ban.BanUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// verifico se l'utente sta cercando di bannare se stesso e ritorno errore in caso sia vero
	if ban.PersonalUserID == ban.BanUserID {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// verifico se l'utente sia già bannato
	isBan, err := rt.db.CheckBan(ban.PersonalUserID, ban.BanUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se è già bannato ritorno errore
	if isBan == true {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// aggiungo il ban al database
	err = rt.db.NewBan(ban.PersonalUserID, ban.BanUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifico l'utente che si è bannato era seguito
	isFollow, err := rt.db.CheckFollow(ban.PersonalUserID, ban.BanUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se era seguito rimuovo il follow
	if isFollow == true {
		err = rt.db.RemoveFollow(ban.PersonalUserID, ban.BanUserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ban)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto ban
	var ban Ban

	// prendo l'username personale
	ban.PersonalUserID = ps.ByName("username")

	// estraggo l'id dal database
	userID, err := rt.db.CheckUserExists(ban.PersonalUserID)
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

	// aggiungo all'oggetto ban l'username personale dell'utente
	ban.BanUserID = ps.ByName("banid")

	// verifico se l'utente al quale si vuole levare il ban esista
	userID, err = rt.db.CheckUserExists(ban.BanUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// se non esiste l'utente da sbannare ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// verifico se l'utente che si vuole sbannare si bannato
	isBan, err := rt.db.CheckBan(ban.PersonalUserID, ban.BanUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non è bannato ritorno l'errore
	if isBan == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// rimuovo dal database il ban
	err = rt.db.RemoveBan(ban.PersonalUserID, ban.BanUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
