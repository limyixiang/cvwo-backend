package posts

import (
    "encoding/json"
	"time"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func ListPosts(db *database.Database) ([]models.Post, error) {
	var posts []models.Post
	rows, err := db.Query("SELECT `id`, `user_id`, `category_id`, `title`, `content`, `created_at`, `updated_at`, `likes`, `dislikes` FROM `post`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
        var createdAt, updatedAt []uint8

        err := rows.Scan(&post.ID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &createdAt, &updatedAt, &post.Likes, &post.Dislikes)
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
    rows, err := db.Query("SELECT `id`, `user_id`, `category_id`, `title`, `content`, `created_at`, `updated_at`, `likes`, `dislikes` FROM `post` WHERE `category_id` = ?", categoryID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var post models.Post
        var createdAt, updatedAt []uint8

        err := rows.Scan(&post.ID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &createdAt, &updatedAt, &post.Likes, &post.Dislikes)
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

    err := db.QueryRow("SELECT `id`, `user_id`, `category_id`, `title`, `content`, `created_at`, `updated_at`, `likes`, `dislikes` FROM `post` WHERE `id` = ?", id).Scan(&post.ID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &createdAt, &updatedAt, &post.Likes, &post.Dislikes)
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

func LikePost(db *database.Database, postID int, likesUsersID []int) error {
    _, err := db.Exec("UPDATE `post` SET `likes` = `likes` + 1 WHERE `id` = ?", postID)
    if err != nil {
        return err
    }

    likesUsersIDJSON, err := json.Marshal(likesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `post` SET `likes_users_id` = ? WHERE `id` = ?", likesUsersIDJSON, postID)
    return err
}

func UnlikePost(db *database.Database, postID int, likesUsersID []int) error {
    _, err := db.Exec("UPDATE `post` SET `likes` = `likes` - 1 WHERE `id` = ?", postID)
    if err != nil {
        return err
    }

    likesUsersIDJSON, err := json.Marshal(likesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `post` SET `likes_users_id` = ? WHERE `id` = ?", likesUsersIDJSON, postID)
    return err
}

func DislikePost(db *database.Database, postID int, dislikesUsersID []int) error {
    _, err := db.Exec("UPDATE `post` SET `dislikes` = `dislikes` + 1 WHERE `id` = ?", postID)
    if err != nil {
        return err
    }

    dislikesUsersIDJSON, err := json.Marshal(dislikesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `post` SET `dislikes_users_id` = ? WHERE `id` = ?", dislikesUsersIDJSON, postID)
    return err
}

func UndislikePost(db *database.Database, postID int, dislikesUsersID []int) error {
    _, err := db.Exec("UPDATE `post` SET `dislikes` = `dislikes` - 1 WHERE `id` = ?", postID)
    if err != nil {
        return err
    }

    dislikesUsersIDJSON, err := json.Marshal(dislikesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `post` SET `dislikes_users_id` = ? WHERE `id` = ?", dislikesUsersIDJSON, postID)
    return err
}

func CheckPostLikedByUser(db *database.Database, postID, userID int) (bool, error) {
    var likesUsersIDJSON []byte
    err := db.QueryRow("SELECT `likes_users_id` FROM `post` WHERE `id` = ?", postID).Scan(&likesUsersIDJSON)
    if err != nil {
        return false, err
    }

    if (len(likesUsersIDJSON) == 0) {
        return false, nil
    }

    var likesUsersID []int
    err = json.Unmarshal(likesUsersIDJSON, &likesUsersID)
    if err != nil {
        return false, err
    }

    for _, id := range likesUsersID {
        if id == userID {
            return true, nil
        }
    }
    return false, nil
}

func CheckPostDislikedByUser(db *database.Database, postID, userID int) (bool, error) {
    var dislikesUsersIDJSON []byte
    err := db.QueryRow("SELECT `dislikes_users_id` FROM `post` WHERE `id` = ?", postID).Scan(&dislikesUsersIDJSON)
    if err != nil {
        return false, err
    }

    if (len(dislikesUsersIDJSON) == 0) {
        return false, nil
    }

    var dislikesUsersID []int
    err = json.Unmarshal(dislikesUsersIDJSON, &dislikesUsersID)
    if err != nil {
        return false, err
    }

    for _, id := range dislikesUsersID {
        if id == userID {
            return true, nil
        }
    }
    return false, nil
}
