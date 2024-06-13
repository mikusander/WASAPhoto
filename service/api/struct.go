package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// User rappresenta un utente del sistema
type User struct {
	ID       uint64 `json:"ID"`
	Username string `json:"Username"`
}

// Photo rappresenta una foto del sistema
type Photo struct {
	ID             uint64 `json:"ID"`
	Date           string `json:"Date"`
	Text           string `json:"Text"`
	URL            []byte `json:"URL"`
	LikeCounter    uint64 `json:"LikeCounter"`
	CommentCounter uint64 `json:"CommentCounter"`
	UserID         uint64 `json:"UserID"`
	ListComment	   []database.Comment `json:"ListComment"`
}

// Follow rappresenta un follow del sistema
type Follow struct {
	PersonalUserID string `json:"PersonalUserId"`
	FollowUserID   string `json:"FollowUserId"`
}

// Ban rappresenta un ban del sistema
type Ban struct {
	PersonalUserID string `json:"PersonalUserId"`
	BanUserID      string `json:"BanUserId"`
}

// Comment rappresenta un commento del sistema
type Comment struct {
	ID      uint64 `json:"ID"`
	Date    string `json:"Date"`
	Text    string `json:"Text"`
	UserID  uint64 `json:"User_id"`
	PhotoID uint64 `json:"Photo_id"`
}

// Like rappresenta un like del sistema
type Like struct {
	UserID        uint64 `json:"User_id"`
	Username      string `json:"Username"`
	PhotoID       uint64 `json:"Photo_id"`
	Ownerphoto    uint64 `json:"Owner_photo"`
	Usernameowner string `json:"Username_owner"`
}

// MyProfile rappresenta il profile dell'utente
type MyProfile struct {
	NumFollow    uint64           `json:"NumFollow"`
	NumFollowing uint64           `json:"NumFollowing"`
	NumPhoto     uint64           `json:"NumPhoto"`
	ListPhoto    []database.Photo `json:"ListPhoto"`
	UserOwner    User             `json:"UserOwner"`
}

// MyStream rappresenta lo stream dell'utente
type MyStream struct {
	ListPhoto []database.Photo `json:"ListPhoto"`
	UserOwner User             `json:"UserOwner"`
}
