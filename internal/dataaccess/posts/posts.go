package posts

import (
	"time"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func ListPosts(db *database.Database) ([]models.Post, error) {
	var posts []models.Post
	rows, err := db.Query("SELECT * FROM `post`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
        var createdAt, updatedAt []uint8

        err := rows.Scan(&post.ID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &createdAt, &updatedAt)
        if err != nil {
            return nil, err
        }

        layout := "2006-01-02 15:04:05"
        post.CreatedAt, err = time.Parse(layout, string(createdAt))
        if err != nil {
            return nil, err
        }

        post.UpdatedAt, err = time.Parse(layout, string(updatedAt))
        if err != nil {
            return nil, err
        }

        posts = append(posts, post)
	}
	return posts, nil
}

func ListPostsByCategory(db *database.Database, categoryID int) ([]models.Post, error) {
    var posts []models.Post
    rows, err := db.Query("SELECT * FROM `post` WHERE `category_id` = ?", categoryID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var post models.Post
        var createdAt, updatedAt []uint8

        err := rows.Scan(&post.ID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &createdAt, &updatedAt)
        if err != nil {
            return nil, err
        }

        layout := "2006-01-02 15:04:05"
        post.CreatedAt, err = time.Parse(layout, string(createdAt))
        if err != nil {
            return nil, err
        }

        post.UpdatedAt, err = time.Parse(layout, string(updatedAt))
        if err != nil {
            return nil, err
        }

        posts = append(posts, post)
    }
    return posts, nil
}

func GetByID(db *database.Database, id int) (*models.Post, error) {
	var post models.Post
    var createdAt, updatedAt []uint8

    err := db.QueryRow("SELECT * FROM `post` WHERE `id` = ?", id).Scan(&post.ID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &createdAt, &updatedAt)
    if err != nil {
        return nil, err
    }

    layout := "2006-01-02 15:04:05"
    post.CreatedAt, err = time.Parse(layout, string(createdAt))
    if err != nil {
        return nil, err
    }

    post.UpdatedAt, err = time.Parse(layout, string(updatedAt))
    if err != nil {
        return nil, err
    }

    return &post, nil
}

func CreatePost(db *database.Database, post *models.Post) error {
    _, err := db.Exec("INSERT INTO `post` (`user_id`, `category_id`, `title`, `content`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?)", post.UserID, post.CategoryID, post.Title, post.Content, post.CreatedAt, post.UpdatedAt)
    return err
}

func UpdatePost(db *database.Database, post *models.Post) error {
    _, err := db.Exec("UPDATE `post` SET `title` = ?, `content` = ?, `updated_at` = ? WHERE `id` = ?", post.Title, post.Content, post.UpdatedAt, post.ID)
    return err
}

func DeletePost(db *database.Database, id int) error {
    _, err := db.Exec("DELETE FROM `post` WHERE `id` = ?", id)
    return err
}
