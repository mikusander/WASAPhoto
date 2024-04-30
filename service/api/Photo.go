package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto user
	var user User

	// prendo l'username dal path
	user.Username = ps.ByName("username")

	// verifico se l'uuser esiste
	userID, err := rt.db.CheckUserExists(user.Username)
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

	// assegno all'oggetto user il suo id
	user.ID = userID

	// creo l'oggetto Photo
	var photo Photo

	photo.UserID = userID

	// decodifico il body e metto i campi dentro photo
	err = json.NewDecoder(r.Body).Decode(&photo)
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

	photo.Date = dateString

	photo.ID, err = rt.db.AddPhoto(photo.Date, photo.Text, photo.URL, photo.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.LikeCounter = 0
	photo.CommentCounter = 0

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// creo l'oggetto photo
	var photo Photo

	// prendo l'username personale
	username := ps.ByName("username")

	// estraggo l'id dal database
	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.UserID = userID
	// verifico se l'utente esiste
	if photo.UserID == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// eseguo il controllo di sicurezza
	err = autentification(r.Header.Get("Authorization"), photo.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// estraggo dal path l'id della photo
	photo.ID, err = strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifico se la photo esiste
	exist, err := rt.db.CheckPhotoExists(photo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// se non esiste la photo ritorno errore
	if exist == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// cancello la photo
	err = rt.db.DeletePhoto(photo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
