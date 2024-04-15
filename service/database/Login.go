package database

import(
    "database/sql"
)

func (db *appdbimpl) CheckUserExists(username string) (uint64, error){
	// Eseguire una query per verificare se l'username esiste già nel database
    query := `SELECT Id FROM User WHERE Username = ?`
    var id uint64
    err := db.c.QueryRow(query, username).Scan(&id)
    if err != nil {
        if err == sql.ErrNoRows {
            // L'utente non esiste nel database
            return 0, nil
        }
        // Si è verificato un errore durante l'esecuzione della query
        return 0, err
    }

    // L'utente esiste nel database
    return id, nil
}

func (db *appdbimpl) CreateUser(username string) (uint64, error){
    // Esegui una query per inserire il nuovo utente nella tabella degli utenti
    query := `INSERT INTO User (username) VALUES (?)`
    result, err := db.c.Exec(query, username)
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