package models 

import (
	"database/sql"
	"fmt"
	"time"
	_ "modernc.org/sqlite"
)

/*
CREATE TABLE connections (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    offer_client_sdp TEXT,
    accept_client_sdp TEXT,
		access_key TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
*/

type Connection struct {
    ID              int64
    OfferClientSdp  string
		AcceptClientSdp string
		AccessKey       string
		CreatedAt				time.Time 
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite", "my-database.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	if db == nil {
		fmt.Println("Databse connection is nil")
	}
}

func GetConnectionByAccessKey(accessKey string) Connection {
	rows, err := db.Query("SELECT * FROM connections WHERE access_key = ?;", accessKey)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var firstRow Connection

	rows.Next()
	rows.Scan(&firstRow.ID, &firstRow.OfferClientSdp, &firstRow.AcceptClientSdp, &firstRow.AccessKey, &firstRow.CreatedAt); if err != nil {
		fmt.Println("Error", err)
	}
	return firstRow
}

func UpdateConnectionRecordByAccessKey(accessKey string, sdpAccept string) {
	_, err := db.Exec("UPDATE connections set accept_client_sdp = ? WHERE access_key = ?", sdpAccept , accessKey)
	if err != nil {
		fmt.Println(err)
	}
}

func InsertConnectionRecord(newConnection Connection) {
	_ , err := db.Exec(
		"INSERT INTO connections (offer_client_sdp, accept_client_sdp, access_key) VALUES  (?, ?, ?)",
		newConnection.OfferClientSdp, newConnection.AcceptClientSdp, newConnection.AccessKey,
	)
	if err != nil {
		fmt.Println(err)
	}
}
