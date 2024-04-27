package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// creo l'oggetto comment
	var comment Comment

	// prendo l'username dal path
	username := ps.ByName("username")

	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if userID == 0 {
		http.Error(w, "L'utente non esiste", http.StatusBadRequest)
		return
	}
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, "Errore di autentificazione", http.StatusBadRequest)
		return
	}

	comment.user_id = userID

	// prendo l'id della photo dal path
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, "Errore nella conversione", http.StatusBadRequest)
		return
	}

	id, err := rt.db.CheckPhotoExists(photoid)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	if id == false {
		http.Error(w, "La photo non esiste", http.StatusBadRequest)
		return
	}

	comment.photo_id = photoid

	// decodifico il body e metto i campi dentro user
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Impossibile leggere il corpo della richiesta", http.StatusBadRequest)
		return
	}

	// Chiudere il corpo della richiesta dopo la lettura
	defer r.Body.Close()

	// Ottieni la data e l'ora attuali
	now := time.Now()

	// Converti la data in una stringa utilizzando un formato specifico
	dateString := now.Format("2006-01-02") // Formato: "YYYY-MM-DD"

	comment.Date = dateString

	comment.ID, err = rt.db.AddComment(comment.Date, comment.Text, comment.user_id, comment.photo_id)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// prendo il nuovo username
	var comment Comment

	// prendo l'username personale
	username := ps.ByName("username")

	// estraggo l'id dal database
	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}
	comment.user_id = userID
	// verifico se l'utente esiste
	if comment.user_id == 0 {
		http.Error(w, "L'utente non esiste", http.StatusBadRequest)
		return
	}
	// eseguo il controllo di sicurezza
	err = autentification(r.Header.Get("Authorization"), comment.user_id)
	if err != nil {
		http.Error(w, "Errore di autentificazione", http.StatusBadRequest)
		return
	}

	// aggiungo all'oggetto follow l'username personale dell'utente
	comment.photo_id, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, "Errore nella conversione", http.StatusBadRequest)
		return
	}

	// verifico se la foto che vuole seguire esiste
	exist, err := rt.db.CheckPhotoExists(comment.photo_id)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	// se non esiste la photo ritorno errore
	if exist == false {
		http.Error(w, "La photo non esiste", http.StatusBadRequest)
		return
	}

	// aggiungo all'oggetto follow l'username personale dell'utente
	comment.ID, err = strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	if err != nil {
		http.Error(w, "Errore nella conversione", http.StatusBadRequest)
		return
	}

	// verifico se la foto che vuole seguire esiste
	exist, err = rt.db.CheckCommentExists(comment.ID)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	// se non esiste la photo ritorno errore
	if exist == false {
		http.Error(w, "Il commento non esiste non esiste", http.StatusBadRequest)
		return
	}

	// cancello la photo
	err = rt.db.DeleteComment(comment.ID)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("Delete comment success")
}
