package database

import(
)

func (db *appdbimpl) SetNewUsername(userID uint64, newUsername string) error{
	// Prepara la query di aggiornamento
	query := "UPDATE User SET username = ? WHERE id = ?"

	// Esegui l'istruzione di aggiornamento
	_, err := db.c.Exec(query, newUsername, userID)
	if err != nil {
		return err
	}
	return nil
}