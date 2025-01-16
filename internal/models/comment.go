package models

import "time"

type Comment struct {
    ID                  int         `json:"id"`
    PostID              int         `json:"post_id"`
    UserID              int         `json:"user_id"`
    Content             string      `json:"content"`
    CreatedAt           time.Time   `json:"created_at"`
    UpdatedAt           time.Time   `json:"updated_at"`
    Likes               int         `json:"likes"`
    Dislikes            int         `json:"dislikes"`
    LikesUsersID        []int       `json:"likes_users_id"`
    DislikesUsersID     []int       `json:"dislikes_users_id"`
}
