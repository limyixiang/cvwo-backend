package comments

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CVWO/sample-go-app/internal/api"
	"github.com/CVWO/sample-go-app/internal/dataaccess/comments"
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

const (
    ListComments = "comments.HandleList"

    SuccessfulListCommentsMessage = "Successfully listed comments"
    ErrRetrieveDatabase           = "Failed to retrieve database in %s"
    ErrRetrieveComments           = "Failed to retrieve comments in %s"
    ErrEncodeView                 = "Failed to retrieve comments in %s"
)

func HandleList(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postIDStr := chi.URLParam(r, "postID")
        postID, err := strconv.Atoi(postIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid post ID"), http.StatusBadRequest)
            return
        }

        commentList, err := comments.ListByPostID(db, postID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comments"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, commentList, http.StatusOK)
    }
}

func HandleGet(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        comment, err := comments.GetByID(db, commentID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comment"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, comment, http.StatusOK)
    }
}

func HandleCreate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var comment models.Comment
        if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        if err := comments.CreateComment(db, &comment); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to create comment"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, comment, http.StatusCreated)
    }
}

func HandleUpdate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        var comment models.Comment
        if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }
        comment.ID = commentID

        if err := comments.UpdateComment(db, &comment); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update comment"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, comment, http.StatusOK)
    }
}

func HandleDelete(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        if err := comments.DeleteComment(db, commentID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete comment"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, nil, http.StatusNoContent)
    }
}

func HandleLike(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        var userID struct {
            UserID int `json:"user_id"`
        }
        if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        comment, err := comments.GetByID(db, commentID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comment"), http.StatusInternalServerError)
            return
        }

        comment.LikesUsersID = append(comment.LikesUsersID, userID.UserID)
        if err := comments.LikeComment(db, commentID, comment.LikesUsersID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update comment likes"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, comment, http.StatusOK)
    }
}

func HandleUnlike(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        var userID struct {
            UserID int `json:"user_id"`
        }
        if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        comment, err := comments.GetByID(db, commentID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comment"), http.StatusInternalServerError)
            return
        }

        for i, id := range comment.LikesUsersID {
            if id == userID.UserID {
                comment.LikesUsersID = append(comment.LikesUsersID[:i], comment.LikesUsersID[i+1:]...)
                break
            }
        }

        if err := comments.UnlikeComment(db, commentID, comment.LikesUsersID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update comment likes"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, comment, http.StatusOK)
    }
}

func HandleDislike(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        var userID struct {
            UserID int `json:"user_id"`
        }
        if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        comment, err := comments.GetByID(db, commentID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comment"), http.StatusInternalServerError)
            return
        }

        comment.DislikesUsersID = append(comment.DislikesUsersID, userID.UserID)
        if err := comments.DislikeComment(db, commentID, comment.DislikesUsersID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update comment dislikes"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, comment, http.StatusOK)
    }
}

func HandleUndislike(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        var userID struct {
            UserID int `json:"user_id"`
        }
        if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        comment, err := comments.GetByID(db, commentID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comment"), http.StatusInternalServerError)
            return
        }

        for i, id := range comment.DislikesUsersID {
            if id == userID.UserID {
                comment.DislikesUsersID = append(comment.DislikesUsersID[:i], comment.DislikesUsersID[i+1:]...)
                break
            }
        }

        if err := comments.UndislikeComment(db, commentID, comment.DislikesUsersID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update comment dislikes"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, comment, http.StatusOK)
    }
}

func HandleCheckLikedByUser(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        userIDStr := chi.URLParam(r, "userID")
        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid user ID"), http.StatusBadRequest)
            return
        }

        liked, err := comments.CheckCommentLikedByUser(db, commentID, userID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to check if comment is liked by user"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, liked, http.StatusOK)
    }
}

func HandleCheckDislikedByUser(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentIDStr := chi.URLParam(r, "id")
        commentID, err := strconv.Atoi(commentIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
            return
        }

        userIDStr := chi.URLParam(r, "userID")
        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid user ID"), http.StatusBadRequest)
            return
        }

        disliked, err := comments.CheckCommentDislikedByUser(db, commentID, userID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to check if comment is disliked by user"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, disliked, http.StatusOK)
    }
}
