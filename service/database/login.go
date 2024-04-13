package database

import(
    "database/sql"
)

func (db *appdbimpl) CheckUserExists(username string) (string, error){
	// Eseguire una query per verificare se l'username esiste già nel database
    query := `SELECT Id FROM User WHERE Username = ?`
    var id string
    err := db.c.QueryRow(query, username).Scan(&id)
    if err != nil {
        if err == sql.ErrNoRows {
            // L'utente non esiste nel database
            return "", nil
        }
        // Si è verificato un errore durante l'esecuzione della query
        return "", err
    }

    // L'utente esiste nel database
    return id, nil
}

func (db *appdbimpl) CreateUser(username string) (string, error){
    // Esegui una query per inserire il nuovo utente nella tabella degli utenti
    query := `INSERT INTO User (username) VALUES (?)`
    result, err := db.c.Exec(query, username)
    if err != nil {
        // Si è verificato un errore durante l'inserimento dell'utente
        return "", err
    }

    // Ottieni l'ID generato per l'utente appena inserito
    _, err = result.LastInsertId()
    if err != nil {
        // Si è verificato un errore durante il recupero dell'ID generato
        return "", err
    }

    // Restituisci l'ID dell'utente appena inserito
    return "", nil
}