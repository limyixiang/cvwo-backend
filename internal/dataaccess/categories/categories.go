package categories

import (
    "github.com/CVWO/sample-go-app/internal/database"
    "github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.Category, error) {
    var categories []models.Category
    rows, err := db.Query("SELECT * FROM `categories`")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var category models.Category
        if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
            return nil, err
        }
        categories = append(categories, category)
    }
    return categories, nil
}

func GetByName(db *database.Database, name string) (*models.Category, error) {
    var category models.Category
    err := db.QueryRow("SELECT `id`, `name`, `description` FROM `category` WHERE `name` = ?", name).Scan(&category.ID, &category.Name, &category.Description)
    if err != nil {
        return nil, err
    }
    return &category, nil
}

func CreateCategory(db *database.Database, category *models.Category) error {
    _, err := db.Exec("INSERT INTO `category` (`name`, `description`) VALUES (?, ?)", category.Name, category.Description)
    return err
}

func UpdateCategory(db *database.Database, newName string, newDescription string, oldName string) error {
    _, err := db.Exec("UPDATE `category` SET `name` = ?, `description` = ? WHERE `name` = ?", newName, newDescription, oldName)
    return err
}

func DeleteCategory(db *database.Database, name string) error {
    _, err := db.Exec("DELETE FROM `category` WHERE `name` = ?", name)
    return err
}