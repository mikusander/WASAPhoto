package database

import (
	"database/sql"
)

func (db *appdbimpl) SetNewUsername(userID uint64, newUsername string) error {
	// Prepara la query di aggiornamento
	query := "UPDATE User SET username = ? WHERE id = ?"

	// Esegui l'istruzione di aggiornamento
	_, err := db.c.Exec(query, newUsername, userID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) CountFollow(username string) (uint64, error) {
	// Eseguire una query per verificare se l'username segue già quell'utente
	query := `SELECT COUNT(*) FROM Follow WHERE personal_user_id = ?`
	var numFollow uint64
	err := db.c.QueryRow(query, username).Scan(&numFollow)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return 0, err
	}
	return numFollow, err
}

func (db *appdbimpl) CountFollowing(username string) (uint64, error) {
	// Eseguire una query per verificare se l'username segue già quell'utente
	query := `SELECT COUNT(*) FROM Follow WHERE follow_user_id = ?`
	var numFollowing uint64
	err := db.c.QueryRow(query, username).Scan(&numFollowing)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return 0, err
	}
	return numFollowing, err
}

func (db *appdbimpl) GetFollowPerson(userID string) ([]string, error) {
	// Eseguire una query per verificare se l'username segue già quell'utente
	query := `SELECT follow_user_id FROM Follow WHERE personal_user_id = ?`
	var follow []string
	rows, err := db.c.Query(query, userID)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return nil, err
	}
	for rows.Next() {
		var x string
		err := rows.Scan(&x)
		if err != nil {
			return nil, err
		}

		_, err = db.CheckUserExists(x)
		if err != nil {
			return nil, err
		}

		follow = append(follow, x)
	}

	// Controlla se ci sono errori dopo l'iterazione dei risultati
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return follow, nil
}

func (db *appdbimpl) GetUsernameFromID(id uint64) (string, error) {
	// Eseguire una query per verificare se l'username esiste già nel database
	query := `SELECT Username FROM User WHERE Id = ?`
	var username string
	err := db.c.QueryRow(query, id).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			// L'utente non esiste nel database
			return "", nil
		}
		// Si è verificato un errore durante l'esecuzione della query
		return "", err
	}

	// L'utente esiste nel database
	return username, nil
}

func (db *appdbimpl) GetListPhoto(userID uint64) ([]Photo, error) {
	// Eseguire una query per verificare se l'username segue già quell'utente
	query := `SELECT * FROM Photo WHERE user_id = ? ORDER BY date DESC`
	rows, err := db.c.Query(query, userID)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return nil, err
	}

	defer rows.Close()

	var listPhoto []Photo
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
