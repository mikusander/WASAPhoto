package api

import (
	"errors"
	"strconv"
	"strings"
)

func autentification(bearer string, id uint64) error {
	token := strings.TrimPrefix(bearer, "Bearer ")
	idString := strconv.FormatUint(id, 10)
	if token == idString {
		return nil
	}
	return errors.New("autenticazione fallita")
}
