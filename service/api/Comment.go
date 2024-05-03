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

	// verifico se l'utente esiste
	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se l'utente non esiste ritorno errore
	if userID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// effettuo il controllo sull'autentificazione
	err = autentification(r.Header.Get("Authorization"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// assegno l'user id all'oggetto comment
	comment.UserID = userID

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

	// assegno l'id della photo all'oggetto comment
	comment.PhotoID = photoid

	// decodifico il body e metto i campi dentro comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Chiudere il corpo della richiesta dopo la lettura
	defer r.Body.Close()

	// Ottieni la data e l'ora attuali
	now := time.Now()

	// Converti la data in una stringa utilizzando un formato specifico
	dateString := now.Format("2006-01-02") // Formato: "YYYY-MM-DD"

	// assegno la data all'oggetto comment
	comment.Date = dateString

	// aggiungo al database il nuovo commento
	comment.ID, err = rt.db.AddComment(comment.Date, comment.Text, comment.UserID, comment.PhotoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto comment
	var comment Comment

	// prendo l'username personale
	username := ps.ByName("username")

	// estraggo l'id dal database
	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// assegno l'id dell'utebte all'oggetto comment
	comment.UserID = userID

	// verifico se l'utente esiste
	if comment.UserID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// eseguo il controllo di sicurezza
	err = autentification(r.Header.Get("Authorization"), comment.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// prendo dal path l'id della photo
	comment.PhotoID, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifico se la photo che vuole commentare esiste
	exist, err := rt.db.CheckPhotoExists(comment.PhotoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// se non esiste la photo ritorno errore
	if exist == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// prendo l'id del commento
	comment.ID, err = strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifico se il commento esiste
	exist, err = rt.db.CheckCommentExists(comment.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// se non esiste ritorno errore
	if exist == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// cancello il commento
	err = rt.db.DeleteComment(comment.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
