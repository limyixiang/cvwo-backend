package categories

import (
    "encoding/json"
    "net/http"

    "github.com/CVWO/sample-go-app/internal/api"
    "github.com/CVWO/sample-go-app/internal/dataaccess/categories"
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

    categoryList, err := categories.List(db)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve categories"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, categoryList, http.StatusOK)
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
    categoryName := chi.URLParam(r, "name")

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    category, err := categories.GetByName(db, categoryName)
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve category"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, category, http.StatusOK)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := categories.CreateCategory(db, &category); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to create category"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, category, http.StatusCreated)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
    oldName := chi.URLParam(r, "name")

    var category models.Category
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
        return
    }

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := categories.UpdateCategory(db, category.Name, category.Description, oldName); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to update category"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, category, http.StatusOK)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
    categoryName := chi.URLParam(r, "name")

    db, err := database.GetDB()
    if err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve database"), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    if err := categories.DeleteCategory(db, categoryName); err != nil {
        api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete category"), http.StatusInternalServerError)
        return
    }

    api.WriteResponse(w, nil, http.StatusNoContent)
}
