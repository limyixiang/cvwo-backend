package posts

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func ListByThreadID(db *database.Database, threadID int) ([]models.Post, error) {
	var posts []models.Post
	rows, err := db.Query("SELECT `id`, `thread_id`, `content` FROM `post` WHERE `thread_id` = ?", threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.ThreadID, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetByID(db *database.Database, id int) (*models.Post, error) {
	var post models.Post
	err := db.QueryRow("SELECT `id`, `thread_id`, `user_id`, `content`, `created_at`, `updated_at` FROM `post` WHERE `id` = ?", id).Scan(&post.ID, &post.ThreadID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func CreatePost(db *database.Database, post *models.Post) error {
    _, err := db.Exec("INSERT INTO `post` (`thread_id`, `user_id`, `content`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?)", post.ThreadID, post.UserID, post.Content, post.CreatedAt, post.UpdatedAt)
    return err
}

func UpdatePost(db *database.Database, post *models.Post) error {
    _, err := db.Exec("UPDATE `post` SET `content` = ?, `updated_at` = ? WHERE `id` = ?", post.Content, post.UpdatedAt, post.ID)
    return err
}

func DeletePost(db *database.Database, id int) error {
    _, err := db.Exec("DELETE FROM `post` WHERE `id` = ?", id)
    return err
}
