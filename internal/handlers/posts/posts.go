package posts

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/CVWO/sample-go-app/internal/api"
    "github.com/CVWO/sample-go-app/internal/dataaccess/posts"
    "github.com/CVWO/sample-go-app/internal/database"
    "github.com/CVWO/sample-go-app/internal/models"
    "github.com/go-chi/chi/v5"
    "github.com/pkg/errors"
)

const (
    ListPosts = "posts.HandleList"

    SuccessfulListPostsMessage = "Successfully listed posts"
    ErrRetrieveDatabase        = "Failed to retrieve database in %s"
    ErrRetrievePosts           = "Failed to retrieve posts in %s"
    ErrEncodeView              = "Failed to retrieve posts in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListPosts)), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    postList, err := posts.ListPosts(db)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts)), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, postList, http.StatusOK)
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
    postIDStr := chi.URLParam(r, "id")
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

    post, err := posts.GetByID(db, postID)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve post"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, post, http.StatusOK)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := posts.CreatePost(db, &post); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to create post"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, post, http.StatusCreated)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
    postIDStr := chi.URLParam(r, "id")
    postID, err := strconv.Atoi(postIDStr)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "invalid post ID"), http.StatusBadRequest)
        return
    }

    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
        return
    }
    post.ID = postID

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := posts.UpdatePost(db, &post); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to update post"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, post, http.StatusOK)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
    postIDStr := chi.URLParam(r, "id")
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

    if err := posts.DeletePost(db, postID); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete post"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, nil, http.StatusNoContent)
}
