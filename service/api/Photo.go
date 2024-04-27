package api

import (
	"encoding/json"
	"net/http"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// prendo il nuovo username
	var user User

	// prendo l'username dal path
	user.Username = ps.ByName("username")

	userID, err := rt.db.CheckUserExists(user.Username)
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

	user.ID = userID

	// creo l'oggetto Photo
	var photo Photo

	photo.userID = userID

	// decodifico il body e metto i campi dentro user
	err = json.NewDecoder(r.Body).Decode(&photo)
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

	photo.Date = dateString

	photo.ID, err = rt.db.AddPhoto(photo.Date, photo.Text, photo.URL, photo.userID)
	if err != nil {
		http.Error(w, "Errore nel db", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}
