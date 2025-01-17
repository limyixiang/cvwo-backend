package comments

import (
	"encoding/json"
	"time"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func ListByPostID(db *database.Database, postID int) ([]models.Comment, error) {
    var comments []models.Comment
    rows, err := db.Query("SELECT `id`, `post_id`, `user_id`, `content`, `created_at`, `updated_at`, `likes`, `dislikes` FROM `comment` WHERE `post_id` = ?", postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var comment models.Comment
        var createdAt, updatedAt []uint8

        if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &createdAt, &updatedAt, &comment.Likes, &comment.Dislikes); err != nil {
            return nil, err
        }

        layout := "2006-01-02 15:04:05"
        comment.CreatedAt, err = time.Parse(layout, string(createdAt))
        if err != nil {
            return nil, err
        }

        comment.UpdatedAt, err = time.Parse(layout, string(updatedAt))
        if err != nil {
            return nil, err
        }

        comments = append(comments, comment)
    }
    return comments, nil
}

func GetByID(db *database.Database, id int) (*models.Comment, error) {
    var comment models.Comment
    var createdAt, updatedAt []uint8

    err := db.QueryRow("SELECT `id`, `post_id`, `user_id`, `content`, `created_at`, `updated_at`, `likes`, `dislikes` FROM `comment` WHERE `id` = ?", id).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &createdAt, &updatedAt, &comment.Likes, &comment.Dislikes)
    if err != nil {
        return nil, err
    }

    layout := "2006-01-02 15:04:05"
    comment.CreatedAt, err = time.Parse(layout, string(createdAt))
    if err != nil {
        return nil, err
    }

    comment.UpdatedAt, err = time.Parse(layout, string(updatedAt))
    if err != nil {
        return nil, err
    }

    return &comment, nil
}

func CreateComment(db *database.Database, comment *models.Comment) error {
    _, err := db.Exec("INSERT INTO `comment` (`post_id`, `user_id`, `content`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?)", comment.PostID, comment.UserID, comment.Content, comment.CreatedAt, comment.UpdatedAt)
    return err
}

func UpdateComment(db *database.Database, comment *models.Comment) error {
    _, err := db.Exec("UPDATE `comment` SET `content` = ?, `updated_at` = ? WHERE `id` = ?", comment.Content, comment.UpdatedAt, comment.ID)
    return err
}

func DeleteComment(db *database.Database, id int) error {
    _, err := db.Exec("DELETE FROM `comment` WHERE `id` = ?", id)
    return err
}

func LikeComment(db *database.Database, commentID int, likesUsersID []int) error {
    _, err := db.Exec("UPDATE `comment` SET `likes` = `likes` + 1 WHERE `id` = ?", commentID)
    if err != nil {
        return err
    }

    likesUsersIDJSON, err := json.Marshal(likesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `comment` SET `likes_users_id` = ? WHERE `id` = ?", likesUsersIDJSON, commentID)
    return err
}

func UnlikeComment(db *database.Database, commentID int, likesUsersID []int) error {
    _, err := db.Exec("UPDATE `comment` SET `likes` = `likes` - 1 WHERE `id` = ?", commentID)
    if err != nil {
        return err
    }

    likesUsersIDJSON, err := json.Marshal(likesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `comment` SET `likes_users_id` = ? WHERE `id` = ?", likesUsersIDJSON, commentID)
    return err
}

func DislikeComment(db *database.Database, commentID int, dislikesUsersID []int) error {
    _, err := db.Exec("UPDATE `comment` SET `dislikes` = `dislikes` + 1 WHERE `id` = ?", commentID)
    if err != nil {
        return err
    }

    dislikesUsersIDJSON, err := json.Marshal(dislikesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `comment` SET `dislikes_users_id` = ? WHERE `id` = ?", dislikesUsersIDJSON, commentID)
    return err
}

func UndislikeComment(db *database.Database, commentID int, dislikesUsersID []int) error {
    _, err := db.Exec("UPDATE `comment` SET `dislikes` = `dislikes` - 1 WHERE `id` = ?", commentID)
    if err != nil {
        return err
    }

    dislikesUsersIDJSON, err := json.Marshal(dislikesUsersID)
    if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE `comment` SET `dislikes_users_id` = ? WHERE `id` = ?", dislikesUsersIDJSON, commentID)
    return err
}

func CheckCommentLikedByUser(db *database.Database, commentID, userID int) (bool, error) {
    var likesUsersIDJSON []byte
    err := db.QueryRow("SELECT `likes_users_id` FROM `comment` WHERE `id` = ?", commentID).Scan(&likesUsersIDJSON)
    if err != nil {
        return false, err
    }

    if len(likesUsersIDJSON) == 0 {
        return false, nil
    }

    var likesUsersID []int
    if err := json.Unmarshal(likesUsersIDJSON, &likesUsersID); err != nil {
        return false, err
    }

    for _, id := range likesUsersID {
        if id == userID {
            return true, nil
        }
    }

    return false, nil
}

func CheckCommentDislikedByUser(db *database.Database, commentID, userID int) (bool, error) {
    var dislikesUsersIDJSON []byte
    err := db.QueryRow("SELECT `dislikes_users_id` FROM `comment` WHERE `id` = ?", commentID).Scan(&dislikesUsersIDJSON)
    if err != nil {
        return false, err
    }

    if len(dislikesUsersIDJSON) == 0 {
        return false, nil
    }

    var dislikesUsersID []int
    if err := json.Unmarshal(dislikesUsersIDJSON, &dislikesUsersID); err != nil {
        return false, err
    }

    for _, id := range dislikesUsersID {
        if id == userID {
            return true, nil
        }
    }

    return false, nil
}
