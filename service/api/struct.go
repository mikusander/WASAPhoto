package api

import(
	"image"
)

type User struct{
	ID 			string `json:"ID"`
	Username	string `json:"username"`
}

type photo struct{
	ID			string `json:"ID"`
	Date 		string
	Text 		string `json:"text"`
	Image 		image.Image `json:"image"`
}