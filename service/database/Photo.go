package database

func (db *appdbimpl) AddPhoto(Date string, Text string, URL []byte, userID uint64, username string) (uint64, error) {
	// Esegui una query per inserire la nuova photo nella tabella delle photo
	query := `INSERT INTO Photo (Date, Text, image, user_id, user_username) VALUES (?, ?, ?, ?, ?)`
	result, err := db.c.Exec(query, Date, Text, URL, userID, username)
	if err != nil {
		// Si è verificato un errore durante l'inserimento dell'utente
		return 0, err
	}

	// Ottieni l'ID generato per la photo appena inserita
	idInt, err := result.LastInsertId()
	if err != nil {
		// Si è verificato un errore durante il recupero dell'ID generato
		return 0, err
	}
	id := uint64(idInt)

	// Restituisci l'ID della photo appena inserito
	return id, nil
}

func (db *appdbimpl) CheckPhotoExists(photoid uint64) (bool, error) {
	// Eseguire una query per verificare se la photo esiste
	query := `SELECT COUNT(*) FROM Photo WHERE Id = ?`
	var exist int
	err := db.c.QueryRow(query, photoid).Scan(&exist)
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

func (db *appdbimpl) DeletePhoto(id uint64) error {
	// Esegui la query DELETE per rimuovere la photo
	query := "DELETE FROM Photo WHERE Id = ?"

	// Esegui la query
	_, err := db.c.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) CountPhoto(userID uint64) (uint64, error) {
	// Eseguire una query per contare il numero di photo che ha l'utente
	query := `SELECT COUNT(*) FROM Photo WHERE user_id = ?`
	var exist int
	err := db.c.QueryRow(query, userID).Scan(&exist)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return 0, err
	}

	// La photo esiste
	return uint64(exist), nil
}

func (db *appdbimpl) GetAllPhoto(username string) ([]Photo, error) {

	var listPhoto []Photo

	// Eseguire una query per restituire tutte le photo
	query := `SELECT Photo.Id, Photo.Date, Photo.Text, Photo.Image, Photo.user_id, Photo.user_username FROM Photo, User WHERE user_id = User.Id AND Username IN 
	(SELECT follow_user_id FROM Follow WHERE personal_user_id=? AND follow_user_id NOT IN 
	(SELECT personal_user_id FROM Ban WHERE ban_user_id=?)) ORDER BY date DESC`
	rows, err := db.c.Query(query, username, username)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var photo Photo
		err := rows.Scan(&photo.ID, &photo.Date, &photo.Text, &photo.URL, &photo.UserID, &photo.UserUsername)
		if err != nil {
			return nil, err
		}

		photo.LikeCounter, err = db.LikeCounterPhoto(photo.ID)
		if err != nil {
			return nil, err
		}

		photo.CommentCounter, err = db.CommentCounterPhoto(photo.ID)
		if err != nil {
			return nil, err
		}

		photo.ListComment, err = db.GetCommentList(photo.ID)
		if err != nil {
			return nil, err
		}

		listPhoto = append(listPhoto, photo)
	}

	// Controlla se ci sono errori dopo l'iterazione dei risultati
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listPhoto, nil

}
