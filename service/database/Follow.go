package database

func (db *appdbimpl) CheckFollow(PersonaleUserID string, FollowUserID string) (bool, error) {
	// Eseguire una query per verificare se l'utente segue l'altro utente
	query := `SELECT COUNT(*) FROM Follow WHERE personal_user_id = ? and follow_user_id = ?`
	var isFollow int
	err := db.c.QueryRow(query, PersonaleUserID, FollowUserID).Scan(&isFollow)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return false, err
	}
	if isFollow == 0 {
		return false, nil
	}

	// L'utente segue già quell'utente
	return true, nil
}

func (db *appdbimpl) NewFollow(PersonalUserID string, FollowUserID string) error {
	// Esegui una query per inserire il nuovo follow nella tabella dei follow
	query := `INSERT INTO Follow (personal_user_id, follow_user_id) VALUES (?, ?)`
	_, err := db.c.Exec(query, PersonalUserID, FollowUserID)
	if err != nil {
		// Si è verificato un errore durante l'inserimento del follow
		return err
	}

	return nil
}

func (db *appdbimpl) RemoveFollow(PersonalUserID string, FollowUserID string) error {
	// Esegui la query DELETE per rimuovere il follow
	query := "DELETE FROM Follow WHERE personal_user_id = ? and follow_user_id = ?"

	// Esegui la query
	_, err := db.c.Exec(query, PersonalUserID, FollowUserID)
	if err != nil {
		return err
	}
	return nil
}
