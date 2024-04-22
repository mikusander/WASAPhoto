package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// prendo il nuovo username
	var ban Ban

	// prendo l'username personale
	ban.PersonalUserId = ps.ByName("username")

	userID, err := rt.db.CheckUserExists(ban.PersonalUserId)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if userID == 0{
		http.Error(w, "L'utente non esiste", http.StatusBadRequest)
		return
	}
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil{
		http.Error(w, "Errore di autentificazione", http.StatusBadRequest)
		return
	}

	ban.BanUserId = ps.ByName("banid")

	userID, err = rt.db.CheckUserExists(ban.BanUserId)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if userID == 0{
		http.Error(w, "L'utente che si vuole seguire non esiste", http.StatusBadRequest)
		return
	}

	isBan, err := rt.db.CheckBan(ban.PersonalUserId, ban.BanUserId)
	if isBan == true{
		http.Error(w, "Utente già bannato", http.StatusBadRequest)
		return
	}
	if err != nil{
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	

	err = rt.db.NewBan(ban.PersonalUserId, ban.BanUserId)
	if err != nil{
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	// verifico se segue già l'utente che vuole non seguire più
	isFollow, err := rt.db.CheckFollow(ban.PersonalUserId, ban.BanUserId)
	if err != nil{
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if isFollow == true{
		err = rt.db.RemoveFollow(ban.PersonalUserId, ban.BanUserId)
		if err != nil{
			http.Error(w, "Errore nel db", http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ban)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// prendo il nuovo username
	var ban Ban

	// prendo l'username personale
	ban.PersonalUserId = ps.ByName("username")

	// estraggo l'id dal database
	userID, err := rt.db.CheckUserExists(ban.PersonalUserId)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	// verifico se l'utente esiste
	if userID == 0{
		http.Error(w, "L'utente non esiste", http.StatusBadRequest)
		return
	}
	// eseguo il controllo di sicurezza
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil{
		http.Error(w, "Errore di autentificazione", http.StatusBadRequest)
		return
	}

	// aggiungo all'oggetto ban l'username personale dell'utente
	ban.BanUserId = ps.ByName("banid")

	// verifico se l'utente che vuole seguire esiste ed estraggo l'utente
	userID, err = rt.db.CheckUserExists(ban.BanUserId)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	// se non esiste l'utente da seguire ritorno errore
	if userID == 0{
		http.Error(w, "L'utente che si vuole seguire non esiste", http.StatusBadRequest)
		return
	}

	// verifico se segue già l'utente che vuole non seguire più
	isBan, err := rt.db.CheckBan(ban.PersonalUserId, ban.BanUserId)
	if err != nil{
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if isBan == false{
		http.Error(w, "L'utente non era bannato", http.StatusBadRequest)
		return
	}

	// rimuovo dal database il ban
	err = rt.db.RemoveBan(ban.PersonalUserId, ban.BanUserId)
	if err != nil{
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("unban user success")
}