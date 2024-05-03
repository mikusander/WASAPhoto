package database

import ()

func (db *appdbimpl) AddLike(userid uint64, photoid uint64) error {
	// Esegui una query per inserire il nuovo like nella tabella dei like
	query := `INSERT INTO like (user_id, photo_id) VALUES (?, ?)`
	_, err := db.c.Exec(query, userid, photoid)
	if err != nil {
		// Si è verificato un errore durante l'inserimento del like
		return err
	}

	return nil
}

func (db *appdbimpl) CheckLikeExists(userid uint64, photoid uint64) (bool, error) {
	// Eseguire una query per verificare se il like esiste
	query := `SELECT COUNT(*) FROM Like WHERE user_id = ? and photo_id = ?`
	var exist int
	err := db.c.QueryRow(query, userid, photoid).Scan(&exist)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return false, err
	}
	if exist == 0 {
		return false, nil
	}

	// il like esiste
	return true, nil
}

func (db *appdbimpl) DeleteLike(userid uint64, photoid uint64) error {
	// Esegui la query DELETE per rimuovere il like
	query := "DELETE FROM Like WHERE user_id = ? and photo_id = ?"

	// Esegui la query
	_, err := db.c.Exec(query, userid, photoid)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) LikeCounterPhoto(photoid uint64) (uint64, error) {
	// Eseguire una query per verificare contare quanti like ha una photo
	query := `SELECT COUNT(*) FROM Like WHERE photo_id = ?`
	var exist int
	err := db.c.QueryRow(query, photoid).Scan(&exist)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return 0, err
	}
	if exist == 0 {
		return 0, nil
	}

	return uint64(exist), nil
}
