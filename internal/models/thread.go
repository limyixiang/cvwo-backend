package models

import "time"

type Thread struct {
    ID        int       `json:"id"`
    CategoryID int      `json:"category_id"`
    Title     string    `json:"title"`
    PostID   int        `json:"post_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
