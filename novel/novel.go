package novel

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Novel struct {
	Id int64
	Title string
	Url string
	ImageUrl string
	Summary string
	Favorite bool
}

func New() *Novel {
	return &Novel{}
}

func FetchList() []Novel {
	novels := make([]Novel, 0)
	db := dbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, url, image_url, summary, favorite FROM novels")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		novels = append(novels, rowMapper(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return novels
}

func rowMapper(rows *sql.Rows) Novel {
	var id int64
	var title string
	var url string
	var imageUrl string
	var summary string
	var favorite bool

	if err := rows.Scan(&id, &title, &url, &imageUrl, &summary, &favorite); err != nil {
		log.Fatal(err)
	}
	return Novel{
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
		panic(err)
	}
	return db
}
