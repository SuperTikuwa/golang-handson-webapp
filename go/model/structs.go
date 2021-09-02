package model

// {
// 	"title": "Shopping",
// 	"description": "egg milk book",
// 	"date": "2021-09-01"
// }

type TaskPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

//   `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
//   `title` varchar(255) NOT NULL,
//   `description` varchar(255) NOT NULL,
//   `date` date NOT NULL,
//   `done` tinyint(1) NOT NULL DEFAULT '0'

type Task struct {
	ID          int    `gorm:"id"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Date        string `gorm:"date"`
	Done        bool   `gorm:"done"`
}

// {
// 	"id": 0,
// 	"title": "Shopping",
// 	"description": "egg milk book",
// 	"date": "2021-09-01"
// }

type TaskGetResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type TaskPutRequest struct {
	ID int `json:"id"`
}
