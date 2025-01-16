package comments

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "time"

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

func parseTime(value interface{}) (time.Time, error) {
    switch v := value.(type) {
    case time.Time:
        return v, nil
    case []uint8:
        return time.Parse("2006-01-02 15:04:05", string(v))
    case string:
        return time.Parse("2006-01-02 15:04:05", v)
    case nil:
        return time.Time{}, fmt.Errorf("nil value provided")
    default:
        return time.Time{}, fmt.Errorf("unsupported type: %T", v)
    }
}

func HandleList(w http.ResponseWriter, r *http.Request) {
    postIDStr := chi.URLParam(r, "postID")
    postID, err := strconv.Atoi(postIDStr)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "invalid post ID"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    commentList, err := comments.ListByPostID(db, postID)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comments"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, commentList, http.StatusOK)
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
    commentIDStr := chi.URLParam(r, "id")
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    comment, err := comments.GetByID(db, commentID)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve comment"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, comment, http.StatusOK)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
    var comment models.Comment
    if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := comments.CreateComment(db, &comment); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to create comment"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, comment, http.StatusCreated)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
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

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := comments.UpdateComment(db, &comment); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to update comment"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, comment, http.StatusOK)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
    commentIDStr := chi.URLParam(r, "id")
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "invalid comment ID"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := comments.DeleteComment(db, commentID); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete comment"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, nil, http.StatusNoContent)
}

func HandleLike(w http.ResponseWriter, r *http.Request) {
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

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

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
