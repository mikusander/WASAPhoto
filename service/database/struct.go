package database

// User rappresenta un utente nel sistema.
type User struct {
	ID       uint64 `json:"ID"`
	Username string `json:"username"`
}

// Photo rappresenta una foto del sistema
type Photo struct {
	ID             uint64    `json:"ID"`
	Date           string    `json:"Date"`
	Text           string    `json:"Text"`
	URL            []byte    `json:"URL"`
	LikeCounter    uint64    `json:"likeCounter"`
	CommentCounter uint64    `json:"commentCounter"`
	ListComment    []Comment `json:"listComment"`
	UserID         uint64    `json:"userID"`
}

// Comment rappresenta la struttura di un commento
// ID è l'ID del commento
// PhotoID è l'ID della foto associata
// UserID è l'ID dell'utente che ha postato il commento
// Text è il testo del commento
// CreatedAt è la data e l'ora di creazione del commento
type Comment struct {
	ID           uint64 `json:"ID"`
	Date         string `json:"Date"`
	Text         string `json:"Text"`
	UserID       uint64 `json:"User_id"`
	UserUsername string `json:"UserUsername"`
	PhotoID      uint64 `json:"Photo_id"`
}
