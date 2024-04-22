package database

import (
)

func (db *appdbimpl) CheckBan(PersonaleUserId string, BanUserId string) (bool, error) {
	// Eseguire una query per verificare se l'username ha già bannato l'altro utente
	query := `SELECT COUNT(*) FROM Ban WHERE personal_user_id = ? and ban_user_id = ?`
	var isBan int
	err := db.c.QueryRow(query, PersonaleUserId, BanUserId).Scan(&isBan)
	if err != nil {
		// Si è verificato un errore durante l'esecuzione della query
		return false, err
	}
	if isBan == 0{
		// L'utente può essere bannato
		return false, nil
	}

	// L'utente ha già bannato quell'utente
	return true, nil
}

func (db *appdbimpl) NewBan(PersonalUserId string, BanUserId string) error {
	// Esegui una query per inserire il nuovo utente nella tabella degli utenti
	query := `INSERT INTO Ban (personal_user_id, ban_user_id) VALUES (?, ?)`
	_, err := db.c.Exec(query, PersonalUserId, BanUserId)
	if err != nil {
		// Si è verificato un errore durante l'inserimento del Ban
		return err
	}

	return nil
}

func (db *appdbimpl) RemoveBan(PersonalUserId string, BanUserId string) error {
	// Esegui la query DELETE per rimuovere l'utente
    query := "DELETE FROM Ban WHERE personal_user_id = ? and ban_user_id = ?"

    // Esegui la query
    _, err := db.c.Exec(query, PersonalUserId, BanUserId)
    if err != nil {
        return err
    }
	return nil
}