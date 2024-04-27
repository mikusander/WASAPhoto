package api

import ()

type User struct {
	ID       uint64 `json:"ID"`
	Username string `json:"username"`
}

type Photo struct {
	ID             uint64 `json:"ID"`
	Date           string `json:"Date"`
	Text           string `json:"Text"`
	URL            []byte `json:"URL"`
	likeCounter    uint64 `json:"likeCounter"`
	commentCounter uint64 `json:"commentCounter"`
	userID         uint64 `json:"userID"`
}

type Follow struct {
	PersonalUserId string `json:"PersonalUserId"`
	FollowUserId   string `json:"FollowUserId"`
}

type Ban struct {
	PersonalUserId string `json:"PersonalUserId"`
	BanUserId      string `json:"BanUserId"`
}
