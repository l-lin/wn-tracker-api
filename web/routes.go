package web

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Availables routes
var routes = Routes{
	Route{
		"Novels",
		"GET",
		"/novels",
		Novels,
	},
	Route{
		"SaveNovel",
		"POST",
		"/novels",
		SaveNovel,
	},
	Route{
		"Novel",
		"GET",
		"/novels/{novelId}",
		Novel,
	},
	Route{
		"UpdateNovel",
		"PUT",
		"/novels/{novelId}",
		UpdateNovel,
	},
	Route{
		"DeleteNovel",
		"DELETE",
		"/novels/{novelId}",
		DeleteNovel,
	},
	Route{
		"Notifications",
		"GET",
		"/notifications",
		Notifications,
	},
	Route{
		"Notification",
		"GET",
		"/notifications/{notificationId}",
		Notification,
	},
	Route{
		"DeleteNotification",
		"DELETE",
		"/notifications/{notificationId}",
		DeleteNotification,
	},
}
