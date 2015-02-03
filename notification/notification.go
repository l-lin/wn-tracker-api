package notification

import (
	"github.com/l-lin/wn-tracker-api/db"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Notification struct {
	NotificationId string `json:"notificationId`
	FeedId         string `json:"feedId`
	Title          string `json:"title"`
	Link           string `json:"link"`
	Description    string `json:"description"`
	PubDate        time.Time `json:"-"`
}

// Instanciate a new Notification
func New() *Notification {
	return &Notification{}
}

// Check with the given feedId and link, there is already a notification or not
func Exists(feedId, link string) bool {
	database := db.Connect()
	defer database.Close()

	row := database.QueryRow("SELECT CASE WHEN EXISTS(SELECT 1 FROM notifications WHERE feed_id = $1 AND link = $2) THEN 1 ELSE 0 END", feedId, link)
	var exists int64
	if err := row.Scan(&exists); err != nil {
		log.Fatalf("[x] Could not check if there is existing notifications for feedId '%s' and link '%s'. Reason: %s", feedId, link, err.Error())
	}
	return exists == 1;
}

// Fetch all notifications from the db
func GetList(userId string) []*Notification {
	feeds := make([]*Notification, 0)
	database := db.Connect()
	defer database.Close()

	rows, err := database.Query(`
		SELECT n.notification_id, n.feed_id, n.title, n.link, n.description, n.pub_date
		FROM notifications n
		JOIN feeds f ON f.feed_id = n.feed_id
		JOIN novels no ON no.novel_id = f.novel_id
		WHERE no.user_id = $1
	`, userId)
	if err != nil {
		log.Fatalf("[x] Error when getting the list of feeds. Reason: %s", err.Error())
	}
	for rows.Next() {
		feeds = append(feeds, toNotification(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("[x] Error when getting the list of feeds. Reason: %s", err.Error())
	}
	return feeds
}

// Save the novel in the database
func (n *Notification) Save() {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Fatalf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	row := tx.QueryRow("INSERT INTO notifications (feed_id, title, link, description, pub_date) VALUES ($1, $2, $3, $4, $5) RETURNING notification_id",
		n.FeedId, n.Title, n.Link, n.Description, n.PubDate)
	var lastId string
	if err := row.Scan(&lastId); err != nil {
		tx.Rollback()
		log.Fatalf("[x] Could not fetch the novel_id of the newly created novel. Reason: %s", err.Error())
	}
	n.NotificationId = lastId
	if err := tx.Commit(); err != nil {
		log.Fatalf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
}

// Delete a notification
func (n *Notification) Delete() {
	database := db.Connect()
	defer database.Close()
	tx, err := database.Begin()
	if err != nil {
		log.Fatalf("[x] Could not start the transaction. Reason: %s", err.Error())
	}
	_, err = tx.Exec("DELETE FROM notification WHERE novel_id = $1", n.NotificationId)
	if err != nil {
		tx.Rollback()
		log.Fatalf("[x] Could not delete the notification. Reason: %s", err.Error())
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("[x] Could not commit the transaction. Reason: %s", err.Error())
	}
}

// Fetch the content of the rows and build a new default notification
func toNotification(rows db.RowMapper) *Notification {
	notification := New()
	err := rows.Scan(
		&notification.NotificationId,
		&notification.FeedId,
		&notification.Title,
		&notification.Link,
		&notification.Description,
		&notification.PubDate,
	)
	if err != nil {
		log.Printf("[-] Could not scan the notification. Reason: %s", err.Error())
	}
	return notification
}