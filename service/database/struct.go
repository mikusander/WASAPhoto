package database

import (
/*
"encoding/json"
"fmt"
"net/http"
"strconv"
*/
)

// User rappresenta un utente nel sistema.
type User struct {
	ID       uint64 `json:"ID"`
	Username string `json:"username"`
}

// Photo rappresenta una foto del sistema
type Photo struct {
	ID             uint64 `json:"ID"`
	Date           string `json:"Date"`
	Text           string `json:"Text"`
	URL            []byte `json:"URL"`
	LikeCounter    uint64 `json:"likeCounter"`
	CommentCounter uint64 `json:"commentCounter"`
	UserID         uint64 `json:"userID"`
}
