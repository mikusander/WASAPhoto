package database

func (db *appdbimpl) AddComment(Date string, Text string, userID uint64, photoID uint64) (uint64, error) {
	// Esegui una query per inserire il nuovo commento nella tabella dei commenti
	query := `INSERT INTO Comment (Date, Text, user_id, photo_id) VALUES (?, ?, ?, ?)`
	result, err := db.c.Exec(query, Date, Text, userID, photoID)
	if err != nil {
		// Si è verificato un errore durante l'inserimento della photo
		return 0, err
	}

	// Ottieni l'ID generato per la photo appena inserito
	idInt, err := result.LastInsertId()
	if err != nil {
		// Si è verificato un errore durante il recupero dell'ID generato
		return 0, err
	}
	id := uint64(idInt)

	// Restituisci l'ID dell'utente appena inserito
	return id, nil
}

func (db *appdbimpl) CheckCommentExists(commentid uint64) (bool, error) {
	// Eseguire una query per verificare se il commento esiste
	query := `SELECT COUNT(*) FROM Comment WHERE Id = ?`
	var exist int
	err := db.c.QueryRow(query, commentid).Scan(&exist)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return false, err
	}
	if exist == 0 {
		return false, nil
	}

	// La photo esiste
	return true, nil
}

func (db *appdbimpl) DeleteComment(id uint64) error {
	// Esegui la query DELETE per rimuovere il commento
	query := "DELETE FROM Comment WHERE Id = ?"

	// Esegui la query
	_, err := db.c.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) CommentCounterPhoto(photoid uint64) (uint64, error) {
	// Eseguire una query per contare il numero di commenti ha una photo
	query := `SELECT COUNT(*) FROM Comment WHERE photo_id = ?`
	var exist int
	err := db.c.QueryRow(query, photoid).Scan(&exist)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return 0, err
	}
	if exist == 0 {
		return 0, nil
	}

	// La photo esiste
	return uint64(exist), nil
}
