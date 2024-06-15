/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error

	// login
	CheckUserExists(username string) (uint64, error)
	CreateUser(username string) (uint64, error)

	// User
	SetNewUsername(userID uint64, newUsername string) error
	GetListPhoto(userID uint64) ([]Photo, error)
	GetFollowPerson(userID string) ([]uint64, error)

	// Photo
	AddPhoto(Date string, Text string, URL []byte, userID uint64) (uint64, error)
	CheckPhotoExists(photoid uint64) (bool, error)
	DeletePhoto(id uint64) error
	CountPhoto(userUD uint64) (uint64, error)
	GetAllPhoto(listFollow []uint64) ([]Photo, error)

	// follow
	NewFollow(PersonaleUserID string, FollowUserID string) error
	CheckFollow(PersonaleUserID string, FollowUserID string) (bool, error)
	RemoveFollow(PersonalUserID string, FollowUserID string) error

	// Ban
	NewBan(PersonaleUserID string, BanUserID string) error
	CheckBan(PersonaleUserID string, BanUserID string) (bool, error)
	RemoveBan(PersonalUserID string, BanUserID string) error

	// like
	AddLike(userid uint64, photoid uint64) error
	CheckLikeExists(userid uint64, photoid uint64) (bool, error)
	DeleteLike(userid uint64, photoid uint64) error
	LikeCounterPhoto(photoid uint64) (uint64, error)

	// Comment
	AddComment(Date string, Text string, userID uint64, photoID uint64) (uint64, error)
	CheckCommentExists(commentid uint64) (bool, error)
	DeleteComment(id uint64) error
	CommentCounterPhoto(photoid uint64) (uint64, error)

	// My Profile
	CountFollow(username string) (uint64, error)
	CountFollowing(username string) (uint64, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	_, err := db.Exec("PRAGMA foreign_key = ON")
	if err != nil {
		return nil, err
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure

	sqlStmt := `CREATE TABLE IF NOT EXISTS User(
		Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		Username TEXT NOT NULL UNIQUE
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Photo(
		Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		Date TEXT NOT NULL, 
		Text TEXT NOT NULL, 
		Image BLOB NOT NULL, 
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES User(id)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Comment(
		Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Text TEXT NOT NULL,
		Date TEXT NOT NULL,
		user_id INTEGERE NOT NULL,
		photo_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES User(id),
		FOREIGN KEY (photo_id) REFERENCES photo(id)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Like(
		user_id INTEGER NOT NULL,
		photo_id INTEGER NON TULL,
		PRIMARY KEY (user_id, photo_id),
		FOREIGN KEY (user_id) REFERENCES User(id),
		FOREIGN KEY (photo_id) REFERENCES photo(id)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Follow(
		personal_user_id  INTEGER NOT NULL,
		follow_user_id INTEGER NOT NULL,
		PRIMARY KEY(personal_user_id, follow_user_id),
		FOREIGN KEY (personal_user_id) REFERENCES User(id),
		FOREIGN KEY (follow_user_id) REFERENCES User(id)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS Ban(
		personal_user_id INTEGER NOT NULL,
		ban_user_id INTEGER NOT NULL,
		PRIMARY KEY(personal_user_id, ban_user_id)
		FOREIGN KEY(personal_user_id) REFERENCES User(id)
		FOREIGN KEY(ban_user_id) REFERENCES User(id)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
