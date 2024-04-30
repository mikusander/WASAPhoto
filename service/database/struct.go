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
	ID             uint64 `json:"ID"`
	Date           string `json:"Date"`
	Text           string `json:"Text"`
	URL            []byte `json:"URL"`
	LikeCounter    uint64 `json:"likeCounter"`
	CommentCounter uint64 `json:"commentCounter"`
	UserID         uint64 `json:"userID"`
}
