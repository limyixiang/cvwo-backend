package threads

import (
    "github.com/CVWO/sample-go-app/internal/database"
    "github.com/CVWO/sample-go-app/internal/models"
)

func ListByCategoryID(db *database.Database, categoryID int) ([]models.Thread, error) {
    var threads []models.Thread
    rows, err := db.Query("SELECT `id`, `category_id`, `title`, `post_id`, `created_at`, `updated_at` FROM `thread` WHERE `category_id` = ?", categoryID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var thread models.Thread
        if err := rows.Scan(&thread.ID, &thread.CategoryID, &thread.Title, &thread.PostID, &thread.CreatedAt, &thread.UpdatedAt); err != nil {
            return nil, err
        }
        threads = append(threads, thread)
    }
    return threads, nil
}

func GetByID(db *database.Database, id int) (*models.Thread, error) {
    var thread models.Thread
    err := db.QueryRow("SELECT `id`, `category_id`, `title`, `post_id`, `created_at`, `updated_at` FROM `thread` WHERE `id` = ?", id).Scan(&thread.ID, &thread.CategoryID, &thread.Title, &thread.PostID, &thread.CreatedAt, &thread.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &thread, nil
}

func CreateThread(db *database.Database, thread *models.Thread) error {
    _, err := db.Exec("INSERT INTO `thread` (`category_id`, `title`, `post_id`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?)", thread.CategoryID, thread.Title, thread.PostID, thread.CreatedAt, thread.UpdatedAt)
    return err
}

func UpdateThread(db *database.Database, thread *models.Thread) error {
    _, err := db.Exec("UPDATE `thread` SET `title` = ?, `updated_at` = ? WHERE `id` = ?", thread.Title, thread.UpdatedAt, thread.ID)
    return err
}

func DeleteThread(db *database.Database, id int) error {
    _, err := db.Exec("DELETE FROM `thread` WHERE `id` = ?", id)
    return err
}
