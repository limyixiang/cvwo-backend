package users

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/CVWO/sample-go-app/internal/api"
    "github.com/CVWO/sample-go-app/internal/dataaccess/users"
    "github.com/CVWO/sample-go-app/internal/database"
    "github.com/CVWO/sample-go-app/internal/models"
    "github.com/go-chi/chi/v5"
    "github.com/pkg/errors"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    userList, err := users.List(db)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve users"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, userList, http.StatusOK)
}

func HandleGetByName(w http.ResponseWriter, r *http.Request) {
    username := chi.URLParam(r, "username")

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    user, err := users.GetByName(db, username)
    if err != nil {
        if err == sql.ErrNoRows {
            api.WriteErrorResponse(w, errors.Wrap(err, "user not found"), http.StatusNotFound)
        } else {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve user"), http.StatusInternalServerError)
        }
        return
    }

    api.WriteResponse(w, user, http.StatusOK)
}

func HandleGetByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "invalid user ID"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    user, err := users.GetByID(db, id)
    if err != nil {
        if err == sql.ErrNoRows {
            api.WriteErrorResponse(w, errors.Wrap(err, "user not found"), http.StatusNotFound)
        } else {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve user"), http.StatusInternalServerError)
        }
        return
    }

    api.WriteResponse(w, user, http.StatusOK)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    createdUser, err := users.CreateUser(db, &user)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to create user"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, createdUser, http.StatusCreated)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
    username := chi.URLParam(r, "username")

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := users.UpdateUser(db, username, &user); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to update user"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, user, http.StatusOK)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
    username := chi.URLParam(r, "username")

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := users.DeleteUser(db, username); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete user"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, nil, http.StatusNoContent)
}
