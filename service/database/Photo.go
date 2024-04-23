package database

import ()

func (db *appdbimpl) AddPhoto(Date string, Text string, URL []byte, userID uint64) (uint64, error) {
	// Esegui una query per inserire la nuova photo nella tabella delle photo
	query := `INSERT INTO Photo (Date, Text, image, user_id) VALUES (?, ?, ?, ?)`
	result, err := db.c.Exec(query, Date, Text, URL, userID)
	if err != nil {
		// Si è verificato un errore durante l'inserimento dell'utente
		return 0, err
	}

	// Ottieni l'ID generato per l'utente appena inserito
	idInt, err := result.LastInsertId()
	if err != nil {
		// Si è verificato un errore durante il recupero dell'ID generato
		return 0, err
	}
	id := uint64(idInt)

	// Restituisci l'ID dell'utente appena inserito
	return id, nil
}
