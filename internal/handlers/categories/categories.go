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

func HandleList(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        categoryList, err := categories.List(db)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve categories"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, categoryList, http.StatusOK)
    }
}

func HandleGet(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        categoryName := chi.URLParam(r, "name")

        category, err := categories.GetByName(db, categoryName)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve category"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, category, http.StatusOK)
    }
}

func HandleCreate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var category models.Category
        if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        if err := categories.CreateCategory(db, &category); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to create category"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, category, http.StatusCreated)
    }
}

func HandleUpdate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        oldName := chi.URLParam(r, "name")

        var category models.Category
        if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        if err := categories.UpdateCategory(db, category.Name, oldName); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update category"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, category, http.StatusOK)
    }
}

func HandleDelete(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        categoryName := chi.URLParam(r, "name")

        if err := categories.DeleteCategory(db, categoryName); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete category"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, nil, http.StatusNoContent)
    }
}
