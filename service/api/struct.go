package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type User struct {
	ID       uint64 `json:"ID"`
	Username string `json:"Username"`
}

type Photo struct {
	ID             uint64 `json:"ID"`
	Date           string `json:"Date"`
	Text           string `json:"Text"`
	URL            []byte `json:"URL"`
	LikeCounter    uint64 `json:"LikeCounter"`
	CommentCounter uint64 `json:"CommentCounter"`
	UserID         uint64 `json:"UserID"`
}

type Follow struct {
	PersonalUserId string `json:"PersonalUserId"`
	FollowUserId   string `json:"FollowUserId"`
}

type Ban struct {
	PersonalUserId string `json:"PersonalUserId"`
	BanUserId      string `json:"BanUserId"`
}

type Comment struct {
	ID       uint64 `json:"ID"`
	Date     string `json:"Date"`
	Text     string `json:"Text"`
	User_id  uint64 `json:"User_id"`
	Photo_id uint64 `json:"Photo_id"`
}

type Like struct {
	User_id        uint64 `json:"User_id"`
	Username       string `json:"Username"`
	Photo_id       uint64 `json:"Photo_id"`
	Owner_photo    uint64 `json:"Owner_photo"`
	Username_owner string `json:"Username_owner"`
}

type MyProfile struct {
	NumFollow    uint64           `json:"NumFollow"`
	NumFollowing uint64           `json:"NumFollowing"`
	NumPhoto     uint64           `json:"NumPhoto"`
	ListPhoto    []database.Photo `json:"ListPhoto"`
	UserOwner    User             `json:"UserOwner"`
}

type MyStream struct {
	ListPhoto []database.Photo `json:"ListPhoto"`
	UserOwner User             `json:"UserOwner"`
}
