package novel

import (
	"github.com/l-lin/wn-tracker-api/db"
	_ "github.com/lib/pq"
	"log"
)

// The Novel representation
type Novel struct {
	Id       string
	Title    string
	Url      string
	ImageUrl string
	Summary  string
	Favorite bool
}

// Instanciate a new Novel
func New() *Novel {
	return &Novel{}
}

// Get the Novel given an id
func Get(id string) *Novel {
	database := db.Connect()
	defer database.Close()

	log.Printf("Executing query: SELECT id, title, url, image_url, summary, favorite FROM novels WHERE id = %s", id)
	row := database.QueryRow("SELECT id, title, url, image_url, summary, favorite FROM novels WHERE id = $1", id)
	return toNovel(row)
}

// Fetch the list of novels
func GetList() []*Novel {
	novels := make([]*Novel, 0)
	database := db.Connect()
	defer database.Close()

	rows, err := database.Query("SELECT id, title, url, image_url, summary, favorite FROM novels")
	if err != nil {
		log.Fatalf("[x] Error when getting the list of novels. Reason: %s", err.Error())
	}
	for rows.Next() {
		n := toNovel(rows)
		if n.IsValid() {
			novels = append(novels, n)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("[x] Error when getting the list of novels. Reason: %s", err.Error())
	}
	return novels
}

// Save the novel in the database
func (n *Novel) Save() {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Fatalf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	row := tx.QueryRow("INSERT INTO novels (title, url, image_url, summary, favorite) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		n.Title, n.Url, n.ImageUrl, n.Summary, n.Favorite)
	var lastId string
	if err := row.Scan(&lastId); err != nil {
		log.Fatalf("[x] Could not fetch the id of the newly created novel. Reason: %s", err.Error())
		tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
	n.Id = lastId
}

// Update the novel in the database
func (n *Novel) Update() {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Fatalf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	_, err = tx.Exec("UPDATE novels SET title = $1, url = $2, image_url = $3, summary = $4, favorite = $5 WHERE id = $6",
		n.Title, n.Url, n.ImageUrl, n.Summary, n.Favorite, n.Id)
	if err != nil {
		log.Fatalf("[x] Could not update the novel. Reason: %s", err.Error())
		tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
}

func (n *Novel) Delete() {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Fatalf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	_, err = tx.Exec("DELETE FROM novels WHERE id = $1", n.Id)
	if err != nil {
		log.Fatalf("[x] Could not delete the novel. Reason: %s", err.Error())
		tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
}

// Check if the novel has valid attributes
func (n *Novel) IsValid() bool {
	return n.Title != ""
}

// Fetch the content of the rows and build a new novel
func toNovel(rows db.RowMapper) *Novel {
	var id string
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
