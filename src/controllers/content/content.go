package content

import (
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/skip2/go-qrcode"
)

type CTRL struct {
	Database *sql.DB
}

func NewCtrl(db *sql.DB) *CTRL {
	return &CTRL{
		Database: db,
	}
}

type Request struct {
	ClientID string `json:"client_id"`
	Content  string `json:"content"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Random func
func Random(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}

// Create func
func (ctrl *CTRL) Create(w http.ResponseWriter, r *http.Request) {
	var req Request
	var png []byte

	body := json.NewDecoder(r.Body)
	if err := body.Decode(&req); err != nil {
		log.Println(err)
	}

	png, err := qrcode.Encode(req.Content, qrcode.Medium, 256)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	stmt, err := ctrl.Database.Prepare("INSERT INTO `content` (client_id, content, code) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
	}

	code := Random(10)

	_, err = stmt.Exec(req.ClientID, req.Content, code)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(png))
}
