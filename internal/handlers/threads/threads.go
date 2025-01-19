package threads

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/CVWO/sample-go-app/internal/api"
    "github.com/CVWO/sample-go-app/internal/dataaccess/threads"
    "github.com/CVWO/sample-go-app/internal/database"
    "github.com/CVWO/sample-go-app/internal/models"
    "github.com/go-chi/chi/v5"
    "github.com/pkg/errors"
)

const (
    ListThreads = "threads.HandleList"

    SuccessfulListThreadsMessage = "Successfully listed threads"
    ErrRetrieveDatabase          = "Failed to retrieve database in %s"
    ErrRetrieveThreads           = "Failed to retrieve threads in %s"
    ErrEncodeView                = "Failed to retrieve threads in %s"
)

func HandleList(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        categoryIDStr := chi.URLParam(r, "categoryID")
        categoryID, err := strconv.Atoi(categoryIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid category ID"), http.StatusBadRequest)
            return
        }
        
        threadList, err := threads.ListByCategoryID(db, categoryID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, fmt.Sprintf(ErrRetrieveThreads, ListThreads)), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, threadList, http.StatusOK)
    }
}

func HandleGet(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        threadIDStr := chi.URLParam(r, "id")
        threadID, err := strconv.Atoi(threadIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid thread ID"), http.StatusBadRequest)
            return
        }

        thread, err := threads.GetByID(db, threadID)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to retrieve thread"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, thread, http.StatusOK)
    }
}

func HandleCreate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var thread models.Thread
        if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }

        if err := threads.CreateThread(db, &thread); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to create thread"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, thread, http.StatusCreated)
    }
}

func HandleUpdate(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        threadIDStr := chi.URLParam(r, "id")
        threadID, err := strconv.Atoi(threadIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid thread ID"), http.StatusBadRequest)
            return
        }

        var thread models.Thread
        if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to decode request body"), http.StatusBadRequest)
            return
        }
        thread.ID = threadID

        if err := threads.UpdateThread(db, &thread); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to update thread"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, thread, http.StatusOK)
    }
}

func HandleDelete(db *database.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        threadIDStr := chi.URLParam(r, "id")
        threadID, err := strconv.Atoi(threadIDStr)
        if err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "invalid thread ID"), http.StatusBadRequest)
            return
        }

        if err := threads.DeleteThread(db, threadID); err != nil {
            api.WriteErrorResponse(w, errors.Wrap(err, "failed to delete thread"), http.StatusInternalServerError)
            return
        }

        api.WriteResponse(w, nil, http.StatusNoContent)
    }
}
