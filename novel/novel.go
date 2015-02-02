package novel

import (
	"github.com/l-lin/wn-tracker-api/db"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// The Novel representation
type Novel struct {
	Id          string        `json:"id"`
	UserId      string        `json:"-"`
	Title       string        `json:"title"`
	Url         string        `json:"url"`
	ImageUrl    string        `json:"imageUrl"`
	Summary     string        `json:"summary"`
	Favorite    bool        `json:"favorite"`
	Feed
}

type Feed struct {
	FeedUrl     string        `json:"feedUrl"`
	LastUpdated time.Time    `json:"-"`
}

// Instanciate a new Novel
func New() *Novel {
	return &Novel{}
}

// Check if the given user has novels
func Exists(userId string) bool {
	database := db.Connect()
	defer database.Close()

	row := database.QueryRow("SELECT CASE WHEN EXISTS(SELECT 1 FROM novels WHERE user_id = $1) THEN 1 ELSE 0 END", userId)
	var exists int64
	if err := row.Scan(&exists); err != nil {
		log.Fatalf("[x] Could not check if there is existing novels for user '%s'. Reason: %s", userId, err.Error())
	}
	return exists == 1;
}

// Copy the default novels to the newly subscribed user
func CopyDefaultFor(userId string) {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Fatalf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	_, err = tx.Exec("INSERT INTO novels (user_id, title, url, image_url, summary, favorite) SELECT $1, title, url, image_url, summary, favorite FROM default_novels", userId)
	if err != nil {
		tx.Rollback()
		log.Fatalf("[x] Could not copy the default novels. Reason: %s", err.Error())
	}
	_, err = tx.Exec("INSERT INTO feeds (id, url) SELECT n.id, dn.url FROM novels n JOIN default_novels dn ON dn.url = n.url WHERE user_id = $1", userId)
	if err != nil {
		tx.Rollback()
		log.Fatalf("[x] Could not copy the default feeds. Reason: %s", err.Error())
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
}

// Get the Novel given an id
func Get(id, userId string) *Novel {
	database := db.Connect()
	defer database.Close()

	row := database.QueryRow(`
	SELECT n.id, n.user_id, n.title, n.url, n.image_url, n.summary, n.favorite, f.url, f.last_updated
	FROM novels n
	 JOIN feeds f on f.id = n.id
	WHERE id = $1 AND user_id = $2`,
		id, userId)
	return toNovel(row)
}

// Fetch the list of novels
func GetList(userId string) []*Novel {
	novels := make([]*Novel, 0)
	database := db.Connect()
	defer database.Close()

	rows, err := database.Query(`
	SELECT n.id, n.user_id, n.title, n.url, n.image_url, n.summary, n.favorite, f.url, f.last_updated
	FROM novels n
	 JOIN feeds f on f.id = n.id
	WHERE user_id = $1`,
		userId)
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
	row := tx.QueryRow("INSERT INTO novels (user_id, title, url, image_url, summary, favorite) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		n.UserId, n.Title, n.Url, n.FeedUrl, n.ImageUrl, n.Summary, n.Favorite)
	var lastId string
	if err := row.Scan(&lastId); err != nil {
		tx.Rollback()
		log.Fatalf("[x] Could not fetch the id of the newly created novel. Reason: %s", err.Error())
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
	_, err = tx.Exec(`
	UPDATE novels SET title = $1, url = $2, image_url = $5, summary = $6, favorite = $7
	WHERE id = $8 AND user_id = $9`,
		n.Title, n.Url, n.ImageUrl, n.Summary, n.Favorite, n.Id, n.UserId)
	if err != nil {
		log.Fatalf("[x] Could not update the novel. Reason: %s", err.Error())
		tx.Rollback()
	}
	_, err = tx.Exec(`
	UPDATE feeds SET url = $1, last_updated = $2
	WHERE id = $3 AND user_id = $4`,
		n.FeedUrl, time.Now(), n.Id, n.UserId)
	if err != nil {
		tx.Rollback()
		log.Fatalf("[x] Could not update the feeds. Reason: %s", err.Error())
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
	_, err = tx.Exec("DELETE FROM novels WHERE id = $1 AND user_id = $2", n.Id, n.UserId)
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
	return n.UserId != "" && n.Title != "" && n.Url != ""
}

// Fetch the content of the rows and build a new default novel
func toNovel(rows db.RowMapper) *Novel {
	novel := New()
	err := rows.Scan(
		&novel.Id,
		&novel.UserId,
		&novel.Title,
		&novel.Url,
		&novel.ImageUrl,
		&novel.Summary,
		&novel.Favorite,
		&novel.FeedUrl,
		&novel.LastUpdated,
	)
	if err != nil {
		log.Printf("[-] Could not scan the novel. Reason: %s", err.Error())
	}
	return novel
}
