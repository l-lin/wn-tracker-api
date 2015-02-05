package token

import (
	"github.com/l-lin/wn-tracker-api/db"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"os"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
)

const oauth2RefreshEndPoint  = "https://www.googleapis.com/oauth2/v3/token"

// The feed
type Token struct {
	UserId       string `json:"-"`
	RefreshToken string `json:"-"`
}

func (t Token) String() string {
	return fmt.Sprintf("UserId = %s, RefreshToken = %s", t.UserId, t.RefreshToken)
}

type ResfreshTokenConfig struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// Instanciate a new feed
func New() *Token {
	return &Token{}
}

// Check if the given user has already stored his token
func Exists(userId string) bool {
	database := db.Connect()
	defer database.Close()

	row := database.QueryRow("SELECT CASE WHEN EXISTS(SELECT 1 FROM tokens WHERE user_id = $1) THEN 1 ELSE 0 END", userId)
	var exists int64
	if err := row.Scan(&exists); err != nil {
		log.Printf("[x] Could not check if there is existing token for user '%s'. Reason: %s", userId, err.Error())
	}
	return exists == 1;
}

// Get the Novel given an novelId
func Get(userId string) *Token {
	database := db.Connect()
	defer database.Close()

	row := database.QueryRow("SELECT user_id, refresh_token FROM tokens WHERE user_id = $1", userId)
	return toToken(row)
}

// Get the Novel given an novelId
func GetByAccessToken(accessToken string) *Token {
	database := db.Connect()
	defer database.Close()

	row := database.QueryRow("SELECT user_id, refresh_token FROM tokens WHERE access_token = $1", accessToken)
	return toToken(row)
}

// Save the token in the database
func (token *Token) Save() {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Printf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	_, err = tx.Exec("INSERT INTO tokens (user_id, refresh_token) VALUES ($1, $2)", token.UserId, token.RefreshToken)
	if err != nil {
		tx.Rollback()
		log.Printf("[x] Could not save the token. Reason: %s", err.Error())
	}
	if err := tx.Commit(); err != nil {
		log.Printf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
}

// Update the token
func (t *Token) Update() {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Printf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	_, err = tx.Exec("UPDATE tokens SET refresh_token = $1 WHERE user_id = $2", t.RefreshToken, t.UserId)
	if err != nil {
		tx.Rollback()
		log.Printf("[x] Could not update the token. Reason: %s", err.Error())
	}

	if err := tx.Commit(); err != nil {
		log.Printf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
}

// Refresh the given token
func (t *Token) Refresh() bool {
	c := &ResfreshTokenConfig{os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "refresh_token", t.RefreshToken}
	buf, _ := json.Marshal(c)
	body := bytes.NewBuffer(buf)
	r, err := http.Post(oauth2RefreshEndPoint, "application/json", body)
	if err != nil {
		log.Printf("[x] Could not refresh the token. Reason: %s", err.Error())
		return false
	}
	defer r.Body.Close()
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[x] Error reading content of %s. Reson: %s", oauth2RefreshEndPoint, err.Error())
		return false
	}
	var oauthResponse RefreshTokenResponse
	if err := json.Unmarshal(response, &oauthResponse); err != nil {
		log.Printf("[x] Could not read the JSON of the response after refreshing the token. Reason: %s", err.Error())
		return false
	}
	return true
}

// Fetch the content of the rows and build a new token
func toToken(rows db.RowMapper) *Token {
	token := New()
	err := rows.Scan(
		&token.UserId,
		&token.RefreshToken,
	)
	if err != nil {
		log.Printf("[-] Could not scan the token. Reason: %s", err.Error())
	}
	return token
}
