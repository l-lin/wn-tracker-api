package novel

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// The Novel representation
type Novel struct {
	Id int64
	Title string
	Url string
	ImageUrl string
	Summary string
	Favorite bool
}

// Instanciate a new Novel
func New(title, url, imageUrl, summary string, favorite bool) *Novel {
	return &Novel{
		Title: title,
		Url: url,
		ImageUrl: imageUrl,
		Summary: summary,
		Favorite: favorite,
	}
}

// Get the Novel given an id
func Get(id int64) *Novel {
	db := dbConnect()
	defer db.Close()

	log.Printf("Executing query: SELECT id, title, url, image_url, summary, favorite FROM novels WHERE id = %d", id)
	row := db.QueryRow("SELECT id, title, url, image_url, summary, favorite FROM novels WHERE id = $1", id)
	return rowMapper(row)
}

// Fetch the list of novels
func GetList() []*Novel {
	novels := make([]*Novel, 0)
	db := dbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, url, image_url, summary, favorite FROM novels")
	if err != nil {
		log.Fatalf("[x] Error when getting the list of novels. Reason: %s", err.Error())
	}
	for rows.Next() {
		n := rowMapper(rows)
		if n.IsValid() {
			novels = append(novels, n)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("[x] Error when getting the list of novels. Reason: %s", err.Error())
	}
	return novels
}

// Save the novel in the db
func (n *Novel) Save() {
	db := dbConnect()
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	row := tx.QueryRow("INSERT INTO novels (title, url, image_url, summary, favorite) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		n.Title, n.Url, n.ImageUrl, n.Summary, n.Favorite)
	var lastId int64
	if err := row.Scan(&lastId); err != nil {
		log.Fatalf("[x] Could not fetch the id of the newly created novel. Reason: %s", err.Error())
		tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
	n.Id = lastId
}

// Check if the novel has valid attributes
func (n *Novel) IsValid() bool {
	return n.Title != ""
}

// TODO: Put it somewhere else?

type NovelRowMapper interface {
	Scan(dest ...interface{}) error
}

func rowMapper(rows NovelRowMapper) *Novel {
	var id int64
	var title string
	var url string
	var imageUrl string
	var summary string
	var favorite bool

	rows.Scan(&id, &title, &url, &imageUrl, &summary, &favorite)

	return &Novel{
		Id: id,
		Title: title,
		Url: url,
		ImageUrl: imageUrl,
		Summary: summary,
		Favorite: favorite,
	}
}

func dbConnect() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("[x] Could not open the connection to the database. Reason: %s", err.Error())
	}
	return db
}
