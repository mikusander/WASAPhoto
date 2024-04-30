package api

import (
	"errors"
	"strconv"
	"strings"
)

func autentification(bearer string, id uint64) error {
	// prendo il messaggio ed elimino dal messaggio la parola bearer
	token := strings.TrimPrefix(bearer, "Bearer ")
	// lo conventro in intero
	idString := strconv.FormatUint(id, 10)
	// verifico se l'utente è autorizzato
	if token == idString {
		return nil
	}
	return errors.New("autenticazione fallita")
}
