package database

import (
/*
"encoding/json"
"fmt"
"net/http"
"strconv"
*/
)

type User struct {
	ID       uint64 `json:"ID"`
	Username string `json:"username"`
}

type Photo struct {
	ID             string `json:"ID"`
	Date           string `json:"Date"`
	Text           string `json:"Text"`
	URL            []byte `json:"URL"`
	likeCounter    uint64 `json:"likeCounter"`
	commentCounter uint64 `json:"commentCounter"`
	userID         uint64 `json:"userID"`
}
