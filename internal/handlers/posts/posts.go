package posts

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "time"

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

func HandleList(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postList, err := posts.ListPosts(db)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts)), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, postList, http.StatusOK)
    }
}

func HandleListByCategory(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        categoryIDStr := chi.URLParam(r, "id")
        categoryID, err := strconv.Atoi(categoryIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid category ID"), http.StatusBadRequest)
            return
        }

        postList, err := posts.ListPostsByCategory(db, categoryID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve posts"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, postList, http.StatusOK)
    }
}

func HandleGet(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postIDStr := chi.URLParam(r, "id")
        postID, err := strconv.Atoi(postIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid post ID"), http.StatusBadRequest)
            return
        }

        post, err := posts.GetByID(db, postID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve post"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, post, http.StatusOK)
    }
}

func HandleCreate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var post models.Post
        if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        if err := posts.CreatePost(db, &post); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to create post"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, post, http.StatusCreated)
    }
}

func HandleUpdate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
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

        if err := posts.UpdatePost(db, &post); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update post"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, post, http.StatusOK)
    }
}

func HandleUpdateLastUpdated(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postIDStr := chi.URLParam(r, "id")
        postID, err := strconv.Atoi(postIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid post ID"), http.StatusBadRequest)
            return
        }

        post, err := posts.GetByID(db, postID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve post"), http.StatusInternalServerError)
            return
        }

        post.UpdatedAt = time.Now()

        if err := posts.UpdatePost(db, post); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update post"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, post, http.StatusOK)
    }
}

func HandleDelete(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postIDStr := chi.URLParam(r, "id")
        postID, err := strconv.Atoi(postIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid post ID"), http.StatusBadRequest)
            return
        }

        if err := posts.DeletePost(db, postID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete post"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, nil, http.StatusNoContent)
    }
}
