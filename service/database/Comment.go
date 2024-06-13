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

func (db *appdbimpl) GetCommentList(photoid uint64) ([]Comment, error) {
	// Eseguire una query per ottenere tutti i commenti associati alla foto
	query := `SELECT * FROM Comment WHERE photo_id = ?`

	rows, err := db.c.Query(query, photoid)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return nil, err
	}
	defer rows.Close()

	// Creare una lista per contenere i commenti
	var comments []Comment

	// Iterare sui risultati della query
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.Text, &comment.Date, &comment.UserID, &comment.PhotoID)
		if err != nil {
			// Si è verificato un errore durante la scansione dei risultati
			return nil, err
		}

		comment.UserUsername, err = db.GetUsernameFromID(comment.UserID)
		if err != nil && comment.UserUsername == "" {
			// Si è verificato un errore durante la scansione dei risultati
			return nil, err
		}

		// Aggiungere il commento alla lista
		comments = append(comments, comment)
	}

	// Controllare se si è verificato un errore durante l'iterazione
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Restituire la lista dei commenti
	return comments, nil
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
